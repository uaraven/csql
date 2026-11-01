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
