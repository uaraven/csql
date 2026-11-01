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
	"sort"
)

type OrderByField struct {
	name  string
	index int
	desc  bool
	loc   *errors.SourceLocation
}

func NewOrderBy(column string, desc bool) OrderByField {
	return OrderByField{
		name:  column,
		index: -1,
		desc:  desc,
	}
}

func NewOrderByL(column string, desc bool, l *errors.SourceLocation) OrderByField {
	return OrderByField{
		name:  column,
		index: -1,
		desc:  desc,
		loc:   l,
	}
}

func NewOrderByIndex(index int, desc bool) OrderByField {
	return OrderByField{
		name:  "",
		index: index,
		desc:  desc,
	}
}

func NewOrderByIndexL(index int, desc bool, l *errors.SourceLocation) OrderByField {
	return OrderByField{
		name:  "",
		index: index,
		desc:  desc,
		loc:   l,
	}
}

// NoLimit is the limit value that disables limiting of the rows
const NoLimit = -1

// CanOrderBy returns true if all fields can be resolved against the header
func CanOrderBy(header DataSourceHeader, orderBy []OrderByField) bool {
	for _, order := range orderBy {
		if order.index > 0 {
			if order.index > header.ColumnCount() {
				return false
			}
		} else if header.IndexByName(order.name) == InvalidFieldIndex {
			return false
		}
	}
	return true
}

// NewOrderedDatasource sorts the rows of ds (without modifying ds) and keeps at most limit of them.
// The sort is stable. NULLs sort before all other values, numbers before strings.
func NewOrderedDatasource(ds DataSource, orderBy []OrderByField, limit int) DataSource {
	if !CanOrderBy(ds.Header(), orderBy) {
		for _, order := range orderBy {
			if !CanOrderBy(ds.Header(), []OrderByField{order}) {
				panic(errors.NewError(order.loc, fmt.Sprintf("unknown column: %v", orderColumnName(order))))
			}
		}
	}
	rows := make([]Row, len(ds.GetRows()))
	copy(rows, ds.GetRows())

	if len(orderBy) > 0 {
		sort.SliceStable(rows, func(i, j int) bool {
			return rowSorter(rows[i], rows[j], orderBy) < 0
		})
	}

	if limit >= 0 && limit < len(rows) {
		rows = rows[:limit]
	}
	return NewMemDataSource(ds.GetName(), ds.Header().ColumnsMetadata(), rows)
}

func orderColumnName(order OrderByField) string {
	if order.index > 0 {
		return fmt.Sprintf("%d", order.index)
	}
	return order.name
}

// typeRank defines the order between values of different types
func typeRank(v Value) int {
	switch v.Type() {
	case TypeNull:
		return 0
	case TypeInt, TypeFloat:
		return 1
	case TypeString:
		return 2
	default:
		return 3
	}
}

// compareForOrder defines a total order of values: NULL < numbers < strings < booleans
func compareForOrder(a Value, b Value) int {
	ra, rb := typeRank(a), typeRank(b)
	if ra != rb {
		return compare(ra, rb)
	}
	switch ra {
	case 1:
		if a.Type() == TypeInt && b.Type() == TypeInt {
			return compare(a.Value().(int64), b.Value().(int64))
		}
		return compare(a.AsFloat().Value().(float64), b.AsFloat().Value().(float64))
	case 2:
		return compare(a.Value().(string), b.Value().(string))
	case 3:
		return compare(boolRank(a), boolRank(b))
	}
	return 0
}

func boolRank(v Value) int {
	if v.Value().(bool) {
		return 1
	}
	return 0
}

// rowSorter compares two rows using the order by fields, returns negative, zero or positive number
func rowSorter(r1 Row, r2 Row, orderBy []OrderByField) int {
	for _, order := range orderBy {
		var r1v, r2v Value
		if order.index > 0 {
			r1v = r1.GetByIndex(order.index)
			r2v = r2.GetByIndex(order.index)
		} else {
			r1v = r1.Get(order.name).OrPanic(errors.NewError(order.loc, fmt.Sprintf("unknown column: %v", order.name)))
			r2v = r2.Get(order.name).OrPanic(errors.NewError(order.loc, fmt.Sprintf("unknown column: %v", order.name)))
		}
		if c := compareForOrder(r1v, r2v); c != 0 {
			if order.desc {
				return -c
			}
			return c
		}
	}
	return 0
}
