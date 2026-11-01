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
	"github.com/uaraven/ansie"
	"github.com/uaraven/csql/core"
)

func TestRuneIndex(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(RuneIndex("this is a \"string\";", ';')).To(Equal(18))
	g.Expect(RuneIndex("this is a \"string;\"", ';')).To(Equal(-1))
	g.Expect(RuneIndex("this is a \"string\\\";\"", ';')).To(Equal(-1))

	g.Expect(RuneIndex("this is a \"string;\";dw", ';')).To(Equal(19))
}

func TestSplitArgs(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(splitArgs("")).To(BeEmpty())
	g.Expect(splitArgs("   ")).To(BeEmpty())
	g.Expect(splitArgs("ls")).To(Equal([]string{"ls"}))
	g.Expect(splitArgs("ls\t-l  dir")).To(Equal([]string{"ls", "-l", "dir"}))
	g.Expect(splitArgs(`cd "my dir"`)).To(Equal([]string{"cd", "my dir"}))
	g.Expect(splitArgs(`cd 'my dir'/sub`)).To(Equal([]string{"cd", "my dir/sub"}))
	g.Expect(splitArgs(`csv --null ""`)).To(Equal([]string{"csv", "--null", ""}))
	g.Expect(splitArgs(`cd "unclosed dir`)).To(Equal([]string{"cd", "unclosed dir"}))
}

func TestShellIsCommand(t *testing.T) {
	g := NewGomegaWithT(t)
	shell := &CsqlShell{Commands: map[string]*Command{"ls": LsCmd(), "cd": CdCmd()}}

	g.Expect(shell.IsCommand("ls")).To(BeTrue())
	g.Expect(shell.IsCommand("ls\t-l")).To(BeTrue())
	g.Expect(shell.IsCommand("  cd   \"my dir\"")).To(BeTrue())
	g.Expect(shell.IsCommand("select 1")).To(BeFalse())
	g.Expect(shell.IsCommand("")).To(BeFalse())
}

func TestShellIsTerminalLine(t *testing.T) {
	g := NewGomegaWithT(t)
	shell := &CsqlShell{}

	g.Expect(shell.isTerminalLine("select 1;")).To(BeTrue())
	g.Expect(shell.isTerminalLine("select 1;  \t")).To(BeTrue())
	g.Expect(shell.isTerminalLine("select 1")).To(BeFalse())
	g.Expect(shell.isTerminalLine("select ';'")).To(BeFalse())
	g.Expect(shell.isTerminalLine("")).To(BeFalse())
}

func testShell() *CsqlShell {
	return &CsqlShell{C: ansie.NewAnsiFor(os.Stdout)}
}

func TestCdCommandChecksArguments(t *testing.T) {
	g := NewGomegaWithT(t)
	dir, err := os.Getwd()
	g.Expect(err).ToNot(HaveOccurred())
	t.Cleanup(func() { _ = os.Chdir(dir) })
	target := t.TempDir()

	g.Expect(func() { CdCmd().Func(testShell(), nil) }).ToNot(Panic())
	g.Expect(func() { CdCmd().Func(testShell(), []string{target, "extra"}) }).ToNot(Panic())
	now, _ := os.Getwd()
	g.Expect(now).To(Equal(dir), "cd with a wrong number of arguments must not change the directory")

	CdCmd().Func(testShell(), []string{target})
	now, _ = os.Getwd()
	resolvedTarget, _ := filepath.EvalSymlinks(target)
	resolvedNow, _ := filepath.EvalSymlinks(now)
	g.Expect(resolvedNow).To(Equal(resolvedTarget))
}

func TestClearCommandWithoutTerminal(t *testing.T) {
	g := NewGomegaWithT(t)

	// tests do not run in a terminal
	g.Expect(func() { ClearCmd().Func(testShell(), nil) }).ToNot(Panic())
}

func TestHistoryFileName(t *testing.T) {
	g := NewGomegaWithT(t)
	shell := testShell()

	t.Setenv("HOME", "/home/someone")
	g.Expect(shell.historyFileName()).To(Equal("/home/someone/.csql_history"))

	t.Setenv("HOME", "")
	g.Expect(shell.historyFileName()).To(BeEmpty())
	g.Expect(func() { shell.loadHistory() }).ToNot(Panic())
	g.Expect(func() { shell.saveHistory() }).ToNot(Panic())
}

func TestCsvCommandValidatesIgnoreMalformedRows(t *testing.T) {
	g := NewGomegaWithT(t)
	saved := core.IgnoreMalformedCsvRows
	t.Cleanup(func() { core.IgnoreMalformedCsvRows = saved })
	core.IgnoreMalformedCsvRows = true

	CsvCommand().Func(testShell(), []string{"--ignore-malformed-rows=maybe"})
	g.Expect(core.IgnoreMalformedCsvRows).To(BeTrue(), "an invalid value must not change the setting")

	CsvCommand().Func(testShell(), []string{"--ignore-malformed-rows=false"})
	g.Expect(core.IgnoreMalformedCsvRows).To(BeFalse())

	CsvCommand().Func(testShell(), []string{"--ignore-malformed-rows=TRUE"})
	g.Expect(core.IgnoreMalformedCsvRows).To(BeTrue())
}
