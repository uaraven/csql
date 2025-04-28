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

package main

import (
	"log"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/uaraven/csql/cli"
	"github.com/uaraven/csql/core"
)

var Version = "v0.4.1"

var opts struct {
	NullValue *string `short:"n" long:"null" description:"Sets the string value to be treated as NULL"`
	RunQuery  bool    `short:"c" long:"command" description:"Runs the query passed as command line arguments"`
	Separator *string `short:"s" long:"separator" description:"CSV separator character"`
	Output    *string `short:"o" long:"output" description:"Output results into a file"`
	NoCache   bool    `long:"no-cache" description:"Disable datasource caching"`
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("Invalid command line arguments: %v", err)
	}
	if opts.NullValue != nil {
		core.NullValueString = *opts.NullValue
	}
	if opts.Separator != nil {
		err = core.SetCsvSeparator(*opts.Separator)
		if err != nil {
			log.Fatalf("Failed to set CSV separator %v", err)
		}
	}
	core.CacheEnabled = !(opts.RunQuery || opts.NoCache)
	if opts.RunQuery {
		query := strings.Join(args, " ")
		if opts.Output == nil {
			cli.ExecuteQuery(query)
		} else {
			cli.ExecuteToCsv(query, *opts.Output)

		}
	} else {
		shell := cli.NewCsqlShell(Version)
		shell.Start()
	}
}
