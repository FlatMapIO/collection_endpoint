// Code generated from Filter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Filter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FilterListener is a complete listener for a parse tree produced by FilterParser.
type FilterListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterArrayValue is called when entering the arrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// EnterNumberArrary is called when entering the numberArrary production.
	EnterNumberArrary(c *NumberArraryContext)

	// EnterStringArray is called when entering the stringArray production.
	EnterStringArray(c *StringArrayContext)

	// EnterPrimaryValue is called when entering the primaryValue production.
	EnterPrimaryValue(c *PrimaryValueContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitArrayValue is called when exiting the arrayValue production.
	ExitArrayValue(c *ArrayValueContext)

	// ExitNumberArrary is called when exiting the numberArrary production.
	ExitNumberArrary(c *NumberArraryContext)

	// ExitStringArray is called when exiting the stringArray production.
	ExitStringArray(c *StringArrayContext)

	// ExitPrimaryValue is called when exiting the primaryValue production.
	ExitPrimaryValue(c *PrimaryValueContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)
}
