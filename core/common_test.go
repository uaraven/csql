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
	"github.com/uaraven/csql/funky"
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
