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
	"os"
	"path/filepath"
	"testing"

	"github.com/uaraven/csql/core"

	. "github.com/onsi/gomega"
)

func writeCsv(t *testing.T, name string, content string) string {
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestSelectRowIdColumn(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,name\n1,x\n2,y\n")
	b := writeCsv(t, "b.csv", "aid,v\n1,p\n2,q\n")

	rows := ExecuteSql(`select __row_id, name from "` + a + `"`).GetRows()
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[1].GetByIndex(1).Value()).To(Equal(int64(1)))
	g.Expect(rows[1].GetByIndex(2).Value()).To(Equal("y"))

	rows = ExecuteSql(`select a.id, __row_id from "` + a + `" a join "` + b + `" b on a.id = b.aid`).GetRows()
	g.Expect(rows).To(HaveLen(2))
}

func TestLimitZero(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n2\n3\n")

	g.Expect(ExecuteSql(`select id from "` + a + `" limit 0`).GetRows()).To(BeEmpty())
	g.Expect(ExecuteSql(`select id from "` + a + `" limit 2`).GetRows()).To(HaveLen(2))
	g.Expect(ExecuteSql(`select id from "` + a + `"`).GetRows()).To(HaveLen(3))
}

func TestOuterJoinRowIds(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n2\n3\n")
	b := writeCsv(t, "b.csv", "aid\n1\n1\n3\n")

	for _, kind := range []string{"left", "right", "full"} {
		rows := ExecuteSql(`select __row_id from "` + a + `" a ` + kind + ` join "` + b + `" b on a.id = b.aid`).GetRows()
		g.Expect(len(rows)).To(BeNumerically(">=", 3), kind)
		for idx, row := range rows {
			g.Expect(row.GetByIndex(1).Value()).To(Equal(int64(idx)), kind)
		}
	}
}

func TestFullJoinKeepsDuplicateUnmatchedRightRows(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n")
	b := writeCsv(t, "b.csv", "aid,v\n9,x\n9,x\n1,y\n")

	rows := ExecuteSql(`select * from "` + a + `" a full join "` + b + `" b on a.id = b.aid`).GetRows()
	g.Expect(rows).To(HaveLen(3))
}

func TestFloatPrecision(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "a\n0.1234567\n0.1234568\n1.5\n")

	rows := ExecuteSql(`select a from "` + a + `"`).GetRows()
	g.Expect(rows[0].GetByIndex(1).String()).To(Equal("0.1234567"))
	g.Expect(rows[2].GetByIndex(1).String()).To(Equal("1.5"))
	g.Expect(ExecuteSql(`select distinct a from "` + a + `"`).GetRows()).To(HaveLen(3))
	g.Expect(ExecuteSql(`select count(distinct a) from "` + a + `"`).GetRows()[0].GetByIndex(1).Value()).To(Equal(int64(3)))
}

func firstValue(ds interface{ GetRows() []core.Row }) interface{} {
	return ds.GetRows()[0].GetByIndex(1).Value()
}

func TestComparisonDoesNotHideErrors(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n")

	g.Expect(func() { ExecuteSql(`select id from "` + a + `" where nope = 1`) }).
		To(PanicWith(MatchError(ContainSubstring("unknown column"))))
	g.Expect(func() { ExecuteSql(`select id from "` + a + `" where id = 1/0`) }).
		ToNot(PanicWith(MatchError(ContainSubstring("Incompatible operand types"))))
}

func TestKeysDoNotCollide(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "a,b\n\"x-'y\",z\nx,\"y'-'z\"\n1,2\n\"1-2\",\n")

	g.Expect(ExecuteSql(`select distinct a, b from "` + a + `"`).GetRows()).To(HaveLen(4))
	g.Expect(ExecuteSql(`select a, b, count(*) from "` + a + `" group by a, b`).GetRows()).To(HaveLen(4))
}

func TestAggregatesSkipNulls(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "g,v,s\nx,1,b\nx,,a\nx,null,c\nx,3,\n")
	q := func(expr string) interface{} {
		return firstValue(ExecuteSql(`select ` + expr + ` from "` + a + `"`))
	}

	g.Expect(q("count(v)")).To(Equal(int64(2)))
	g.Expect(q("count(*)")).To(Equal(int64(4)))
	g.Expect(q("sum(v)")).To(Equal(int64(4)))
	g.Expect(q("avg(v)")).To(Equal(2.0))
	g.Expect(q("min(v)")).To(Equal(int64(1)))
	g.Expect(q("max(v)")).To(Equal(int64(3)))
	g.Expect(q("min(s)")).To(Equal("a"))
	g.Expect(q("max(s)")).To(Equal("c"))
}

