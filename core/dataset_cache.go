package core

import "time"

type DatasetCacheEntry struct {
	AddedTime  time.Time
	Size       int
	ModTime    time.Time
	DataSource DataSource
}

const MaxSize = 1_000_000

type DatasetCache map[string]*DatasetCacheEntry

var TableCache = make(DatasetCache)

func (ds DatasetCache) AddToCache(name string, dsEntry DataSource, modTime time.Time) {
	entry := &DatasetCacheEntry{
		AddedTime:  time.Now(),
		Size:       len(dsEntry.GetRows()),
		ModTime:    modTime,
		DataSource: dsEntry,
	}
	totalSize := 0
	for _, v := range ds {
		totalSize += v.Size
	}

	for totalSize+entry.Size > MaxSize {
		oldest := ""
		oldestTime := time.Now()
		for k, v := range ds {
			if v.AddedTime.Before(oldestTime) {
				oldestTime = v.AddedTime
				oldest = k
			}
		}
		delete(ds, oldest)
	}
	ds[name] = entry
}

func (ds DatasetCache) ClearCache() (datasets int, rows int) {
	for k, v := range TableCache {
		rows += len(v.DataSource.GetRows())
		delete(TableCache, k)
		datasets++
	}
	return
}

func (ds DatasetCache) Touch(name string) {
	if v, ok := ds[name]; ok {
		v.AddedTime = time.Now()
	}
}
