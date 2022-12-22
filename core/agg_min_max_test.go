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

func TestMinFunc_GroupBy(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		{source: newMinFunction(NewRowValue("population"), nil)},
		NewColumn("city")}
	groupBy := []GroupColumn{{name: "city"}}

	aggDs := NewAggregationDs(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"MIN(population)", "city"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewIntValue(42615), NewStringValue("Dubrovnik")),
		HavingRowValues(columns, NewIntValue(936), NewStringValue("London")),
		HavingRowValues(columns, NewIntValue(1015826), NewStringValue("Odesa")),
		HavingRowValues(columns, NewIntValue(8080), NewStringValue("Odessa")),
		HavingRowValues(columns, NewIntValue(2209), NewStringValue("Warsaw")),
		HavingRowValues(columns, NewIntValue(670), NewStringValue("Waterloo")),
	))
}

func TestMaxFunc_GroupBy(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		{source: newMaxFunction(NewRowValue("population"), nil)},
		NewColumn("city")}
	groupBy := []GroupColumn{{name: "city"}}

	aggDs := NewAggregationDs(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"MAX(population)", "city"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewIntValue(42615), NewStringValue("Dubrovnik")),
		HavingRowValues(columns, NewIntValue(8799800), NewStringValue("London")),
		HavingRowValues(columns, NewIntValue(1015826), NewStringValue("Odesa")),
		HavingRowValues(columns, NewIntValue(16221), NewStringValue("Odessa")),
		HavingRowValues(columns, NewIntValue(3053104), NewStringValue("Warsaw")),
		HavingRowValues(columns, NewIntValue(121436), NewStringValue("Waterloo")),
	))
}
