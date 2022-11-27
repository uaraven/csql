package core

import (
	"fmt"
	"github.com/uaraven/csql/util"
)

type BinaryOperator int

const (
	addition       = '+'
	subtraction    = '-'
	multiplication = '*'
	division       = '/'
	modulo         = '%'
)

func wrapInFloat(f float64, err error) (Value, error) {
	return NewFloatValue(f), err
}

func wrapInInt(i int64, err error) (Value, error) {
	return NewIntValue(i), err
}

type binaryExpression struct {
	operation BinaryOperator
	left      Evaluator
	right     Evaluator
}

func operationExecutor[T int64 | float64](oper BinaryOperator, a T, b T) (T, error) {
	switch oper {
	case addition:
		return a + b, nil
	case subtraction:
		return a - b, nil
	case multiplication:
		return a * b, nil
	case division:
		return a / b, nil
	case modulo:
		return T(int64(a) % int64(b)), nil
	default:
		var empty T
		return empty, fmt.Errorf("unsupported operation: %v", oper)
	}
}

func (ae binaryExpression) Evaluate(context EvaluationContext) Value {
	left := ae.left.Evaluate(context)
	right := ae.right.Evaluate(context)

	if !IsNumeric(left) || !IsNumeric(right) {
		panic(fmt.Errorf("both operands must be numeric: %v and %v", left, right))
	}
	if left.Type() == TypeFloat || right.Type() == TypeFloat {
		a := left.AsFloat().Value().(float64)
		b := right.AsFloat().Value().(float64)
		return util.Must(wrapInFloat(operationExecutor(ae.operation, a, b)))
	} else {
		a := left.AsInt().Value().(int64)
		b := right.AsInt().Value().(int64)
		return util.Must(wrapInInt(operationExecutor(ae.operation, a, b)))
	}
}

func NewExpression(operation BinaryOperator, left Evaluator, right Evaluator) Evaluator {
	return &binaryExpression{
		operation: operation,
		left:      left,
		right:     right,
	}
}

type concatExpression struct {
	left  Evaluator
	right Evaluator
}

func NewConcat(left Evaluator, right Evaluator) Evaluator {
	return &concatExpression{
		left:  left,
		right: right,
	}
}

func (c concatExpression) Evaluate(context EvaluationContext) Value {
	ls := EnsureString(c.left.Evaluate(context)).AsString().Value().(string)
	rs := EnsureString(c.right.Evaluate(context)).AsString().Value().(string)
	return NewStringValue(ls + rs)
}
