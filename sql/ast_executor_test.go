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
	"github.com/uaraven/csql/funky"
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
	g.Expect(rows[0].GetByIndex(1)).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].GetByIndex(2)).To(Equal(core.NewFloatValue(2.0)))
	g.Expect(rows[1].GetByIndex(1)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[1].GetByIndex(2)).To(Equal(core.NewFloatValue(1.2 * 2)))
	g.Expect(rows[3].GetByIndex(1)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[3].GetByIndex(2)).To(Equal(core.NewFloatValue(1.6 * 2)))

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

func TestAstExecutor_FullOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select * from "../test-data/authors.csv" 
    full join "../test-data/books.csv"
	on authors.id = books.author_id`)

	jds := Transform(&sql)

	rows := core.ReadAllRows(jds)
	g.Expect(rows).To(HaveLen(6))
	g.Expect(rows[0].Values()).To(Equal([]core.Value{
		core.NewIntValue(1),
		core.NewStringValue("J.R.R. Tolkien"),
		core.NewIntValue(1),
		core.NewIntValue(1),
		core.NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[1].Values()).To(Equal([]core.Value{
		core.NewIntValue(2),
		core.NewStringValue("Stephen King"),
		core.NewIntValue(2),
		core.NewIntValue(2),
		core.NewStringValue("Carrie"),
	}))

	g.Expect(rows[4].Values()).To(Equal([]core.Value{
		core.NewNullValue(),
		core.NewNullValue(),
		core.NewIntValue(4),
		core.NewIntValue(6),
		core.NewStringValue("Unwritten Book"),
	}))

	g.Expect(rows[5].Values()).To(Equal([]core.Value{
		core.NewNullValue(),
		core.NewNullValue(),
		core.NewIntValue(5),
		core.NewIntValue(6),
		core.NewStringValue("Unwritten Book 2"),
	}))
}

func TestAstExecutor_Nulls(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/nulls.csv"`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[1].Get("property").Value()).To(Equal(core.NewNullValue()))

	rows = ExecuteSql(`select * from "../test-data/nulls.csv" where property is null`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("name").Value()).To(Equal(core.NewStringValue("Billy-Bob")))

	rows = ExecuteSql(`select * from "../test-data/nulls.csv" where property is not null`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("name").Value()).To(Equal(core.NewStringValue("Brbra")))

	rows = ExecuteSql(`select id, null from "../test-data/nulls.csv" where name = 'Brbra'`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].GetByIndex(1)).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[0].GetByIndex(2).Value()).To(BeNil())
}

func TestAstExecutor_Logical(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name = 'James' and last_name = 'Snow'`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name = 'John' or last_name = 'Clarkson'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               not (first_name = 'John' or last_name = 'Clarkson')`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

}

func TestAstExecutor_Condition(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id > 3`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id < 2`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id >= 3`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               id <= 2`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name != 'James'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               last_name <> 'Snow'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_Like(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name like 'Jam%'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name not like 'Jam%'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name like 'J' || 'am%'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_Match(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name match 'Jam.*'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name not match 'Jam.*'`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name match 'J[ae].*' || 'm.*'`).GetRows()
	g.Expect(rows).To(HaveLen(3))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(4)))
}

func TestAstExecutor_In(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               first_name in ('James', 'John')`).GetRows()

	g.Expect(rows).To(HaveLen(3))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               dept_id in (11, 1+1)`).GetRows()

	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" where 
                                               dept_id not in (2, 3)`).GetRows()

	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
}

func TestAstTransformer_Limit(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" limit 2`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
}

func TestAstTransformer_OrderBy(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`select * from "../test-data/employees.csv" order by width`).GetRows()
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(1)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(4)))
	g.Expect(rows[3].Get("id").Value()).To(Equal(core.NewIntValue(3)))

	rows = ExecuteSql(`select * from "../test-data/employees.csv" order by last_name asc, 1 desc`).GetRows()
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Get("id").Value()).To(Equal(core.NewIntValue(3)))
	g.Expect(rows[1].Get("id").Value()).To(Equal(core.NewIntValue(4)))
	g.Expect(rows[2].Get("id").Value()).To(Equal(core.NewIntValue(2)))
	g.Expect(rows[3].Get("id").Value()).To(Equal(core.NewIntValue(1)))
}

func TestAstTransformer_UnionAll(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select first_name, last_name from "../test-data/employees.csv" 
		union all 
		select first_name, last_name from "../test-data/people.csv"`).GetRows()

	g.Expect(rows).To(HaveLen(8))

	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal("John"))
	g.Expect(rows[0].GetByIndex(2).Value()).To(Equal("Snow"))

	g.Expect(rows[4].GetByIndex(1).Value()).To(Equal("John"))
	g.Expect(rows[4].GetByIndex(2).Value()).To(Equal("Snow"))
}

func TestAstTransformer_Union(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select first_name, last_name from "../test-data/employees.csv" 
		union 
		select first_name, last_name from "../test-data/people.csv"`).GetRows()

	g.Expect(rows).To(HaveLen(4))
}

