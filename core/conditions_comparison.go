/*
 *    Copyright 2022 Oleksiy Voronin
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 *    SPDX-License-Identifier: Apache-2.0
 *    SPDX-FileCopyrightText: (c) 2022 Oleksiy Voronin <ovoronin@gmail.com>
 */

package core

import (
	"fmt"
	"github.com/uaraven/csql/errors"
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
	left     Evaluator
	right    Evaluator
	loc      *errors.SourceLocation
}

func (cc ComparisonCondition) Evaluate(ctx EvaluationContext) Value {
	lv := cc.left.Evaluate(ctx)
	rv := cc.right.Evaluate(ctx)
	return NewBoolValue(cc.compareValues(lv, rv))
}

// numericOperand returns v as a number, converting strings that look like numbers.
func numericOperand(v Value) (Value, bool) {
	switch v.Type() {
	case TypeInt, TypeFloat:
		return v, true
	case TypeString:
		return parseNumber(v.Value().(string))
	}
	return nil, false
}

// compareValues applies the operator to two values. Numbers compare as numbers (as floats if either is a float),
// strings compare lexically, and a string compared to a number is coerced if it parses as a number. Any comparison
// with NULL is false. Operands that cannot be compared are never equal and never ordered.
func (cc ComparisonCondition) compareValues(lv Value, rv Value) bool {
	if lv.Type() == TypeNull || rv.Type() == TypeNull {
		return false
	}
	if lv.Type() == TypeString && rv.Type() == TypeString {
		return getComparator(lv.Value().(string), cc.operator)(rv.Value().(string))
	}
	ln, lok := numericOperand(lv)
	rn, rok := numericOperand(rv)
	if lok && rok {
		if ln.Type() == TypeInt && rn.Type() == TypeInt {
			return getComparator(ln.Value().(int64), cc.operator)(rn.Value().(int64))
		}
		return getComparator(ln.AsFloat().Value().(float64), cc.operator)(rn.AsFloat().Value().(float64))
	}
	if lv.Type() == TypeBoolean && rv.Type() == TypeBoolean {
		switch cc.operator {
		case equal:
			return lv.Equals(rv)
		case notEqual:
			return !lv.Equals(rv)
		}
		panic(errors.NewError(cc.loc, fmt.Sprintf("unsupported operator for boolean operands in %s", cc)))
	}
	return cc.operator == notEqual
}

func (cc ComparisonCondition) String() string {
	return fmt.Sprintf("%s %s %s", cc.left, operatorStr[cc.operator], cc.right)
}

func WithLoc(condition Condition, loc *errors.SourceLocation) Condition {
	if compCond, ok := condition.(ComparisonCondition); ok {
		return ComparisonCondition{
			operator: compCond.operator,
			left:     compCond.left,
			right:    compCond.right,
			loc:      loc,
		}
	} else {
		return condition
	}
}

func NewEq(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: equal,
		left:     left,
		right:    right,
	}
}

func NewNeq(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: notEqual,
		left:     left,
		right:    right,
	}
}

func NewLt(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: lessThan,
		left:     left,
		right:    right,
	}
}

func NewGt(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: greaterThan,
		left:     left,
		right:    right,
	}
}

func NewLte(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: lessEqual,
		left:     left,
		right:    right,
	}
}

func NewGte(left Evaluator, right Evaluator) Condition {
	return ComparisonCondition{
		operator: greaterEqual,
		left:     left,
		right:    right,
	}
}
