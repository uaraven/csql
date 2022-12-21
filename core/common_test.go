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
	"fmt"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/uaraven/csql/funky"
	"strings"
)

func loadTestDatasource() DataSource {
	return NewCsvDataSource("../test-data/employees.csv")
}

func loadDefaultTestMemDatasource() DataSource {
	return NewCsvDataSource("../test-data/employees.csv")
}

func loadTestMemDatasource(name string) DataSource {
	return NewCsvDataSource(fmt.Sprintf("../test-data/%s.csv", name))
}

func loadDsWithAlias(name string, alias string) DataSource {
	return NewCsvDataSourceWithAlias(fmt.Sprintf("../test-data/%s.csv", name), alias)
}

func loadTestMemDatasourceWithAlias(alias string) DataSource {
	return NewCsvDataSourceWithAlias("../test-data/employees.csv", alias)
}

type mapContext struct {
	EvaluationContext
	id     int
	values map[string]string
}

func newMapContext() *mapContext {
	return &mapContext{
		id:     0,
		values: make(map[string]string),
	}
}

func (ec *mapContext) Get(name string) funky.Option[Value] {
	if val, present := ec.values[name]; present {
		return funky.SomeOf(decodeToValue(val))
	}
	return funky.NoneOf[Value]()
}

func (ec *mapContext) Id() int {
	return ec.id
}

func (ec *mapContext) SetId(newId int) {
	ec.id = newId
}

func (ec *mapContext) SetValue(name string, value string) {
	ec.values[name] = value
}

func (ec *mapContext) Clear() {
	ec.values = make(map[string]string)
}

type columnValuePair struct {
	column string
	value  Value
}

func (cvp columnValuePair) String() string {
	return fmt.Sprintf("%s = %s", cvp.column, cvp.value.String())
}

type RowMatcher struct {
	Expected []columnValuePair
}

func (rowMatcher RowMatcher) String() string {
	return strings.Join(funky.Map(rowMatcher.Expected, func(p columnValuePair) string {
		return p.String()
	}), ", ")
}

func HavingRowValues(columns []string, values ...Value) types.GomegaMatcher {
	if len(columns) != len(values) {
		panic(fmt.Errorf("Number of columns must be equal to number of values. \nColumns:%v\n Values:%v\n", columns, values))
	}
	pairs := make([]columnValuePair, len(columns))
	for i := range columns {
		pairs[i] = columnValuePair{
			column: columns[i],
			value:  values[i],
		}
	}
	return &RowMatcher{Expected: pairs}
}

func (rowMatcher *RowMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && rowMatcher.Expected == nil {
		return false, fmt.Errorf("Both actual and expected must not be nil.")
	}

	convertedActual, ok := actual.(Row)
	if !ok {
		return false, fmt.Errorf("%v is not a Row", actual)
	}

	for _, colVal := range rowMatcher.Expected {
		optValue := convertedActual.Get(colVal.column)
		if optValue.IsEmpty() {
			return false, fmt.Errorf("row doesn't contain column '%s'", colVal.column)
		}
		if !colVal.value.Equals(optValue.Value()) {
			return false, nil
		}
	}
	return true, nil
}

func (rowMatcher *RowMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have columns with values equal to", rowMatcher.Expected)
}

func (rowMatcher *RowMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have columns with values equal to", rowMatcher.Expected)
}
