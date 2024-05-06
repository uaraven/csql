/*
 *    Copyright 2022-2023 Oleksiy Voronin
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
	"encoding/csv"
	"fmt"
	. "github.com/uaraven/ansie"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/funky"
	"github.com/uaraven/csql/sql"
	"os"
	"time"
)

func RunQuery(query string) (result core.DataSource) {
	defer func() {
		if err := recover(); err != nil {
			switch sqle := err.(type) {
			case *errors.CsqlError:
				if sqle.Location == nil {
					fmt.Println(Ansi.A("[unknown location] ").Attr(Bold).A(sqle.Message).Reset().String())
				} else {
					fmt.Println(Ansi.S("[%d:%d] ", sqle.Location.Line, sqle.Location.Column+1).Attr(Bold).A(sqle.Message).Reset().String())
				}
			default:
				fmt.Println(Ansi.Attr(Bold).S("%v", err).Reset().String())
			}
			result = nil
		}
	}()
	result = sql.ExecuteSql(query)
	return
}

func ExecuteQuery(query string) {
	start := time.Now()
	ds := RunQuery(query)
	completed := time.Now().Sub(start)
	if ds == nil {
		return
	}
	table := InitTable(ds, -1)
	fmt.Println(table.PrintHeader())
	fmt.Println(table.PrintData(ds.GetRows()))
	fmt.Println(fmt.Sprintf("Rows: %d, elapsed time: %5.3f sec", len(ds.GetRows()), completed.Seconds()))
}

func ExecuteToCsv(query string, outputFile string) {
	start := time.Now()
	ds := RunQuery(query)
	if ds == nil {
		return
	}
	completed := time.Now().Sub(start)
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write(funky.Map(ds.Header().ColumnsMetadata(), func(c core.ColumnMetadata) string {
		return c.Name()
	}))
	for _, row := range ds.GetRows() {
		rowValues := funky.Map(row.Values(), func(v core.Value) string {
			return v.AsString().Value().(string)
		})
		csvWriter.Write(rowValues)
	}
	fmt.Println(fmt.Sprintf("Rows: %d, elapsed time: %5.3f sec", len(ds.GetRows()), completed.Seconds()))
}
