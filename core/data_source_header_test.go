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

func TestRowImpl_GetRowId(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row := ds.NextRow()

	val := row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(0)))

	row = ds.NextRow()

	val = row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(1)))

	g.Expect(func() { ds.Rewind() }).ToNot(Panic())
	row = ds.NextRow()

	val = row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(0)))
}

func TestRowImpl_Satisfies(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row := ds.NextRow()

	condT := NewAnd(
		NewEq(NewRowValue("first_name"), NewStringValue("John")),
		NewLt(NewRowValue("dept_id"), NewIntValue(3)))

	condF := NewEq(NewRowValue("first_name"), NewStringValue("Connor"))

	g.Expect(row.Satisfies(condT)).To(BeTrue())
	g.Expect(row.Satisfies(condF)).To(BeFalse())
}
