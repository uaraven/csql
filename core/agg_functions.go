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

type AggregationFunction interface {
	Evaluator
	GetArgs() []Evaluator
	Location() *errors.SourceLocation
	EvaluateAgg(rows []Row) Value
}

type AggregationFunctionImplementor func(rows []Row, function AggregationFunction) Value

type aggFunctionImpl struct {
	args        []Evaluator
	loc         *errors.SourceLocation
	implementor AggregationFunctionImplementor
}

func (sf aggFunctionImpl) GetArgs() []Evaluator {
	return sf.args
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

func newAggFunction(implementor AggregationFunctionImplementor, loc *errors.SourceLocation, args ...Evaluator) AggregationFunction {
	return &aggFunctionImpl{
		args:        args,
		loc:         loc,
		implementor: implementor,
	}
}

func newCountFunction(loc *errors.SourceLocation, distinct bool, args ...Evaluator) AggregationFunction {
	return &aggFunctionImpl{
		args: args,
		loc:  loc,
		implementor: func(rows []Row, function AggregationFunction) Value {
			arg := onlyOneArg("COUNT", function)
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

func onlyOneArg(name string, f AggregationFunction) Evaluator {
	if len(f.GetArgs()) > 1 {
		panic(errors.NewError(f.Location(), fmt.Sprintf("%s expects one argument", name)))
	}
	return f.GetArgs()[0]
}

func isCountAll(evaluator Evaluator) bool {
	if rowId, ok := evaluator.(Identifiable); ok {
		_, sourceColumn := qualified(rowId.Identifier())
		return sourceColumn == AllColumns
	}
	return false
}
