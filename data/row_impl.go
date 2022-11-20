package data

import (
	"csql/funky"
	"fmt"
	"strconv"
)

type rowImpl struct {
	id     int64
	parent DataSource
	values []Value
}

func newRowWithId(id int64, parent DataSource, values []string) Row {
	return &rowImpl{
		id:     id,
		parent: parent,
		values: funky.Map(values, func(s string) Value {
			return decodeToValue(s)
		}),
	}
}

func copyRowWithId(id int64, row Row) Row {
	return &rowImpl{
		id:     id,
		parent: row.Parent(),
		values: row.Values(),
	}
}

func joinRows(id int64, parent DataSource, row1 Row, row2 Row) Row {
	rowData := make([]Value, len(row1.Values()))
	copy(rowData, row1.Values())
	rowData = append(rowData, row2.Values()...)

	return &rowImpl{
		id:     id,
		parent: parent,
		values: rowData,
	}
}

func (r rowImpl) Parent() DataSource {
	return r.parent
}

func (r rowImpl) Values() []Value {
	return r.values
}

func (r rowImpl) Get(column string) funky.Option[Value] {
	index := r.Parent().Header().IndexByName(column)
	if index == RowIdIndex {
		return funky.SomeOf(NewIntValue(r.id))
	} else if index < RowIdIndex {
		return funky.NoneOf[Value]()
	} else {
		return funky.SomeOf(r.values[index])
	}
}

func decodeToValue(value string) Value {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return NewStringValue(value)
		}
		return NewFloatValue(f64)
	}
	return NewIntValue(i64)
}

func (r rowImpl) Satisfies(cond Condition) bool {
	return cond.Evaluate(r).AsBool().Value().(bool)
}

func (r rowImpl) Id() int64 {
	return r.id
}

func (r rowImpl) String() string {
	return fmt.Sprintf("%v", r.values)
}
