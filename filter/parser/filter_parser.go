// Code generated from Filter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Filter

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FilterParser struct {
	*antlr.BaseParser
}

var filterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func filterParserInit() {
	staticData := &filterParserStaticData
	staticData.literalNames = []string{
		"", "'['", "']'", "'='", "'('", "','", "')'", "'eq'", "'ne'", "'gt'",
		"'ge'", "'lt'", "'le'", "'in'", "'nin'", "'like'", "'nlike'", "'between'",
		"'nbetween'", "", "", "'|'", "'&'", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "BOOL_TRUR", "BOOL_FALSE", "OR", "AND", "NULL", "STRING", "WORD",
		"NUMBER", "WS",
	}
	staticData.ruleNames = []string{
		"expr", "filter", "field", "value", "arrayValue", "numberArrary", "stringArray",
		"primaryValue", "condition", "operator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 89, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 0, 5, 0, 25, 8, 0, 10, 0, 12, 0, 28, 9, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 42, 8, 3, 1,
		4, 1, 4, 3, 4, 46, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 52, 8, 5, 10, 5,
		12, 5, 55, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 63, 8, 6, 10,
		6, 12, 6, 66, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 87,
		8, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 2, 3, 0, 19,
		20, 24, 24, 26, 26, 1, 0, 21, 22, 96, 0, 20, 1, 0, 0, 0, 2, 29, 1, 0, 0,
		0, 4, 36, 1, 0, 0, 0, 6, 41, 1, 0, 0, 0, 8, 45, 1, 0, 0, 0, 10, 47, 1,
		0, 0, 0, 12, 58, 1, 0, 0, 0, 14, 69, 1, 0, 0, 0, 16, 71, 1, 0, 0, 0, 18,
		86, 1, 0, 0, 0, 20, 26, 3, 2, 1, 0, 21, 22, 3, 16, 8, 0, 22, 23, 3, 2,
		1, 0, 23, 25, 1, 0, 0, 0, 24, 21, 1, 0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24,
		1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0,
		29, 30, 3, 4, 2, 0, 30, 31, 5, 1, 0, 0, 31, 32, 3, 18, 9, 0, 32, 33, 5,
		2, 0, 0, 33, 34, 5, 3, 0, 0, 34, 35, 3, 6, 3, 0, 35, 3, 1, 0, 0, 0, 36,
		37, 5, 25, 0, 0, 37, 5, 1, 0, 0, 0, 38, 42, 3, 14, 7, 0, 39, 42, 5, 23,
		0, 0, 40, 42, 3, 8, 4, 0, 41, 38, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 40,
		1, 0, 0, 0, 42, 7, 1, 0, 0, 0, 43, 46, 3, 10, 5, 0, 44, 46, 3, 12, 6, 0,
		45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 9, 1, 0, 0, 0, 47, 48, 5, 4,
		0, 0, 48, 53, 5, 26, 0, 0, 49, 50, 5, 5, 0, 0, 50, 52, 5, 26, 0, 0, 51,
		49, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0,
		0, 54, 56, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 57, 5, 6, 0, 0, 57, 11,
		1, 0, 0, 0, 58, 59, 5, 4, 0, 0, 59, 64, 5, 24, 0, 0, 60, 61, 5, 5, 0, 0,
		61, 63, 5, 24, 0, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67,
		68, 5, 6, 0, 0, 68, 13, 1, 0, 0, 0, 69, 70, 7, 0, 0, 0, 70, 15, 1, 0, 0,
		0, 71, 72, 7, 1, 0, 0, 72, 17, 1, 0, 0, 0, 73, 87, 1, 0, 0, 0, 74, 87,
		5, 7, 0, 0, 75, 87, 5, 8, 0, 0, 76, 87, 5, 9, 0, 0, 77, 87, 5, 10, 0, 0,
		78, 87, 5, 11, 0, 0, 79, 87, 5, 12, 0, 0, 80, 87, 5, 13, 0, 0, 81, 87,
		5, 14, 0, 0, 82, 87, 5, 15, 0, 0, 83, 87, 5, 16, 0, 0, 84, 87, 5, 17, 0,
		0, 85, 87, 5, 18, 0, 0, 86, 73, 1, 0, 0, 0, 86, 74, 1, 0, 0, 0, 86, 75,
		1, 0, 0, 0, 86, 76, 1, 0, 0, 0, 86, 77, 1, 0, 0, 0, 86, 78, 1, 0, 0, 0,
		86, 79, 1, 0, 0, 0, 86, 80, 1, 0, 0, 0, 86, 81, 1, 0, 0, 0, 86, 82, 1,
		0, 0, 0, 86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87,
		19, 1, 0, 0, 0, 6, 26, 41, 45, 53, 64, 86,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FilterParserInit initializes any static state used to implement FilterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFilterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterParserInit() {
	staticData := &filterParserStaticData
	staticData.once.Do(filterParserInit)
}

// NewFilterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFilterParser(input antlr.TokenStream) *FilterParser {
	FilterParserInit()
	this := new(FilterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &filterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Filter.g4"

	return this
}

// FilterParser tokens.
const (
	FilterParserEOF        = antlr.TokenEOF
	FilterParserT__0       = 1
	FilterParserT__1       = 2
	FilterParserT__2       = 3
	FilterParserT__3       = 4
	FilterParserT__4       = 5
	FilterParserT__5       = 6
	FilterParserT__6       = 7
	FilterParserT__7       = 8
	FilterParserT__8       = 9
	FilterParserT__9       = 10
	FilterParserT__10      = 11
	FilterParserT__11      = 12
	FilterParserT__12      = 13
	FilterParserT__13      = 14
	FilterParserT__14      = 15
	FilterParserT__15      = 16
	FilterParserT__16      = 17
	FilterParserT__17      = 18
	FilterParserBOOL_TRUR  = 19
	FilterParserBOOL_FALSE = 20
	FilterParserOR         = 21
	FilterParserAND        = 22
	FilterParserNULL       = 23
	FilterParserSTRING     = 24
	FilterParserWORD       = 25
	FilterParserNUMBER     = 26
	FilterParserWS         = 27
)

// FilterParser rules.
const (
	FilterParserRULE_expr         = 0
	FilterParserRULE_filter       = 1
	FilterParserRULE_field        = 2
	FilterParserRULE_value        = 3
	FilterParserRULE_arrayValue   = 4
	FilterParserRULE_numberArrary = 5
	FilterParserRULE_stringArray  = 6
	FilterParserRULE_primaryValue = 7
	FilterParserRULE_condition    = 8
	FilterParserRULE_operator     = 9
)

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllFilter() []IFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterContext); ok {
			len++
		}
	}

	tst := make([]IFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterContext); ok {
			tst[i] = t.(IFilterContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Filter(i int) IFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *ExprContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *FilterParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterParserRULE_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Filter()
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterParserOR || _la == FilterParserAND {
		{
			p.SetState(21)
			p.Condition()
		}

		{
			p.SetState(22)
			p.Filter()
		}

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FilterContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *FilterContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *FilterParser) Filter() (localctx IFilterContext) {
	this := p
	_ = this

	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FilterParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Field()
	}
	{
		p.SetState(30)
		p.Match(FilterParserT__0)
	}
	{
		p.SetState(31)
		p.Operator()
	}
	{
		p.SetState(32)
		p.Match(FilterParserT__1)
	}
	{
		p.SetState(33)
		p.Match(FilterParserT__2)
	}
	{
		p.SetState(34)
		p.Value()
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) WORD() antlr.TerminalNode {
	return s.GetToken(FilterParserWORD, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *FilterParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(FilterParserWORD)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) PrimaryValue() IPrimaryValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryValueContext)
}

func (s *ValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(FilterParserNULL, 0)
}

func (s *ValueContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *FilterParser) Value() (localctx IValueContext) {
	this := p
	_ = this

	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterParserBOOL_TRUR, FilterParserBOOL_FALSE, FilterParserSTRING, FilterParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.PrimaryValue()
		}

	case FilterParserNULL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(FilterParserNULL)
		}

	case FilterParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.ArrayValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_arrayValue
	return p
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) NumberArrary() INumberArraryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberArraryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberArraryContext)
}

