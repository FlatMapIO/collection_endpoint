// Code generated from Filter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Filter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFilterListener is a complete listener for a parse tree produced by FilterParser.
type BaseFilterListener struct{}

var _ FilterListener = &BaseFilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFilterListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFilterListener) ExitExpr(ctx *ExprContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseFilterListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseFilterListener) ExitFilter(ctx *FilterContext) {}

// EnterField is called when production field is entered.
func (s *BaseFilterListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseFilterListener) ExitField(ctx *FieldContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFilterListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFilterListener) ExitValue(ctx *ValueContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseFilterListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseFilterListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterNumberArrary is called when production numberArrary is entered.
func (s *BaseFilterListener) EnterNumberArrary(ctx *NumberArraryContext) {}

// ExitNumberArrary is called when production numberArrary is exited.
func (s *BaseFilterListener) ExitNumberArrary(ctx *NumberArraryContext) {}

// EnterStringArray is called when production stringArray is entered.
func (s *BaseFilterListener) EnterStringArray(ctx *StringArrayContext) {}

// ExitStringArray is called when production stringArray is exited.
func (s *BaseFilterListener) ExitStringArray(ctx *StringArrayContext) {}

// EnterPrimaryValue is called when production primaryValue is entered.
func (s *BaseFilterListener) EnterPrimaryValue(ctx *PrimaryValueContext) {}

// ExitPrimaryValue is called when production primaryValue is exited.
func (s *BaseFilterListener) ExitPrimaryValue(ctx *PrimaryValueContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseFilterListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseFilterListener) ExitCondition(ctx *ConditionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseFilterListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseFilterListener) ExitOperator(ctx *OperatorContext) {}
