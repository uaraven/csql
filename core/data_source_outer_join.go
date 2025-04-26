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
	"github.com/uaraven/csql/collection"
)

func NewLeftOuterJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s LEFT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = leftOuterJoin(left, right, mds.headers, joinOn)
	return mds
}

func NewRightOuterJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s RIGHT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = rightOuterJoin(left, right, mds.headers, joinOn)
	return mds
}

func leftOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	result := make([]Row, 0)
	counter := 0
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return result
		}
		joined, counter := selectRight(counter, header, leftRow, right, joinOn)
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

func selectRight(counter int, header DataSourceHeader, leftRow Row, right DataSource, joinOn Condition) (collection.LinkedList[Row], int) {
	var rightSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	right.Rewind()
	for {
		rightRow := right.NextRow()
		if rightRow == nil {
			return rightSelected, counter
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			rightSelected.Append(newRowWithId(counter, joined.Header(), joined.Values()))
			counter++
		}
	}
}

func rightOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	result := make([]Row, 0)
	counter := 0
	for {
		rightRow := right.NextRow()
		if rightRow == nil {
			return result
		}
		joined, counter := selectLeft(counter, header, rightRow, left, joinOn)
		if joined.Empty() {
			joinedRow := joinRows(counter, header, nullRow(left.Header()), rightRow)
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

func selectLeft(counter int, header DataSourceHeader, rightRow Row, left DataSource, joinOn Condition) (collection.LinkedList[Row], int) {
	var leftSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	left.Rewind()
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return leftSelected, counter
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			leftSelected.Append(newRowWithId(counter, header, joined.Values()))
			counter++
		}
	}
}
