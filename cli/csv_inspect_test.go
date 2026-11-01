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

package cli

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/core"
)

func csvFile(t *testing.T, content string) string {
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestInspectCsv(t *testing.T) {
	g := NewGomegaWithT(t)

	// type comes from the first value that is not NULL
	columns, err := InspectCsv(csvFile(t, "\xEF\xBB\xBFa,b,c\n,x,\n5,y,\n"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(columns).To(Equal([]CsvColumn{
		{Name: "a", Type: "Number(Int)"}, {Name: "b", Type: "String"}, {Name: "c", Type: "Unknown"},
	}))

	// header only
	columns, err = InspectCsv(csvFile(t, "a,b\n"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(columns).To(Equal([]CsvColumn{{Name: "a", Type: "Unknown"}, {Name: "b", Type: "Unknown"}}))

	_, err = InspectCsv(csvFile(t, ""))
	g.Expect(err).To(MatchError(ContainSubstring("empty")))
}

func TestInspectCsvUsesSeparator(t *testing.T) {
	g := NewGomegaWithT(t)
	saved := core.CSVSeparator
	t.Cleanup(func() { core.CSVSeparator = saved })
	g.Expect(core.SetCsvSeparator(";")).To(Succeed())

	columns, err := InspectCsv(csvFile(t, "a;b\n1;2.5\n"))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(columns).To(HaveLen(2))
	g.Expect(columns[1].Type).To(Equal("Number(Float)"))
}
