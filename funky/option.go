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

import "fmt"

type Option[T any] interface {
	IsEmpty() bool
	IsPresent() bool
	Value() T
	OrGet(v T) T
	OrPanic(err error) T
	IfPresent(func(T))
}

type Some[T any] struct {
	value T
}

func (s Some[T]) IsEmpty() bool {
	return false
}

func (s Some[T]) IsPresent() bool {
	return true
}

func (s Some[T]) Value() T {
	return s.value
}

func (s Some[T]) OrGet(_ T) T {
	return s.value
}

func (s Some[T]) OrPanic(_ error) T {
	return s.value
}

func (s Some[T]) IfPresent(consumer func(v T)) {
	consumer(s.value)
}

func (s Some[T]) String() string {
	return fmt.Sprintf("Some(%v)", s.value)
}

type None[T any] struct {
}

func (n None[T]) IsEmpty() bool {
	return true
}

func (n None[T]) IsPresent() bool {
	return false
}

func (n None[T]) Value() T {
	panic(fmt.Errorf("none doesn't contain a value"))
}

func (n None[T]) OrGet(v T) T {
	return v
}

func (n None[T]) OrPanic(err error) T {
	panic(err)
}

func (n None[T]) IfPresent(_ func(v T)) {
	// do nothing
}

func (n None[T]) String() string {
	return "None()"
}

func SomeOf[T any](v T) Option[T] {
	return Some[T]{value: v}
}

func NoneOf[T any]() Option[T] {
	return None[T]{}
}

func OptionMap[T any, R any](in Option[T], mapper func(T) R) Option[R] {
	if in.IsPresent() {
		return SomeOf(mapper(in.Value()))
	} else {
		return NoneOf[R]()
	}
}
