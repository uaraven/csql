package data

import "csql/funky"

// ColumnMetadata describes a column
type ColumnMetadata interface {
	// Parent returns the datasource to which this column belongs
	Parent() DataSource
	// Index is the numeric index of the column
	Index() int
	// MatchesName returns true if the requested name matches the name or alias (or index) of this column
	MatchesName(string) bool
}

// DataSourceHeader describes all the columns in the datasource
type DataSourceHeader interface {
	// Parent returns the DataSource containing this header
	Parent() DataSource
	// ColumnCount returns the number of the columns in the DataSource
	ColumnCount() int
	// ColumnsMetadata is a slice of all the columns' metadata
	ColumnsMetadata() []ColumnMetadata
	// IndexByName resolves a name/alias to column index
	IndexByName(string) int
}

type Row interface {
	Parent() DataSource
	Values() []string
	GetByName(string) funky.Option[string]
	Satisfies(c Condition) bool
}

// DataSource is an representation of rows of columns
type DataSource interface {
	Header() DataSourceHeader
	// MatchesName returns true if the provided parameter matches name or alias of the DataSource
	MatchesName(string) bool
	// NextRow returns the next row of values or nil
	NextRow() (Row, error)
	// Rewind the dataset to ahead of the first row
	Rewind() error
}

type QueryContext interface {
	Projection() []string
	GetDataSources() []DataSource
	GetDataSourceByName(string) DataSource
}
