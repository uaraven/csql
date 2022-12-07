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
	"github.com/uaraven/csql/collection"
	"github.com/uaraven/csql/funky"
)

func NewFullJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s FULL JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = fullJoin(left, right, mds.headers, joinOn)
	return mds

}

func fullJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	leftJoined, usedRightRows := leftJoin(left, right, header, joinOn)
	unusedRightRows := funky.Filter(right.(SliceDataSource).GetRows(), func(row Row) bool {
		return !usedRightRows.Contains(row.Key())
	})
	counter := len(leftJoined) - 1
	rightJoined := funky.Map(unusedRightRows, func(row Row) Row {
		counter++
		return joinRows(counter, header, nullRow(left.Header()), row)
	})
	return append(leftJoined, rightJoined...)
}

func selectJoinedLeft(counter int, header DataSourceHeader, leftRow Row, right DataSource, joinOn Condition, usedRight mapset.Set[string]) (collection.LinkedList[Row], int) {
	var rightSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	right.Rewind()
	for {
		rightRow := right.NextRow()
		if rightRow == nil {
			return rightSelected, counter
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			usedRight.Add(rightRow.Key())
			rightSelected.Append(newRowWithId(counter, header, joined.Values()))
			counter++

		}
	}
}

func leftJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) ([]Row, mapset.Set[string]) {
	result := make([]Row, 0)
	usedRightRows := mapset.NewSet[string]()
	counter := 0
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return result, usedRightRows
		}
		joined, counter := selectJoinedLeft(counter, header, leftRow, right, joinOn, usedRightRows)
		if joined.Empty() {
			joinedRow := joinRows(counter, header, leftRow, nullRow(right.Header()))
			counter++
			result = append(result, joinedRow)
		} else {
			iter := joined.Iterator()
			for iter.HasNext() {
				result = append(result, iter.Next())
				counter++
			}
		}
	}

}
