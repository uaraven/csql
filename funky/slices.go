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
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/sync/errgroup"
	"runtime"
	"sync"
)

func Map[T any, R any](data []T, transformer func(T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(elem)
	}
	return result
}

func ParallelMap[T any, R any](data []T, transformer func(T) R) []R {
	result := make([]R, len(data))
	var wg sync.WaitGroup
	for idx := range data {
		i := idx
		wg.Add(1)
		go func() {
			result[i] = transformer(data[i])
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}

func MapWithIndex[T any, R any](data []T, transformer func(int, T) R) []R {
	result := make([]R, len(data))
	for idx, elem := range data {
		result[idx] = transformer(idx, elem)
	}
	return result
}

func ParallelMapWithIndex[T any, R any](data []T, transformer func(int, T) R) []R {
	result := make([]R, len(data))
	var wg sync.WaitGroup
	for idx := range data {
		i := idx
		wg.Add(1)
		go func() {
			result[i] = transformer(i, data[i])
			wg.Done()
		}()
	}
	wg.Wait()
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

func ParallelFilter[T any](data []T, matcher func(T) bool) []T {
	size := len(data)
	cpuCount := runtime.NumCPU()
	if size <= 4*cpuCount {
		return Filter(data, matcher)
	}
	chunkSize := size / cpuCount
	count := cpuCount
	if size%cpuCount != 0 {
		count++
	}
	eg := new(errgroup.Group)
	start := 0
	resultChunks := make([][]T, count)
	for i := 0; i < count; i++ {
		l := start + chunkSize
		if l > len(data) {
			l = len(data)
		}
		iter := i
		if start < l {
			dt := data[start:l]
			eg.Go(func() (err error) {
				defer func() {
					fail := recover()
					if fail != nil {
						if ferr, ok := fail.(error); ok {
							err = ferr
							return
						} else {
							err = fmt.Errorf("%v", fail)
							return
						}
					}
				}()
				resultChunks[iter] = Filter(dt, matcher)
				return
			})
		}
		start = l
	}
	err := eg.Wait()
	if err != nil {
		panic(err)
	}
	result := make([]T, 0)
	for _, d := range resultChunks {
		result = append(result, d...)
	}
	return result
}
func Fold[T any](data []T, zero T, folding func(T, T) T) T {
	result := zero
	for _, elem := range data {
		result = folding(result, elem)
	}
	return result
}

func Max[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		panic(fmt.Errorf("cannot find max value in empty slice"))
	}
	if len(data) == 1 {
		return data[0]
	}
	return Fold(data[1:], data[0], func(t1, t2 T) T {
		if t1 > t2 {
			return t1
		} else {
			return t2
		}
	})
}

func Sum[T constraints.Integer](data []T) T {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	return Fold(data, 0, func(t1, t2 T) T {
		return t1 + t2
	})
}
