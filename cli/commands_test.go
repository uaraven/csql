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
	. "github.com/onsi/gomega"
	"testing"
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
