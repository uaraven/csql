package data

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type compareOperator int

const (
	equal compareOperator = iota
	notEqual
	greaterThan
	lessThan
	greaterEqual
	lessEqual
)

var operatorStr = map[compareOperator]string{
	equal:        "=",
	notEqual:     "<>",
	greaterEqual: ">=",
	greaterThan:  ">",
	lessEqual:    "<=",
	lessThan:     "<",
}

func getComparator[T constraints.Ordered](l T, operator compareOperator) func(T) bool {
	switch operator {
	case equal:
		return func(r T) bool { return compare(l, r) == 0 }
	case notEqual:
		return func(r T) bool { return compare(l, r) != 0 }
	case greaterThan:
		return func(r T) bool { return compare(l, r) > 0 }
	case lessThan:
		return func(r T) bool { return compare(l, r) < 0 }
	case greaterEqual:
		return func(r T) bool { return compare(l, r) >= 0 }
	case lessEqual:
		return func(r T) bool { return compare(l, r) <= 0 }
	default:
		panic(fmt.Errorf("unsupported operator code: %v", operator))
	}
}

func compare[T constraints.Ordered](left T, right T) int {
	if left > right {
		return 1
	} else if left < right {
		return -1
	} else {
		return 0
	}
}

type ComparisonCondition struct {
	Condition
	operator compareOperator
	left     Value
	right    Value
}

func (cc ComparisonCondition) Evaluate() Value {
	switch cc.left.Type() {
	case StringType:
		return NewBoolValue(getComparator(cc.left.Value().(string), cc.operator)(cc.right.Value().(string)))
	case IntType:
		return NewBoolValue(getComparator(cc.left.Value().(int64), cc.operator)(cc.right.Value().(int64)))
	case FloatType:
		return NewBoolValue(getComparator(cc.left.Value().(float64), cc.operator)(cc.right.Value().(float64)))
	default:
		panic(fmt.Errorf("unsupported type in expression %v", cc))
	}
}

func (cc ComparisonCondition) String() string {
	return fmt.Sprintf("%s %s %s", cc.left, operatorStr[cc.operator], cc.right)
}

func NewEq(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: equal,
		left:     left,
		right:    right,
	}
}

func NewNeq(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: notEqual,
		left:     left,
		right:    right,
	}
}

func NewLt(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: lessThan,
		left:     left,
		right:    right,
	}
}

func NewGt(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: greaterThan,
		left:     left,
		right:    right,
	}
}

func NewLte(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: lessEqual,
		left:     left,
		right:    right,
	}
}

func NewGte(left Value, right Value) Condition {
	return ComparisonCondition{
		operator: greaterEqual,
		left:     left,
		right:    right,
	}
}
