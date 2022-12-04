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

package util

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEqualsIgnoreCase(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(EqualsIgnoreCase("string1", "StRiNg1")).To(BeTrue())
	g.Expect(EqualsIgnoreCase("string1", "StRiNg2")).To(BeFalse())
}

func TestMust(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(Must(15, nil)).To(Equal(15))

	g.Expect(func() { Must(15, fmt.Errorf("error")) }).To(PanicWith(fmt.Errorf("error")))
}
