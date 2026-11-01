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

func NewLeftOuterJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s LEFT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data, _ = outerJoin(left.GetRows(), right.GetRows(), true, nullRow(right.Header()), mds.headers, joinOn)
	return mds
}

func NewRightOuterJoin(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s RIGHT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data, _ = outerJoin(right.GetRows(), left.GetRows(), false, nullRow(left.Header()), mds.headers, joinOn)
	return mds
}

// outerJoin joins every row of outerRows with the matching rows of innerRows. An outer row without a match is joined
// with nullInner. outerIsLeft tells which side of the result the outer rows go to. Row ids are sequential in the
// result. The returned slice tells which inner rows had at least one match.
func outerJoin(outerRows []Row, innerRows []Row, outerIsLeft bool, nullInner Row, header DataSourceHeader, joinOn Condition) ([]Row, []bool) {
	result := make([]Row, 0, len(outerRows))
	usedInner := make([]bool, len(innerRows))
	pair := func(outer Row, inner Row) Row {
		if outerIsLeft {
			return joinRows(len(result), header, outer, inner)
		}
		return joinRows(len(result), header, inner, outer)
	}
	for _, outerRow := range outerRows {
		matched := false
		for idx, innerRow := range innerRows {
			joined := pair(outerRow, innerRow)
			if joined.Satisfies(joinOn) {
				result = append(result, joined)
				usedInner[idx] = true
				matched = true
			}
		}
		if !matched {
			result = append(result, pair(outerRow, nullInner))
		}
	}
	return result, usedInner
}
