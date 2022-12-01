package core

import (
	"fmt"
	"sort"
)

type OrderByField struct {
	name  string
	index int
	desc  bool
}

func NewOrderBy(column string, desc bool) OrderByField {
	return OrderByField{
		name:  column,
		index: -1,
		desc:  desc,
	}
}

func NewOrderByIndex(index int, desc bool) OrderByField {
	return OrderByField{
		name:  "",
		index: index,
		desc:  desc,
	}
}

func NewOrderedDatasource(ds DataSource, orderBy []OrderByField, limit int) DataSource {
	var rows []Row
	if mds, ok := ds.(SliceDataSource); ok {
		rows = mds.GetRows()
	} else {
		rows = ReadAllRows(ds)
	}

	if len(orderBy) > 0 {
		sort.Slice(rows, func(i, j int) bool {
			return rowSorter(i, j, rows, orderBy)
		})
	}

	if limit > 0 {
		rows = rows[:limit]
	}
	return NewMemDataSource(ds.GetName(), ds.Header().ColumnsMetadata(), rows)
}

func rowSorter(i, j int, rows []Row, orderBy []OrderByField) bool {
	r1 := rows[i]
	r2 := rows[j]

	for _, order := range orderBy {
		var r1v, r2v Value
		if order.index > 0 {
			r1v = r1.GetByIndex(order.index)
			r2v = r2.GetByIndex(order.index)
		} else {
			r1v = r1.Get(order.name).OrPanic(fmt.Errorf("unknown column: %v", order.name))
			r2v = r2.Get(order.name).OrPanic(fmt.Errorf("unknown column: %v", order.name))
		}
		if !r1v.Equals(r2v) {
			less := r1v.Less(r2v)
			if order.desc {
				return !less
			} else {
				return less
			}
		}
	}
	return false
}
