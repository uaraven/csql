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
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func writeTestCsv(t *testing.T, content string) string {
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestCsvByteOrderMarkIsSkipped(t *testing.T) {
	g := NewGomegaWithT(t)
	ds := NewCsvDataSource(writeTestCsv(t, "\xEF\xBB\xBFid,name\n1,x\n"))

	g.Expect(ds.Header().ColumnsMetadata()[0].Name()).To(Equal("id"))
	g.Expect(ds.GetRows()[0].Get("id").Value().Value()).To(Equal(int64(1)))
}

func TestCsvMalformedRows(t *testing.T) {
	g := NewGomegaWithT(t)
	saved := IgnoreMalformedCsvRows
	t.Cleanup(func() { IgnoreMalformedCsvRows = saved })
	content := "id,name\n1,a\n2\n3,b\"c\n4,d\n"

	IgnoreMalformedCsvRows = false
	g.Expect(func() { NewCsvDataSource(writeTestCsv(t, content)) }).To(PanicWith(MatchError(ContainSubstring("data.csv"))))

	IgnoreMalformedCsvRows = true
	rows := NewCsvDataSource(writeTestCsv(t, content)).GetRows()
	g.Expect(rows).To(HaveLen(2))
	// skipped rows do not leave gaps in row ids
	g.Expect(rows[0].Id()).To(Equal(0))
	g.Expect(rows[1].Id()).To(Equal(1))
	g.Expect(rows[1].GetByIndex(1).Value()).To(Equal(int64(4)))
}

func TestCsvEmptyFileMentionsFileName(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { NewCsvDataSource(writeTestCsv(t, "")) }).
		To(PanicWith(MatchError(And(ContainSubstring("data.csv"), ContainSubstring("empty")))))
}
