/*
 *    Copyright 2025 Oleksiy Voronin
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
 *    SPDX-FileCopyrightText: (c) 2025 Oleksiy Voronin <ovoronin@gmail.com>
 */

package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/uaraven/csql/funky"
)

func WriteDataSourceToFile(ds DataSource, outputFile string) error {
	var errBuffer strings.Builder
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	csvWriter := csv.NewWriter(file)
	err = csvWriter.Write(funky.Map(ds.Header().ColumnsMetadata(), func(c ColumnMetadata) string {
		return c.Name()
	}))
	if err != nil {
		return err
	}
	for _, row := range ds.GetRows() {
		rowValues := funky.Map(row.Values(), func(v Value) string {
			return v.AsString().Value().(string)
		})
		err = csvWriter.Write(rowValues)
		if err != nil {
			fmt.Fprintf(&errBuffer, "Failed to write row %d: %s\n", row.Id(), err)
		}
	}
	csvWriter.Flush()
	if errBuffer.Len() > 0 {
		return fmt.Errorf("errors occurred while writing CSV: %v", errBuffer.String())
	}
	return nil
}

func NewStatusDataSource(columnNames []string, values []string) DataSource {
	columns := funky.MapWithIndex(columnNames, func(index int, name string) ColumnMetadata {
		return NewSimpleColumnMetadata(name, index)
	})
	headers := NewHeadersFromOtherHeaders(columns)
	data := funky.Map(values, func(value string) Value {
		return ParseStringToValue(value)
	})
	row := newRowWithId(0, headers, data)
	return NewMemDataSource("DROP_TEMP_TABLE", columns, []Row{row})
}
