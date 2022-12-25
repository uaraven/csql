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
	"strconv"
	"strings"
)

type rowImpl struct {
	id     int
	header DataSourceHeader
	values []Value
}

func parseRowWithId(id int, headers DataSourceHeader, values []string) Row {
	return &rowImpl{
		id:     id,
		header: headers,
		values: funky.Map(values, func(s string) Value {
			return decodeToValue(s)
		}),
	}
}

func updateRowId(id int, row Row) Row {
	return &rowImpl{
		id:     id,
		header: row.Header(),
		values: row.Values(),
	}
}

func newRowWithId(id int, headers DataSourceHeader, values []Value) Row {
	return &rowImpl{
		id:     id,
		header: headers,
		values: values,
	}
}

func joinRows(id int, header DataSourceHeader, row1 Row, row2 Row) Row {
	rowData := make([]Value, len(row1.Values()))
	copy(rowData, row1.Values())
	rowData = append(rowData, row2.Values()...)

	return &rowImpl{
		id:     id,
		header: header,
		values: rowData,
	}
}

func nullRow(header DataSourceHeader) Row {
	rowData := make([]Value, header.ColumnCount())
	for i := range rowData {
		rowData[i] = NewNullValue()
	}
	return newRowWithId(0, header, rowData)
}

func (r rowImpl) Header() DataSourceHeader {
	return r.header
}

func (r rowImpl) Values() []Value {
	return r.values
}

func (r rowImpl) Count() int {
	return len(r.values)
}

func (r rowImpl) GetByIndex(index int) Value {
	if index < 1 || index > len(r.values) {
		panic(errors.NewError(nil, fmt.Sprintf("Invalid column index: %d", index)))
	}
	return r.values[index-1]
}

func (r rowImpl) Get(column string) funky.Option[Value] {
	index := r.Header().IndexByName(column)
	if index == RowIdIndex {
		return funky.SomeOf(NewIntValue(int64(r.id)))
	} else if index < RowIdIndex {
		return funky.NoneOf[Value]()
	} else {
		return funky.SomeOf(r.values[index])
	}
}

func (r rowImpl) Satisfies(cond Condition) bool {
	return cond.Evaluate(r).AsBool().Value().(bool)
}

func (r rowImpl) Id() int {
	return r.id
}

func (r rowImpl) String() string {
	return fmt.Sprintf("%v", r.values)
}

func (r rowImpl) Key() string {
	var sb strings.Builder
	for idx, v := range r.values {
		if idx != 0 {
			sb.WriteRune('-')
		}
		sb.WriteString(v.Evaluate(r).String())
	}
	return sb.String()
}

func decodeToValue(value string) Value {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			if value == NullValueString {
				return NewNullValue()
			}
			return NewStringValue(value)
		}
		return NewFloatValue(f64)
	}
	return NewIntValue(i64)
}

func ReadAllRows(ds DataSource) []Row {
	return ds.GetRows()
}

func RowsToSlice(rows []Row, columns ...string) [][]interface{} {
	results := make([][]interface{}, len(rows))
	for idx, row := range rows {
		values := make([]interface{}, len(columns))
		for cidx, column := range columns {
			value := row.Get(column).OrPanic(fmt.Errorf("no column '%s' in the row", column)).Value()
			values[cidx] = value
		}
		results[idx] = values
	}
	return results
}
