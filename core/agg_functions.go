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
	"strings"

	"github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/funky"
	"github.com/uaraven/csql/util"
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
	aggFunctionImpl
	arg         Evaluator
	distinct    bool
	loc         *errors.SourceLocation
	implementor AggregateFunctionImplementor
}

// isMissing reports whether an aggregate function has to skip the value: NULL or an empty CSV cell
func isMissing(v Value) bool {
	return v.Type() == TypeNull || (v.Type() == TypeString && v.Value().(string) == "")
}

// sumNumeric adds up all numeric values of arg, skipping missing ones, and returns the sum with the number of values added
func sumNumeric(name string, arg Evaluator, rows []Row, loc *errors.SourceLocation) (Value, int) {
	sum := NewIntValue(0)
	count := 0
	for _, row := range rows {
		value := arg.Evaluate(row)
		if isMissing(value) {
			continue
		}
		if !IsNumeric(value) {
			panic(errors.NewError(loc, fmt.Sprintf("non-numeric value for %s: %v in row %s", name, value, row)))
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
		count++
	}
	return sum, count
}

// extremum finds the smallest or the largest value of arg, skipping missing ones. Numbers and strings can be
// compared, but not with each other. The result is NULL if there are no values.
func extremum(name string, arg Evaluator, rows []Row, loc *errors.SourceLocation, wantMax bool) Value {
	var best Value
	for _, row := range rows {
		value := arg.Evaluate(row)
		if isMissing(value) {
			continue
		}
		if best == nil {
			best = value
			continue
		}
		var cmp int
		switch {
		case IsNumeric(best) && IsNumeric(value):
			if best.Type() == TypeInt && value.Type() == TypeInt {
				cmp = compare(value.Value().(int64), best.Value().(int64))
			} else {
				cmp = compare(value.AsFloat().Value().(float64), best.AsFloat().Value().(float64))
			}
		case best.Type() == TypeString && value.Type() == TypeString:
			cmp = compare(value.Value().(string), best.Value().(string))
		default:
			panic(errors.NewError(loc, fmt.Sprintf("cannot compare %v and %v in %s", best, value, name)))
		}
		if (wantMax && cmp > 0) || (!wantMax && cmp < 0) {
			best = value
		}
	}
	if best == nil {
		return NewNullValue()
	}
	return best
}

func NewSumFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &aggFunctionImpl{
		arg:  arg,
		name: "SUM",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			sum, count := sumNumeric("sum", arg, rows, loc)
			if count == 0 {
				return NewNullValue()
			}
			return sum
		},
	}
}

func NewAvgFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &aggFunctionImpl{
		arg:  arg,
		name: "AVG",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			sum, count := sumNumeric("avg", arg, rows, loc)
			if count == 0 {
				return NewNullValue()
			}
			return NewFloatValue(sum.AsFloat().Value().(float64) / float64(count))
		},
	}
}

func NewMinFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &aggFunctionImpl{
		arg:  arg,
		name: "MIN",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			return extremum("MIN", arg, rows, loc, false)
		},
	}
}

func NewMaxFunction(arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &aggFunctionImpl{
		arg:  arg,
		name: "MAX",
		loc:  loc,
		implementor: func(rows []Row, function AggregateFunction) Value {
			return extremum("MAX", arg, rows, loc, true)
		},
	}
}

// NewCountFunction creates COUNT. COUNT(*) counts all rows, COUNT(column) skips NULLs and empty cells
func NewCountFunction(distinct bool, arg Evaluator, loc *errors.SourceLocation) AggregateFunction {
	return &countFunction{
		arg:      arg,
		loc:      loc,
		distinct: distinct,
		implementor: func(rows []Row, function AggregateFunction) Value {
			countAll := isCountAll(arg)
			if countAll && !distinct {
				return NewIntValue(int64(len(rows)))
			}
			seen := make(map[string]bool)
			counter := 0
			for _, row := range rows {
				var key string
				if countAll {
					key = row.Key()
				} else {
					value := arg.Evaluate(row)
					if isMissing(value) {
						continue
					}
					key = valueKey(value)
				}
				if distinct {
					if seen[key] {
						continue
					}
					seen[key] = true
				}
				counter++
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

func (cf countFunction) Evaluate(ctx EvaluationContext) Value {
	if aggCtx, isAgg := ctx.(AggregateContext); isAgg {
		return cf.AggregateEvaluate(aggCtx)
	}
	panic(errors.NewError(cf.loc, "COUNT function cannot be used as a scalar"))
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
