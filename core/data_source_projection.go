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
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/uaraven/csql/errors"
)

const AllColumns = "*"

type ProjectionColumn struct {
	source Evaluator
	loc    *errors.SourceLocation
	alias  string
}

func NewColumn(name string) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name)}
}

func NewColumnL(name string, location *errors.SourceLocation) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name), loc: location}
}

func NewColumnWithAlias(name string, alias string) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name), alias: alias}
}

func NewColumnWithAliasL(name string, alias string, location *errors.SourceLocation) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name), alias: alias, loc: location}
}

func NewExpressionColumn(src Evaluator, alias string) ProjectionColumn {
	return ProjectionColumn{source: src, alias: alias}
}

func NewExpressionColumnL(src Evaluator, alias string, location *errors.SourceLocation) ProjectionColumn {
	return ProjectionColumn{source: src, alias: alias, loc: location}
}

func NewSimpleProjection(sourceNames []string, aliases []string) []ProjectionColumn {
	if len(sourceNames) != len(aliases) {
		panic(fmt.Errorf("projection requires equal number of names and aliases.\nNames:%v\nAliases:%v", sourceNames, aliases))
	}
	projection := make([]ProjectionColumn, len(sourceNames))
	for idx, src := range sourceNames {
		projection[idx] = NewColumnWithAlias(src, aliases[idx])
	}
	return projection
}

func NewProjectionDataSource(src DataSource, projection []ProjectionColumn, distinct bool) DataSource {
	projection = ExpandProjection(projection, src.Header())
	headers := NewHeaderWithProjection(src, projection)

	var rows []Row
	if distinct {
		rows = performProjectionDistinct(src, headers, projection)
	} else {
		rows = performProjection(src, headers, projection)
	}

	return NewMemDataSource(src.GetName(), headers.ColumnsMetadata(), rows)
}

func performProjection(ds DataSource, header DataSourceHeader, projection []ProjectionColumn) []Row {
	result := make([]Row, 0)
	srcRow := ds.NextRow()

	for srcRow != nil {
		newRow := ProjectRow(projection, header, srcRow)
		result = append(result, newRow)
		srcRow = ds.NextRow()
	}
	return result
}

func performProjectionDistinct(ds DataSource, header DataSourceHeader, projection []ProjectionColumn) []Row {
	keySet := mapset.NewThreadUnsafeSet[string]()
	result := make([]Row, 0)
	srcRow := ds.NextRow()

	for srcRow != nil {
		newRow := ProjectRow(projection, header, srcRow)
		if !keySet.Contains(newRow.Key()) {
			result = append(result, newRow)
			keySet.Add(newRow.Key())
		}
		srcRow = ds.NextRow()
	}
	return result
}
