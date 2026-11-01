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
	"errors"
	"fmt"
	"github.com/peterh/liner"
	"github.com/uaraven/ansie"
	"golang.org/x/term"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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
		"help":    HelpCmd(),
		"clear":   ClearCmd(),
		"exit":    ExitCmd(),
		"ls":      LsCmd(),
		"pwd":     PwdCmd(),
		"cd":      CdCmd(),
		"csv":     CsvCommand(),
		"inspect": InspectCommand(),
		"status":  StatusCommand(),
		"purge":   PurgeCommand(),
	}

	csql.line.SetCtrlCAborts(true)
	csql.loadHistory()
	fmt.Println(csql.C.A("csql version ").Attr(ansie.Bold).A(version).Reset().String())

	return csql
}

// TerminalSize returns the size of the terminal, or -1, -1 if the output is not a terminal
func (s *CsqlShell) TerminalSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return -1, -1
	}
	return w, h
}

func (s *CsqlShell) PrintMessage(text string) {
	fmt.Println(text)
}

func (s *CsqlShell) PrintError(text string) {
	fmt.Println(s.C.Fg(ansie.Red).A(text).Reset().String())
}

// historyFileName returns the path of the history file. It is empty if the home directory is not known
func (s *CsqlShell) historyFileName() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".csql_history")
}

func (s *CsqlShell) loadHistory() {
	name := s.historyFileName()
	if name == "" {
		return
	}
	if f, err := os.Open(name); err == nil {
		_, _ = s.line.ReadHistory(f)
		_ = f.Close()
	}
}

func (s *CsqlShell) saveHistory() {
	name := s.historyFileName()
	if name == "" {
		return
	}
	// history contains queries, keep it private
	if f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		_ = f.Chmod(0600)
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
		if errors.Is(err, liner.ErrPromptAborted) {
			if inMultiline {
				buffer.Reset()
				continue
			} else {
				continue
			}
		} else if err != nil {
			panic(err)
		} else if inMultiline {
			// shell commands are not recognised here, because a query line can start with a command name
			// (for example a column called "status"). Ctrl-C abandons the query.
			buffer.WriteString(input)
			buffer.WriteRune('\n')
			if len(input) > 0 {
				s.line.AppendHistory(input)
			}
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
	trimmed := strings.TrimRight(input, " \t")
	return len(trimmed) > 0 && RuneIndex(trimmed, ';') == len([]rune(trimmed))-1
}

func (s *CsqlShell) IsCommand(input string) bool {
	inputData := splitArgs(input)
	if len(inputData) == 0 {
		return false
	}
	_, ok := s.Commands[inputData[0]]
	return ok
}

// splitArgs splits the input on whitespace. Single or double quotes group words with whitespace into one
// argument and are removed. A quote that is never closed runs to the end of the input.
func splitArgs(input string) []string {
	args := make([]string, 0)
	var current strings.Builder
	inArg := false
	var quote rune
	for _, r := range input {
		switch {
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				current.WriteRune(r)
			}
		case r == '"' || r == '\'':
			quote = r
			inArg = true
		case r == ' ' || r == '\t':
			if inArg {
				args = append(args, current.String())
				current.Reset()
				inArg = false
			}
		default:
			current.WriteRune(r)
			inArg = true
		}
	}
	if inArg {
		args = append(args, current.String())
	}
	return args
}

func (s *CsqlShell) ExecuteQuery(query string) {
	start := time.Now()
	ds := RunQuery(query)
	completed := time.Since(start)
	if ds == nil {
		return
	}
	table := InitTable(ds, -1)
	s.PrintMessage(table.PrintHeader())
	s.PrintMessage(table.PrintData(ds.GetRows()))
	s.PrintMessage(fmt.Sprintf("Rows: %d, elapsed time: %5.3f sec", len(ds.GetRows()), completed.Seconds()))
}

func (s *CsqlShell) ExecuteCommand(input string) {
	inputData := splitArgs(input)
	if len(inputData) == 0 {
		return
	}
	cmd, ok := s.Commands[inputData[0]]
	if !ok {
		return
	}
	cmd.Func(s, inputData[1:])
}
