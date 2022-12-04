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

func TestLeftOuterJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.Header().ColumnCount()).To(Equal(5))
	rows := ReadAllRows(jds)

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(Equal([]Value{
		NewIntValue(1),
		NewStringValue("J.R.R. Tolkien"),
		NewIntValue(1),
		NewIntValue(1),
		NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[3].Values()).To(Equal([]Value{
		NewIntValue(3),
		NewStringValue("Lazy Bum"),
		NewNullValue(),
		NewNullValue(),
		NewNullValue(),
	}))
}

func TestLeftOuterJoinDatasource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.GetName()).To(Equal("(authors LEFT JOIN books)"))
}

func TestLeftOuterJoinDatasource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.CurrentRow()).To(BeNil())

	row := jds.NextRow()
	g.Expect(jds.CurrentRow()).To(Equal(row))
}

func TestLeftOuterJoinDatasource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	row := jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))

	g.Expect(func() { jds.Rewind() }).ToNot(Panic())

	row = jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))
}
