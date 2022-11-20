package data

import (
	"sync"
	"sync/atomic"
)

type filteredDataSource struct {
	lock       sync.Mutex
	datasource DataSource
	condition  Condition
	counter    atomic.Int64
}

func (f *filteredDataSource) Header() DataSourceHeader {
	return f.datasource.Header()
}

func (f *filteredDataSource) MatchesName(s string) bool {
	return f.datasource.MatchesName(s)
}

func (f *filteredDataSource) GetName() string {
	return f.datasource.GetName()
}

func (f *filteredDataSource) NextRow() (Row, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	for {
		row, err := f.datasource.NextRow()
		if err != nil {
			return nil, err
		}
		if row == nil {
			return nil, nil
		}
		if row.Satisfies(f.condition) {
			return copyRowWithId(f.counter.Add(1), row), nil
		}
	}
}

func (f *filteredDataSource) Rewind() error {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.counter.Store(-1)
	return f.datasource.Rewind()
}

func NewFilteredDataSource(source DataSource, condition Condition) DataSource {
	fds := &filteredDataSource{
		datasource: source,
		condition:  condition,
	}
	fds.counter.Store(-1)
	return fds
}
