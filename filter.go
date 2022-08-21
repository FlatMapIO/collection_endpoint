package collection_endpoint

import (
	"errors"
	"github.com/FlatMapIO/collection_endpoint/filter/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type filterVisitor struct {
	parser.BaseFilterListener

	antlr.DefaultErrorListener

	specs       map[string]FilterSpec
	nodes       map[string]*filterNode
	prevCond    string
	prevNode    *filterNode
	sqlAppender *strings.Builder

	completed bool

	err error
}
type filterNode struct {
	field string
	op    string

	stmt string
}

var operatorMap = map[string]string{
	"eq":       "=",
	"ne":       "!=",
	"gt":       ">",
	"ge":       ">=",
	"lt":       "<",
	"le":       "<=",
	"in":       "in",
	"nin":      "not in",
	"like":     "like",
	"nlike":    "not like",
	"between":  "between",
	"nbetween": "not between",
}

func parseAndWriteConditions(rawQuery string, specs map[string]FilterSpec, appender *strings.Builder) error {
	stream := antlr.NewInputStream(rawQuery)
	lex := parser.NewFilterLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)

	p := parser.NewFilterParser(tokenStream)

	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	visitor := filterVisitor{
		specs:       specs,
		nodes:       map[string]*filterNode{},
		sqlAppender: appender,
	}
	p.AddErrorListener(&visitor)

	antlr.ParseTreeWalkerDefault.Walk(&visitor, p.Expr())

	if visitor.err != nil {
		return visitor.err
	}
	return nil
}

func (f *filterVisitor) EnterExpr(ctx *parser.ExprContext) {

}
func (f *filterVisitor) EnterCondition(ctx *parser.ConditionContext) {
	if ctx.OR() != nil {
		f.prevCond = "or"
	}
	if ctx.AND() != nil {
		f.prevCond = "and"
	}
}
func (f *filterVisitor) EnterFilter(ctx *parser.FilterContext) {
	if f.err != nil {
		return
	}
	field := ctx.GetChild(0).(*parser.FieldContext).GetText()
	operator := ctx.GetChild(2).(*parser.OperatorContext).GetText()
	value := ctx.GetChild(5).(*parser.ValueContext)

	if spec, ok := f.specs[field]; ok {
		if !spec.opSet.has(operator) {
			f.err = errors.New("operator " + operator + " is not allowed for field " + field)
			return
		}
	}

	err := errors.New("value is not specified for operator " + operator)

	switch operator {
	case "eq", "ne":
		if value.NULL() != nil {
			switch operator {
			case "eq":
				node := filterNode{field: field, op: operator, stmt: "is null"}
				f.nodes[field] = &node
				f.appendCondition()
			case "ne":
				node := filterNode{field: field, op: operator, stmt: "is not null"}
				f.nodes[field] = &node
				f.prevNode = &node
			}
			break
		}
		if value.PrimaryValue() != nil {
			pv := value.PrimaryValue().(*parser.PrimaryValueContext)
			if pv.STRING() != nil {
				txt := pv.GetText()
				node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " '" + txt + "'"}
				f.nodes[field] = &node
				f.prevNode = &node
				break
			}
			if pv.NUMBER() != nil {
				txt := pv.NUMBER().GetText()
				node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " " + txt}
				f.nodes[field] = &node
				f.prevNode = &node
				break
			}
			if pv.BOOL_TRUR() != nil {
				node := filterNode{field: field, op: operator, stmt: "= true"}
				f.nodes[field] = &node
				f.prevNode = &node
				break
			}
			if pv.BOOL_FALSE() != nil {
				node := filterNode{field: field, op: operator, stmt: "= false"}
				f.nodes[field] = &node
				f.prevNode = &node
				break
			}
		}

		goto errret
	case "gt", "ge", "lt", "le":
		if value.PrimaryValue() == nil {
			goto errret
		}
		pv := value.PrimaryValue().(*parser.PrimaryValueContext)
		if pvs := pv.STRING(); pvs != nil {
			txt := pvs.GetText()
			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " '" + txt + "'"}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}

		if pvn := pv.NUMBER(); pvn != nil {
			txt := pvn.GetText()
			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " " + txt}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}
		goto errret
	case "in", "nin":
		if value.ArrayValue() == nil {
			goto errret
		}
		arr := value.ArrayValue().(*parser.ArrayValueContext)
		if arrs := arr.StringArray(); arrs != nil {
			var list []string
			for _, s := range arrs.(*parser.StringArrayContext).AllSTRING() {
				list = append(list, "'"+s.GetText()+"'")
			}
			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " (" + strings.Join(list, ", ") + ")"}
			f.nodes[field] = &node
			f.prevNode = &node

			break
		}
		if arrn := arr.NumberArrary(); arrn != nil {
			var list []string
			for _, n := range arrn.(*parser.NumberArraryContext).AllNUMBER() {
				list = append(list, n.GetText())
			}

			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " (" + strings.Join(list, ", ") + ")"}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}
		goto errret
	case "like", "nlike":
		if value.PrimaryValue() == nil {
			goto errret
		}
		if str := value.PrimaryValue().(*parser.PrimaryValueContext).STRING(); str != nil {
			txt := str.GetText()
			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " '%" + txt + "%'"}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}
		goto errret
	case "between", "nbetween":
		if value.ArrayValue() == nil {
			goto errret
		}
		arr := value.ArrayValue().(*parser.ArrayValueContext)
		if arr.StringArray() != nil {
			strs := arr.StringArray().(*parser.StringArrayContext).AllSTRING()
			if len(strs) != 2 {
				err = errors.New("between operator requires two string values")
				goto errret
			}
			s1, s2 := strs[0].GetText(), strs[1].GetText()

			if len(s1) <= 2 || len(s2) <= 2 {
				err = errors.New("value is empty")
				goto errret
			}
			s1 = s1[1 : len(s1)-1]
			s2 = s2[1 : len(s2)-1]

			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " '" + s1 + "' and '" + s2 + "'"}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}
		if arr.NumberArrary() != nil {
			nums := arr.NumberArrary().(*parser.NumberArraryContext).AllNUMBER()
			if len(nums) != 2 {
				err = errors.New("between operator requires two number values")
				goto errret
			}
			node := filterNode{field: field, op: operator, stmt: operatorMap[operator] + " " + nums[0].GetText() + " and " + nums[1].GetText()}
			f.nodes[field] = &node
			f.prevNode = &node
			break
		}
		goto errret
	default:
		err = errors.New("operator " + operator + " is not supported")
		goto errret
	}

	f.appendCondition()
	return
errret:
	f.err = err
	return
}

func (f *filterVisitor) appendCondition() {
	if f.prevCond == "" {
		f.sqlAppender.WriteString(" and " + f.prevNode.field + " " + f.prevNode.stmt)
	} else {
		f.sqlAppender.WriteString(" " + f.prevCond + " " + f.prevNode.field + " " + f.prevNode.stmt)
	}
}

func (f *filterVisitor) ExitExpr(ctx *parser.ExprContext) {
	if f.err != nil {
		return
	}
	for field, spec := range f.specs {
		if spec.Required {
			if f.nodes[field] == nil {
				f.err = errors.New("required field " + field + " is not specified")
				return
			}
		}
	}
	f.completed = true
}

func (f *filterVisitor) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	f.err = errors.New(msg)
}