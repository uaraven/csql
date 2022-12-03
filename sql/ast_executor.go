package sql

import (
	"fmt"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"strings"
)

type AstExecutor struct {
}

func Execute(ast *Select) core.DataSource {
	ae := NewAstExecutor()
	return ae.VisitSelect(ast)
}

func NewAstExecutor() AstExecutor {
	return AstExecutor{}
}

func (ae AstExecutor) VisitSelect(ast *Select) core.DataSource {
	ds := ae.VisitDataSource(&ast.From)
	if ast.Filter != nil {
		where := ae.VisitExpression(ast.Filter)
		ds = core.NewFilteredDataSource(ds, where)
	}

	projection := ae.VisitProjection(ast.Projection.ProjectionFields)

	return core.NewProjectionDataSource(ds, projection, ast.Projection.Distinct)
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

func (ae AstExecutor) VisitJoinedSource(ctx *JoinedSource) core.DataSource {
	left := ae.VisitDataSource(&ctx.Source)
	right := ae.VisitDataSource(&ctx.Target)
	switch ctx.Type {
	case CrossJoin:
		return core.NewCrossJoin(left, right)
	case InnerJoin:
		return core.NewInnerJoin(left, right, ae.VisitExpression(ctx.Condition))
	case LeftOuterJoin:
		return core.NewLeftOuterJoinDatasource(left, right, ae.VisitExpression(ctx.Condition))
	case RightOuterJoin:
		return core.NewRightOuterJoinDatasource(left, right, ae.VisitExpression(ctx.Condition))
	}
	// todo: Implement a proper panic
	panic("Unsupported join type")
}

func (ae AstExecutor) VisitDataSource(ctx *DataSource) core.DataSource {
	if ctx.TableName != nil {
		return ae.VisitSourceName(ctx.TableName)
	} else {
		return ae.VisitJoinedSource(ctx.Join)
	}
}

func (ae AstExecutor) VisitExpression(condition Expression) core.Condition {
	if condition == nil {
		return nil
	}
	switch cond := condition.(type) {
	case Term:
		return ae.VisitTerm(cond)
	case AndExpression:
		return ae.VisitAndExpression(cond)
	case OrExpression:
		return ae.VisitOrExpression(cond)
	case NotExpression:
		return ae.VisitNotExpression(cond)
	case ComparisonExpression:
		return ae.VisitComparisionExpression(cond)
	case BinaryExpression:
		return ae.VisitBinaryExpression(cond)
	case LikeExpression:
		return ae.VisitLikeExpression(cond)
	case MatchExpression:
		return ae.VisitMatchExpression(cond)
	case IsNullExpression:
		return ae.VisitIsNull(cond)
	}
	panic(fmt.Errorf("unsupported expression: %v", condition))
	return nil
}

func (ae AstExecutor) VisitAndExpression(cond AndExpression) core.Condition {
	left := ae.VisitExpression(cond.LHS)
	right := ae.VisitExpression(cond.RHS)
	return core.NewAnd(left, right)
}

func (ae AstExecutor) VisitOrExpression(cond OrExpression) core.Condition {
	left := ae.VisitExpression(cond.LHS)
	right := ae.VisitExpression(cond.RHS)
	return core.NewOr(left, right)
}

func (ae AstExecutor) VisitNotExpression(cond NotExpression) core.Condition {
	expr := ae.VisitExpression(cond.Child)
	return core.NewNot(expr)
}

func (ae AstExecutor) VisitComparisionExpression(cond ComparisonExpression) core.Condition {
	left := ae.VisitExpression(cond.LHS)
	right := ae.VisitExpression(cond.RHS)
	switch cond.Operator {
	case "=":
		return core.NewEq(left, right)
	case "<>":
		return core.NewNeq(left, right)
	case "!=":
		return core.NewNeq(left, right)
	case ">":
		return core.NewGt(left, right)
	case "<":
		return core.NewLt(left, right)
	case ">=":
		return core.NewGte(left, right)
	case "<=":
		return core.NewLte(left, right)
	}
	panic(fmt.Errorf("unsupported expression: %v", cond))
}

func (ae AstExecutor) VisitLikeExpression(cond LikeExpression) core.Condition {
	what := ae.VisitExpression(cond.What)

	var like core.Condition
	if cond.Pattern.Expr != nil {
		like = core.NewLikeCondition(what, ae.VisitExpression(cond.Pattern.Expr))
	} else {
		like = core.NewLikeCondition(what, core.NewStringValue(cond.Pattern.Text))
	}
	if cond.NotLike {
		return core.NewNot(like)
	} else {
		return like
	}
}

func (ae AstExecutor) VisitMatchExpression(cond MatchExpression) core.Condition {
	what := ae.VisitExpression(cond.What)

	var match core.Condition
	if cond.Pattern.Expr != nil {
		match = core.NewMatchCondition(what, ae.VisitExpression(cond.Pattern.Expr))
	} else {
		match = core.NewMatchCondition(what, core.NewStringValue(cond.Pattern.Text))
	}
	if cond.Not {
		return core.NewNot(match)
	} else {
		return match
	}
}

func (ae AstExecutor) VisitProjection(projection []ProjectionField) []core.ProjectionColumn {
	return funky.Map(projection, func(p ProjectionField) core.ProjectionColumn {
		return ae.VisitProjectionField(p)
	})
}

func (ae AstExecutor) VisitProjectionField(p ProjectionField) core.ProjectionColumn {
	if p.NamedField != nil {
		return ae.VisitNamedProjectionField(p.NamedField, p.Alias)
	} else {
		return ae.VisitEvaluatedField(p.EvaluatedField, p.Alias)
	}
}

func (ae AstExecutor) VisitNamedProjectionField(field *NamedProjectionField, alias *Identifier) core.ProjectionColumn {
	var sb strings.Builder
	if field.TableName != nil {
		sb.WriteString(string(*field.TableName))
		sb.WriteRune('.')
	}
	sb.WriteString(string(field.Name))
	if alias != nil {
		return core.NewColumnWithAlias(sb.String(), string(*alias))
	} else {
		return core.NewColumn(sb.String())
	}
}

func (ae AstExecutor) VisitEvaluatedField(field *EvaluatedProjectionField, alias *Identifier) core.ProjectionColumn {
	expr := ae.VisitExpression(field.Expr)
	var aliasStr string
	if alias != nil {
		aliasStr = string(*alias)
	}
	return core.NewExpressionColumn(expr, aliasStr)
}

func (ae AstExecutor) VisitTerm(cond Term) core.Condition {
	if cond.Name != nil {
		return ae.VisitCompoundName(cond.Name)
	} else {
		return ae.VisitLiteral(*cond.Value)
	}
}

func (ae AstExecutor) VisitCompoundName(name *CompoundName) core.Condition {
	var sb strings.Builder
	if name.Qualifier != nil {
		sb.WriteString(string(*name.Qualifier))
		sb.WriteRune('.')
	}
	sb.WriteString(string(name.Name))
	return core.NewRowValue(sb.String())
}

func (ae AstExecutor) VisitLiteral(value Literal) core.Condition {
	if value.IsNull {
		return core.NewNullValue()
	} else if value.StringValue != nil {
		return core.NewStringValue(*value.StringValue)
	} else {
		return core.DecodeNumeric(*value.NumericValue)
	}
}

func (ae AstExecutor) VisitBinaryExpression(cond BinaryExpression) core.Condition {
	left := ae.VisitExpression(cond.LHS)
	right := ae.VisitExpression(cond.RHS)

	if cond.Operator == "||" {
		return core.NewConcat(left, right)
	} else {
		var op core.BinaryOperator = core.BinaryOperator(cond.Operator[0])

		return core.NewExpression(op, left, right)
	}
}

func (ae AstExecutor) VisitIsNull(cond IsNullExpression) core.Condition {
	expr := ae.VisitExpression(cond.What)
	return core.NewIsNull(expr, cond.Not)
}
