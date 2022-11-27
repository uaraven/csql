package core

import "fmt"

type SourceLocation struct {
	line int
	pos  int
}

func NewSourceLocation(line, pos int) SourceLocation {
	return SourceLocation{line: line, pos: pos}
}

type CsqlError struct {
	message string
}

func (ce CsqlError) Error() string {
	return ce.message
}

func NewError(message string) error {
	return &CsqlError{message: message}
}

func UnknownColumnError(columnName string) error {
	return &CsqlError{message: fmt.Sprintf("Unknown column name: %s", columnName)}
}
