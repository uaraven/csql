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
	"regexp"
	"strings"
)

// matchCondition matches left expression against the right expression, which
// must resolve to a string value and is treated as regular expression
type matchCondition struct {
	Condition
	loc   *errors.SourceLocation
	left  Evaluator
	right Evaluator
}

func NewMatchCondition(left Evaluator, right Evaluator) Condition {
	return &matchCondition{
		left:  left,
		right: right,
	}
}

func NewMatchConditionL(left Evaluator, right Evaluator, l *errors.SourceLocation) Condition {
	return &matchCondition{
		left:  left,
		right: right,
		loc:   l,
	}
}

func (mc matchCondition) Evaluate(context EvaluationContext) Value {
	expression := mc.right.Evaluate(context).AsString()
	value := mc.left.Evaluate(context).AsString()

	matched, err := regexp.MatchString(expression.Value().(string), value.Value().(string))
	if err != nil {
		panic(errors.NewError(mc.loc, fmt.Sprintf("failure while evaluating %v MATCH %v: %v", mc.left, mc.right, err)))
	}

	return NewBoolValue(matched)
}

// likeCondition matches left expression against the right expression, which
// must resolve to a string value and is treated as SQL like syntax
type likeCondition struct {
	Condition
	left  Evaluator
	right Evaluator
	loc   *errors.SourceLocation
}

func NewLikeCondition(left Evaluator, right Evaluator) Condition {
	return &likeCondition{
		left:  left,
		right: right,
	}
}

func NewLikeConditionL(left Evaluator, right Evaluator, location *errors.SourceLocation) Condition {
	return &likeCondition{
		left:  left,
		right: right,
		loc:   location,
	}
}

func (lc likeCondition) Evaluate(context EvaluationContext) Value {
	expression := lc.right.Evaluate(context).AsString()
	value := lc.left.Evaluate(context).AsString()

	likeExpr := translateLike(expression)

	matched, err := regexp.MatchString(likeExpr, value.Value().(string))
	if err != nil {
		panic(errors.NewError(lc.loc, fmt.Sprintf("Failure while evaluating %v LIKE %v: %v", lc.left, lc.right, err)))
	}

	return NewBoolValue(matched)
}

func translateLike(value Value) string {
	val := "^" + regexp.QuoteMeta(value.Value().(string)) + "$"
	return strings.ReplaceAll(strings.ReplaceAll(val, "%", ".*?"), "_", ".")
}
