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

type GroupColumn struct {
	name  string
	index int
	loc   *errors.SourceLocation
}

func NewAggregationDs(source DataSource, projection []ProjectionColumn, groupBy []GroupColumn) DataSource {
	rowGroups := groupRows(source, groupBy)

	aggregateFunctions := filterAggregateFunctions(projection)

	// if there are aggregate functions in the projection then build a new header with
	// additional columns for the aggregate function results and add those results to
	// each row generated from grouped sets.
	// original projection should be modified to include column references in place of aggregate functions
	// and then applied to the resulting rows (using ProjectionDatasource)

	var result []Row
	if len(aggregateFunctions) == 0 {
		result = applyGroupingNoAggregates(rowGroups)
		intermediate := NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), result)
		return NewProjectionDataSource(intermediate, projection, false)
	} else {
		result = applyAggregates(rowGroups)
		return NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), result)
	}
}

func applyAggregates(groups [][]Row) []Row {
	return nil
}

func applyGroupingNoAggregates(rowGroups [][]Row) []Row {
	result := make([]Row, 0)
	for _, rows := range rowGroups {
		if len(rows) == 0 {
			continue
		}
		result = append(result, rows[0])
	}
	return result
}

func filterAggregateFunctions(projection []ProjectionColumn) []AggregationFunction {
	result := make([]AggregationFunction, 0)
	for _, column := range projection {
		if aggF, ok := column.source.(AggregationFunction); ok {
			result = append(result, aggF)
		}
	}
	return result
}

func groupRows(ds DataSource, by []GroupColumn) [][]Row {
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

func getByGroupColumn(row Row, column GroupColumn) Value {
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
