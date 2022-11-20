package data

import (
	"csql/util"
	"strings"
)

type columnMetadata struct {
	parent     DataSource
	parentName string
	name       string
	index      int
}

func (cm columnMetadata) Parent() DataSource {
	return cm.parent
}

func (cm columnMetadata) Index() int {
	return cm.index
}

func (cm columnMetadata) MatchesName(name string) bool {
	table, column := qualified(name)
	if table != "" && !util.EqualsIgnoreCase(table, cm.parentName) {
		return false

	}
	return util.EqualsIgnoreCase(column, cm.name)
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
	_, c := qualified(s)
	if c == RowId {
		return RowIdIndex
	}
	for idx, column := range d.columns {
		if column.MatchesName(s) {
			return idx
		}
	}
	return InvalidFieldIndex
}

func NewHeadersFromSlice(parent DataSource, headers []string) DataSourceHeader {
	columns := make([]ColumnMetadata, 0)
	for idx, name := range headers {
		columns = append(columns, &columnMetadata{
			parent:     parent,
			parentName: parent.GetName(),
			name:       name,
			index:      idx,
		})
	}
	return &dataSourceHeader{parent: parent, columns: columns}
}

func qualified(name string) (string, string) {
	p := strings.Split(name, ".")
	if len(p) == 1 {
		return "", name
	} else {
		return p[0], p[1]
	}
}
