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
	"github.com/mattn/go-runewidth"
	"github.com/uaraven/ansie"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"golang.org/x/term"
	"os"
	"strings"
)

var (
	columnFgColour = [][]ansie.Colour{
		{5, 0, 0},
		{0, 5, 0},
		{0, 0, 5},
		{4, 4, 0},
		{4, 0, 4},
		{0, 4, 4},
	}
)

type Column struct {
	Name              string
	DataWidth         int
	FgColour          ansie.Colour
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

func columnFgColorIndex(colorIdx int, offset int) ansie.Colour {
	r := max(columnFgColour[colorIdx][0]+offset, 0)
	g := max(columnFgColour[colorIdx][1]+offset, 0)
	b := max(columnFgColour[colorIdx][2]+offset, 0)
	return ansie.RgbTo216Colours(uint(r), uint(g), uint(b))
}

func InitTable(ds core.DataSource, maxWidth int) *Table {
	columnIdx := 0
	columns := funky.Map(ds.Header().ColumnsMetadata(), func(t core.ColumnMetadata) Column {
		c := Column{
			Name:              t.Name(),
			DataWidth:         runewidth.StringWidth(t.Name()),
			minWidth:          runewidth.StringWidth(t.Name()) + 2,
			FgColour:          columnFgColorIndex(columnIdx, 0),
			BgColour:          ansie.Black,
			AlternateBgColour: ansie.Grey11,
			HeaderFormat:      ansie.Ansi.Attr(ansie.Underline).Attr(ansie.Bold).String(),
		}
		columnIdx++
		if columnIdx >= len(columnFgColour) {
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
			l := runewidth.StringWidth(rows[i].GetByIndex(idx + 1).Repr())
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
		Zebra:               true,
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

// calculateWidths calculates widths of the columns taking into account the actual width of the data and
// max width of the terminal. If the columns do not fit, the widest columns are shrunk first, but no column
// becomes narrower than its minimum width, so the table can still be wider than the terminal.
//
// If maxWidth <= 0, then default width of 80 columns is assumed
func (t *Table) calculateWidths() {
	separators := 0
	if t.SeparateColumns && len(t.Columns) > 1 {
		separators = len(t.Columns) - 1
	}
	total := separators
	for idx, c := range t.Columns {
		t.Columns[idx].columnWidth = max(c.minWidth, c.DataWidth)
		total += t.Columns[idx].columnWidth
	}
	for ; total > t.TerminalWidth; total-- {
		widest := -1
		for idx, c := range t.Columns {
			if c.columnWidth > c.minWidth && (widest < 0 || c.columnWidth > t.Columns[widest].columnWidth) {
				widest = idx
			}
		}
		if widest < 0 {
			return
		}
		t.Columns[widest].columnWidth--
	}
}

func (t *Table) PrintHeader() string {
	var sb strings.Builder
	if t.SeparateRows {
		sb.WriteString(strings.Repeat(string(t.HorizontalSeparator), t.TerminalWidth))
	}
	for idx, col := range t.Columns {
		var suffix string
		suffixLen := col.columnWidth - runewidth.StringWidth(col.Name)
		if suffixLen > 0 {
			suffix = strings.Repeat(" ", suffixLen)
		}
		sb.WriteString(ansie.Ansi.Fg(ansie.BrightWhite).
			S("%s%s", col.HeaderFormat, col.Name).Reset().Fg(ansie.BrightWhite).A(suffix).Reset().String())
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
		for idx, col := range t.Columns {
			value := runewidth.Truncate(row.GetByIndex(idx+1).Repr(), col.columnWidth, "…")
			var bgColour ansie.Colour
			if t.Zebra && rowIdx%2 != 0 {
				bgColour = col.AlternateBgColour
			} else {
				bgColour = col.BgColour
			}
			sb.WriteString(ansie.Ansi.Fg(col.FgColour).Bg(bgColour).
				A(runewidth.FillRight(value, col.columnWidth)).Reset().String())
			if idx < len(t.Columns)-1 && t.SeparateColumns {
				sb.WriteRune(t.VerticalSeparator)
			}
		}
		sb.WriteRune('\n')

		if t.SeparateRows && rowIdx != len(data) {
			sb.WriteString(strings.Repeat(string(t.HorizontalSeparator), t.TerminalWidth))
		}
	}
	return sb.String()
}