func TestAggregatesOverNoRows(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v\n1,5\n")

	rows := ExecuteSql(`select count(*), sum(v), min(v), avg(v) from "` + a + `" where id = 99`).GetRows()
	g.Expect(rows).To(HaveLen(1))
	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal(int64(0)))
	g.Expect(rows[0].GetByIndex(2).Type()).To(Equal(core.TypeNull))
	g.Expect(rows[0].GetByIndex(3).Type()).To(Equal(core.TypeNull))
	g.Expect(rows[0].GetByIndex(4).Type()).To(Equal(core.TypeNull))

	// GROUP BY over no rows has no groups
	g.Expect(ExecuteSql(`select id, count(*) from "` + a + `" where id = 99 group by id`).GetRows()).To(BeEmpty())
}

func TestDistinctWithGroupBy(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "g,v\nx,1\ny,1\nz,2\n")

	rows := ExecuteSql(`select distinct count(*) from "` + a + `" group by g`).GetRows()
	g.Expect(rows).To(HaveLen(1))
}

func TestTrailingGarbageIsRejected(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n2\n")
	src := `"` + a + `"`

	for _, q := range []string{
		`select id from ` + src + ` where id=1 garbage`,
		`select id from ` + src + `; garbage here`,
		`select id from ` + src + ` limit 1 where id = 1`,
		`select id from ` + src + ` ; ;`,
	} {
		g.Expect(func() { ExecuteSql(q) }).To(Panic(), q)
	}
	for _, q := range []string{
		`select id from ` + src,
		`select id from ` + src + `;`,
		`select id from ` + src + ` where id = 1;`,
	} {
		g.Expect(func() { ExecuteSql(q) }).ToNot(Panic(), q)
	}
}

func column(ds interface{ GetRows() []core.Row }, index int) []interface{} {
	result := make([]interface{}, 0)
	for _, row := range ds.GetRows() {
		result = append(result, row.GetByIndex(index).Value())
	}
	return result
}

func TestEmptyCellsAreNull(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v,s\n1,10,a\n2,,\n3,null,c\n")
	src := `"` + a + `"`

	g.Expect(column(ExecuteSql(`select id from `+src+` where v is null`), 1)).To(Equal([]interface{}{int64(2), int64(3)}))
	g.Expect(column(ExecuteSql(`select id from `+src+` where s is null`), 1)).To(Equal([]interface{}{int64(2)}))
	g.Expect(column(ExecuteSql(`select v + 1 from `+src), 1)).To(Equal([]interface{}{int64(11), nil, nil}))
	g.Expect(column(ExecuteSql(`select s || 'x' from `+src), 1)).To(Equal([]interface{}{"ax", nil, "cx"}))
	g.Expect(column(ExecuteSql(`select to_upper(s) from `+src), 1)).To(Equal([]interface{}{"A", nil, "C"}))
	g.Expect(column(ExecuteSql(`select coalesce(v, 0) from `+src), 1)).To(Equal([]interface{}{int64(10), int64(0), int64(0)}))
	g.Expect(firstValue(ExecuteSql(`select avg(v) from ` + src))).To(Equal(10.0))
}

func TestThreeValuedLogic(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v\n1,1\n2,\n3,5\n")
	src := `"` + a + `"`
	ids := func(where string) []interface{} {
		return column(ExecuteSql(`select id from `+src+` where `+where), 1)
	}

	g.Expect(ids(`v = 1`)).To(Equal([]interface{}{int64(1)}))
	g.Expect(ids(`not v = 1`)).To(Equal([]interface{}{int64(3)}))
	g.Expect(ids(`v <> 1`)).To(Equal([]interface{}{int64(3)}))
	g.Expect(ids(`not (v = 1 or v = 5)`)).To(BeEmpty())
	g.Expect(ids(`v = 1 or v is null`)).To(Equal([]interface{}{int64(1), int64(2)}))
	g.Expect(ids(`v = 99 and v is null`)).To(BeEmpty())
	g.Expect(ids(`not (v = 99 and v = 1)`)).To(Equal([]interface{}{int64(1), int64(3)}))
	g.Expect(ids(`not (v = 1 and 1 = 2)`)).To(Equal([]interface{}{int64(1), int64(2), int64(3)}))
	g.Expect(ids(`v in (1, 2.0)`)).To(Equal([]interface{}{int64(1)}))
	g.Expect(ids(`v not in (1, 2)`)).To(Equal([]interface{}{int64(3)}))
	g.Expect(ids(`v in (1, null)`)).To(Equal([]interface{}{int64(1)}))
	g.Expect(ids(`v not in (1, null)`)).To(BeEmpty())
	g.Expect(ids(`v between 1 and 5`)).To(Equal([]interface{}{int64(1), int64(3)}))
}

