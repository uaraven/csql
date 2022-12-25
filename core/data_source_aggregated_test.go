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

func TestGroupByCity(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{NewColumn("*")}
	groupBy := []GroupByColumn{
		{name: "city"},
	}

	aggDs := NewAggregationDataSource(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(6))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"city", "country", "population"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewStringValue("Dubrovnik"), NewStringValue("Croatia"), NewIntValue(42615)),
		HavingRowValues(columns, NewStringValue("Odesa"), NewStringValue("Ukraine"), NewIntValue(1015826)),
		HavingRowValues(columns, NewStringValue("Odessa"), NewStringValue("Canada"), NewIntValue(16221)),
		HavingRowValues(columns, NewStringValue("London"), NewStringValue("UK"), NewIntValue(8799800)),
		HavingRowValues(columns, NewStringValue("Waterloo"), NewStringValue("Belgium"), NewIntValue(30174)),
		HavingRowValues(columns, NewStringValue("Warsaw"), NewStringValue("Poland"), NewIntValue(3053104)),
	))
}

func TestGroupByCityIndex(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{NewColumn("*")}
	groupBy := []GroupByColumn{
		{index: 2},
	}

	aggDs := NewAggregationDataSource(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(6))
	g.Expect(rows).To(HaveLen(6))
	columns := []string{"city", "country", "population"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewStringValue("Dubrovnik"), NewStringValue("Croatia"), NewIntValue(42615)),
		HavingRowValues(columns, NewStringValue("Odesa"), NewStringValue("Ukraine"), NewIntValue(1015826)),
		HavingRowValues(columns, NewStringValue("Odessa"), NewStringValue("Canada"), NewIntValue(16221)),
		HavingRowValues(columns, NewStringValue("London"), NewStringValue("UK"), NewIntValue(8799800)),
		HavingRowValues(columns, NewStringValue("Waterloo"), NewStringValue("Belgium"), NewIntValue(30174)),
		HavingRowValues(columns, NewStringValue("Warsaw"), NewStringValue("Poland"), NewIntValue(3053104)),
	))
}

func TestGroupByInvalidColumn(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{NewColumn("*")}
	groupBy := []GroupByColumn{
		{name: "county"},
	}
	g.Expect(func() { NewAggregationDataSource(src, projection, groupBy) }).To(Panic())
}

func TestGroupByInvalidIndex(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{NewColumn("*")}
	groupBy := []GroupByColumn{
		{index: 12},
	}
	g.Expect(func() { NewAggregationDataSource(src, projection, groupBy) }).To(Panic())
}

func TestGroupByCityCountry(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{NewColumn("*")}
	groupBy := []GroupByColumn{
		{name: "city"},
		{name: "country"},
	}

	aggDs := NewAggregationDataSource(src, projection, groupBy)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(6))
	g.Expect(rows).To(HaveLen(13))
	columns := []string{"city", "country", "population"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewStringValue("Dubrovnik"), NewStringValue("Croatia"), NewIntValue(42615)),
		HavingRowValues(columns, NewStringValue("Odesa"), NewStringValue("Ukraine"), NewIntValue(1015826)),
		HavingRowValues(columns, NewStringValue("Odessa"), NewStringValue("Canada"), NewIntValue(16221)),
		HavingRowValues(columns, NewStringValue("Odessa"), NewStringValue("USA"), NewIntValue(8080)),
		HavingRowValues(columns, NewStringValue("London"), NewStringValue("USA"), NewIntValue(936)),
		HavingRowValues(columns, NewStringValue("London"), NewStringValue("UK"), NewIntValue(8799800)),
		HavingRowValues(columns, NewStringValue("London"), NewStringValue("Canada"), NewIntValue(422324)),
	))
}

func TestAggregateDataSource_Having(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		NewExpressionColumn(NewMinFunction(NewRowValue("area"), nil), ""),
		NewColumn("country"),
	}
	groupBy := []GroupByColumn{
		{name: "country"},
	}
	having := NewGt(NewSumFunction(NewRowValue("population"), nil), NewIntValue(10000))

	aggDs := NewAggregationDataSourceHaving(src, projection, groupBy, having)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(8))
	columns := []string{"MIN(area)", "country"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewFloatValue(1.1), NewStringValue("Australia")),
		HavingRowValues(columns, NewFloatValue(21.32), NewStringValue("Belgium")),
		HavingRowValues(columns, NewFloatValue(64.06), NewStringValue("Canada")),
		HavingRowValues(columns, NewFloatValue(21.35), NewStringValue("Croatia")),
		HavingRowValues(columns, NewIntValue(6100), NewStringValue("Poland")),
		HavingRowValues(columns, NewFloatValue(1572.03), NewStringValue("UK")),
		HavingRowValues(columns, NewFloatValue(7.11), NewStringValue("USA")),
		HavingRowValues(columns, NewFloatValue(162.42), NewStringValue("Ukraine")),
	))
}

func TestAggregateDataSource_HavingWithoutAggregate(t *testing.T) {
	g := NewGomegaWithT(t)

	src := loadTestMemDatasource("cities")

	projection := []ProjectionColumn{
		NewExpressionColumn(NewMinFunction(NewRowValue("area"), nil), ""),
		NewColumn("country"),
	}
	groupBy := []GroupByColumn{
		{name: "country"},
	}
	having := NewGt(NewRowValue("population"), NewIntValue(10000))

	aggDs := NewAggregationDataSourceHaving(src, projection, groupBy, having)

	rows := aggDs.GetRows()

	g.Expect(aggDs.Header().ColumnsMetadata()).To(HaveLen(2))
	g.Expect(rows).To(HaveLen(7))
	columns := []string{"MIN(area)", "country"}
	g.Expect(rows).To(ContainElements(
		HavingRowValues(columns, NewFloatValue(21.32), NewStringValue("Belgium")),
		HavingRowValues(columns, NewFloatValue(64.06), NewStringValue("Canada")),
		HavingRowValues(columns, NewFloatValue(21.35), NewStringValue("Croatia")),
		HavingRowValues(columns, NewIntValue(6100), NewStringValue("Poland")),
		HavingRowValues(columns, NewFloatValue(1572.03), NewStringValue("UK")),
		HavingRowValues(columns, NewFloatValue(162.42), NewStringValue("Ukraine")),
		HavingRowValues(columns, NewFloatValue(1.1), NewStringValue("Australia")),
	))
}
