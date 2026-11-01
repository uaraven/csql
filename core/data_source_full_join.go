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
)

func NewFullJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s FULL JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = fullJoin(left, right, mds.headers, joinOn)
	return mds

}

func fullJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	rightRows := right.GetRows()
	result, usedRight := outerJoin(left.GetRows(), rightRows, true, nullRow(right.Header()), header, joinOn)
	nullLeft := nullRow(left.Header())
	for idx, row := range rightRows {
		if !usedRight[idx] {
			result = append(result, joinRows(len(result), header, nullLeft, row))
		}
	}
	return result
}