func TestGroupByKeepsFirstSeenOrder(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "g\nz\nb\nz\na\nq\nb\nm\n")

	expected := []interface{}{"z", "b", "a", "q", "m"}
	for i := 0; i < 20; i++ {
		g.Expect(column(ExecuteSql(`select g, count(*) from "`+a+`" group by g`), 1)).To(Equal(expected))
	}
}

func TestOrderBy(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v,s\n1,b,x\n2,,x\n3,2,y\n4,1.5,y\n5,a,z\n6,2,z\n")
	src := `"` + a + `"`
	ids := func(query string) []interface{} {
		return column(ExecuteSql(query), 1)
	}
	i := func(v ...int64) []interface{} {
		r := make([]interface{}, len(v))
		for idx, x := range v {
			r[idx] = x
		}
		return r
	}

	// NULL first, then numbers, then strings; ties keep the input order
	g.Expect(ids(`select id from ` + src + ` order by v`)).To(Equal(i(2, 4, 3, 6, 5, 1)))
	g.Expect(ids(`select id from ` + src + ` order by v desc`)).To(Equal(i(1, 5, 3, 6, 4, 2)))
	g.Expect(ids(`select id from ` + src + ` order by s desc, id desc`)).To(Equal(i(6, 5, 4, 3, 2, 1)))
	// order by a column that is not in the select list
	g.Expect(ids(`select s from ` + src + ` order by id desc limit 2`)).To(Equal([]interface{}{"z", "z"}))
	g.Expect(ids(`select id from ` + src + ` order by s desc, id limit 3`)).To(Equal(i(5, 6, 3)))
	// aliases and positions still work
	g.Expect(ids(`select id as k from ` + src + ` order by k desc limit 2`)).To(Equal(i(6, 5)))
	g.Expect(ids(`select id from ` + src + ` order by 1 desc limit 1`)).To(Equal(i(6)))

	g.Expect(func() { ExecuteSql(`select id from ` + src + ` where id = 1 order by nope`) }).
		To(PanicWith(MatchError(ContainSubstring("unknown column"))))
}

func TestNullsAreExportedAsEmptyCells(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "a,b\n1,\nnull,x\n")
	out := filepath.Join(t.TempDir(), "out.csv")

	g.Expect(core.WriteDataSourceToFile(ExecuteSql(`select * from "`+a+`"`), out)).To(Succeed())
	content, err := os.ReadFile(out)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(string(content)).To(Equal("a,b\n1,\n,x\n"))
}

func TestTempTableSelfJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v\n1,x\n2,y\n3,z\n")
	ExecuteSql(`select * from "` + a + `" into temp self_tmp`)
	t.Cleanup(func() { delete(core.TableCache, "self_tmp") })

	g.Expect(ExecuteSql(`select x.id, y.id from self_tmp x join self_tmp y on x.id = y.id`).GetRows()).To(HaveLen(3))
	g.Expect(ExecuteSql(`select x.id, y.id from self_tmp x join self_tmp y on x.id < y.id`).GetRows()).To(HaveLen(3))
	g.Expect(ExecuteSql(`select * from self_tmp x cross join self_tmp y`).GetRows()).To(HaveLen(9))
	g.Expect(ExecuteSql(`select * from self_tmp x left join self_tmp y on x.id = y.id + 1`).GetRows()).To(HaveLen(3))
}

func TestTempTableNameAndAlias(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id,v\n1,x\n2,y\n")
	b := writeCsv(t, "b.csv", "id\n7\n8\n9\n")
	ExecuteSql(`select * from "` + a + `" into temp alias_tmp`)
	t.Cleanup(func() { delete(core.TableCache, "alias_tmp") })

	// the temp table can be referred to by its name and by an alias
	g.Expect(ExecuteSql(`select alias_tmp.id from alias_tmp`).GetRows()).To(HaveLen(2))
	g.Expect(ExecuteSql(`select t.id from alias_tmp t`).GetRows()).To(HaveLen(2))
	// an alias equal to the name of a temp table does not replace the file
	g.Expect(ExecuteSql(`select id from "` + b + `" alias_tmp`).GetRows()).To(HaveLen(3))
	// ...and a file with the same base name as a temp table is a file
	g.Expect(ExecuteSql(`select id from "` + b + `"`).GetRows()).To(HaveLen(3))
}

func TestLastResultCanBeQueried(t *testing.T) {
	g := NewGomegaWithT(t)
	a := writeCsv(t, "a.csv", "id\n1\n2\n3\n")

	ExecuteSql(`select id from "` + a + `" where id > 1`)
	g.Expect(ExecuteSql(`select id from ".last"`).GetRows()).To(HaveLen(2))
	g.Expect(ExecuteSql(`select id from ".last"`).GetRows()).To(HaveLen(2))
}
