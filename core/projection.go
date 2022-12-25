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
	"github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/funky"
)

func isAsterisk(pc ProjectionColumn) bool {
	if rowId, ok := pc.source.(Identifiable); ok {
		_, sourceColumn := qualified(rowId.Identifier())
		return sourceColumn == AllColumns
	}
	return false
}

func ExpandProjection(projection []ProjectionColumn, headers DataSourceHeader) []ProjectionColumn {
	if funky.NoneMatches(projection, isAsterisk) {
		return projection
	}
	result := make([]ProjectionColumn, 0)
	for _, column := range projection {
		if isAsterisk(column) {
			fullAsterisk := column.source.(Identifiable).Identifier()
			for _, cmd := range headers.ColumnsMetadata() {
				if cmd.MatchesName(fullAsterisk) {
					var newName string
					if cmd.ParentName() != "" {
						newName = cmd.ParentName() + "." + cmd.Name()
					} else {
						newName = cmd.Name()
					}
					result = append(result, NewColumn(newName))
				}
			}
		} else {
			result = append(result, column)
		}
	}
	return result
}

func NewHeaderWithProjection(src DataSource, projection []ProjectionColumn) DataSourceHeader {
	columns := make([]ColumnMetadata, len(projection))
	for idx, prjColumn := range projection {
		var columnMeta columnMetadata
		if rowId, ok := prjColumn.source.(Identifiable); ok {
			// projection column references datasource column
			sourceColumn := rowId.Identifier()
			columnIdx := src.Header().IndexByName(sourceColumn)
			if columnIdx == InvalidFieldIndex {
				panic(errors.UnknownColumnError(prjColumn.loc, sourceColumn))
			}
			sourceColumnMetadata := src.Header().ColumnsMetadata()[columnIdx]
			columnMeta = columnMetadata{
				parentName: sourceColumnMetadata.ParentName(),
				index:      idx,
				name:       sourceColumnMetadata.Name(),
			}
		} else {
			// projection column is an expression
			columnMeta = columnMetadata{
				parentName: "",
				index:      idx,
				name:       prjColumn.source.String(),
			}
		}
		if prjColumn.alias != "" {
			columnMeta.name = prjColumn.alias
		}
		columns[idx] = &columnMeta
	}

	return &dataSourceHeader{
		columns: columns,
	}
}

func ProjectRow(projection []ProjectionColumn, header DataSourceHeader, row Row) Row {
	resultValues := make([]Value, header.ColumnCount())
	for idx, prj := range projection {
		v := prj.source.Evaluate(row)
		resultValues[idx] = v
	}
	return newRowWithId(row.Id(), header, resultValues)
}
