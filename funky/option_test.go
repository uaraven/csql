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

package funky

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSome(t *testing.T) {
	g := NewGomegaWithT(t)
	expected := 12
	v := SomeOf(expected)
	g.Expect(v.IsEmpty()).To(BeFalse(), "expected IsEmpty() to return false")
	g.Expect(v.IsPresent()).To(BeTrue(), "expected IsPresent() to return true")
	g.Expect(v.Value()).To(Equal(expected))
	g.Expect(v.OrGet(10)).To(Equal(expected))
	var executed bool = false
	v.IfPresent(func(value int) {
		g.Expect(value).To(Equal(expected))
		executed = true
	})
	g.Expect(executed).To(BeTrue())
}

func TestNone(t *testing.T) {
	g := NewGomegaWithT(t)
	v := NoneOf[int]()
	g.Expect(v.IsEmpty()).To(BeTrue(), "expected IsEmpty() to return true")
	g.Expect(v.IsPresent()).To(BeFalse(), "expected IsPresent() to return false")
	g.Expect(func() { v.Value() }).To(Panic())

	expected := 10
	g.Expect(v.OrGet(expected)).To(Equal(expected))
	v.IfPresent(func(value int) {
		t.Fatal("Should not execute")
	})
}
