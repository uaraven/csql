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
	"github.com/peterh/liner"
	"github.com/uaraven/ansie"
	"github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/sql"
	"golang.org/x/term"
	"io"
	"log"
	"os"
	"strings"
)

const (
	prompt   = "csql> "
	mlPrompt = "  >>> "
)

type CsqlShell struct {
	C          *ansie.AnsiBuffer
	line       *liner.State
	Commands   map[string]*Command
	terminated bool
}

func NewCsqlShell(version string) *CsqlShell {
	csql := &CsqlShell{
		C:    ansie.NewAnsiFor(os.Stdout),
		line: liner.NewLiner(),
	}

	csql.Commands = map[string]*Command{
		"help":  HelpCmd(),
		"clear": ClearCmd(),
		"exit":  ExitCmd(),
		"ls":    LsCmd(),
		"pwd":   PwdCmd(),
		"cd":    CdCmd(),
		"csv":   CsvCommand(),
	}

	csql.line.SetCtrlCAborts(true)
	csql.loadHistory()
	fmt.Println(csql.C.A("csql version ").Attr(ansie.Bold).A(version).Reset().String())

	return csql
}

func (s *CsqlShell) TerminalSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return w, h
}

func (s *CsqlShell) PrintMessage(text string) {
	fmt.Println(text)
}

func (s *CsqlShell) PrintError(text string) {
	fmt.Println(s.C.Fg(ansie.Red).A(text).Reset().String())
}

func (s *CsqlShell) historyFileName() string {
	home := os.Getenv("HOME")
	return home + "/.csql_history"
}

func (s *CsqlShell) loadHistory() {
	if f, err := os.Open(s.historyFileName()); err == nil {
		_, _ = s.line.ReadHistory(f)
		_ = f.Close()
	}
}

func (s *CsqlShell) saveHistory() {
	if f, err := os.Create(s.historyFileName()); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		_, _ = s.line.WriteHistory(f)
		_ = f.Close()
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
		if err == io.EOF {
			s.PrintMessage("Bye!")
			s.Terminate()
			break
		}
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
		} else if len(input) > 0 {
			buffer.WriteString(input)
			buffer.WriteRune('\n')
			s.line.AppendHistory(input)
		}
		if s.isTerminalLine(input) {
			query := buffer.String()
			buffer.Reset()
			s.ExecuteQuery(query)
		}
	}
}

func (s *CsqlShell) Terminate() {
	s.saveHistory()
	_ = s.line.Close()
	s.terminated = true
}

// isTerminalLine checks whether the input string terminates up with ';', meaning it is the last line in multiline input
func (s *CsqlShell) isTerminalLine(input string) bool {
	return len(input) > 0 && RuneIndex(input, ';') == len(input)-1
}

func (s *CsqlShell) IsCommand(input string) bool {
	inputData := strings.Split(input, " ")
	_, ok := s.Commands[inputData[0]]
	return ok
}

func (s *CsqlShell) ExecuteQuery(query string) {
	defer func() {
		if err := recover(); err != nil {
			switch sqle := err.(type) {
			case *errors.CsqlError:
				s.PrintError(s.C.S("[%d:%d] ", sqle.Location.Line, sqle.Location.Column+1).
					Attr(ansie.Bold).A(sqle.Message).Reset().String())
			default:
				s.PrintError(s.C.Attr(ansie.Bold).S("%v", err).Reset().String())
			}
		}
	}()
	ds := sql.ExecuteSql(query)
	table := InitTable(ds, -1)
	s.PrintMessage(table.PrintHeader())
	s.PrintMessage(table.PrintData(ds.GetRows()))
}

func (s *CsqlShell) ExecuteCommand(input string) {
	inputData := strings.Split(input, " ")
	cmd, ok := s.Commands[inputData[0]]
	if !ok {
		return
	}
	cmd.Func(s, inputData[1:])
}
