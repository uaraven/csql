package core

import "time"

// DatasetCacheEntry contains data for a datasource in memory
type DatasetCacheEntry struct {
	// AccessTime contains the time when the dataset was last accessed
	AccessTime time.Time
	// Size contains the size of datasource (currently in rows)
	Size int
	// ModTime contains the time when the source file was modified
	ModTime time.Time
	// DataSource contains the actual data source
	DataSource DataSource

	TempTable bool
}

// maxCacheSize is the maximum size of the cache in rows
const maxCacheSize = 1_000_000

var CacheEnabled = false

type DatasetCache map[string]*DatasetCacheEntry

var TableCache = make(DatasetCache)

func (ds DatasetCache) AddToCache(name string, dsEntry DataSource, modTime time.Time, tempTable bool) {
	if !CacheEnabled && !tempTable {
		return
	}
	entry := &DatasetCacheEntry{
		AccessTime: time.Now(),
		Size:       len(dsEntry.GetRows()),
		ModTime:    modTime,
		DataSource: dsEntry,
		TempTable:  tempTable,
	}
	totalSize := 0
	for _, v := range ds {
		totalSize += v.Size
	}

	for totalSize+entry.Size > maxCacheSize {
		oldest := ""
		oldestTime := time.Now()
		for k, v := range ds {
			if v.AccessTime.Before(oldestTime) {
				oldestTime = v.AccessTime
				oldest = k
			}
		}
		delete(ds, oldest)
	}
	ds[name] = entry
}

func (ds DatasetCache) ClearCache() (datasets int, rows int) {
	for k, v := range TableCache {
		if v.TempTable {
			rows += len(v.DataSource.GetRows())
			delete(TableCache, k)
			datasets++
		}
	}
	return
}

func (ds DatasetCache) Touch(name string) {
	if v, ok := ds[name]; ok {
		v.AccessTime = time.Now()
	}
}
