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

import "github.com/uaraven/csql/funky"

type inCondition struct {
	left  Evaluator
	right []Evaluator
}

func NewInList(left Evaluator, right []Evaluator) Evaluator {
	return &inCondition{
		left:  left,
		right: right,
	}
}

func (ie inCondition) Evaluate(context EvaluationContext) Value {
	return NewBoolValue(funky.AnyMatches(ie.right, func(v Evaluator) bool {
		return ie.left.Evaluate(context).Equals(v.Evaluate(context))
	}))
}
