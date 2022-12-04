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

func Map[T any, R any](data []T, transformer func(T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(elem)
	}
	return result
}

func MapWithIndex[T any, R any](data []T, transformer func(int, T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(idx, elem)
	}
	return result
}

func Find[T any](data []T, matcher func(T) bool) (int, Option[T]) {
	for idx, elem := range data {
		if matcher(elem) {
			return idx, SomeOf(elem)
		}
	}
	return -1, None[T]{}
}

func AnyMatches[T any](data []T, matcher func(T) bool) bool {
	for _, elem := range data {
		if matcher(elem) {
			return true
		}
	}
	return false
}

func NoneMatches[T any](data []T, matcher func(T) bool) bool {
	return !AnyMatches(data, matcher)
}

func Filter[T any](data []T, matcher func(T) bool) []T {
	result := make([]T, 0)
	for idx, elem := range data {
		if matcher(elem) {
			result = append(result, data[idx])
		}
	}
	return result
}
