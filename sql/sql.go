package sql

import (
	"github.com/uaraven/csql/core"
)

func ExecuteSql(query string) []core.Row {
	ast := ParseSQL(query)
	data := Execute(&ast)

	return core.ReadAllRows(data)
}
