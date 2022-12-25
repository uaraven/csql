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
	"strings"
)

type isNullCondition struct {
	expr   Evaluator
	negate bool
}

func NewIsNull(expr Evaluator, negate bool) Evaluator {
	return &isNullCondition{
		expr:   expr,
		negate: negate,
	}
}

func (inc isNullCondition) String() string {
	var sb strings.Builder
	sb.WriteString(inc.expr.String())
	sb.WriteString(" IS ")
	if inc.negate {
		sb.WriteString("NOT ")
	}
	sb.WriteString("NULL")
	return sb.String()
}

func (inc isNullCondition) Evaluate(context EvaluationContext) Value {
	result := inc.expr.Evaluate(context)
	if inc.negate {
		return NewBoolValue(result != nullV)
	} else {
		return NewBoolValue(result == nullV)
	}
}
