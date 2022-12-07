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

func TestFullJoinDataSource_Rows(t *testing.T) {
	g := NewGomegaWithT(t)
	authors := loadTestMemDatasource("authors")
	books := loadTestMemDatasource("books")

	cj := NewFullJoin(authors, books, NewEq(NewRowValue("authors.id"), NewRowValue("books.author_id")))
	g.Expect(cj).ToNot(BeNil())

	rows := ReadAllRows(cj)
	g.Expect(rows).To(HaveLen(6))

	g.Expect(rows[0].Get("authors.id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[0].Get("books.id").Value()).To(Equal(NewIntValue(1)))

	g.Expect(rows[1].Get("authors.id").Value()).To(Equal(NewIntValue(2)))
	g.Expect(rows[1].Get("books.id").Value()).To(Equal(NewIntValue(2)))

	g.Expect(rows[2].Get("authors.id").Value()).To(Equal(NewIntValue(2)))
	g.Expect(rows[2].Get("books.id").Value()).To(Equal(NewIntValue(3)))

	g.Expect(rows[3].Get("authors.id").Value()).To(Equal(NewIntValue(3)))
	g.Expect(rows[3].Get("books.id").Value()).To(Equal(NewNullValue()))

	g.Expect(rows[4].Get("authors.id").Value()).To(Equal(NewNullValue()))
	g.Expect(rows[4].Get("books.id").Value()).To(Equal(NewIntValue(4)))

	g.Expect(rows[5].Get("authors.id").Value()).To(Equal(NewNullValue()))
	g.Expect(rows[5].Get("books.id").Value()).To(Equal(NewIntValue(5)))
}
