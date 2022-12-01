package core

import (
	"fmt"
	"github.com/uaraven/csql/collection"
)

func NewLeftOuterJoinDatasource(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s LEFT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = leftOuterJoin(left, right, mds.headers, joinOn)
	return mds
}

func NewRightOuterJoinDatasource(left DataSource, right DataSource, joinOn Condition) DataSource {
	mds := &memDataSource{
		name:  fmt.Sprintf("(%s RIGHT JOIN %s)", left.GetName(), right.GetName()),
		index: -1,
	}
	mds.headers = NewHeadersFromOtherHeaders(mds, left.Header().ColumnsMetadata(), right.Header().ColumnsMetadata())
	mds.data = rightOuterJoin(left, right, mds.headers, joinOn)
	return mds
}

func leftOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	result := make([]Row, 0)
	counter := 0
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return result
		}
		joined := selectRight(counter, header, leftRow, right, joinOn)
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

func selectRight(counter int, header DataSourceHeader, leftRow Row, right DataSource, joinOn Condition) collection.LinkedList[Row] {
	var rightSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	right.Rewind()
	for {
		rightRow := right.NextRow()
		if rightRow == nil {
			return rightSelected
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			rightSelected.Append(joined)
		}
	}
}

func rightOuterJoin(left DataSource, right DataSource, header DataSourceHeader, joinOn Condition) []Row {
	result := make([]Row, 0)
	counter := 0
	for {
		rightRow := right.NextRow()
		if rightRow == nil {
			return result
		}
		joined := selectLeft(counter, header, rightRow, left, joinOn)
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

func selectLeft(counter int, header DataSourceHeader, rightRow Row, left DataSource, joinOn Condition) collection.LinkedList[Row] {
	var leftSelected collection.LinkedList[Row] = collection.NewLinkedList[Row]()
	left.Rewind()
	for {
		leftRow := left.NextRow()
		if leftRow == nil {
			return leftSelected
		}
		joined := joinRows(counter, header, leftRow, rightRow)
		if joined.Satisfies(joinOn) {
			leftSelected.Append(joined)
		}
	}
}
