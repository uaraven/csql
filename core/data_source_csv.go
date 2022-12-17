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
	"io"
	"os"
	"path/filepath"
	"strings"
)

const csvExt = ".csv"

func NewCsvDataSource(csvFile string) DataSource {
	return NewCsvDataSourceWithAlias(csvFile, "")
}

func NewCsvDataSourceWithAlias(csvFile string, alias string) DataSource {
	if filepath.Ext(csvFile) == "" {
		_, err := os.Stat(csvFile)
		if err != nil && errors.Is(err, os.ErrNotExist) {
			_, err = os.Stat(csvFile + csvExt)
			if err == nil {
				csvFile = csvFile + csvExt
			}
		}
	}
	_, nameExt := filepath.Split(csvFile)
	var justName string
	if alias != "" {
		justName = alias
	} else {
		justName = strings.TrimSuffix(nameExt, filepath.Ext(nameExt))
	}
	file, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.TrimLeadingSpace = true
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
	return cds
}

func loadCsv(csvReader *csv.Reader, header DataSourceHeader) ([]Row, error) {
	rows := make([]Row, 0)
	index := 0
	for {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			return rows, nil
		} else if err != nil {
			return nil, err
		}
		r := parseRowWithId(index, header, csvRow)
		index++
		rows = append(rows, r)
	}
}
