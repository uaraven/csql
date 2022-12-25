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
	"strings"
)

type GroupByColumn struct {
	name  string
	index int
	loc   *errors.SourceLocation
}

func NewGroupByIndex(index int, loc *errors.SourceLocation) GroupByColumn {
	return GroupByColumn{
		name:  "",
		index: index,
		loc:   loc,
	}
}

func NewGroupByName(name string, loc *errors.SourceLocation) GroupByColumn {
	return GroupByColumn{
		name:  name,
		index: 0,
		loc:   loc,
	}
}

func NewAggregationDataSource(source DataSource,
	projection []ProjectionColumn,
	groupBy []GroupByColumn) DataSource {
	return NewAggregationDataSourceHaving(source, projection, groupBy, nil)
}

func NewAggregationDataSourceHaving(
	source DataSource,
	projection []ProjectionColumn,
	groupBy []GroupByColumn,
	having Evaluator) DataSource {
	rowGroups := groupRows(source, groupBy)

	aggregateFunctions := filterAggregateFunctions(projection)

	if len(aggregateFunctions) == 0 && len(groupBy) == 0 && having != nil {
		panic(errors.NewError(nil, "HAVING clause on non-aggregate query"))
	}

	//aggregateHaving, isAgg := having.(AggregateEvaluator)
	//if !isAgg {
	//	panic(fmt.Errorf("expecting aggregate evaluator"))
	//}

	// if there are aggregate functions in the projection then build a new header with
	// additional columns for the aggregate function results and add those results to
	// each row generated from grouped sets.
	// original projection should be modified to include column references in place of aggregate functions
	// and then applied to the resulting rows (using ProjectionDatasource)

	var result []Row
	if len(aggregateFunctions) == 0 {
		result = foldGroupNoAggregates(rowGroups, having)
		intermediate := NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), result)
		return NewProjectionDataSource(intermediate, projection, false)
	} else {
		projection = ExpandProjection(projection, source.Header())
		newHeader := NewHeaderWithProjection(source, projection)
		result = applyAggregates(rowGroups, projection, newHeader, aggregateFunctions, having)
		return NewMemDataSource(source.GetName(), newHeader.ColumnsMetadata(), result)
	}
}

func applyAggregates(groups [][]Row, projection []ProjectionColumn, headers DataSourceHeader, aggregates map[int]AggregateFunction, having Evaluator) []Row {
	result := make([]Row, 0)
	rowId := 0
	for _, row := range groups {
		if len(row) > 0 {
			aggregateCtx := NewAggregateContext(row)
			if having == nil || having.Evaluate(aggregateCtx).Equals(ValueTrue) {
				projectedRow := projectRow(projection, headers, row[0])
				values := projectedRow.Values()
				for index, v := range aggregates {
					value := v.AggregateEvaluate(aggregateCtx)
					values[index] = value
				}
				nrow := newRowWithId(rowId, headers, values)
				result = append(result, nrow)
				rowId++
			}
		}
	}
	return result
}

func projectRow(projection []ProjectionColumn, header DataSourceHeader, row Row) Row {
	resultValues := make([]Value, header.ColumnCount())
	for idx, prj := range projection {
		var v Value
		if _, isAgg := prj.source.(AggregateFunction); isAgg {
			v = nullV
		} else {
			v = prj.source.Evaluate(row)
		}
		resultValues[idx] = v
	}
	return newRowWithId(row.Id(), header, resultValues)
}

func foldGroupNoAggregates(rowGroups [][]Row, having Evaluator) []Row {
	result := make([]Row, 0)
	rowId := 0
	for _, rows := range rowGroups {
		if len(rows) == 0 {
			continue
		}
		if having == nil || having.Evaluate(NewAggregateContext(rows)).Equals(ValueTrue) {
			result = append(result, updateRowId(rowId, rows[0]))
			rowId++
		}
	}
	return result
}

func filterAggregateFunctions(projection []ProjectionColumn) map[int]AggregateFunction {
	result := make(map[int]AggregateFunction)
	for idx, column := range projection {
		if aggF, ok := column.source.(AggregateFunction); ok {
			result[idx] = aggF
		}
	}
	return result
}

func groupRows(ds DataSource, by []GroupByColumn) [][]Row {
	groupings := make(map[string][]Row)
	for _, row := range ds.GetRows() {
		vals := make([]Value, len(by))
		for idx, groupColumn := range by {
			vals[idx] = getByGroupColumn(row, groupColumn)
		}
		key := strings.Join(funky.Map(vals, func(v Value) string {
			return v.String()
		}), "-")
		groupings[key] = append(groupings[key], row)
	}
	results := make([][]Row, len(groupings))
	index := 0
	for _, vals := range groupings {
		results[index] = vals
		index++
	}
	return results
}

func getByGroupColumn(row Row, column GroupByColumn) Value {
	if column.name != "" {
		optValue := row.Get(column.name)
		if optValue.IsEmpty() {
			panic(errors.NewError(column.loc, fmt.Sprintf("Unknown column: \"%s\"", column.name)))
		}
		return optValue.Value()
	} else {
		return row.GetByIndex(column.index)
	}
}
