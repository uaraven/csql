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
)

type andCondition struct {
	left  Evaluator
	right Evaluator
}

func NewAnd(left Evaluator, right Evaluator) Evaluator {
	return &andCondition{
		left:  left,
		right: right,
	}
}

func (ac andCondition) String() string {
	return fmt.Sprintf("%s AND %s", ac.left, ac.right)
}

func (ac andCondition) Evaluate(ctx EvaluationContext) Value {
	if ac.left.Evaluate(ctx).AsBool().Value().(bool) {
		return ac.right.Evaluate(ctx).AsBool()
	} else {
		return ValueFalse
	}
}

type orCondition struct {
	left  Evaluator
	right Evaluator
}

func NewOr(left Evaluator, right Evaluator) Evaluator {
	return &orCondition{
		left:  left,
		right: right,
	}
}

func (oc orCondition) String() string {
	return fmt.Sprintf("%s OR %s", oc.left, oc.right)
}

func (oc orCondition) Evaluate(ctx EvaluationContext) Value {
	if oc.left.Evaluate(ctx).AsBool().Value().(bool) {
		return ValueTrue
	} else {
		return oc.right.Evaluate(ctx).AsBool()
	}
}

type notCondition struct {
	arg Evaluator
}

func NewNot(arg Evaluator) Evaluator {
	return &notCondition{arg: arg}
}

func (nc notCondition) String() string {
	return fmt.Sprintf("NOT %s", nc.arg)
}

func (nc notCondition) Evaluate(ctx EvaluationContext) Value {
	return NewBoolValue(!nc.arg.Evaluate(ctx).AsBool().Value().(bool))
}
