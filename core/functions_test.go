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

func TestLenFunction(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "abc")
	ctx.SetValue("i", "123")

	l := NewLenFunction(NewStringValue("12345"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))
	l = NewLenFunction(NewRowValue("s"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(3)))

	l = NewLenFunction(NewIntValue(10), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
	l = NewLenFunction(NewRowValue("i"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestRoundFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "abc")
	ctx.SetValue("i", "13")
	ctx.SetValue("f", "123.6")

	l := NewRoundFunction(NewFloatValue(5.2), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))
	l = NewRoundFunction(NewFloatValue(5.5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(6)))
	l = NewRoundFunction(NewIntValue(5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))

	l = NewRoundFunction(NewRowValue("f"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(124)))
	l = NewRoundFunction(NewRowValue("i"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(13)))

	l = NewRoundFunction(NewStringValue("s"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
	l = NewRoundFunction(NewRowValue("s"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestToStringFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "abc")
	ctx.SetValue("i", "13")
	ctx.SetValue("f", "123.6")

	l := NewToStringFunc(NewFloatValue(5.2), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("5.2")))
	l = NewToStringFunc(NewIntValue(5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("5")))

	l = NewToStringFunc(NewRowValue("f"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("123.6")))
	l = NewToStringFunc(NewRowValue("i"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("13")))

	l = NewToStringFunc(NewStringValue("s"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("s")))
}

func TestToFloatFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "0.5")
	ctx.SetValue("i", "13")
	ctx.SetValue("f", "123.6")

	l := NewToFloatFunc(NewStringValue("5.2"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(5.2)))
	l = NewToFloatFunc(NewIntValue(5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(5.0)))
	l = NewToFloatFunc(NewRowValue("i"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(13)))
	l = NewToFloatFunc(NewRowValue("s"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(0.5)))

	l = NewToFloatFunc(NewStringValue("5.str"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestToIntFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewToIntFunc(NewStringValue("5"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))
	l = NewToIntFunc(NewIntValue(10), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(10)))
	l = NewToIntFunc(NewFloatValue(42.42), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(42)))

	l = NewToFloatFunc(NewStringValue("5.str"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestTruncFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "0.5")
	ctx.SetValue("i", "13")
	ctx.SetValue("f", "123.6")

	l := NewTruncFunc(NewFloatValue(5.2), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))
	l = NewTruncFunc(NewIntValue(5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(5)))
	l = NewTruncFunc(NewRowValue("i"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(13)))
	l = NewTruncFunc(NewRowValue("s"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewIntValue(0)))

	l = NewTruncFunc(NewStringValue("5.str"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestFracFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()
	ctx.SetValue("s", "0.5")
	ctx.SetValue("i", "13")
	ctx.SetValue("f", "123.75")

	l := NewFracFunc(NewFloatValue(5.25), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(0.25)))
	l = NewFracFunc(NewIntValue(5), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(0)))
	l = NewFracFunc(NewRowValue("f"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(0.75)))
	l = NewFracFunc(NewRowValue("s"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(0.5)))

	l = NewFracFunc(NewStringValue("5.str"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestSubStringFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewSubStringFunc(nil, NewStringValue("New York"), NewIntValue(4), NewIntValue(4))
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("York")))

	l = NewSubStringFunc(nil, NewStringValue("New York"), NewIntValue(4), NewIntValue(5))
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("York")))

	l = NewSubStringFunc(nil, NewStringValue("New York"), NewIntValue(4))
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())

	l = NewSubStringFunc(nil, NewStringValue("New York"), NewFloatValue(4.0), NewIntValue(5))
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestToUpperFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewToUpperFunction(NewStringValue("New York"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("NEW YORK")))
}

func TestToLowerFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewToLowerFunction(NewStringValue("New York"), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewStringValue("new york")))
}

func TestSqrtFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewSqrtFunction(NewIntValue(16), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(4.0)))

	l = NewSqrtFunction(NewFloatValue(9.0), nil)
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(3.0)))

	l = NewSqrtFunction(NewStringValue("New York"), nil)
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}

func TestPowFunc(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := newMapContext()

	l := NewPowFunction(nil, NewIntValue(4), NewFloatValue(2.0))
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(16.0)))

	l = NewPowFunction(nil, NewFloatValue(3.0), NewIntValue(3))
	g.Expect(l.Evaluate(ctx)).To(BeEquivalentTo(NewFloatValue(27.0)))

	l = NewPowFunction(nil, NewStringValue("New York"), NewFloatValue(1))
	g.Expect(func() { l.Evaluate(ctx) }).To(Panic())
}
