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
	"fmt"
	"reflect"
)

type linkedListIterator[T any] struct {
	current LinkedListElement[T]
}

func (l *linkedListIterator[T]) HasNext() bool {
	v := l.current.Next()
	return !(v == nil || reflect.ValueOf(v).IsNil())
}

func (l *linkedListIterator[T]) Next() T {
	if !l.HasNext() {
		panic(fmt.Sprintf("exhausted iterator"))
	}
	l.current = l.current.Next()
	return l.current.Value()
}

func newLinkedListIterator[T any](element LinkedListElement[T]) Iterator[T] {
	return &linkedListIterator[T]{current: element}
}
