/*
 *    Copyright 2022 Oleksiy Voronin
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 *    SPDX-License-Identifier: Apache-2.0
 *    SPDX-FileCopyrightText: (c) 2022 Oleksiy Voronin <ovoronin@gmail.com>
 */

package sql

import (
	"fmt"
	"github.com/uaraven/csql/core"
	"github.com/uaraven/csql/funky"
	"strings"
)

type AstTransformer struct {
}

// Transform processes the abstract syntax tree and creates executable Datasource that will
// produce results when queried
func Transform(ast *UnionSource) core.DataSource {
	ae := NewAstTransformer()
	return ae.TransformUnion(*ast)
}

func NewAstTransformer() AstTransformer {
	return AstTransformer{}
}

func (ae AstTransformer) TransformUnion(u UnionSource) core.DataSource {
	ds1 := ae.TransformSelect(u.Select)
	if u.Union == nil {
		return ds1
	}
	ds2 := ae.TransformUnion(*u.Union)
	if u.unionAll {
		return core.NewUnionAll(ds1, ds2)
	} else {
		return core.NewUnion(ds1, ds2)
	}
}

func (ae AstTransformer) TransformSelect(ast Select) core.DataSource {
	ds := ae.TransformDataSource(&ast.From)
	if ast.Filter != nil {
		where := ae.TransformExpression(ast.Filter)
		ds = core.NewFilteredDataSource(ds, where)
	}

	var orderBy []core.OrderByField
	if ast.OrderBy != nil {
		orderBy = ae.TransformOrderByExpression(ast.OrderBy)
	}
	if ast.Limit > 0 || len(orderBy) > 0 {
		ds = core.NewOrderedDatasource(ds, orderBy, int(ast.Limit))
	}

	projection := ae.TransformProjection(ast.Projection.ProjectionFields)

	return core.NewProjectionDataSource(ds, projection, ast.Projection.Distinct)
}

func (ae AstTransformer) TransformSourceName(ctx *SourceName) core.DataSource {
	var ds core.DataSource
	if ctx.Alias != nil {
		ds = core.NewMemDataSourceWithAlias(string(ctx.Name), string(*ctx.Alias))
	} else {
		ds = core.NewMemDataSourceFromCsv(string(ctx.Name))
	}
	return ds
}

func (ae AstTransformer) TransformJoinedSource(ctx *JoinedSource) core.DataSource {
	left := ae.TransformDataSource(&ctx.Source)
	right := ae.TransformDataSource(&ctx.Target)
	switch ctx.Type {
	case CrossJoin:
		return core.NewCrossJoin(left, right)
	case InnerJoin:
		return core.NewInnerJoin(left, right, ae.TransformExpression(ctx.Condition))
	case LeftOuterJoin:
		return core.NewLeftOuterJoin(left, right, ae.TransformExpression(ctx.Condition))
	case RightOuterJoin:
		return core.NewRightOuterJoin(left, right, ae.TransformExpression(ctx.Condition))
	case FullJoin:
		return core.NewFullJoin(left, right, ae.TransformExpression(ctx.Condition))
	}
	// todo: Implement a proper panic
	panic("Unsupported join type")
}

func (ae AstTransformer) TransformDataSource(ctx *DataSource) core.DataSource {
	if ctx.TableName != nil {
		return ae.TransformSourceName(ctx.TableName)
	} else {
		return ae.TransformJoinedSource(ctx.Join)
	}
}

func (ae AstTransformer) TransformExpression(condition Expression) core.Condition {
	if condition == nil {
		return nil
	}
	switch cond := condition.(type) {
	case Term:
		return ae.TransformTerm(cond)
	case AndExpression:
		return ae.TransformAndExpression(cond)
	case OrExpression:
		return ae.TransformOrExpression(cond)
	case NotExpression:
		return ae.TransformNotExpression(cond)
	case ComparisonExpression:
		return ae.TransformComparisionExpression(cond)
	case BinaryExpression:
		return ae.TransformBinaryExpression(cond)
	case LikeExpression:
		return ae.TransformLikeExpression(cond)
	case MatchExpression:
		return ae.TransformMatchExpression(cond)
	case IsNullExpression:
		return ae.TransformIsNull(cond)
	case InListExpression:
		return ae.TransformInListExpression(cond)
	}
	panic(fmt.Errorf("unsupported expression: %v", condition))
	return nil
}

func (ae AstTransformer) TransformAndExpression(cond AndExpression) core.Condition {
	left := ae.TransformExpression(cond.LHS)
	right := ae.TransformExpression(cond.RHS)
	return core.NewAnd(left, right)
}

func (ae AstTransformer) TransformOrExpression(cond OrExpression) core.Condition {
	left := ae.TransformExpression(cond.LHS)
	right := ae.TransformExpression(cond.RHS)
	return core.NewOr(left, right)
}

