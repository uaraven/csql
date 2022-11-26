package core

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

type ProjectionColumn struct {
	source Evaluator
	alias  string
}

func NewColumn(name string) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name)}
}

func NewColumnWithAlias(name string, alias string) ProjectionColumn {
	return ProjectionColumn{source: NewRowValue(name), alias: alias}
}

func NewExpressionColumn(src Evaluator, alias string) ProjectionColumn {
	return ProjectionColumn{source: src, alias: alias}
}

func NewSimpleProjection(sourceNames []string, aliases []string) []ProjectionColumn {
	if len(sourceNames) != len(aliases) {
		panic(fmt.Errorf("projection requires equal number of names and aliases.\nNames:%v\nAliases:%v", sourceNames, aliases))
	}
	projection := make([]ProjectionColumn, len(sourceNames))
	for idx, src := range sourceNames {
		projection[idx] = NewColumnWithAlias(src, aliases[idx])
	}
	return projection
}

func NewProjectionDataSource(src DataSource, projection []ProjectionColumn, distinct bool) (DataSource, error) {
	var err error
	headers, err := newHeaderWithProjection(src, projection)
	if err != nil {
		return nil, err
	}
	var rows []Row
	if distinct {
		rows, err = performProjectionDistinct(src, headers, projection)
	} else {
		rows, err = performProjection(src, headers, projection)
	}
	if err != nil {
		return nil, err
	}

	return NewMemDataSource(src.GetName(), headers.ColumnsMetadata(), rows)
}

func newHeaderWithProjection(src DataSource, projection []ProjectionColumn) (DataSourceHeader, error) {
	columns := make([]ColumnMetadata, len(projection))
	for idx, prjColumn := range projection {
		var columnMeta columnMetadata
		if rowId, ok := prjColumn.source.(Identifiable); ok {
			// projection column references datasource column
			sourceColumn := rowId.Identifier()
			columnIdx := src.Header().IndexByName(sourceColumn)
			if columnIdx == InvalidFieldIndex {
				return nil, fmt.Errorf("unknown column: %v", sourceColumn)
			}
			columnMeta = columnMetadata{
				parentName: src.Header().ColumnsMetadata()[columnIdx].ParentName(),
				index:      idx,
				name:       sourceColumn,
			}
		} else {
			// projection column is an expression
			columnMeta = columnMetadata{
				parentName: "",
				index:      idx,
				name:       "",
			}
		}
		if prjColumn.alias != "" {
			columnMeta.name = prjColumn.alias
		}
		columns[idx] = &columnMeta
	}

	return &dataSourceHeader{
		parent:  src,
		columns: columns,
	}, nil
}

func performProjection(ds DataSource, header DataSourceHeader, projection []ProjectionColumn) ([]Row, error) {
	result := make([]Row, 0)
	srcRow, err := ds.NextRow()
	if err != nil {
		return nil, err
	}

	for srcRow != nil {
		newRow := projectRow(projection, header, srcRow)
		result = append(result, newRow)
		srcRow, err = ds.NextRow()
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

func performProjectionDistinct(ds DataSource, header DataSourceHeader, projection []ProjectionColumn) ([]Row, error) {
	keySet := mapset.NewThreadUnsafeSet[string]()
	result := make([]Row, 0)
	srcRow, err := ds.NextRow()
	if err != nil {
		return nil, err
	}

	for srcRow != nil {
		newRow := projectRow(projection, header, srcRow)
		if !keySet.Contains(newRow.Key()) {
			result = append(result, newRow)
			keySet.Add(newRow.Key())
		}
		srcRow, err = ds.NextRow()
		if err != nil {
			return nil, err
		}

	}
	return result, nil
}

func projectRow(projection []ProjectionColumn, header DataSourceHeader, row Row) Row {
	resultValues := make([]Value, len(projection))
	for idx, prj := range projection {
		v := prj.source.Evaluate(row)
		resultValues[idx] = v
	}
	return newRowWithId(row.Id(), header, resultValues)
}
