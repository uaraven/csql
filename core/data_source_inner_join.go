package core

import (
	"fmt"
	"sync"
)

func NewCrossJoin(left DataSource, right DataSource) DataSource {
	return NewInnerJoin(left, right, nil)
}

func NewInnerJoin(left DataSource, right DataSource, joinCondition Condition) DataSource {
	mds := &memDataSource{
		lock:  sync.Mutex{},
		name:  fmt.Sprintf("(%s JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = innerJoin(left, right, mds.headers, joinCondition)
	return mds
}

func innerJoin(left DataSource, right DataSource, header DataSourceHeader, joinCondition Condition) []Row {
	rows := make([]Row, 0)
	counter := 0
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return rows
		}
		for {
			rightRow := right.NextRow()
			if rightRow == nil {
				right.Rewind()
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
