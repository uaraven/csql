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

func TestLoadMemCsv(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := NewCsvDataSource("../test-data/employees.csv")

	hdr := ds.Header()

	g.Expect(hdr.ColumnCount()).To(Equal(5))

	rows := ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
}

func TestLoadMemCsvRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewIntValue(1),
		NewStringValue("John"),
		NewStringValue("Snow"),
		NewFloatValue(1.0)}))
	g.Expect(row.Id()).To(Equal(0))

	row = ds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(2), NewIntValue(1),
		NewStringValue("James"),
		NewStringValue("Snow"),
		NewFloatValue(1.2)}))
	g.Expect(row.Id()).To(Equal(1))
}

func TestLoadMemCsvRewind(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row1 := ds.NextRow()

	g.Expect(func() { ds.Rewind() }).ToNot(Panic())

	g.Expect(ds.NextRow()).To(Equal(row1))
}

func TestMemCsvGetValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	val := row.Get("first_name")

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value().Value()).To(Equal("John"))
}

func TestMemCsvGetValueByInvalidName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	val := row.Get("middle_name")

	g.Expect(val.IsPresent()).To(BeFalse())
}
