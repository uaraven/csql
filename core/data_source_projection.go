package core

import (
	"csql/funky"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
)

const AllColumns = "*"

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
	projection = expandProjection(projection, src.Header())
	headers := newHeaderWithProjection(src, projection)

	var err error
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

func isAsterisk(pc ProjectionColumn) bool {
	if rowId, ok := pc.source.(Identifiable); ok {
		_, sourceColumn := qualified(rowId.Identifier())
		return sourceColumn == AllColumns
	}
	return false
}

func expandProjection(projection []ProjectionColumn, headers DataSourceHeader) []ProjectionColumn {
	if funky.NoneMatches(projection, isAsterisk) {
		return projection
	}
	result := make([]ProjectionColumn, 0)
	for _, column := range projection {
		if isAsterisk(column) {
			fullAsterisk := column.source.(Identifiable).Identifier()
			table, _ := qualified(fullAsterisk)
			for _, cmd := range headers.ColumnsMetadata() {
				if cmd.MatchesName(fullAsterisk) {
					var newName string
					if table != "" {
						newName = table + "." + cmd.Name()
					} else {
						newName = cmd.Name()
					}
					result = append(result, NewColumn(newName))
				}
			}
		} else {
			result = append(result, column)
		}
	}
	return result
}

func newHeaderWithProjection(src DataSource, projection []ProjectionColumn) DataSourceHeader {
	columns := make([]ColumnMetadata, len(projection))
	for idx, prjColumn := range projection {
		var columnMeta columnMetadata
		if rowId, ok := prjColumn.source.(Identifiable); ok {
			// projection column references datasource column
			sourceColumn := rowId.Identifier()
			columnIdx := src.Header().IndexByName(sourceColumn)
			if columnIdx == InvalidFieldIndex {
				panic(UnknownColumnError(sourceColumn))
			}
			sourceColumnMetadata := src.Header().ColumnsMetadata()[columnIdx]
			columnMeta = columnMetadata{
				parentName: sourceColumnMetadata.ParentName(),
				index:      idx,
				name:       sourceColumnMetadata.Name(),
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
	}
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
	resultValues := make([]Value, header.ColumnCount())
	for idx, prj := range projection {
		v := prj.source.Evaluate(row)
		resultValues[idx] = v
	}
	return newRowWithId(row.Id(), header, resultValues)
}
