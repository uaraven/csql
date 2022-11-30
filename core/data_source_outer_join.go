package core

import (
	"fmt"
	"github.com/uaraven/csql/collection"
)

func NewLeftOuterJoinDatasource(left DataSource, right DataSource, joinOn Condition) (DataSource, error) {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s LEFT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	var err error
	mds.data, err = leftOuterJoin(left, right, mds.headers, joinOn)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func NewRightOuterJoinDatasource(left DataSource, right DataSource, joinOn Condition) (DataSource, error) {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s RIGHT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	var err error
	mds.data, err = rightOuterJoin(left, right, mds.headers, joinOn)
	if err != nil {
		return nil, err
	}
	return mds, nil
}

func leftOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) ([]Row, error) {
	result := make([]Row, 0)
	counter := 0
	for {
		leftRow, err := left.NextRow()
		if err != nil {
			return nil, err
		}
		if leftRow == nil {
			return result, nil
		}
		joined, err := selectRight(counter, header, leftRow, right, joinOn)
		if err != nil {
			return nil, err
		}
		if joined.Empty() {
			joinedRow := joinRows(counter, header, leftRow, nullRow(right.Header()))
			counter++
			result = append(result, joinedRow)
		} else {
			iter := joined.Iterator()
			for iter.HasNext() {
				result = append(result, iter.Next())
				counter++
			}
		}
	}

}

func selectRight(counter int, header DataSourceHeader, leftRow Row, right DataSource, joinOn Condition) (collection.LinkedList[Row], error) {
	var rightSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
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
			return rightSelected, nil
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			rightSelected.Append(joined)
		}
	}
}

func rightOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) ([]Row, error) {
	result := make([]Row, 0)
	counter := 0
	for {
		rightRow, err := right.NextRow()
		if err != nil {
			return nil, err
		}
		if rightRow == nil {
			return result, nil
		}
		joined, err := selectLeft(counter, header, rightRow, left, joinOn)
		if err != nil {
			return nil, err
		}
		if joined.Empty() {
			joinedRow := joinRows(counter, header, nullRow(left.Header()), rightRow)
			counter++
			result = append(result, joinedRow)
		} else {
			iter := joined.Iterator()
			for iter.HasNext() {
				result = append(result, iter.Next())
				counter++
			}
		}
	}

}

func selectLeft(counter int, header DataSourceHeader, rightRow Row, left DataSource, joinOn Condition) (collection.LinkedList[Row], error) {
	var leftSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	err := left.Rewind()
	if err != nil {
		return nil, err
	}
	for {
		leftRow, err := left.NextRow()
		if err != nil {
			return nil, err
		}
		if leftRow == nil {
			return leftSelected, nil
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			leftSelected.Append(joined)
		}
	}
}
