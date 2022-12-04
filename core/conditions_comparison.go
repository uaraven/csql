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
}

func (cc ComparisonCondition) Evaluate(ctx EvaluationContext) Value {
	lv := cc.left.Evaluate(ctx)
	rv := cc.right.Evaluate(ctx)
	switch lv.Type() {
	case TypeString:
		return NewBoolValue(getComparator(lv.Value().(string), cc.operator)(rv.Value().(string)))
	case TypeInt:
		return NewBoolValue(getComparator(lv.Value().(int64), cc.operator)(rv.Value().(int64)))
	case TypeFloat:
		return NewBoolValue(getComparator(lv.Value().(float64), cc.operator)(rv.Value().(float64)))
	case TypeNull:
		return NewBoolValue(false)
	default:
		panic(fmt.Errorf("unsupported type in expression %v", cc))
	}
}

func (cc ComparisonCondition) String() string {
	return fmt.Sprintf("%s %s %s", cc.left, operatorStr[cc.operator], cc.right)
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
