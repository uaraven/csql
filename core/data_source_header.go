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

	"github.com/uaraven/csql/util"
)

type columnMetadata struct {
	parentName string
	name       string
	index      int
}

func NewSimpleColumnMetadata(name string, index int) ColumnMetadata {
	return &columnMetadata{
		name:  name,
		index: index,
	}
}

func (cm columnMetadata) ParentName() string {
	return cm.parentName
}

func (cm columnMetadata) Index() int {
	return cm.index
}

func (cm columnMetadata) MatchesName(name string) bool {
	table, column := qualified(name)
	if table != "" && !util.EqualsIgnoreCase(table, cm.parentName) {
		return false

	}
	return column == "*" || util.EqualsIgnoreCase(column, cm.name)
}

func (cm columnMetadata) Name() string {
	return cm.name
}

type dataSourceHeader struct {
	columns []ColumnMetadata
}

func (d dataSourceHeader) ColumnCount() int {
	return len(d.columns)
}

func (d dataSourceHeader) ColumnsMetadata() []ColumnMetadata {
	return d.columns
}

func (d dataSourceHeader) IndexByName(s string) int {
	_, c := qualified(s)
	if c == RowId {
		return RowIdIndex
	}
	index := InvalidFieldIndex
	for idx, column := range d.columns {
		if column.MatchesName(s) {
			if index != InvalidFieldIndex {
				panic(fmt.Errorf("ambiguous column name: %s", s))
			}
			index = idx
		}
	}
	return index
}

func NewHeadersFromSlice(parent DataSource, headers []string) DataSourceHeader {
	columns := make([]ColumnMetadata, 0)
	for idx, name := range headers {
		columns = append(columns, &columnMetadata{
			parentName: parent.GetName(),
			name:       name,
			index:      idx,
		})
	}
	return &dataSourceHeader{ /*parent: parent,*/ columns: columns}
}

func NewHeadersFromOtherHeaders(headers ...[]ColumnMetadata) DataSourceHeader {
	columns := make([]ColumnMetadata, 0)
	for _, otherHeaders := range headers {
		columns = append(columns, otherHeaders...)
	}
	return &dataSourceHeader{
		columns: columns,
	}
}

func qualified(name string) (string, string) {
	p := strings.Split(name, ".")
	if len(p) == 1 {
		return "", name
	} else {
		return p[0], p[1]
	}
}
