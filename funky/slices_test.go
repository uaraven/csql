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

func TestFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	result := Filter(list, func(i int) bool { return (i % 2) == 0 })

	g.Expect(result).To(ContainElements(2, 4, 6))
}

func TestFind_Success(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	idx, result := Find(list, func(i int) bool { return i == 4 })

	g.Expect(idx).To(Equal(3))

	g.Expect(result.IsPresent()).To(BeTrue())
	g.Expect(result.Value()).To(Equal(4))
}

func TestFind_NotFound(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	idx, result := Find(list, func(i int) bool { return i == 10 })

	g.Expect(idx).To(Equal(-1))

	g.Expect(result.IsPresent()).To(BeFalse())
}

func TestMap(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	result := Map(list, func(i int) int { return i * 2 })

	g.Expect(result).To(ContainElements(2, 4, 6, 8, 10, 12))
}

func TestAnyMatches(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}

	result := AnyMatches(list, func(i int) bool { return i == 4 })
	g.Expect(result).To(BeTrue())

	result = AnyMatches(list, func(i int) bool { return i == -1 })
	g.Expect(result).To(BeFalse())
}

func TestNoneMatches(t *testing.T) {
	g := NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}

	result := NoneMatches(list, func(i int) bool { return i == 4 })
	g.Expect(result).To(BeFalse())

	result = NoneMatches(list, func(i int) bool { return i == -1 })
	g.Expect(result).To(BeTrue())
}
