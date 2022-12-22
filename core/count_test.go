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

func TestCountFunc_1(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		{source: newCountFunction(nil, false, NewRowValue("country"))},
		NewColumn("city")}
	groupBy := []GroupColumn{
		{name: "city"},
	}

	aggDs := NewAggregationDs(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"COUNT(country)", "city"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewIntValue(1), NewStringValue("Dubrovnik")),
		HavingRowValues(columns, NewIntValue(5), NewStringValue("London")),
		HavingRowValues(columns, NewIntValue(1), NewStringValue("Odesa")),
		HavingRowValues(columns, NewIntValue(2), NewStringValue("Odessa")),
		HavingRowValues(columns, NewIntValue(4), NewStringValue("Waterloo")),
		HavingRowValues(columns, NewIntValue(2), NewStringValue("Warsaw")),
	))
}

func TestCountFunc_2(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		{source: newCountFunction(nil, false, NewRowValue("country"))},
		NewColumn("country")}
	groupBy := []GroupColumn{
		{name: "city"},
	}

	aggDs := NewAggregationDs(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"COUNT(country)", "country"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewIntValue(1), NewStringValue("Croatia")),
		HavingRowValues(columns, NewIntValue(5), NewStringValue("UK")),
		HavingRowValues(columns, NewIntValue(1), NewStringValue("Ukraine")),
		HavingRowValues(columns, NewIntValue(2), NewStringValue("Canada")),
		HavingRowValues(columns, NewIntValue(4), NewStringValue("Belgium")),
		HavingRowValues(columns, NewIntValue(2), NewStringValue("Poland")),
	))
}