func (ae AstTransformer) TransformNotExpression(cond NotExpression) core.Condition {
	expr := ae.TransformExpression(cond.Child)
	return core.NewNot(expr)
}

func (ae AstTransformer) TransformComparisionExpression(cond ComparisonExpression) core.Condition {
	left := ae.TransformExpression(cond.LHS)
	right := ae.TransformExpression(cond.RHS)
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

func (ae AstTransformer) TransformLikeExpression(cond LikeExpression) core.Condition {
	what := ae.TransformExpression(cond.What)

	like := core.NewLikeCondition(what, ae.TransformExpression(cond.Pattern))
	if cond.NotLike {
		return core.NewNot(like)
	} else {
		return like
	}
}

func (ae AstTransformer) TransformMatchExpression(cond MatchExpression) core.Condition {
	what := ae.TransformExpression(cond.What)

	match := core.NewMatchCondition(what, ae.TransformExpression(cond.Pattern))
	if cond.Not {
		return core.NewNot(match)
	} else {
		return match
	}
}

func (ae AstTransformer) TransformProjection(projection []ProjectionField) []core.ProjectionColumn {
	return funky.Map(projection, func(p ProjectionField) core.ProjectionColumn {
		return ae.TransformProjectionField(p)
	})
}

func (ae AstTransformer) TransformProjectionField(p ProjectionField) core.ProjectionColumn {
	if p.NamedField != nil {
		return ae.TransformNamedProjectionField(p.NamedField, p.Alias)
	} else {
		return ae.TransformEvaluatedField(p.EvaluatedField, p.Alias)
	}
}

func (ae AstTransformer) TransformNamedProjectionField(field *CompoundName, alias *Identifier) core.ProjectionColumn {
	var sb strings.Builder
	if field.Qualifier != nil {
		sb.WriteString(string(*field.Qualifier))
		sb.WriteRune('.')
	}
	sb.WriteString(string(field.Name))
	if alias != nil {
		return core.NewColumnWithAlias(sb.String(), string(*alias))
	} else {
		return core.NewColumn(sb.String())
	}
}

func (ae AstTransformer) TransformEvaluatedField(field *EvaluatedProjectionField, alias *Identifier) core.ProjectionColumn {
	expr := ae.TransformExpression(field.Expr)
	var aliasStr string
	if alias != nil {
		aliasStr = string(*alias)
	}
	return core.NewExpressionColumn(expr, aliasStr)
}

func (ae AstTransformer) TransformTerm(cond Term) core.Evaluator {
	if cond.Name != nil {
		return ae.TransformCompoundName(cond.Name)
	} else {
		return ae.TransformLiteral(*cond.Value)
	}
}

func (ae AstTransformer) TransformCompoundName(name *CompoundName) core.Condition {
	var sb strings.Builder
	if name.Qualifier != nil {
		sb.WriteString(string(*name.Qualifier))
		sb.WriteRune('.')
	}
	sb.WriteString(string(name.Name))
	return core.NewRowValue(sb.String())
}

func (ae AstTransformer) TransformLiteral(value Literal) core.Condition {
	if value.IsNull {
		return core.NewNullValue()
	} else if value.StringValue != nil {
		return core.NewStringValue(*value.StringValue)
	} else {
		return core.DecodeNumeric(*value.NumericValue)
	}
}

func (ae AstTransformer) TransformBinaryExpression(cond BinaryExpression) core.Condition {
	left := ae.TransformExpression(cond.LHS)
	right := ae.TransformExpression(cond.RHS)

	if cond.Operator == "||" {
		return core.NewConcat(left, right)
	} else {
		var op = core.BinaryOperator(cond.Operator[0])
		return core.NewExpression(op, left, right)
	}
}

func (ae AstTransformer) TransformIsNull(cond IsNullExpression) core.Condition {
	expr := ae.TransformExpression(cond.What)
	return core.NewIsNull(expr, cond.Not)
}

func (ae AstTransformer) TransformInListExpression(cond InListExpression) core.Condition {
	expr := ae.TransformExpression(cond.What)
	values := funky.Map(cond.List.Values, func(li Expression) core.Evaluator {
		return ae.TransformExpression(li)
	})
	inExpr := core.NewInList(expr, values)
	if cond.NotIn {
		return core.NewNot(inExpr)
	} else {
		return inExpr
	}
}

func (ae AstTransformer) TransformOrderByExpression(orderBy *OrderByExpression) []core.OrderByField {
	var result []core.OrderByField
	for _, obf := range orderBy.OrderFields {
		result = append(result, ae.TransformOrderByField(obf))
	}
	return result
}

func (ae AstTransformer) TransformOrderByField(obf OrderByField) core.OrderByField {
	if obf.FieldIndex > 0 {
		return core.NewOrderByIndex(int(obf.FieldIndex), obf.Descending)
	} else {
		fn := ae.TransformCompoundName(obf.FieldName).(core.Identifiable)
		return core.NewOrderBy(fn.Identifier(), obf.Descending)
	}
}
