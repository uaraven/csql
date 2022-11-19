package data

import (
	"csql/funky"
	"csql/util"
	"strconv"
)

type columnMetadata struct {
	parent DataSource
	name   string
	alias  string
	index  int
}

func (cm columnMetadata) Parent() DataSource {
	return cm.parent
}

func (cm columnMetadata) Index() int {
	return cm.index
}

func (cm columnMetadata) MatchesName(name string) bool {
	return util.EqualsIgnoreCase(name, cm.alias) ||
		util.EqualsIgnoreCase(name, cm.name)
}

type dataSourceHeader struct {
	parent  DataSource
	columns []ColumnMetadata
}

func (d dataSourceHeader) Parent() DataSource {
	return d.parent
}

func (d dataSourceHeader) ColumnCount() int {
	return len(d.columns)
}

func (d dataSourceHeader) ColumnsMetadata() []ColumnMetadata {
	return d.columns
}

func (d dataSourceHeader) IndexByName(s string) int {
	for idx, column := range d.columns {
		if column.MatchesName(s) {
			return idx
		}
	}
	return -1
}

func NewHeadersFromSlice(parent DataSource, headers []string) DataSourceHeader {
	columns := make([]ColumnMetadata, 0)
	for idx, name := range headers {
		columns = append(columns, &columnMetadata{
			parent: parent,
			name:   name,
			alias:  "",
			index:  idx,
		})
	}
	return &dataSourceHeader{parent: parent, columns: columns}
}

type rowImpl struct {
	id     int64
	parent DataSource
	values []string
}

func newRowWithId(id int64, parent DataSource, values []string) Row {
	return &rowImpl{
		id:     id,
		parent: parent,
		values: values,
	}
}

func copyRowWithId(id int64, row Row) Row {
	return &rowImpl{
		id:     id,
		parent: row.Parent(),
		values: row.Values(),
	}
}

func (r rowImpl) Parent() DataSource {
	return r.parent
}

func (r rowImpl) Values() []string {
	return r.values
}

func (r rowImpl) GetByName(column string) funky.Option[string] {
	if column == RowId {
		return funky.SomeOf(strconv.FormatInt(r.id, 10))
	}
	index := r.Parent().Header().IndexByName(column)
	if index < 0 {
		return funky.NoneOf[string]()
	} else {
		return funky.SomeOf(r.values[index])
	}
}

func (r rowImpl) Get(column string) funky.Option[Value] {
	s := r.GetByName(column)
	return funky.OptMap(s, func(p string) Value {
		return decodeToValue(p)
	})
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
