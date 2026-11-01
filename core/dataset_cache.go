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
const maxCacheSize = 5_000_000

var CacheEnabled = false

type DatasetCache map[string]*DatasetCacheEntry

var TableCache = make(DatasetCache)

// LastResultName is the name of the temp table with the result of the last query
const LastResultName = ".last"

// AddToCache stores the datasource in the cache, replacing the entry with the same name. Temp tables are owned by
// the user: they do not count against maxCacheSize and are never evicted. Other entries are evicted, oldest first,
// to make room for the new one, and entries bigger than maxCacheSize are not cached at all.
func (ds DatasetCache) AddToCache(name string, dsEntry DataSource, modTime time.Time, tempTable bool) {
	if !CacheEnabled && !tempTable {
		return
	}
	delete(ds, name)
	entry := &DatasetCacheEntry{
		AccessTime: time.Now(),
		Size:       len(dsEntry.GetRows()),
		ModTime:    modTime,
		DataSource: dsEntry,
		TempTable:  tempTable,
	}
	if !tempTable {
		if entry.Size > maxCacheSize {
			return
		}
		ds.evictFor(entry.Size)
	}
	ds[name] = entry
}

// SetLastResult stores the result of the query as ".last" temp table. The old result is dropped, and
// a result that is bigger than maxCacheSize is not kept in memory.
func (ds DatasetCache) SetLastResult(dsEntry DataSource) {
	delete(ds, LastResultName)
	if len(dsEntry.GetRows()) > maxCacheSize {
		return
	}
	ds.AddToCache(LastResultName, dsEntry, time.Now(), true)
}

// evictFor removes the least recently used cached (non-temp) datasets until size more rows fit
func (ds DatasetCache) evictFor(size int) {
	totalSize := 0
	for _, v := range ds {
		if !v.TempTable {
			totalSize += v.Size
		}
	}
	for totalSize+size > maxCacheSize {
		oldest := ""
		var oldestTime time.Time
		for k, v := range ds {
			if !v.TempTable && (oldest == "" || v.AccessTime.Before(oldestTime)) {
				oldestTime = v.AccessTime
				oldest = k
			}
		}
		if oldest == "" {
			return
		}
		totalSize -= ds[oldest].Size
		delete(ds, oldest)
	}
}

func (ds DatasetCache) ClearCache() (datasets int, rows int) {
	for k, v := range ds {
		if !v.TempTable {
			rows += len(v.DataSource.GetRows())
			delete(ds, k)
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
