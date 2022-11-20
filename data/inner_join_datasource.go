package data

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// innerJoinDatasource combines two DataSources with a filter
// if filer is nil, then it works as a cross join
type innerJoinDatasource struct {
	lock          sync.Mutex
	id            atomic.Int64
	left          DataSource
	right         DataSource
	header        DataSourceHeader
	joinCondition Condition
	currentRow    Row
}

func NewCrossJoin(left DataSource, right DataSource) DataSource {
	return NewInnerJoin(left, right, nil)
}

func NewInnerJoin(left DataSource, right DataSource, joinCondition Condition) DataSource {
	ds := &innerJoinDatasource{
		left:          left,
		right:         right,
		joinCondition: joinCondition,
		currentRow:    nil,
	}
	ds.id.Store(-1)
	ds.header = NewHeadersFromOtherHeaders(ds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	return ds
}

func (i *innerJoinDatasource) Header() DataSourceHeader {
	return i.header
}

func (i *innerJoinDatasource) GetName() string {
	return i.left.GetName() + " JOIN " + i.right.GetName()
}

func (i *innerJoinDatasource) NextRow() (Row, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	for {
		leftRow, err := i.left.CurrentRow()
		if err != nil {
			return nil, fmt.Errorf("failed to read current row from left dataset: %w", err)
		}
		if leftRow == nil {
			leftRow, err = i.left.NextRow()
			if err != nil {
				return nil, fmt.Errorf("failed to read next row from left dataset: %w", err)
			}
			if leftRow == nil {
				return nil, nil
			}
		}
		rightRow, err := i.right.NextRow()
		if err != nil {
			return nil, fmt.Errorf("failed to read next row from right dataset: %w", err)
		}
		if rightRow == nil {
			leftRow, err = i.left.NextRow()
			if err != nil {
				return nil, fmt.Errorf("failed to read next row from left dataset: %w", err)
			}
			if leftRow == nil {
				return nil, nil
			}
			err = i.right.Rewind()
			if err != nil {
				return nil, fmt.Errorf("failed to rewind right dataset: %w", err)
			}
		} else {
			newRow := joinRows(i.id.Add(1), i.Header(), leftRow, rightRow)
			if i.joinCondition == nil || newRow.Satisfies(i.joinCondition) {
				i.currentRow = newRow
				return newRow, nil
			} else {
				i.id.Add(-1) // reset counter and repeat
			}
		}
	}
}

func (i *innerJoinDatasource) CurrentRow() (Row, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.currentRow, nil
}

func (i *innerJoinDatasource) Rewind() error {
	i.lock.Lock()
	defer i.lock.Unlock()

	err := i.left.Rewind()
	if err != nil {
		return fmt.Errorf("failed to rewind left: %w", err)
	}
	err = i.right.Rewind()
	if err != nil {
		return fmt.Errorf("failed to rewind right: %w", err)
	}
	i.currentRow = nil
	i.id.Store(-1)

	return nil
}
