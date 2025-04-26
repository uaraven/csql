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

package errors_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/sql"
)

func TestTypeErrorComparison(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { sql.ExecuteSql("SELECT * FROM \"../test-data/employees.csv\" WHERE 1 > '4'") }).
		To(PanicWith(SourceLocation(errors.Loc(1, 49))))
}

func TestInvalidSource(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { sql.ExecuteSql("SELECT * FROM 213") }).
		To(PanicWith(SourceLocation(errors.Loc(1, 14))))
}

func TestInvalidInList(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		sql.ExecuteSql(`
			SELECT * FROM "../test-data/employees.csv" 
			         WHERE width between 1 and "3"`)
	}).To(PanicWith(SourceLocation(errors.Loc(3, 18))))
}

func TestInvalidTypeMatch(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		sql.ExecuteSql(`
			SELECT * FROM "../test-data/employees.csv" 
			         WHERE last_name match 11`)
	}).To(PanicWith(SourceLocation(errors.Loc(3, 18))))
}

func TestInvalidTypeLike(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		sql.ExecuteSql(`
			SELECT * FROM "../test-data/employees.csv" 
			         WHERE last_name like 2*11`)
	}).To(PanicWith(SourceLocation(errors.Loc(3, 18))))
}

func TestInvalidTypeConcat(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() {
		sql.ExecuteSql(`
			SELECT * FROM "../test-data/employees.csv" 
			         WHERE last_name = '12' || width`)
	}).To(PanicWith(SourceLocation(errors.Loc(3, 18))))
}

type WithLocationMatcher struct {
	Expected *errors.SourceLocation
}

func SourceLocation(location *errors.SourceLocation) types.GomegaMatcher {
	return &WithLocationMatcher{Expected: location}
}

func (matcher *WithLocationMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && matcher.Expected == nil {
		return false, fmt.Errorf("Both actual and expected must not be nil.")
	}

	convertedActual, ok := actual.(*errors.CsqlError)
	if !ok {
		return false, fmt.Errorf("%v is not a CsqlError", actual)
	}

	return reflect.DeepEqual(*convertedActual.Location, *matcher.Expected), nil
}

func (matcher *WithLocationMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to have location equal to", matcher.Expected)
}

func (matcher *WithLocationMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not to have location equal to", matcher.Expected)
}
