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
	. "github.com/onsi/gomega"
	"testing"
)

func TestRightOuterJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewCsvDataSource("../test-data/authors.csv")
	right := NewCsvDataSource("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewRightOuterJoin(left, right, joinCondition)

	g.Expect(jds.Header().ColumnCount()).To(Equal(5))
	rows := ReadAllRows(jds)

	g.Expect(rows).To(HaveLen(5))
	g.Expect(rows[0].Values()).To(Equal([]Value{
		NewIntValue(1),
		NewStringValue("J.R.R. Tolkien"),
		NewIntValue(1),
		NewIntValue(1),
		NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[4].Values()).To(Equal([]Value{
		NewNullValue(),
		NewNullValue(),
		NewIntValue(5),
		NewIntValue(6),
		NewStringValue("Unwritten Book 2"),
	}))
}

func TestRightOuterJoinDatasource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewCsvDataSource("../test-data/authors.csv")
	right := NewCsvDataSource("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoin(left, right, joinCondition)

	g.Expect(jds.CurrentRow()).To(BeNil())

	row := jds.NextRow()
	g.Expect(jds.CurrentRow()).To(Equal(row))
}

func TestRightOuterJoinDatasource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewCsvDataSource("../test-data/authors.csv")
	right := NewCsvDataSource("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoin(left, right, joinCondition)

	row := jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))

	g.Expect(func() { jds.Rewind() }).ToNot(Panic())

	row = jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))
}

func TestRightOuterJoinDatasource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewCsvDataSource("../test-data/authors.csv")
	right := NewCsvDataSource("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewRightOuterJoin(left, right, joinCondition)

	g.Expect(jds.GetName()).To(Equal("(authors RIGHT JOIN books)"))
}
