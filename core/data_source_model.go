package core

import (
	"csql/funky"
	"fmt"
)

const (
	RowId             = "__row_id"
	RowIdIndex        = -1
	InvalidFieldIndex = -2
)

// ColumnMetadata describes a column
type ColumnMetadata interface {
	// Index is the numeric index of the column
	Index() int
	// MatchesName returns true if the requested name matches the name or alias (or index) of this column
	MatchesName(string) bool
	// ParentName returns the qualified name of the datasource
	ParentName() string
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
	fmt.Stringer
	Id() int
	Header() DataSourceHeader
	Values() []Value
	Get(string) funky.Option[Value]
	Count() int
	GetByIndex(index int) Value
	Satisfies(c Condition) bool
	Key() string
}

// DataSource is a representation of rows of columns
type DataSource interface {
	// Header returns header information for the DataSource
	Header() DataSourceHeader
	// GetName returns the name of the DataSource
	GetName() string
	// NextRow returns the next row of values.
	//
	// If returned row is nil and error is also nil, then this DataSource does not contain any more rows
	NextRow() (Row, error)

	// CurrentRow returns the row of values that internal DataSource pointer is currently pointed to
	//
	// If nil is returned that means that the pointer is positioned before the first row and NextRow() should be called
	CurrentRow() (Row, error)
	// Rewind the dataset to ahead of the first row
	Rewind() error
}

type SliceDataSource interface {
	GetRows() []Row
}
