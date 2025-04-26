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

func NewCrossJoin(left DataSource, right DataSource) DataSource {
	return NewInnerJoin(left, right, nil)
}

func NewInnerJoin(left DataSource, right DataSource, joinCondition Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = innerJoin(left, right, mds.headers, joinCondition)
	return mds
}

func innerJoin(left DataSource, right DataSource, header DataSourceHeader, joinCondition Condition) []Row {
	rows := make([]Row, 0)
	counter := 0
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return rows
		}
		for {
			rightRow := right.NextRow()
			if rightRow == nil {
				right.Rewind()
				break
			}
			newRow := joinRows(counter, header, leftRow, rightRow)
			if joinCondition == nil || newRow.Satisfies(joinCondition) {
				rows = append(rows, newRow)
				counter++
			}
		}
	}
}
