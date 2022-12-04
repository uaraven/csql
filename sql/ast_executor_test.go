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

package sql

import (
	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/core"
	"testing"
)

func TestAstExecutor_Select(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL("select * from \"../test-data/employees.csv\"")

	ds := Transform(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(5))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
}

func TestAstExecutor_SelectWhere(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL("select * from \"../test-data/employees.csv\" where dept_id = 1")

	ds := Transform(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(5))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("width").Value()).To(Equal(core.NewFloatValue(1.0)))
	g.Expect(rows[1].Get("first_name").Value()).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[1].Get("width").Value()).To(Equal(core.NewFloatValue(1.2)))
}

func TestAstExecutor_SelectProjection(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select first_name, width*2 as w2 from "../test-data/employees.csv"`)

	ds := Transform(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(2))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0)).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].GetByIndex(1)).To(Equal(core.NewFloatValue(2.0)))
	g.Expect(rows[1].GetByIndex(0)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[1].GetByIndex(1)).To(Equal(core.NewFloatValue(1.2 * 2)))
	g.Expect(rows[3].GetByIndex(0)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[3].GetByIndex(1)).To(Equal(core.NewFloatValue(1.6 * 2)))

	g.Expect(rows[0].Get("w2").IsPresent()).To(BeTrue())
}

func TestAstExecutor_SelectProjectionAliases(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select emp.first_name as fn, emp.last_name from "../test-data/employees.csv" emp`)

	ds := Transform(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(2))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Get("fn").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("last_name").Value()).To(Equal(core.NewStringValue("Snow")))
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeFalse())
}

func TestAstExecutor_LeftOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select * from "../test-data/authors.csv" 
    left join "../test-data/books.csv"
	on authors.id = books.author_id`)

	jds := Transform(&sql)

	rows := core.ReadAllRows(jds)
	g.Expect(rows[0].Values()).To(Equal([]core.Value{
		core.NewIntValue(1),
		core.NewStringValue("J.R.R. Tolkien"),
		core.NewIntValue(1),
		core.NewIntValue(1),
		core.NewStringValue("Lord Of The Rings")}))
	g.Expect(rows[1].Values()).To(Equal([]core.Value{
		core.NewIntValue(2),
		core.NewStringValue("Stephen King"),
		core.NewIntValue(2),
		core.NewIntValue(2),
		core.NewStringValue("Carrie")}))
	g.Expect(rows[2].Values()).To(Equal([]core.Value{
		core.NewIntValue(2),
		core.NewStringValue("Stephen King"),
		core.NewIntValue(3),
		core.NewIntValue(2),
		core.NewStringValue("The Dark Tower")}))
}

func TestAstExecutor_RightOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select * from "../test-data/authors.csv" 
    right join "../test-data/books.csv"
	on authors.id = books.author_id`)

	jds := Transform(&sql)

	rows := core.ReadAllRows(jds)
	g.Expect(rows).To(HaveLen(5))
	g.Expect(rows[0].Values()).To(Equal([]core.Value{
		core.NewIntValue(1),
		core.NewStringValue("J.R.R. Tolkien"),
		core.NewIntValue(1),
		core.NewIntValue(1),
		core.NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[4].Values()).To(Equal([]core.Value{
		core.NewNullValue(),
		core.NewNullValue(),
		core.NewIntValue(5),
		core.NewIntValue(6),
		core.NewStringValue("Unwritten Book 2"),
	}))
}

func TestAstExecutor_CrossJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	ast := ParseSQL(`select * from "../test-data/employees.csv", "../test-data/dept.csv"`)

	cj := Transform(&ast)

	rows := core.ReadAllRows(cj)
	g.Expect(rows).To(HaveLen(8))

	g.Expect(rows[0].Get("employees.dept_id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].Get("dept.id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("name").Value()).To(Equal(core.NewStringValue("The Wall")))

	g.Expect(rows[1].Get("employees.dept_id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("dept.id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("employees.first_name").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[1].Get("dept.name").Value()).To(Equal(core.NewStringValue("The Grand Tour")))
}

func TestAstExecutor_InnerJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	ast := ParseSQL(`select * from "../test-data/employees.csv" join "../test-data/dept.csv" on dept_id = dept.id`)

	cj := Transform(&ast)

	rows := core.ReadAllRows(cj)
	g.Expect(rows).To(HaveLen(4))

	g.Expect(rows[0].Get("employees.dept_id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].Get("dept.id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].Get("employees.first_name").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("dept.name").Value()).To(Equal(core.NewStringValue("The Wall")))

	g.Expect(rows[3].Get("employees.dept_id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[3].Get("dept.id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[3].Get("employees.first_name").Value()).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[3].Get("dept.name").Value()).To(Equal(core.NewStringValue("The Grand Tour")))
}

func TestAstExecutor_Nulls(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/nulls.csv"`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[1].Get("property").Value()).To(Equal(core.NewNullValue()))

	rows = ExecuteSql(`select * from "../test-data/nulls.csv" where property is null`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("name").Value()).To(Equal(core.NewStringValue("Billy-Bob")))

	rows = ExecuteSql(`select * from "../test-data/nulls.csv" where property is not null`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("name").Value()).To(Equal(core.NewStringValue("Brbra")))

	rows = ExecuteSql(`select id, null from "../test-data/nulls.csv" where name = 'Brbra'`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].GetByIndex(0)).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].GetByIndex(1)).To(Equal(core.NewNullValue()))
}

func TestAstExecutor_Logical(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name = 'James' and last_name = 'Snow'`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name = 'John' or last_name = 'Clarkson'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               not (first_name = 'John' or last_name = 'Clarkson')`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

}

func TestAstExecutor_Condition(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id > 3`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id < 2`)
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id >= 3`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id <= 2`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name != 'James'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               last_name <> 'Snow'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_Like(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name like 'Jam%'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name not like 'Jam%'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name like 'J' || 'am%'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_Match(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name match 'Jam.*'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name not match 'Jam.*'`)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name match 'J[ae].*' || 'm.*'`)
	g.Expect(rows).To(HaveLen(3))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_In(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name in ('James', 'John')`)

	g.Expect(rows).To(HaveLen(3))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               dept_id in (11, 1+1)`)

	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               dept_id not in (2, 3)`)

	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
}
