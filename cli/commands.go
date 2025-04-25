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
	"github.com/jessevdk/go-flags"
	"github.com/uaraven/ansie"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"github.com/uaraven/csql/util"
	"os"
	"strings"
	"time"
)

type CommandFunc func(shell *CsqlShell, parameters []string)

type Command struct {
	Name string
	Help string
	Func CommandFunc
}

func HelpCmd() *Command {
	return &Command{
		Name: "help",
		Help: "Show the help message",
		Func: func(shell *CsqlShell, params []string) {
			for _, cmd := range shell.Commands {
				shell.PrintMessage(shell.C.Attr(ansie.Bold).S("%-10s", cmd.Name).Reset().A(cmd.Help).String())
			}
		},
	}
}

func ClearCmd() *Command {
	return &Command{
		Name: "clear",
		Help: "Clear the screen",
		Func: func(shell *CsqlShell, params []string) {
			_, h := shell.TerminalSize()
			if h < 0 {
				h = 25
			}
			for i := 0; i < h; i += 1 {
				shell.PrintMessage("")
			}
		},
	}
}

func ExitCmd() *Command {
	return &Command{
		Name: "exit",
		Help: "Close tht csql shell",
		Func: func(shell *CsqlShell, _ []string) {
			shell.Terminate()
		},
	}
}

func PwdCmd() *Command {
	return &Command{
		Name: "pwd",
		Help: "Shows current directory",
		Func: func(s *CsqlShell, args []string) {
			path, err := os.Getwd()
			if err != nil {
				s.PrintError(fmt.Sprint(err))
			}
			s.PrintMessage(path)
		},
	}
}

func CdCmd() *Command {
	return &Command{
		Name: "cd",
		Help: "Change current directory",
		Func: func(s *CsqlShell, args []string) {
			if len(args) != 1 {
				s.PrintError(".cd accepts only one parameter - directory name")
			}
			err := os.Chdir(args[0])
			if err != nil {
				s.PrintError(fmt.Sprint(err))
			}
		},
	}
}

func LsCmd() *Command {
	return &Command{
		Name: "ls",
		Help: "List files in the current directory",
		Func: func(s *CsqlShell, _ []string) {
			files, err := os.ReadDir(".")
			if err != nil {
				s.PrintError(fmt.Sprint(err))
			}
			for _, file := range files {
				if strings.HasPrefix(file.Name(), ".") {
					continue
				}
				var ftype string
				if file.IsDir() {
					ftype = s.C.Attr(ansie.Underline).A(file.Name() + "/").Reset().String()
				} else {
					ftype = file.Name()
				}
				s.PrintMessage(ftype)
			}
		},
	}
}

func CsvCommand() *Command {
	return &Command{
		Name: "csv",
		Help: "Set or show csv file parameters. Type 'csv help' to view available options",
		Func: func(s *CsqlShell, params []string) {
			status := func() {
				s.PrintMessage(s.C.A("Null String=").Attr(ansie.Bold).A(core.NullValueString).Reset().String())
				s.PrintMessage(s.C.A("Separator=").Attr(ansie.Bold).S("%q", string(core.CSVSeparator)).Reset().String())
			}
			if params == nil || len(params) == 0 {
				status()
			} else {
				var opts struct {
					Null      *string `long:"null"`
					Separator *string `long:"separator"`
				}
				args, err := flags.ParseArgs(&opts, params)
				if err != nil {
					s.PrintError(fmt.Sprintf("Invalid parameters: %v", params))
					return
				}
				if funky.AnyMatches(args, func(s string) bool {
					return s == "help"
				}) {
					s.PrintMessage("csv --null=\"<null string>\", i.e. csv --null=\"\" to treat empty strings as null")
					s.PrintMessage("csv --separator=\"<csv separator>\", i.e. csv --separator=\\t to use tab as a csv field separator")
					return
				}
				if opts.Null != nil {
					core.NullValueString = util.Unquote(*opts.Null)

				}
				if opts.Separator != nil {
					err = core.SetCsvSeparator(*opts.Separator)
					if err != nil {
						s.PrintError(fmt.Sprintf("Failed to set separator %v", err))
					}
				}

				status()
			}
		},
	}
}

func InspectCommand() *Command {
	return &Command{
		Name: "inspect",
		Help: "Inspect a CSV file. Returns columns and their types",
		Func: func(shell *CsqlShell, parameters []string) {
			var opts struct {
			}
			args, err := flags.ParseArgs(&opts, parameters)
			if err != nil {
				shell.PrintError("Expected file name as a parameter to 'inspect'")
				return
			}
			for _, fName := range args {
				fileName := util.Unquote(fName)
				csvFile := core.ExpandCsvFileName(fileName)
				csvStruct, err := InspectCsv(csvFile)
				if err != nil {
					shell.PrintError(fmt.Sprintf("%v", err))
					return
				}
				width := funky.Max(funky.Map(csvStruct, func(t CsvColumn) int {
					return len(t.Name)
				}))
				shell.PrintMessage(fileName)
				shell.PrintMessage(strings.Repeat("-", len(fileName)))
				for _, column := range csvStruct {
					shell.PrintMessage(fmt.Sprintf("%*s: %s", width, column.Name, column.Type))
				}
				shell.PrintMessage("")
			}
		},
	}
}

func StatusCommand() *Command {
	return &Command{
		Name: "status",
		Help: "Show current CSQL status - number of datasets and rows in memory",
		Func: func(shell *CsqlShell, parameters []string) {
			titleWidth := 0
			for k := range core.TableCache {
				if len(k) > titleWidth {
					titleWidth = len(k)
				}
			}
			shell.PrintMessage(fmt.Sprintf("%*s %10s Age", titleWidth, "Dataset", "Size"))
			for k, v := range core.TableCache {
				diff := time.Now().Sub(v.AccessTime)
				shell.PrintMessage(shell.C.Fg(ansie.Blue).S("%*s ", titleWidth, k).
					Fg(ansie.Green).S("%15d", len(v.DataSource.GetRows())).Reset().A(" rows ").
					Fg(ansie.Yellow).S("%s", util.FormatDuration(diff)).
					String())
			}
		},
	}
}

func PurgeCommand() *Command {
	return &Command{
		Name: "purge",
		Help: "Clears CSQL dataset cache - removes all datasets from memory",
		Func: func(shell *CsqlShell, parameters []string) {
			datasets, rows := core.TableCache.ClearCache()
			shell.PrintMessage(shell.C.A("Removed ").Fg(ansie.Blue).S("%d", datasets).Reset().A(" datasets and ").Fg(ansie.Blue).S("%d", rows).Reset().A(" rows from memory").String())
		},
	}
}

func RuneIndex(s string, r rune) int {
	runes := []rune(s)
	index := 0
	for index < len(runes) {
		if runes[index] == '"' || runes[index] == '\'' {
			closing := runes[index]
			index++
			for index < len(runes) && runes[index] != closing {
				if runes[index] == '\\' {
					index += 2 // skip escaped character
				} else {
					index++
				}
			}
		} else if runes[index] == r {
			return index
		}
		index++
	}
	return -1
}
