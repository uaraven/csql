package data

import (
	"csql/funky"
	"csql/util"
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
	parent DataSource
	values []string
}

func (r rowImpl) Parent() DataSource {
	return r.parent
}

func (r rowImpl) Values() []string {
	return r.values
}

func (r rowImpl) GetByName(column string) funky.Option[string] {
	index := r.Parent().Header().IndexByName(column)
	if index < 0 {
		return funky.NoneOf[string]()
	} else {
		return funky.SomeOf(r.values[index])
	}
}
