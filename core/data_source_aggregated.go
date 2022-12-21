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

func NewAggregationDs(source DataSource, groupBy []GroupColumn) DataSource {
	rowGroups := groupRows(source, groupBy)

	result := make([]Row, 0)
	for _, rows := range rowGroups {
		if len(rows) > 0 {
			result = append(result, rows[0])
		}
	}
	return NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), result)
}

func groupRows(ds DataSource, by []GroupColumn) map[string][]Row {
	result := make(map[string][]Row)
	for _, row := range ds.GetRows() {
		vals := make([]Value, len(by))
		for idx, groupColumn := range by {
			vals[idx] = getByGroupColumn(row, groupColumn)
		}
		key := strings.Join(funky.Map(vals, func(v Value) string {
			return v.String()
		}), "-")
		result[key] = append(result[key], row)
	}
	return result
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
