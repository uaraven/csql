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

package collection

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLinkedList_New(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedList[int]()

	g.Expect(ll.Size()).To(Equal(0))
	g.Expect(func() { ll.First() }).To(Panic())
	g.Expect(func() { ll.Last() }).To(Panic())
	g.Expect(ll.Rest()).To(Equal(ll))
}

func TestLinkedList_Append(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedList[int]()

	ll.Append(1)

	g.Expect(ll.First()).To(Equal(1))
	g.Expect(ll.Last()).To(Equal(1))

	ll.Append(2)
	g.Expect(ll.First()).To(Equal(1))
	g.Expect(ll.Last()).To(Equal(2))

	ll.Append(3)
	g.Expect(ll.First()).To(Equal(1))
	g.Expect(ll.FirstElement().Next().Value()).To(Equal(2))
	g.Expect(ll.Last()).To(Equal(3))
}

func TestLinkedList_Prepend(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedList[int]()

	ll.Prepend(1)

	g.Expect(ll.First()).To(Equal(1))
	g.Expect(ll.Last()).To(Equal(1))

	ll.Prepend(2)
	g.Expect(ll.First()).To(Equal(2))
	g.Expect(ll.Last()).To(Equal(1))

	ll.Prepend(3)
	g.Expect(ll.First()).To(Equal(3))
	g.Expect(ll.FirstElement().Next().Value()).To(Equal(2))
	g.Expect(ll.Last()).To(Equal(1))
}

func TestLinkedList_AsSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3})

	actual := ll.AsSlice()

	g.Expect(actual).To(Equal([]int{1, 2, 3}))
}

func TestLinkedList_Clear(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3})

	g.Expect(ll.Size()).To(Equal(3))
	ll.Clear()
	g.Expect(ll.Size()).To(Equal(0))
	g.Expect(ll.FirstElement()).To(BeNil())
	g.Expect(ll.LastElement()).To(BeNil())
}

func TestLinkedList_RemoveByIndex(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})

	ll.RemoveByIndex(3)
	g.Expect(ll.AsSlice()).To(Equal([]int{1, 2, 3, 5}))

	g.Expect(func() { ll.RemoveByIndex(32) }).To(Panic())

	ll.RemoveByIndex(0)
	g.Expect(ll.AsSlice()).To(Equal([]int{2, 3, 5}))

	ll.RemoveByIndex(2)
	g.Expect(ll.AsSlice()).To(Equal([]int{2, 3}))

	ll = NewLinkedListFromSlice([]int{1})
	ll.RemoveByIndex(0)
	g.Expect(ll.Empty()).To(BeTrue())

}

func TestLinkedList_Remove(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})

	ll.RemoveFirst(func(v int) bool {
		return v == 3
	})
	actual := ll.AsSlice()

	g.Expect(actual).To(Equal([]int{1, 2, 4, 5}))
}

func TestLinkedList_Get(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})

	g.Expect(ll.Get(0)).To(Equal(1))
	g.Expect(ll.Get(4)).To(Equal(5))
	g.Expect(func() { ll.Get(5) }).To(Panic())
}

func TestLinkedList_IndexOf(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})

	idx := ll.IndexOf(func(v int) bool { return v == 3 })
	g.Expect(idx).To(Equal(2))

	idx = ll.IndexOf(func(v int) bool { return v == 5 })
	g.Expect(idx).To(Equal(4))

	idx = ll.IndexOf(func(v int) bool { return v == 15 })
	g.Expect(idx).To(Equal(-1))
}

func TestLinkedList_Rest(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})

	sl := ll.Rest().AsSlice()

	g.Expect(sl).To(Equal([]int{2, 3, 4, 5}))

	sl = ll.Rest().Rest().AsSlice()

	g.Expect(sl).To(Equal([]int{3, 4, 5}))
}

func TestLinkedListElement_Previous(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})
	le := ll.LastElement()

	g.Expect(le.Value()).To(Equal(5))
	g.Expect(le.Previous().Value()).To(Equal(4))

	g.Expect(ll.FirstElement().Previous()).To(BeNil())
}

func TestLinkedList_Pop(t *testing.T) {
	g := NewGomegaWithT(t)

	ll := NewLinkedListFromSlice([]int{1, 2})
	g.Expect(ll.Pop()).To(Equal(1))
	g.Expect(ll.Pop()).To(Equal(2))
	g.Expect(func() { ll.Pop() }).To(Panic())

}
