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

import "github.com/uaraven/csql/funky"

func NewFilteredDataSource(source DataSource, condition Condition) DataSource {
	rows := ReadAllRows(source)
	filteredRows := funky.ParallelFilter(rows, func(r Row) bool {
		return r.Satisfies(condition)
	})
	filteredRows = funky.MapWithIndex(filteredRows, func(id int, row Row) Row {
		return updateRowId(id, row)
	})

	return NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), filteredRows)
}
