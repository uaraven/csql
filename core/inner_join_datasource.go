package core

import (
	"fmt"
	"sync"
)

func NewCrossJoin(left DataSource, right DataSource) (DataSource, error) {
	return NewInnerJoin(left, right, nil)
}

func NewInnerJoin(left DataSource, right DataSource, joinCondition Condition) (DataSource, error) {
	mds := &memDataSource{
		lock:  sync.Mutex{},
		name:  fmt.Sprintf("(%s JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	var err error
	mds.data, err = innerJoin(left, right, mds.headers, joinCondition)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func innerJoin(left DataSource, right DataSource, header DataSourceHeader, joinCondition Condition) ([]Row, error) {
	rows := make([]Row, 0)
	counter := 0
	for {
		leftRow, err := left.NextRow()
		if err != nil {
			return nil, err
		}
		if leftRow == nil {
			return rows, nil
		}
		for {
			rightRow, err := right.NextRow()
			if err != nil {
				return nil, err
			}
			if rightRow == nil {
				err := right.Rewind()
				if err != nil {
					return nil, err
				}
				break
			}
			newRow := joinRows(counter, header, leftRow, rightRow)
			if joinCondition == nil || newRow.Satisfies(joinCondition) {
				rows = append(rows, newRow)
				counter++
			}
		}
	}
}
