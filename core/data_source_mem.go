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

type memDataSource struct {
	name    string
	index   int
	data    []Row
	headers DataSourceHeader
}

func NewMemDataSource(name string, headers []ColumnMetadata, rows []Row) DataSource {
	ds := &memDataSource{
		name:  name,
		index: -1,
		data:  rows,
	}

	ds.headers = dataSourceHeader{
		columns: headers,
	}

	return ds
}

func (cds *memDataSource) Header() DataSourceHeader {
	return cds.headers
}

func (cds *memDataSource) GetName() string {
	return cds.name
}

func (cds *memDataSource) NextRow() Row {
	if cds.index+1 >= len(cds.data) {
		return nil
	}
	cds.index++
	return cds.data[cds.index]
}

func (cds *memDataSource) CurrentRow() Row {
	if cds.index < 0 {
		return nil
	}
	return cds.data[cds.index]
}

func (cds *memDataSource) Rewind() {
	cds.index = -1
}

func (cds *memDataSource) GetRows() []Row {
	return cds.data
}
