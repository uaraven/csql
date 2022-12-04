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

func TestNewStringValue(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("str")

	g.Expect(v.Type()).To(Equal(TypeString))
	g.Expect(v.Value()).To(Equal("str"))
}

func TestStringValue_AsInt(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12")
	vi := v.AsInt()

	g.Expect(vi.Type()).To(Equal(TypeInt))
	g.Expect(vi.Value()).To(Equal(int64(12)))
}

func TestStringValue_AsFloat(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12.44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(Equal(TypeFloat))
	g.Expect(vi.Value()).To(Equal(12.44))
}

func TestStringValue_AsFloatExp(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12e44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(Equal(TypeFloat))
	g.Expect(vi.Value()).To(Equal(12e44))
}
