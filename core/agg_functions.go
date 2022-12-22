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
)

type AggregateFunction interface {
	Evaluator
	Name() string
	GetArg() Evaluator
	Location() *errors.SourceLocation
	EvaluateAgg(rows []Row) Value
}

type AggregateFunctionImplementor func(rows []Row, function AggregateFunction) Value

type aggFunctionImpl struct {
	arg         Evaluator
	name        string
	loc         *errors.SourceLocation
	implementor AggregateFunctionImplementor
}

func (sf aggFunctionImpl) Name() string {
	return sf.name
}

func (sf aggFunctionImpl) GetArg() Evaluator {
	return sf.arg
}

func (sf aggFunctionImpl) Location() *errors.SourceLocation {
	return sf.loc
}

func (sf aggFunctionImpl) Evaluate(ctx EvaluationContext) Value {
	panic(errors.NewError(sf.loc, "Aggregate function cannot be used as a scalar"))
}

func (sf aggFunctionImpl) EvaluateAgg(rows []Row) Value {
	return sf.implementor(rows, sf)
}

func (sf aggFunctionImpl) String() string {
	return fmt.Sprintf("%s(%s)", sf.name, sf.arg.String())
}

func newAggFunction(name string, implementor AggregateFunctionImplementor, loc *errors.SourceLocation, args Evaluator) AggregateFunction {
	return &aggFunctionImpl{
		name:        name,
		arg:         args,
		loc:         loc,
		implementor: implementor,
	}
}

func newCountFunction(loc *errors.SourceLocation, distinct bool, arg Evaluator) AggregateFunction {
	return &aggFunctionImpl{
		name: "COUNT",
		arg:  arg,
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			countAll := isCountAll(arg)
			duplicates := make(map[string]bool)
			counter := 0
			for _, row := range rows {
				var value Value
				if countAll {
					value = NewStringValue(row.Key())
				} else {
					value = arg.Evaluate(row)
				}
				if distinct {
					if !duplicates[value.String()] {
						counter += 1
						duplicates[value.String()] = true
					}
				} else {
					counter += 1
				}
			}
			return NewIntValue(int64(counter))
		},
	}
}

func isCountAll(evaluator Evaluator) bool {
	if rowId, ok := evaluator.(Identifiable); ok {
		_, sourceColumn := qualified(rowId.Identifier())
		return sourceColumn == AllColumns
	}
	return false
}
