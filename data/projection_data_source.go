package data

import "fmt"

type ProjectionColumn struct {
	name  string
	alias string
}

func NewProjectionColumn(name string) ProjectionColumn {
	return ProjectionColumn{name: name}
}

func NewProjectionColumnWithAlias(name string, alias string) ProjectionColumn {
	return ProjectionColumn{name: name, alias: alias}
}

func NewProjection(sourceNames []string, aliases []string) []ProjectionColumn {
	if len(sourceNames) != len(aliases) {
		panic(fmt.Errorf("projection requires equal number of names and aliases.\nNames:%v\nAliases:%v", sourceNames, aliases))
	}
	projection := make([]ProjectionColumn, len(sourceNames))
	for idx, src := range sourceNames {
		projection[idx] = NewProjectionColumnWithAlias(src, aliases[idx])
	}
	return projection
}

type projectionDataSource struct {
	source     DataSource
	header     DataSourceHeader
	projection []ProjectionColumn
}

func NewProjectionDataSource(src DataSource, projection []ProjectionColumn) DataSource {
	pds := &projectionDataSource{
		source:     src,
		header:     newHeaderWithProjection(src, projection),
		projection: projection,
	}

	return pds
}

func newHeaderWithProjection(src DataSource, projection []ProjectionColumn) DataSourceHeader {
	columns := make([]ColumnMetadata, len(projection))
	for idx, prjColumn := range projection {
		columnIdx := src.Header().IndexByName(prjColumn.name)
		if columnIdx == InvalidFieldIndex {
			panic(fmt.Errorf("unknown column: %v", prjColumn.name))
		}
		columnMeta := columnMetadata{
			parent:     src.Header().ColumnsMetadata()[columnIdx].Parent(),
			parentName: src.Header().ColumnsMetadata()[columnIdx].ParentName(),
			index:      idx,
		}
		if prjColumn.alias != "" {
			columnMeta.name = prjColumn.alias
		} else {
			columnMeta.name = prjColumn.name
		}
		columns[idx] = &columnMeta
	}

	return &dataSourceHeader{
		parent:  src,
		columns: columns,
	}
}

func (p *projectionDataSource) Header() DataSourceHeader {
	return p.header
}

func (p *projectionDataSource) GetName() string {
	return ""
}

func (p *projectionDataSource) MatchesName(s string) bool {
	return false
}

func (p *projectionDataSource) NextRow() (Row, error) {
	row, err := p.source.NextRow()
	if row == nil {
		return row, err
	}
	if err != nil {
		return nil, err
	}
	return p.Project(row)
}

func (p *projectionDataSource) CurrentRow() (Row, error) {
	row, err := p.source.CurrentRow()
	if err != nil {
		return nil, err
	}
	return p.Project(row)
}

func (p *projectionDataSource) Rewind() error {
	return p.source.Rewind()
}

func (p *projectionDataSource) Project(row Row) (Row, error) {
	resultValues := make([]Value, len(p.projection))
	for idx, prj := range p.projection {
		v := row.Get(prj.name)
		if v.IsEmpty() {
			return nil, fmt.Errorf("expected column '%v' is not found in the dataset", prj.name)
		}
		resultValues[idx] = v.Value()
	}
	return newRowWithIdAndValues(row.Id(), p, resultValues), nil
}
