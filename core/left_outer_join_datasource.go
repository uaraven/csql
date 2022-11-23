package core

import (
	"csql/data"
	"fmt"
	"sync"
)

type leftOuterJoinDatasource struct {
	lock       sync.Mutex
	joinedRows []Row
	header     DataSourceHeader
	name       string
	currentRow Row
	counter    int
}

func NewLeftOuterJoinDatasource(left DataSource, right DataSource, joinOn Condition) (DataSource, error) {
	ds := &leftOuterJoinDatasource{
		name:    fmt.Sprintf("(%s LEFT JOIN %s)", left.GetName(), right.GetName()),
		counter: -1,
	}
	ds.header = NewHeadersFromOtherHeaders(ds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	var err error
	ds.joinedRows, err = ds.innerJoin(left, right, joinOn)
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func (l *leftOuterJoinDatasource) Header() DataSourceHeader {
	return l.header
}

func (l *leftOuterJoinDatasource) GetName() string {
	return l.name
}

func (l *leftOuterJoinDatasource) innerJoin(left DataSource, right DataSource, joinOn Condition) ([]Row, error) {
	result := make([]Row, 0)
	counter := int64(0)
	for {
		leftRow, err := left.NextRow()
		if err != nil {
			return nil, err
		}
		if leftRow == nil {
			return result, nil
		}
		joined, err := l.selectRight(leftRow, right, joinOn)
		if err != nil {
			return nil, err
		}
		if joined.Empty() {
			joinedRow := joinRows(counter, l.header, leftRow, nullRow(right.Header()))
			result = append(result, joinedRow)
		} else {
			for _, row := range joined.AsSlice() {
				result = append(result, newRowWithId(counter, l.header, row.Values()))
				counter++
			}
		}
	}

}

func (l *leftOuterJoinDatasource) selectRight(leftRow Row, right DataSource, joinOn Condition) (data.LinkedList[Row], error) {
	var rightSelected data.LinkedList[Row] = data.NewLinkedList[Row]()
	err := right.Rewind()
	if err != nil {
		return nil, err
	}
	for {
		rightRow, err := right.NextRow()
		if err != nil {
			return nil, err
		}
		if rightRow == nil {
			// todo check if results are empty and join left side with nulled right
			return rightSelected, nil
		}
		joined := joinRows(0, l.Header(), leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			rightSelected.Append(joined)
		}
	}
}

func (l *leftOuterJoinDatasource) NextRow() (Row, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.counter++
	if l.counter >= len(l.joinedRows) {
		return nil, nil
	}
	return l.joinedRows[l.counter], nil
}

func (l *leftOuterJoinDatasource) CurrentRow() (Row, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.joinedRows[l.counter], nil
}

func (l *leftOuterJoinDatasource) Rewind() error {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.counter = -1
	return nil
}
