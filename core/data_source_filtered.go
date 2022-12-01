package core

import "github.com/uaraven/csql/funky"

func NewFilteredDataSource(source DataSource, condition Condition) DataSource {
	rows := ReadAllRows(source)
	filteredRows := funky.Filter(rows, func(r Row) bool {
		return r.Satisfies(condition)
	})
	filteredRows = funky.MapWithIndex(filteredRows, func(id int, row Row) Row {
		return updateRowId(id, row)
	})

	return NewMemDataSource(source.GetName(), source.Header().ColumnsMetadata(), filteredRows)
}
