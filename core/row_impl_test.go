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

func TestRowValue_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	firstName := NewRowValue("first_name")
	g.Expect(firstName.Evaluate(row).Type()).To(Equal(TypeString))
	g.Expect(firstName.Evaluate(row).Value()).To(Equal("John"))

	lastName := NewRowValue("last_name")
	g.Expect(lastName.Evaluate(row).Type()).To(Equal(TypeString))
	g.Expect(lastName.Evaluate(row).Value()).To(Equal("Snow"))

	deptId := NewRowValue("dept_id")
	g.Expect(deptId.Evaluate(row).Type()).To(Equal(TypeInt))
	g.Expect(deptId.Evaluate(row).Value()).To(Equal(int64(1)))

	width := NewRowValue("width")
	g.Expect(width.Evaluate(row).Type()).To(Equal(TypeFloat))
	g.Expect(width.Evaluate(row).Value()).To(Equal(1.0))
}

func TestRowImpl_Get_QualifiedAccess(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	v := row.Get("first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("employees.first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("employee.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

	v = row.Get("emp.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())
}

func TestRowImpl_Get_QualifiedAccessAlias(t *testing.T) {
	g := NewGomegaWithT(t)

	dsa := loadTestMemDatasourceWithAlias("emp")

	row := dsa.NextRow()

	v := row.Get("first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("emp.first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("emp.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

	v = row.Get("employee.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

}
