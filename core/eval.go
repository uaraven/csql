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
	"github.com/uaraven/csql/funky"
)

type DataType int

const (
	TypeInt DataType = iota
	TypeFloat
	TypeString
	TypeBoolean
	TypeNull
)

// EvaluationContext provides a way to retrieve value by name during evaluation
type EvaluationContext interface {
	Get(string) funky.Option[Value]
}

type AggregateContext interface {
	EvaluationContext
	RowSet() []Row
}

type Evaluator interface {
	fmt.Stringer
	Evaluate(EvaluationContext) Value
}

type AggregateEvaluator interface {
	Evaluator
	AggregateEvaluate(ctx AggregateContext) Value
}

type Condition interface {
	Evaluator
}

type Value interface {
	Evaluator
	fmt.Stringer
	Type() DataType
	Value() interface{}
	AsString() Value
	AsInt() Value
	AsFloat() Value
	AsBool() Value
	Equals(Value) bool
	Less(Value) bool
	Repr() string
}

func IsNumeric(v Value) bool {
	return v.Type() == TypeFloat || v.Type() == TypeInt
}
