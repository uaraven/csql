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
	"reflect"
	"strings"
)

func NewUnion(dataSources []DataSource) DataSource {
	unionDs := NewUnionAll(dataSources)
	return NewMemDataSource(unionDs.GetName(), unionDs.Header().ColumnsMetadata(), performDistinct(unionDs))
}

func NewUnionAll(dataSources []DataSource) DataSource {
	first := dataSources[0]
	ensureSameProjection(first, dataSources[1:])

	return combineRows(dataSources)
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
		}
	}
	names := funky.Map(datasources, func(ds DataSource) string {
		return ds.GetName()
	})
	unionName := strings.Join(names, " UNION ")

	return NewMemDataSource(unionName, datasources[0].Header().ColumnsMetadata(), result)
}

func ensureSameProjection(first DataSource, others []DataSource) {
	projectionFirst := getProjection(first.Header().ColumnsMetadata())
	for _, prj := range others {
		projectionOther := getProjection(prj.Header().ColumnsMetadata())
		if !reflect.DeepEqual(projectionFirst, projectionOther) {
			panic(fmt.Errorf("different columns in UNION: %v", projectionOther))
		}
	}
}

func getProjection(columns []ColumnMetadata) []string {
	return funky.Map(columns, func(cm ColumnMetadata) string {
		return cm.Name()
	})
}
