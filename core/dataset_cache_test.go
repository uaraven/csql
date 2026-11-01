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
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func memDsWithRows(n int) DataSource {
	return NewMemDataSource("big", []ColumnMetadata{NewSimpleColumnMetadata("a", 0)}, make([]Row, n))
}

func withCleanCache(t *testing.T) {
	saved, savedEnabled := TableCache, CacheEnabled
	TableCache = make(DatasetCache)
	CacheEnabled = true
	t.Cleanup(func() { TableCache, CacheEnabled = saved, savedEnabled })
}

func TestCache_OversizedEntriesAreNotCached(t *testing.T) {
	g := NewGomegaWithT(t)
	withCleanCache(t)

	done := make(chan struct{})
	go func() {
		defer close(done)
		TableCache.AddToCache("huge", memDsWithRows(maxCacheSize+1), time.Now(), false)
		TableCache.SetLastResult(memDsWithRows(maxCacheSize + 1))
	}()
	g.Eventually(done, 5*time.Second).Should(BeClosed())
	g.Expect(TableCache).To(BeEmpty())
}

func TestCache_LastResultIsReplaced(t *testing.T) {
	g := NewGomegaWithT(t)
	withCleanCache(t)

	TableCache.SetLastResult(memDsWithRows(3))
	TableCache.SetLastResult(memDsWithRows(maxCacheSize + 1))
	g.Expect(TableCache).ToNot(HaveKey(LastResultName))

	TableCache.SetLastResult(memDsWithRows(4))
	g.Expect(TableCache[LastResultName].Size).To(Equal(4))
}

func TestCache_EvictsOldestCachedButNeverTempTables(t *testing.T) {
	g := NewGomegaWithT(t)
	withCleanCache(t)

	TableCache.AddToCache("mine", memDsWithRows(maxCacheSize), time.Now(), true)
	TableCache.AddToCache("a", memDsWithRows(3_000_000), time.Now(), false)
	TableCache["a"].AccessTime = time.Now().Add(-time.Hour)
	TableCache.AddToCache("b", memDsWithRows(1_000_000), time.Now(), false)
	TableCache.AddToCache("c", memDsWithRows(2_000_000), time.Now(), false)

	g.Expect(TableCache).To(HaveKey("mine"))
	g.Expect(TableCache).ToNot(HaveKey("a"))
	g.Expect(TableCache).To(HaveKey("b"))
	g.Expect(TableCache).To(HaveKey("c"))

	// replacing an entry does not count the old one
	TableCache.AddToCache("c", memDsWithRows(4_000_000), time.Now(), false)
	g.Expect(TableCache).To(HaveKey("b"))
	g.Expect(TableCache["c"].Size).To(Equal(4_000_000))
}
