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

func TestMatchCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()

	match := NewMatchCondition(NewStringValue("John Connor"), NewStringValue("^J[oh]+n .o.*r$"))

	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))

	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("^J[oh]+n .o.*r$"))

	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
}

func TestLikeCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()

	match := NewLikeCondition(NewStringValue("John Connor"), NewStringValue("J__n%nor"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))
	match = NewLikeCondition(NewStringValue("Jean Tenor"), NewStringValue("J__n%nor"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))

	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("Joh%Con%"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("John%"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
}
