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
	"regexp"
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/core"
)

var ansiPattern = regexp.MustCompile("\x1b\\[[0-9;]*m")

func renderedLines(t *testing.T, content string, width int) (*Table, []string) {
	ds := core.NewCsvDataSource(csvFile(t, content))
	table := InitTable(ds, width)
	data := table.PrintData(ds.GetRows())
	return table, strings.Split(strings.TrimRight(data, "\n"), "\n")
}

func TestTableManyNarrowColumnsDoesNotPanic(t *testing.T) {
	g := NewGomegaWithT(t)
	var header, row []string
	for i := 0; i < 60; i++ {
		header = append(header, "c")
		row = append(row, "1")
	}
	content := strings.Join(header, ",") + "\n" + strings.Join(row, ",") + "\n"

	g.Expect(func() { renderedLines(t, content, 40) }).ToNot(Panic())
}

func TestTableShrinksWidestColumnsFirst(t *testing.T) {
	g := NewGomegaWithT(t)
	table, _ := renderedLines(t, "id,description\n1,"+strings.Repeat("x", 200)+"\n", 40)

	total := 0
	for _, c := range table.Columns {
		total += c.columnWidth
	}
	g.Expect(total + 1).To(Equal(40))
	g.Expect(table.Columns[0].columnWidth).To(Equal(table.Columns[0].minWidth))
}

func TestTableNonAsciiText(t *testing.T) {
	g := NewGomegaWithT(t)
	_, lines := renderedLines(t, "name,city\nЖанна,Київ\nabcdefghijklmnopqrstuvwxyzабвгдежзийклмнопрстуфхцчшщ,Львів\n", 30)

	g.Expect(lines).To(HaveLen(2))
	for _, line := range lines {
		plain := ansiPattern.ReplaceAllString(line, "")
		g.Expect(runewidth.StringWidth(plain)).To(Equal(runewidth.StringWidth(
			ansiPattern.ReplaceAllString(lines[0], ""))))
		g.Expect(plain).To(MatchRegexp(`^[^\x{FFFD}]*$`))
	}
}
