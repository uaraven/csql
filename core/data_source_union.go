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
	"github.com/uaraven/csql/funky"
	"strings"
)

func NewUnion(ds1, ds2 DataSource) DataSource {
	unionDs := NewUnionAll(ds1, ds2)
	return NewMemDataSource(unionDs.GetName(), unionDs.Header().ColumnsMetadata(), performDistinct(unionDs))
}

func NewUnionAll(ds1, ds2 DataSource) DataSource {
	first := ds1
	ensureSameProjection(first, ds2)

	return combineRows([]DataSource{ds1, ds2})
}

// NewDistinct removes duplicate rows from the datasource
func NewDistinct(ds DataSource) DataSource {
	return NewMemDataSource(ds.GetName(), ds.Header().ColumnsMetadata(), performDistinct(ds))
}

func performDistinct(ds DataSource) []Row {
	keySet := mapset.NewThreadUnsafeSet[string]()
	result := make([]Row, 0)
	srcRow := ds.NextRow()
	counter := 0

	for srcRow != nil {
		if !keySet.Contains(srcRow.Key()) {
			result = append(result, newRowWithId(counter, ds.Header(), srcRow.Values()))
			counter++
			keySet.Add(srcRow.Key())
		}
		srcRow = ds.NextRow()
	}
	return result
}

func combineRows(datasources []DataSource) DataSource {
	counter := 0
	result := make([]Row, 0)
	for _, ds := range datasources {
		rows := ReadAllRows(ds)
		for _, row := range rows {
			result = append(result, newRowWithId(counter, ds.Header(), row.Values()))
			counter++
		}
	}
	names := funky.Map(datasources, func(ds DataSource) string {
		return ds.GetName()
	})
	unionName := strings.Join(names, " UNION ")

	return NewMemDataSource(unionName, datasources[0].Header().ColumnsMetadata(), result)
}

func ensureSameProjection(first DataSource, other DataSource) {
	// like in standard SQL only the number of columns has to match, result takes the names from the first select
	projectionFirst := getProjection(first.Header().ColumnsMetadata())
	projectionOther := getProjection(other.Header().ColumnsMetadata())
	if len(projectionFirst) != len(projectionOther) {
		panic(fmt.Errorf("different columns in UNION: %v", projectionOther))
	}
}

func getProjection(columns []ColumnMetadata) []string {
	return funky.Map(columns, func(cm ColumnMetadata) string {
		return cm.Name()
	})
}
