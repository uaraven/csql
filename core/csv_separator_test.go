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

func TestSetCsvSeparator(t *testing.T) {
	g := NewGomegaWithT(t)
	saved := CSVSeparator
	t.Cleanup(func() { CSVSeparator = saved })

	for input, expected := range map[string]rune{
		";": ';', `"|"`: '|', `'\t'`: '\t', `\t`: '\t', `"\t"`: '\t', "é": 'é', " ": ' ', `\u00e9`: 'é',
	} {
		CSVSeparator = ','
		g.Expect(SetCsvSeparator(input)).To(Succeed(), input)
		g.Expect(CSVSeparator).To(Equal(expected), input)
	}

	for _, input := range []string{"", `""`, "ab", `"ab"`, `"`, `\n`, `\r`, `\x00`, `�`} {
		CSVSeparator = ','
		g.Expect(SetCsvSeparator(input)).ToNot(Succeed(), input)
		g.Expect(CSVSeparator).To(Equal(','), "separator is unchanged after an error: "+input)
	}
}

func TestSeparatorIsUsedForReadingAndWriting(t *testing.T) {
	g := NewGomegaWithT(t)
	saved := CSVSeparator
	t.Cleanup(func() { CSVSeparator = saved })
	dir := t.TempDir()
	in := filepath.Join(dir, "in.csv")
	out := filepath.Join(dir, "out.csv")
	g.Expect(os.WriteFile(in, []byte("a;b\n1;x,y\n"), 0600)).To(Succeed())

	g.Expect(SetCsvSeparator(";")).To(Succeed())
	ds := NewCsvDataSource(in)
	g.Expect(ds.Header().ColumnCount()).To(Equal(2))
	g.Expect(ds.GetRows()[0].GetByIndex(2).Value()).To(Equal("x,y"))

	g.Expect(WriteDataSourceToFile(ds, out)).To(Succeed())
	content, err := os.ReadFile(out)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(string(content)).To(Equal("a;b\n1;x,y\n"))
}
