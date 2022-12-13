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
	"fmt"
	"github.com/uaraven/ansie"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"golang.org/x/exp/constraints"
	"golang.org/x/term"
	"os"
	"strings"
)

var (
	columnBgColour = [][]ansie.Colour{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
	}
)

type Column struct {
	Name              string
	DataWidth         int
	BgColour          ansie.Colour
	AlternateBgColour ansie.Colour
	HeaderFormat      string
	minWidth          int
	columnWidth       int
}

type Table struct {
	Columns             []Column
	Zebra               bool
	SeparateColumns     bool
	SeparateRows        bool
	VerticalSeparator   rune
	HorizontalSeparator rune
	CrossSeparator      rune
	TerminalWidth       int
}

func columnBgColorIndex(colorIdx int, offset int) ansie.Colour {
	r := max(columnBgColour[colorIdx][0]+offset, 0)
	g := max(columnBgColour[colorIdx][1]+offset, 0)
	b := max(columnBgColour[colorIdx][2]+offset, 0)
	return ansie.RgbTo216Colours(uint(r), uint(g), uint(b))
}

func InitTable(ds core.DataSource, maxWidth int) *Table {
	columnIdx := 0
	columns := funky.Map(ds.Header().ColumnsMetadata(), func(t core.ColumnMetadata) Column {
		c := Column{
			Name:              t.Name(),
			DataWidth:         len(t.Name()),
			minWidth:          len(t.Name()) + 2,
			BgColour:          columnBgColorIndex(columnIdx, 0),
			AlternateBgColour: columnBgColorIndex(columnIdx, 1),
			HeaderFormat:      ansie.Ansi.Attr(ansie.Underline).Attr(ansie.Bold).String(),
		}
		columnIdx++
		if columnIdx >= len(columnBgColour) {
			columnIdx = 0
		}
		return c
	})

	rows := ds.GetRows()
	readRows := len(rows)
	if readRows > 100 {
		readRows = 100
	}

	for i := 0; i < readRows; i++ {
		for idx := range columns {
			l := len(rows[i].GetByIndex(idx + 1).Repr())
			if columns[idx].DataWidth < l {
				columns[idx].DataWidth = l
			}
		}
	}

	var mw = maxWidth
	if mw < 0 && term.IsTerminal(int(os.Stdout.Fd())) {
		w, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err == nil {
			mw = w
		}
	}
	if mw <= 0 {
		mw = 80
	}
	t := &Table{
		Columns:             columns,
		Zebra:               false,
		SeparateColumns:     true,
		SeparateRows:        false,
		VerticalSeparator:   '│',
		HorizontalSeparator: '─',
		CrossSeparator:      '┼',
		TerminalWidth:       mw,
	}
	t.calculateWidths()
	return t
}

// CalculateWidths calculates widths of the columns taking into account the actual width of the data and
// max width of the terminal. If terminal width cannot be determined, supplied maxWidth is used.
//
// If maxWidth <= 0, then default width of 80 columns is assumed
func (t *Table) calculateWidths() {
	totalMinWidth := funky.Sum(funky.Map(t.Columns, func(c Column) int {
		return c.minWidth
	}))
	if t.SeparateColumns {
		totalMinWidth += len(t.Columns) - 1
	}

	if totalMinWidth < t.TerminalWidth {
		for idx, cw := range t.Columns {
			t.Columns[idx].columnWidth = max(cw.minWidth, cw.DataWidth)
		}
	}
}

func max[T constraints.Ordered](t1 T, t2 T) T {
	if t1 >= t2 {
		return t1
	}
	return t2
}

func (t *Table) PrintHeader() string {
	var sb strings.Builder
	if t.SeparateRows {
		sb.WriteString(strings.Repeat(string(t.HorizontalSeparator), t.TerminalWidth))
	}
	for idx, col := range t.Columns {
		var suffix string
		suffixLen := col.columnWidth - len(col.Name)
		if suffixLen > 0 {
			suffix = strings.Repeat(" ", suffixLen)
		}
		sb.WriteString(ansie.Ansi.Bg(col.BgColour).
			S("%s%s", col.HeaderFormat, col.Name).Reset().Bg(col.BgColour).A(suffix).Reset().String())
		if idx < len(t.Columns)-1 && t.SeparateColumns {
			sb.WriteRune(t.VerticalSeparator)
		}
	}
	if t.SeparateRows {
		sb.WriteString(strings.Repeat(string(t.HorizontalSeparator), t.TerminalWidth))
	}
	return sb.String()
}

func (t *Table) PrintData(data []core.Row) string {
	var sb strings.Builder
	for rowIdx, row := range data {
		var rowSb strings.Builder

		for idx, col := range t.Columns {
			value := row.GetByIndex(idx + 1).Repr()
			var bgColour ansie.Colour
			if t.Zebra && rowIdx%2 != 0 {
				bgColour = col.AlternateBgColour
			} else {
				bgColour = col.BgColour
			}
			if len(value) > col.columnWidth {
				value = value[:col.columnWidth-1] + "…"
			}
			sb.WriteString(ansie.Ansi.Bg(bgColour).
				S("%*s", -col.columnWidth, value).Reset().String())
			rowSb.WriteString(fmt.Sprintf("%*s", -col.columnWidth, value))
			if idx < len(t.Columns)-1 && t.SeparateColumns {
				rowSb.WriteRune(t.VerticalSeparator)
				sb.WriteRune(t.VerticalSeparator)
			}
		}

		rowString := rowSb.String()

		if len(rowString) < t.TerminalWidth {
			sb.WriteRune('\n')
		}

		if t.SeparateRows && rowIdx != len(data) {
			sb.WriteString(strings.Repeat(string(t.HorizontalSeparator), t.TerminalWidth))
		}
	}
	return sb.String()
}
