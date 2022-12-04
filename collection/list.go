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

type List[T any] interface {
	Append(T)
	AppendAll([]T)
	Clear()
	RemoveByIndex(int) T
	RemoveFirst(func(T) bool)
	IndexOf(func(T) bool) int
	Size() int
	Get(int) T
	Empty() bool

	First() T
	Last() T
	Rest() List[T]
	AsSlice() []T
	Iterator() Iterator[T]
}

type LinkedListElement[T any] interface {
	Next() LinkedListElement[T]
	Previous() LinkedListElement[T]
	Value() T
}

type LinkedList[T any] interface {
	List[T]
	FirstElement() LinkedListElement[T]
	LastElement() LinkedListElement[T]
	Prepend(T)
	Pop() T
}

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}