func TestAstTransformer_UnionRenaming(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select first_name from "../test-data/employees.csv" 
		union 
		select first_name from "../test-data/people.csv"
		union 
		select name as first_name from "../test-data/dept.csv"`).GetRows()

	g.Expect(rows).To(HaveLen(5))
	names := funky.Map(rows, func(row core.Row) string {
		return row.GetByIndex(1).Value().(string)
	})
	g.Expect(names).To(Equal([]string{"John", "James", "Jeremy", "The Wall", "The Grand Tour"}))
}

func TestAstTransformer_UnionDifferentProjections(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		ExecuteSql(`
		select first_name, last_name from "../test-data/employees.csv" 
		union 
		select first_name, first_name from "../test-data/people.csv"`)
	}).To(Panic())
}

func TestAstTransformer_FunctionCall(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select to_string(id), to_upper(first_name) from "../test-data/employees.csv" where to_lower(last_name) = 'snow'`).
		GetRows()
	g.Expect(rows).To(HaveLen(2))

	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal("1"))
	g.Expect(rows[0].GetByIndex(2).Value()).To(Equal("JOHN"))
	g.Expect(rows[1].GetByIndex(1).Value()).To(Equal("2"))
	g.Expect(rows[1].GetByIndex(2).Value()).To(Equal("JAMES"))
}

func TestAstTransformer_CoalesceFunctionCall(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select id, coalesce(property, 'None') from "../test-data/nulls.csv"`).
		GetRows()
	g.Expect(rows).To(HaveLen(2))

	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal(int64(1)))
	g.Expect(rows[0].GetByIndex(2).Value()).To(Equal("Mansion"))
	g.Expect(rows[1].GetByIndex(1).Value()).To(Equal(int64(2)))
	g.Expect(rows[1].GetByIndex(2).Value()).To(Equal("None"))
}

func TestAstTransformer_Count(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select count(distinct city) from "../test-data/cities.csv"`).
		GetRows()
	g.Expect(rows).To(HaveLen(1))

	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal(int64(6)))
}

func TestAstTransformer_CountGrouping(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select count(city), country from "../test-data/cities.csv"
		group by country`).
		GetRows()
	g.Expect(rows).To(HaveLen(8))
	g.Expect(core.RowsToSlice(rows, "COUNT(city)", "country")).To(ContainElements(
		[]interface{}{int64(1), "Belgium"},
		[]interface{}{int64(1), "Australia"},
		[]interface{}{int64(1), "Poland"},
		[]interface{}{int64(1), "Ukraine"},
		[]interface{}{int64(1), "Croatia"},
		[]interface{}{int64(1), "UK"},
		[]interface{}{int64(3), "Canada"},
		[]interface{}{int64(6), "USA"},
	))
}

func TestAstTransformer_MaxGroupBy(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select max(area), country from "../test-data/cities.csv"
		group by country`).
		GetRows()
	g.Expect(rows).To(HaveLen(8))
	g.Expect(core.RowsToSlice(rows, "MAX(area)", "country")).To(ContainElements(
		[]interface{}{21.32, "Belgium"},
		[]interface{}{1.1, "Australia"},
		[]interface{}{int64(6100), "Poland"},
		[]interface{}{162.42, "Ukraine"},
		[]interface{}{21.35, "Croatia"},
		[]interface{}{1572.03, "UK"},
		[]interface{}{420.5, "Canada"},
		[]interface{}{581.58, "USA"},
	))
}

func TestAstTransformer_AvgGroupBy(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select avg(population), city from "../test-data/cities.csv"
		group by city`).
		GetRows()
	g.Expect(rows).To(HaveLen(6))
	g.Expect(core.RowsToSlice(rows, "AVG(population)", "city")).To(ContainElements(
		[]interface{}{42615.0, "Dubrovnik"},
		[]interface{}{1848278.4, "London"},
		[]interface{}{1015826.0, "Odesa"},
		[]interface{}{12150.5, "Odessa"},
		[]interface{}{1527656.5, "Warsaw"},
		[]interface{}{42164.75, "Waterloo"},
	))
}

func TestAstTransformer_Having(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select min(area), country from "../test-data/cities.csv"
		group by country
		having sum(population) > 10000`).
		GetRows()

	g.Expect(rows).To(HaveLen(8))
	columns := []string{"MIN(area)", "country"}
	g.Expect(core.RowsToSlice(rows, columns...)).To(ContainElements(
		[]interface{}{1.1, "Australia"},
		[]interface{}{21.32, "Belgium"},
		[]interface{}{64.06, "Canada"},
		[]interface{}{21.35, "Croatia"},
		[]interface{}{int64(6100), "Poland"},
		[]interface{}{1572.03, "UK"},
		[]interface{}{7.11, "USA"},
		[]interface{}{162.42, "Ukraine"},
	))
}

func TestAstTransformer_AggregateFunctionInWhere(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		ExecuteSql(`
		select * from "../test-data/cities.csv"
		where sum(population) > 10000`).
			GetRows()
	}).To(Panic())
}

func TestAstTransformer_AggregateFunctionInProjectionInExpression(t *testing.T) {
	g := NewGomegaWithT(t)

	rows := ExecuteSql(`
		select round(avg(population) / 2) from "../test-data/cities.csv"`).
		GetRows()
	g.Expect(rows).To(HaveLen(1))
	columns := []string{"ROUND(AVG(population) / 2)"}
	sl := core.RowsToSlice(rows, columns...)
	g.Expect(sl[0]).To(ContainElements(
		[]interface{}{int64(451604)},
	))
}
