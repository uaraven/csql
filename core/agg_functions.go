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
	"github.com/uaraven/csql/funky"
	"github.com/uaraven/csql/util"
	"math"
	"strings"
)

type aggregateContextImpl struct {
	rows []Row
}

func NewAggregateContext(rows []Row) AggregateContext {
	return &aggregateContextImpl{rows: rows}
}

func (aci aggregateContextImpl) Get(s string) funky.Option[Value] {
	if len(aci.rows) > 0 {
		return aci.rows[0].Get(s)
	} else {
		return funky.SomeOf[Value](nullV)
	}
}

func (aci aggregateContextImpl) RowSet() []Row {
	return aci.rows
}

type AggregateFunction interface {
	Evaluator
	Name() string
	GetArg() Evaluator
	Location() *errors.SourceLocation
	AggregateEvaluate(ctx AggregateContext) Value
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
	if aggCtx, isAgg := ctx.(AggregateContext); isAgg {
		return sf.AggregateEvaluate(aggCtx)
	}
	if rowCtx, isRow := ctx.(*rowImpl); isRow {
		aggCtx := aggregateContextImpl{
			rows: []Row{rowCtx},
		}
		return sf.AggregateEvaluate(aggCtx)
	}
	panic(errors.NewError(sf.loc, "Aggregate function cannot be used as a scalar"))
}

func (sf aggFunctionImpl) AggregateEvaluate(ctx AggregateContext) Value {
	return sf.implementor(ctx.RowSet(), sf)
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

type countFunction struct {
	AggregateFunction
	arg         Evaluator
	distinct    bool
	loc         *errors.SourceLocation
	implementor AggregateFunctionImplementor
}

func NewSumFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return aggFunctionImpl{
		arg:  arg,
		name: "SUM",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			sum := NewIntValue(0)
			for _, row := range rows {
				value := arg.Evaluate(row)
				if !IsNumeric(value) {
					panic(errors.NewError(loc, fmt.Sprintf("non-numeric value for sum: %v", value)))
				}
				if sum.Type() == TypeFloat || value.Type() == TypeFloat {
					a := sum.AsFloat().Value().(float64)
					b := value.AsFloat().Value().(float64)
					sum = util.Must(wrapInFloat(operationExecutor(addition, a, b)))
				} else {
					a := sum.AsInt().Value().(int64)
					b := value.AsInt().Value().(int64)
					sum = util.Must(wrapInInt(operationExecutor(addition, a, b)))
				}
			}
			return sum
		},
	}
}

func NewAvgFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return aggFunctionImpl{
		arg:  arg,
		name: "AVG",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			sum := NewIntValue(0)
			for _, row := range rows {
				value := arg.Evaluate(row)
				if !IsNumeric(value) {
					panic(errors.NewError(loc, fmt.Sprintf("non-numeric value for sum: %v", value)))
				}
				if sum.Type() == TypeFloat || value.Type() == TypeFloat {
					a := sum.AsFloat().Value().(float64)
					b := value.AsFloat().Value().(float64)
					sum = util.Must(wrapInFloat(operationExecutor(addition, a, b)))
				} else {
					a := sum.AsInt().Value().(int64)
					b := value.AsInt().Value().(int64)
					sum = util.Must(wrapInInt(operationExecutor(addition, a, b)))
				}
			}
			switch sum.Type() {
			case TypeInt:
				return NewFloatValue(float64(sum.Value().(int64)) / float64(len(rows)))
			case TypeFloat:
				return NewFloatValue(sum.Value().(float64) / float64(len(rows)))
			default:
				panic(fmt.Errorf("unexpected type in AVG"))
			}
		},
	}
}

func NewMinFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return aggFunctionImpl{
		arg:  arg,
		name: "MIN",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			min := arg.Evaluate(rows[0])
			for _, row := range rows {
				value := arg.Evaluate(row)
				if !IsNumeric(value) {
					panic(errors.NewError(loc, fmt.Sprintf("non-numeric value for min: %v", value)))
				}
				if min.Type() == TypeFloat || value.Type() == TypeFloat {
					a := min.AsFloat().Value().(float64)
					b := value.AsFloat().Value().(float64)
					min = NewFloatValue(math.Min(a, b))
				} else {
					a := min.AsInt().Value().(int64)
					b := value.AsInt().Value().(int64)
					if a < b {
						min = NewIntValue(a)
					} else {
						min = NewIntValue(b)
					}
				}
			}
			return min
		},
	}
}

func NewMaxFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return aggFunctionImpl{
		arg:  arg,
		name: "MAX",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			max := arg.Evaluate(rows[0])
			for _, row := range rows[1:] {
				value := arg.Evaluate(row)
				if !IsNumeric(value) {
					panic(errors.NewError(loc, fmt.Sprintf("non-numeric value for max: %v", value)))
				}
				if max.Type() == TypeFloat || value.Type() == TypeFloat {
					a := max.AsFloat().Value().(float64)
					b := value.AsFloat().Value().(float64)
					max = NewFloatValue(math.Max(a, b))
				} else {
					a := max.AsInt().Value().(int64)
					b := value.AsInt().Value().(int64)
					if a > b {
						max = NewIntValue(a)
					} else {
						max = NewIntValue(b)
					}
				}
			}
			return max
		},
	}
}

func NewCountFunction(distinct bool, arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &countFunction{
		arg:      arg,
		loc:      loc,
		distinct: distinct,
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

func (cf countFunction) Name() string {
	return "COUNT"
}

func (cf countFunction) GetArg() Evaluator {
	return cf.arg
}

func (cf countFunction) Location() *errors.SourceLocation {
	return cf.loc
}

func (cf countFunction) AggregateEvaluate(ctx AggregateContext) Value {
	return cf.implementor(ctx.RowSet(), cf)
}

func (cf countFunction) String() string {
	var sb strings.Builder
	sb.WriteString("COUNT(")
	if cf.distinct {
		sb.WriteString("DISTINCT ")
	}
	sb.WriteString(cf.arg.String())
	sb.WriteRune(')')
	return sb.String()
}

func isCountAll(evaluator Evaluator) bool {
	if rowId, ok := evaluator.(Identifiable); ok {
		_, sourceColumn := qualified(rowId.Identifier())
		return sourceColumn == AllColumns
	}
	return false
}
