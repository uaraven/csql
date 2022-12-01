package ast

import "github.com/uaraven/csql/core"

type AstExecutor struct {
}

func (ae AstExecutor) VisitSourceName(ctx *SourceName) core.DataSource {
	var ds core.DataSource
	if ctx.Alias != nil {
		ds = core.NewMemDataSourceWithAlias(string(ctx.Name), string(*ctx.Alias))
	} else {
		ds = core.NewMemDataSourceFromCsv(string(ctx.Name))
	}
	return ds
}
