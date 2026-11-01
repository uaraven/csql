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

package core

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/uaraven/csql/util"
)

var CSVSeparator = ','
var IgnoreMalformedCsvRows = false

// SetCsvSeparator sets the CSV separator used for reading and writing files. The separator can be quoted and can
// contain Go escape sequences, like \t. It must be a single character that is not a quote or a line break.
func SetCsvSeparator(separator string) error {
	sep := util.Unquote(separator)
	if unescaped, err := strconv.Unquote("\"" + sep + "\""); err == nil {
		sep = unescaped
	}
	if utf8.RuneCountInString(sep) != 1 {
		return fmt.Errorf("separator must be a single character, got %q", separator)
	}
	r, _ := utf8.DecodeRuneInString(sep)
	if r == 0 || r == '"' || r == '\r' || r == '\n' || r == utf8.RuneError {
		return fmt.Errorf("invalid separator %q: it cannot be a quote, a line break or an invalid character", separator)
	}
	CSVSeparator = r
	return nil
}

func NewCsvDataSource(csvFile string) DataSource {
	return NewCsvDataSourceWithAlias(csvFile, "")
}

// NewCsvDataSourceWithAlias returns a datasource for a temp table or a CSV file. A temp table is used when the
// source is a plain name (no directory, no extension) of an existing temp table. The columns of the datasource are
// qualified with the alias, or with the table/file name if there is no alias.
//
// Every call returns a datasource with its own cursor, even if the data comes from the cache.
func NewCsvDataSourceWithAlias(csvFile string, alias string) DataSource {
	if tempTable, ok := TableCache[csvFile]; ok && tempTable.TempTable {
		TableCache.Touch(csvFile)
		name := alias
		if name == "" {
			name = csvFile
		}
		return newDataSourceView(tempTable.DataSource, name)
	}

	csvFile = ExpandCsvFileName(csvFile)
	_, nameExt := filepath.Split(csvFile)
	justName := strings.TrimSuffix(nameExt, filepath.Ext(nameExt))
	if alias != "" {
		justName = alias
	}
	cacheKey := csvFile
	if abs, err := filepath.Abs(csvFile); err == nil {
		cacheKey = abs
	}

	stat, err := os.Stat(csvFile)
	if cached, ok := TableCache[cacheKey]; ok && !cached.TempTable {
		if err == nil && stat.ModTime().Equal(cached.ModTime) {
			TableCache.Touch(cacheKey)
			return newDataSourceView(cached.DataSource, justName)
		}
		// the file is gone or was changed, delete the cache entry
		delete(TableCache, cacheKey)
	}

	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
	csvReader.Comma = CSVSeparator
	csvReader.ReuseRecord = true
	headers, err := csvReader.Read()
	if err != nil {
		panic(err)
	}
	cds := &memDataSource{
		name:  justName,
		index: -1,
	}
	cds.headers = NewHeadersFromSlice(cds, headers)
	cds.data, err = loadCsv(csvReader, cds.Header())
	if err != nil {
		panic(err)
	}
	TableCache.AddToCache(cacheKey, cds, stat.ModTime(), false)
	if _, cached := TableCache[cacheKey]; cached {
		// the cached datasource must not be handed out, its cursor would be shared
		return newDataSourceView(cds, justName)
	}
	return cds
}

// newDataSourceView creates a datasource with the rows of src and the columns qualified with the name.
// The values are shared with src, but the cursor and the headers are not.
func newDataSourceView(src DataSource, name string) DataSource {
	columns := make([]ColumnMetadata, src.Header().ColumnCount())
	for idx, column := range src.Header().ColumnsMetadata() {
		columns[idx] = &columnMetadata{parentName: name, name: column.Name(), index: column.Index()}
	}
	header := &dataSourceHeader{columns: columns}
	rows := make([]Row, len(src.GetRows()))
	for idx, row := range src.GetRows() {
		rows[idx] = newRowWithId(row.Id(), header, row.Values())
	}
	return &memDataSource{name: name, index: -1, data: rows, headers: header}
}

func loadCsv(csvReader *csv.Reader, header DataSourceHeader) ([]Row, error) {
	rows := make([]Row, 0)
	index := 0
	for {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			return rows, nil
		} else if err != nil {
			if errors.Is(err, csv.ErrFieldCount) && IgnoreMalformedCsvRows {
				index++
				continue
			} else {
				return nil, err
			}
		}
		r := parseRowWithId(index, header, csvRow)
		index++
		rows = append(rows, r)
	}
}
