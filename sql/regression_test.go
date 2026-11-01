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
