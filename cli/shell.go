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

import "C"
import (
	"fmt"
	"github.com/peterh/liner"
	"github.com/uaraven/ansi"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"github.com/uaraven/csql/sql"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	prompt   = "csql> "
	mlPrompt = "  >>> "
)

type CsqlShell struct {
	C          *ansi.AnsiPrinter
	line       *liner.State
	commands   map[string]func([]string)
	terminated bool
}

func NewCsqlShell(version string) *CsqlShell {
	csql := &CsqlShell{
		C:    ansi.NewAnsiFor(os.Stdout),
		line: liner.NewLiner(),
	}

	csql.commands = map[string]func([]string){
		"help":      csql.Help,
		"clear":     csql.Clear,
		"exit":      csql.Exit,
		"ls":        csql.ListFiles,
		"pwd":       csql.Pwd,
		"cd":        csql.Cd,
		"nullValue": csql.SetNullString,
	}

	csql.line.SetCtrlCAborts(true)
	csql.loadHistory()
	fmt.Println(csql.C.A("csql version ").Attr(ansi.Bold).A(version).Reset().CR().String())

	return csql
}

func (s *CsqlShell) PrintMessage(text string) {
	fmt.Println(text)
}

func (s *CsqlShell) PrintError(text string) {
	fmt.Println(s.C.Fg(ansi.SysRed).A(text).Reset().String())
}

func (s *CsqlShell) SetNullString(args []string) {
	if args == nil || len(args) == 0 {
		s.PrintMessage(s.C.A("Null String=").Attr(ansi.Bold).A(core.NullValueString).Reset().String())
	} else {
		core.NullValueString = args[0]
	}
}

func (s *CsqlShell) Pwd(args []string) {
	path, err := os.Getwd()
	if err != nil {
		s.PrintError(fmt.Sprint(err))
	}
	s.PrintMessage(path)
}

func (s *CsqlShell) Cd(args []string) {
	if len(args) != 1 {
		s.PrintError(".cd accepts only one parameter - directory name")
	}
	err := os.Chdir(args[0])
	if err != nil {
		s.PrintError(fmt.Sprint(err))
	}
}

func (s *CsqlShell) ListFiles(_ []string) {
	files, err := os.ReadDir(".")
	if err != nil {
		s.PrintError(fmt.Sprint(err))
	}
	maxWidth := funky.Max(funky.Map(files, func(f os.DirEntry) int {
		return len(f.Name())
	}))
	if maxWidth < 50 {
		maxWidth = 50
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		var ftype string
		if file.IsDir() {
			ftype = s.C.Attr(ansi.Underline).A(file.Name() + "/").Reset().String()
		} else {
			ftype = file.Name()
		}
		s.PrintMessage(ftype)
	}
}

func (s *CsqlShell) Help(_ []string) {

}

func (s *CsqlShell) Clear(_ []string) {
	for i := 0; i < 100; i += 1 {
		fmt.Println()
	}
}

func (s *CsqlShell) Exit(_ []string) {
	s.Terminate()
}

func (s *CsqlShell) historyFileName() string {
	home := os.Getenv("HOME")
	return home + "/.csql_history"
}

func (s *CsqlShell) loadHistory() {
	if f, err := os.Open(s.historyFileName()); err == nil {
		s.line.ReadHistory(f)
		f.Close()
	}
}

func (s *CsqlShell) saveHistory() {
	if f, err := os.Create(s.historyFileName()); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		s.line.WriteHistory(f)
		f.Close()
	}
}

func (s *CsqlShell) Start() {
	defer s.Terminate()
	var buffer strings.Builder
	for !s.terminated {
		inMultiline := buffer.Len() > 0
		var pr string
		if inMultiline {
			pr = mlPrompt
		} else {
			pr = prompt
		}
		input, err := s.line.Prompt(pr)
		if err == liner.ErrPromptAborted {
			if inMultiline {
				buffer.Reset()
				continue
			} else {
				continue
			}
		} else if err != nil {
			panic(err)
		} else if inMultiline {
			buffer.WriteString(input)
		} else if s.IsCommand(input) {
			s.ExecuteCommand(input)
			s.line.AppendHistory(input)
			continue
		} else {
			buffer.WriteString(input)
			buffer.WriteRune('\n')
		}
		if s.isTerminalLine(input) {
			query := buffer.String()
			buffer.Reset()
			s.line.AppendHistory(query)
			s.ExecuteQuery(query)
		}
	}
}

func (s *CsqlShell) Terminate() {
	s.saveHistory()
	s.line.Close()
	s.terminated = true
}

func (s *CsqlShell) isTerminalLine(input string) bool {
	r := regexp.MustCompile("\".*\"")
	input = r.ReplaceAllString(input, "")
	r = regexp.MustCompile("'.*'")
	input = r.ReplaceAllString(input, "")
	return strings.HasSuffix(strings.TrimSpace(input), ";")
}

func (s *CsqlShell) IsCommand(input string) bool {
	inputData := strings.Split(input, " ")
	_, ok := s.commands[inputData[0]]
	return ok
}

func (s *CsqlShell) ExecuteQuery(query string) {
	ds := sql.ExecuteSql(query)
	rows := core.ReadAllRows(ds)
	for _, r := range rows {
		for _, cell := range r.Values() {
			fmt.Printf("%20v", cell.Repr())
		}
		fmt.Println()
	}
}

func (s *CsqlShell) ExecuteCommand(input string) {
	inputData := strings.Split(input, " ")
	cmd, ok := s.commands[inputData[0]]
	if !ok {
		return
	}
	cmd(inputData[1:])
}
