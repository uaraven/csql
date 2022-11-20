package core

import (
	"sync"
	"sync/atomic"
)

type filteredDataSource struct {
	lock       sync.Mutex
	datasource DataSource
	condition  Condition
	counter    atomic.Int64
	currentRow Row
}

func (f *filteredDataSource) Header() DataSourceHeader {
	return f.datasource.Header()
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
			f.currentRow = copyRowWithId(f.counter.Add(1), row)
			return f.currentRow, nil
		}
	}
}

func (f *filteredDataSource) CurrentRow() (Row, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.currentRow == nil {
		return nil, nil
	}
	return f.currentRow, nil
}

func (f *filteredDataSource) Rewind() error {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.counter.Store(-1)
	f.currentRow = nil
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
