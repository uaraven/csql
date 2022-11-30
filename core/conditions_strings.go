package core

import (
	"fmt"
	"regexp"
	"strings"
)

// matchCondition matches left expression against the right expression, which
// must resolve to a string value and is treated as regular expression
type matchCondition struct {
	Condition
	left  Evaluator
	right Evaluator
}

func NewMatchCondition(left Evaluator, right Evaluator) Condition {
	return &matchCondition{
		left:  left,
		right: right,
	}
}

func (mc matchCondition) Evaluate(context EvaluationContext) Value {
	expression := mc.right.Evaluate(context).AsString()
	value := mc.left.Evaluate(context).AsString()

	matched, err := regexp.MatchString(expression.Value().(string), value.Value().(string))
	if err != nil {
		panic(fmt.Sprintf("failure while evaluating %v MATCH %v: %v", mc.left, mc.right, err))
	}

	return NewBoolValue(matched)
}

// likeCondition matches left expression against the right expression, which
// must resolve to a string value and is treated as SQL like syntax
type likeCondition struct {
	Condition
	left  Evaluator
	right Evaluator
}

func NewLikeCondition(left Evaluator, right Evaluator) Condition {
	return &likeCondition{
		left:  left,
		right: right,
	}
}

func (lc likeCondition) Evaluate(context EvaluationContext) Value {
	expression := lc.right.Evaluate(context).AsString()
	value := lc.left.Evaluate(context).AsString()

	likeExpr := translateLike(expression)

	matched, err := regexp.MatchString(likeExpr, value.Value().(string))
	if err != nil {
		panic(fmt.Sprintf("failure while evaluating %v LIKE %v: %v", lc.left, lc.right, err))
	}

	return NewBoolValue(matched)
}

func translateLike(value Value) string {
	val := "^" + regexp.QuoteMeta(value.Value().(string)) + "$"
	return strings.ReplaceAll(strings.ReplaceAll(val, "%", ".*?"), "_", ".")
}