func (s *ArrayValueContext) StringArray() IStringArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringArrayContext)
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (p *FilterParser) ArrayValue() (localctx IArrayValueContext) {
	this := p
	_ = this

	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterParserRULE_arrayValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.NumberArrary()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.StringArray()
		}

	}

	return localctx
}

// INumberArraryContext is an interface to support dynamic dispatch.
type INumberArraryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberArraryContext differentiates from other interfaces.
	IsNumberArraryContext()
}

type NumberArraryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberArraryContext() *NumberArraryContext {
	var p = new(NumberArraryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_numberArrary
	return p
}

func (*NumberArraryContext) IsNumberArraryContext() {}

func NewNumberArraryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberArraryContext {
	var p = new(NumberArraryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_numberArrary

	return p
}

func (s *NumberArraryContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberArraryContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(FilterParserNUMBER)
}

func (s *NumberArraryContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(FilterParserNUMBER, i)
}

func (s *NumberArraryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberArraryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberArraryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterNumberArrary(s)
	}
}

func (s *NumberArraryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitNumberArrary(s)
	}
}

func (p *FilterParser) NumberArrary() (localctx INumberArraryContext) {
	this := p
	_ = this

	localctx = NewNumberArraryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterParserRULE_numberArrary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(FilterParserT__3)
	}
	{
		p.SetState(48)
		p.Match(FilterParserNUMBER)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterParserT__4 {
		{
			p.SetState(49)
			p.Match(FilterParserT__4)
		}
		{
			p.SetState(50)
			p.Match(FilterParserNUMBER)
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(FilterParserT__5)
	}

	return localctx
}

// IStringArrayContext is an interface to support dynamic dispatch.
type IStringArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringArrayContext differentiates from other interfaces.
	IsStringArrayContext()
}

type StringArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringArrayContext() *StringArrayContext {
	var p = new(StringArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_stringArray
	return p
}

func (*StringArrayContext) IsStringArrayContext() {}

func NewStringArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringArrayContext {
	var p = new(StringArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_stringArray

	return p
}

func (s *StringArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *StringArrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(FilterParserSTRING)
}

func (s *StringArrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(FilterParserSTRING, i)
}

func (s *StringArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterStringArray(s)
	}
}

func (s *StringArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitStringArray(s)
	}
}

func (p *FilterParser) StringArray() (localctx IStringArrayContext) {
	this := p
	_ = this

	localctx = NewStringArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FilterParserRULE_stringArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(FilterParserT__3)
	}
	{
		p.SetState(59)
		p.Match(FilterParserSTRING)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FilterParserT__4 {
		{
			p.SetState(60)
			p.Match(FilterParserT__4)
		}
		{
			p.SetState(61)
			p.Match(FilterParserSTRING)
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(FilterParserT__5)
	}

	return localctx
}

// IPrimaryValueContext is an interface to support dynamic dispatch.
type IPrimaryValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryValueContext differentiates from other interfaces.
	IsPrimaryValueContext()
}

type PrimaryValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryValueContext() *PrimaryValueContext {
	var p = new(PrimaryValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_primaryValue
	return p
}

func (*PrimaryValueContext) IsPrimaryValueContext() {}

func NewPrimaryValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryValueContext {
	var p = new(PrimaryValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_primaryValue

	return p
}

func (s *PrimaryValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterParserSTRING, 0)
}

func (s *PrimaryValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FilterParserNUMBER, 0)
}

func (s *PrimaryValueContext) BOOL_TRUR() antlr.TerminalNode {
	return s.GetToken(FilterParserBOOL_TRUR, 0)
}

func (s *PrimaryValueContext) BOOL_FALSE() antlr.TerminalNode {
	return s.GetToken(FilterParserBOOL_FALSE, 0)
}

func (s *PrimaryValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterPrimaryValue(s)
	}
}

func (s *PrimaryValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitPrimaryValue(s)
	}
}

func (p *FilterParser) PrimaryValue() (localctx IPrimaryValueContext) {
	this := p
	_ = this

	localctx = NewPrimaryValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FilterParserRULE_primaryValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FilterParserBOOL_TRUR)|(1<<FilterParserBOOL_FALSE)|(1<<FilterParserSTRING)|(1<<FilterParserNUMBER))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) OR() antlr.TerminalNode {
	return s.GetToken(FilterParserOR, 0)
}

func (s *ConditionContext) AND() antlr.TerminalNode {
	return s.GetToken(FilterParserAND, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *FilterParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FilterParserRULE_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FilterParserOR || _la == FilterParserAND) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FilterParserRULE_operator
	return p
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *FilterParser) Operator() (localctx IOperatorContext) {
	this := p
	_ = this

	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FilterParserRULE_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FilterParserT__1:
		p.EnterOuterAlt(localctx, 1)

	case FilterParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(FilterParserT__6)
		}

	case FilterParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(FilterParserT__7)
		}

	case FilterParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(76)
			p.Match(FilterParserT__8)
		}

	case FilterParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)
			p.Match(FilterParserT__9)
		}

	case FilterParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(78)
			p.Match(FilterParserT__10)
		}

	case FilterParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(79)
			p.Match(FilterParserT__11)
		}

	case FilterParserT__12:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(80)
			p.Match(FilterParserT__12)
		}

	case FilterParserT__13:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(81)
			p.Match(FilterParserT__13)
		}

	case FilterParserT__14:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(82)
			p.Match(FilterParserT__14)
		}

	case FilterParserT__15:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(83)
			p.Match(FilterParserT__15)
		}

	case FilterParserT__16:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(84)
			p.Match(FilterParserT__16)
		}

	case FilterParserT__17:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(85)
			p.Match(FilterParserT__17)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
