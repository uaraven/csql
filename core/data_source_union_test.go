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
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewUnionAll(t *testing.T) {
	g := NewGomegaWithT(t)

	ds1 := loadTestMemDatasource("employees")
	ds2 := loadTestMemDatasource("people")

	unionAll := NewUnionAll(ds1, ds2)

	rows := ReadAllRows(unionAll)
	g.Expect(rows).To(HaveLen(8))
}

func TestNewUnion(t *testing.T) {
	g := NewGomegaWithT(t)

	ds1 := NewProjectionDataSource(loadTestMemDatasource("employees"),
		NewSimpleProjection([]string{"first_name", "last_name"}, []string{"", ""}), false)
	ds2 := NewProjectionDataSource(loadTestMemDatasource("people"),
		NewSimpleProjection([]string{"first_name", "last_name"}, []string{"", ""}), false)

	union := NewUnion(ds1, ds2)

	rows := ReadAllRows(union)
	g.Expect(rows).To(HaveLen(4))
}

func TestNewUnionDifferentProjections(t *testing.T) {
	g := NewGomegaWithT(t)

	ds1 := NewProjectionDataSource(loadTestMemDatasource("employees"),
		NewSimpleProjection([]string{"first_name", "last_name"}, []string{"", ""}), false)
	ds2 := NewProjectionDataSource(loadTestMemDatasource("people"),
		NewSimpleProjection([]string{"first_name", "last_name", "width"}, []string{"", "", ""}), false)

	g.Expect(func() { NewUnion(ds1, ds2) }).To(Panic())
}
