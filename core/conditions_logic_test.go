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

var c = newMapContext()

func TestAndCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewAnd(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
}

func TestOrCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewOr(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
}

func TestNotCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewNot(ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewNot(ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
}
