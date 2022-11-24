package core

import (
	"csql/funky"
	"fmt"
	"strconv"
)

type rowImpl struct {
	id     int
	header DataSourceHeader
	values []Value
}

func parseRowWithId(id int, headers DataSourceHeader, values []string) Row {
	return &rowImpl{
		id:     id,
		header: headers,
		values: funky.Map(values, func(s string) Value {
			return decodeToValue(s)
		}),
	}
}

func updateRowId(id int, row Row) Row {
	return &rowImpl{
		id:     id,
		header: row.Header(),
		values: row.Values(),
	}
}

func newRowWithId(id int, headers DataSourceHeader, values []Value) Row {
	return &rowImpl{
		id:     id,
		header: headers,
		values: values,
	}
}

func copyRowWithId(id int, row Row) Row {
	return &rowImpl{
		id:     id,
		header: row.Header(),
		values: row.Values(),
	}
}

func joinRows(id int, header DataSourceHeader, row1 Row, row2 Row) Row {
	rowData := make([]Value, len(row1.Values()))
	copy(rowData, row1.Values())
	rowData = append(rowData, row2.Values()...)

	return &rowImpl{
		id:     id,
		header: header,
		values: rowData,
	}
}

func nullRow(header DataSourceHeader) Row {
	rowData := make([]Value, header.ColumnCount())
	for i := range rowData {
		rowData[i] = NewNullValue()
	}
	return newRowWithId(0, header, rowData)
}

func (r rowImpl) Header() DataSourceHeader {
	return r.header
}

func (r rowImpl) Values() []Value {
	return r.values
}

func (r rowImpl) Get(column string) funky.Option[Value] {
	index := r.Header().IndexByName(column)
	if index == RowIdIndex {
		return funky.SomeOf(NewIntValue(int64(r.id)))
	} else if index < RowIdIndex {
		return funky.NoneOf[Value]()
	} else {
		return funky.SomeOf(r.values[index])
	}
}

func (r rowImpl) Satisfies(cond Condition) bool {
	return cond.Evaluate(r).AsBool().Value().(bool)
}

func (r rowImpl) Id() int {
	return r.id
}

func (r rowImpl) String() string {
	return fmt.Sprintf("%v", r.values)
}

func decodeToValue(value string) Value {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			if value == NullValueString {
				return NewNullValue()
			}
			return NewStringValue(value)
		}
		return NewFloatValue(f64)
	}
	return NewIntValue(i64)
}

func ReadAllRows(ds DataSource) ([]Row, error) {
	result := make([]Row, 0)
	row, err := ds.NextRow()
	if err != nil {
		return nil, err
	}
	for row != nil {
		result = append(result, row)
		row, err = ds.NextRow()
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
