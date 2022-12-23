/*
 *    Copyright 2021 Oleksiy Voronin
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
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	. "github.com/uaraven/csql/errors"
	"github.com/uaraven/csql/parser"
	"strconv"
	"strings"
)

type ParsingError struct {
	line    int
	column  int
	message string
}

type Errors struct {
	antlr.ErrorListener
	errors []ParsingError
}

func newErrors() *Errors {
	return &Errors{
		errors: make([]ParsingError, 0),
	}
}

func (e *Errors) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	e.errors = append(e.errors, ParsingError{
		line:    line,
		column:  column,
		message: msg,
	})
}

func (e *Errors) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
}
func (e *Errors) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
}
func (e *Errors) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ antlr.ATNConfigSet) {
}

// ParseSQL reads the sql statement and converts it into an AST
func ParseSQL(sql string) UnionSource {
	fs := antlr.NewInputStream(sql)
	lex := parser.NewCsqlLexer(fs)
	lex.RemoveErrorListeners()
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCsqlParser(tokens)
	p.RemoveErrorListeners()
	errors := newErrors()
	p.AddErrorListener(errors)
	visitor := NewCsqlVisitor()
	tree := p.Query()
	if len(errors.errors) != 0 {
		err1 := errors.errors[0]
		srcLoc := Loc(err1.line, err1.column)
		panic(NewError(srcLoc, err1.message))
	}
	return *visitor.Visit(tree).(*UnionSource)
}

// CsqlVisitorImpl is a visitor for the antlr CST
type CsqlVisitorImpl struct {
	*parser.BaseCsqlVisitor
}

func removeStringQuotes(text string) string {
	return strings.Trim(text, "'")
}

func isQuotedIdentifier(text string) bool {
	return text[0] == '"' && text[len(text)-1] == '"'
}

func removeIdentifierQuotes(text string) string {
	return strings.Trim(text, `"`)
}

func parseIdentifier(identifier string) Identifier {
	if isQuotedIdentifier(identifier) {
		return Identifier(removeIdentifierQuotes(identifier))
	}
	return Identifier(identifier)
}

func parseUnsignedIntValue(token antlr.Token, text string) interface{} {
	limitValue, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		panic(ParsingError{
			line:    token.GetLine(),
			column:  token.GetColumn(),
			message: fmt.Sprintf("Unexpected: %v, expected unsigned integer", text),
		})
	}
	return int32(limitValue)
}

// NewCsqlVisitor creates new instance of the visitor
func NewCsqlVisitor() parser.CsqlVisitor {
	baseVisitor := parser.BaseCsqlVisitor{}
	visitor := CsqlVisitorImpl{
		BaseCsqlVisitor: &baseVisitor,
	}
	return &visitor
}

// Visit starts visiting of the Antlr's parse tree
func (c CsqlVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *parser.QueryContext:
		return val.Accept(c)
	}
	return nil
}

// VisitQuery visits a parse tree produced by CsqlParser#query.
func (c CsqlVisitorImpl) VisitQuery(ctx *parser.QueryContext) interface{} {
	return ctx.UnionSelects().Accept(c).(*UnionSource)
}

// VisitSelectStatement visits a parse tree produced by CsqlParser#selectStatement.
func (c CsqlVisitorImpl) VisitSelectStatement(ctx *parser.SelectStatementContext) interface{} {
	astSelect := Select{}
	astSelect.Projection = ctx.Projection().Accept(c).(SelectProjection)
	astSelect.From = ctx.Sources().Accept(c).(DataSource)
	if ctx.Where() != nil {
		e := ctx.Where().Accept(c).(Expression)
		astSelect.Filter = e
	}
	if ctx.Limit() != nil {
		astSelect.Limit = ctx.Limit().Accept(c).(int32)
	}
	if ctx.OrderBy() != nil {
		obe := ctx.OrderBy().Accept(c).(OrderByExpression)
		astSelect.OrderBy = &obe
	}
	if ctx.GroupBy() != nil {
		gbe := ctx.GroupBy().Accept(c).(GroupByExpression)
		astSelect.GroupBy = &gbe
	}
	return astSelect
}

func (c CsqlVisitorImpl) VisitUnionSelects(ctx *parser.UnionSelectsContext) interface{} {
	astUnion := UnionSource{
		Select:   ctx.SelectStatement().Accept(c).(Select),
		unionAll: ctx.K_ALL(0) != nil,
	}
	if len(ctx.AllUnionSelects()) > 0 {
		astUnion.Union = ctx.UnionSelects(0).Accept(c).(*UnionSource)
	}
	return &astUnion
}

func (c CsqlVisitorImpl) VisitLimit(ctx *parser.LimitContext) interface{} {
	return ctx.LimitValue().Accept(c)
}

func (c CsqlVisitorImpl) VisitLimitValue(ctx *parser.LimitValueContext) interface{} {
	return parseUnsignedIntValue(ctx.GetStart(), ctx.GetText())
}

func (c CsqlVisitorImpl) VisitOrderBy(ctx *parser.OrderByContext) interface{} {
	obe := OrderByExpression{OrderFields: []OrderByField{}, Location: SLFromToken(ctx.GetStart())}
	for idx, _ := range ctx.AllOrderByField() {
		obe.OrderFields = append(obe.OrderFields, ctx.OrderByField(idx).Accept(c).(OrderByField))
	}
	return obe
}

func (c CsqlVisitorImpl) VisitOrderByField(ctx *parser.OrderByFieldContext) interface{} {
	obf := OrderByField{
		Location: SLFromToken(ctx.GetStart()),
	}
	if ctx.CompoundName() != nil {
		fieldName := ctx.CompoundName().Accept(c).(CompoundName)
		obf.FieldName = &fieldName
	} else {
		obf.FieldIndex = ctx.FieldIndex().Accept(c).(int32)
	}
	if ctx.K_DESC() != nil {
		obf.Descending = true
	}
	return obf
}

func (c CsqlVisitorImpl) VisitFieldIndex(ctx *parser.FieldIndexContext) interface{} {
	return parseUnsignedIntValue(ctx.GetStart(), ctx.GetText())
}

// VisitProjection visits a parse tree produced by CsqlParser#projection.
func (c CsqlVisitorImpl) VisitProjection(ctx *parser.ProjectionContext) interface{} {
	astProjection := SelectProjection{}
	astProjection.Distinct = ctx.Distinct() != nil
	astProjection.ProjectionFields = make([]ProjectionField, 0)
	for _, field := range ctx.AllProjectionField() {
		astField := field.Accept(c).(ProjectionField)
		astProjection.ProjectionFields = append(astProjection.ProjectionFields, astField)
	}
	return astProjection
}

// VisitComparisonOperator visits a parse tree produced by CsqlParser#logicalOperator.
func (c CsqlVisitorImpl) VisitComparisonOperator(ctx *parser.ComparisonOperatorContext) interface{} {
	operator := ctx.GetText()
	return operator
}

// VisitBinaryOperation visits a parse tree produced by CsqlParser#binaryOperation.
func (c CsqlVisitorImpl) VisitBinaryOperation(ctx *parser.BinaryOperationContext) interface{} {
	operator := ctx.GetText()
	return operator
}

// VisitList visits a parse tree produced by CsqlParser#list.
func (c CsqlVisitorImpl) VisitList(ctx *parser.ListContext) interface{} {
	result := make([]Expression, 0)
	for _, item := range ctx.AllValueExpr() {
		result = append(result, item.Accept(c).(Expression))
	}
	return ListLiteral{
		Values: result,
	}
}

func (c CsqlVisitorImpl) VisitValueBinaryExpr(ctx *parser.ValueBinaryExprContext) interface{} {
	return BinaryExpression{
		Location: SLFromToken(ctx.GetStart()),
		LHS:      ctx.ValueExpr(0).Accept(c).(Expression),
		RHS:      ctx.ValueExpr(1).Accept(c).(Expression),
		Operator: ctx.BinaryOperation().Accept(c).(string),
	}
}

func (c CsqlVisitorImpl) VisitValueParensExpr(ctx *parser.ValueParensExprContext) interface{} {
	return ctx.ValueExpr().Accept(c)
}

// VisitTerm visits a parse tree produced by CsqlParser#term.
func (c CsqlVisitorImpl) VisitTerm(ctx *parser.TermContext) interface{} {
	if ctx.LiteralValue() != nil {
		l := ctx.LiteralValue().Accept(c).(Literal)
		return Term{
			Value: &l,
		}
	}
	n := ctx.CompoundName().Accept(c).(CompoundName)
	return Term{
		Name: &n,
	}
}

func (c CsqlVisitorImpl) VisitValueExprItem(ctx *parser.ValueExprItemContext) interface{} {
	return ctx.ValueExpr().Accept(c)
}

// VisitCondition visits a parse tree produced by CsqlParser#condition.
func (c CsqlVisitorImpl) VisitCondition(ctx *parser.ConditionContext) interface{} {
	return ComparisonExpression{
		Location: SLFromToken(ctx.GetStart()),
		LHS:      ctx.WhereExpr(0).Accept(c).(Expression),
		RHS:      ctx.WhereExpr(1).Accept(c).(Expression),
		Operator: ctx.ComparisonOperator().Accept(c).(string),
	}
}

// VisitBetweenExpr visits a parse tree produced by CsqlParser#betweenExpr.
func (c CsqlVisitorImpl) VisitBetweenExpr(ctx *parser.BetweenExprContext) interface{} {
	return BetweenExpression{
		Location: SLFromToken(ctx.GetStart()),
		What:     ctx.ValueExpr(0).Accept(c).(Expression),
		Low:      ctx.ValueExpr(1).Accept(c).(Expression),
		High:     ctx.ValueExpr(2).Accept(c).(Expression),
	}
}

func (c CsqlVisitorImpl) VisitTermItem(ctx *parser.TermItemContext) interface{} {
	return ctx.Term().Accept(c).(Term)
}

// VisitNotExpr visits a parse tree produced by CsqlParser#notExpr.
func (c CsqlVisitorImpl) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return NotExpression{
		Location: SLFromToken(ctx.GetStart()),
		Child:    ctx.WhereExpr().Accept(c).(Expression),
	}
}

// VisitInExpr visits a parse tree produced by CsqlParser#inExpr.
func (c CsqlVisitorImpl) VisitInExpr(ctx *parser.InExprContext) interface{} {
	return InListExpression{
		Location: SLFromToken(ctx.GetStart()),
		What:     ctx.ValueExpr().Accept(c).(Expression),
		List:     ctx.List().Accept(c).(ListLiteral),
		NotIn:    ctx.K_NOT() != nil,
	}
}

func (c CsqlVisitorImpl) VisitIsNullExpr(ctx *parser.IsNullExprContext) interface{} {
	return IsNullExpression{
		What: ctx.ValueExpr().Accept(c).(Expression),
		Not:  ctx.K_NOT() != nil,
	}
}

// VisitOrExpr visits a parse tree produced by CsqlParser#orExpr.
func (c CsqlVisitorImpl) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	return OrExpression{
		Location: SLFromToken(ctx.GetStart()),
		LHS:      ctx.WhereExpr(0).Accept(c).(Expression),
		RHS:      ctx.WhereExpr(1).Accept(c).(Expression),
	}
}

// VisitParensExpr visits a parse tree produced by CsqlParser#parensExpr.
func (c CsqlVisitorImpl) VisitParensExpr(ctx *parser.ParensExprContext) interface{} {
	return ctx.WhereExpr().Accept(c).(Expression)
}

// VisitAndExpr visits a parse tree produced by CsqlParser#andExpr.
func (c CsqlVisitorImpl) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	return AndExpression{
		Location: SLFromToken(ctx.GetStart()),
		LHS:      ctx.WhereExpr(0).Accept(c).(Expression),
		RHS:      ctx.WhereExpr(1).Accept(c).(Expression),
	}
}

// VisitLikeExpr visits a parse tree produced by CsqlParser#likeExpr.
func (c CsqlVisitorImpl) VisitLikeExpr(ctx *parser.LikeExprContext) interface{} {
	return LikeExpression{
		What:     ctx.ValueExpr(0).Accept(c).(Expression),
		Pattern:  ctx.ValueExpr(1).Accept(c).(Expression),
		NotLike:  ctx.K_NOT() != nil,
		Location: SLFromToken(ctx.GetStart()),
	}
}

// VisitWhere visits a parse tree produced by CsqlParser#where.
func (c CsqlVisitorImpl) VisitWhere(ctx *parser.WhereContext) interface{} {
	return ctx.WhereExpr().Accept(c).(Expression)
}

// VisitProjectionField visits a parse tree produced by CsqlParser#projectionField.
func (c CsqlVisitorImpl) VisitProjectionField(ctx *parser.ProjectionFieldContext) interface{} {
	var projectionField ProjectionField
	if ctx.ProjectionFieldName() != nil {
		field := ctx.ProjectionFieldName().Accept(c).(CompoundName)
		projectionField.NamedField = &field
	} else if ctx.ValueExpr() != nil {
		expr := ctx.ValueExpr().Accept(c).(Expression)
		projectionField.EvaluatedField = &EvaluatedProjectionField{
			Expr: expr,
		}
	} else {
		expr := ctx.AggregateFunCall().Accept(c).(Expression)
		projectionField.EvaluatedField = &EvaluatedProjectionField{Expr: expr}
	}
	if ctx.Alias() != nil {
		alias := ctx.Alias().Accept(c).(Identifier)
		projectionField.Alias = &alias
	}
	return projectionField
}

// VisitProjectionFieldName visits a parse tree produced by CsqlParser#projectionFieldName.
func (c CsqlVisitorImpl) VisitProjectionFieldName(ctx *parser.ProjectionFieldNameContext) interface{} {
	field := CompoundName{
		Location: SLFromToken(ctx.GetStart()),
		Name:     ctx.FieldName().Accept(c).(Identifier),
	}
	if ctx.Qualifier() != nil {
		tableName := ctx.Qualifier().Accept(c).(Identifier)
		field.Qualifier = &tableName
	}
	return field
}

// VisitFieldName visits a parse tree produced by CsqlParser#fieldName.
func (c CsqlVisitorImpl) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	return parseIdentifier(ctx.GetText())
}

func (c CsqlVisitorImpl) VisitValueTerm(ctx *parser.ValueTermContext) interface{} {
	return ctx.Term().Accept(c).(Term)
}

func (c CsqlVisitorImpl) VisitMatchExpr(ctx *parser.MatchExprContext) interface{} {
	return MatchExpression{
		Location: SLFromToken(ctx.GetStart()),
		What:     ctx.ValueExpr(0).Accept(c).(Expression),
		Pattern:  ctx.ValueExpr(1).Accept(c).(Expression),
		Not:      ctx.K_NOT() != nil,
	}
}

type joinTarget struct {
	dataSource DataSource
	condition  Expression
}

// VisitConditionalJoinTarget visits a parse tree produced by CsqlParser#joinTarget.
func (c CsqlVisitorImpl) VisitConditionalJoinTarget(ctx *parser.ConditionalJoinTargetContext) interface{} {
	result := joinTarget{}
	result.dataSource = ctx.DataSource().Accept(c).(DataSource)

	if ctx.WhereExpr() != nil {
		e := ctx.WhereExpr().Accept(c).(Expression)
		result.condition = e
	}
	return result
}

func (c CsqlVisitorImpl) VisitUnconditionalJoinTarget(ctx *parser.UnconditionalJoinTargetContext) interface{} {
	return ctx.DataSource().Accept(c).(DataSource)
}

func (c CsqlVisitorImpl) VisitConditionalJoinType(ctx *parser.ConditionalJoinTypeContext) interface{} {
	if ctx.InnerJoin() != nil {
		return InnerJoin
	} else if ctx.RightJoin() != nil {
		return RightOuterJoin
	} else if ctx.LeftJoin() != nil {
		return LeftOuterJoin
	} else if ctx.FullJoin() != nil {
		return FullJoin
	}
	return InvalidJoin
}

func (c CsqlVisitorImpl) VisitDataSourceNamed(ctx *parser.DataSourceNamedContext) interface{} {
	dataSource := DataSource{}
	v := ctx.SourceName().Accept(c).(SourceName)
	dataSource.TableName = &v
	return dataSource
}

func (c CsqlVisitorImpl) VisitDataSourceConditionalJoin(ctx *parser.DataSourceConditionalJoinContext) interface{} {
	dataSource := DataSource{}
	joinType := ctx.ConditionalJoinType().Accept(c).(JoinType)
	joinTarget := ctx.ConditionalJoinTarget().Accept(c).(joinTarget)
	dataSource.Join = &JoinedSource{
		Type:      joinType,
		Source:    ctx.DataSource().Accept(c).(DataSource),
		Target:    joinTarget.dataSource,
		Condition: joinTarget.condition,
	}
	return dataSource
}

func (c CsqlVisitorImpl) VisitDataSourceCrossJoin(ctx *parser.DataSourceCrossJoinContext) interface{} {
	dataSource := DataSource{}
	dataSource.Join = &JoinedSource{
		Type: CrossJoin,
	}
	dataSource.Join.Source = ctx.DataSource().Accept(c).(DataSource)
	dataSource.Join.Target = ctx.UnconditionalJoinTarget().Accept(c).(DataSource)
	return dataSource
}

func (c CsqlVisitorImpl) VisitDataSourceItem(ctx *parser.DataSourceItemContext) interface{} {
	return ctx.DataSource().Accept(c).(DataSource)
}

// VisitSources visits a parse tree produced by CsqlParser#sources.
func (c CsqlVisitorImpl) VisitSources(ctx *parser.SourcesContext) interface{} {
	dataSources := make([]DataSource, 0)
	for _, cstDataSource := range ctx.AllDataSource() {
		dataSource := cstDataSource.Accept(c).(DataSource)
		dataSources = append(dataSources, dataSource)
	}
	if len(dataSources) == 1 {
		return dataSources[0]
	}
	// all datasouces in the list are implicit cross join, we need to convert it to such
	topSource := DataSource{
		Join: &JoinedSource{
			Type:   CrossJoin,
			Source: dataSources[0],
		},
	}
	parent := topSource
	for i, src := range dataSources[1:] {
		if i == len(dataSources)-2 {
			parent.Join.Target.TableName = src.TableName
		} else {
			parent.Join.Target = DataSource{
				Join: &JoinedSource{
					Type:   CrossJoin,
					Source: src,
				},
			}
		}
		parent = parent.Join.Target
	}
	return topSource
}

// VisitSignedNumber visits a parse tree produced by CsqlParser#signedNumber.
func (c CsqlVisitorImpl) VisitSignedNumber(ctx *parser.SignedNumberContext) interface{} {
	return ctx.GetText()
}

// VisitStringValue visits a parse tree produced by CsqlParser#stringValue.
func (c CsqlVisitorImpl) VisitStringValue(ctx *parser.StringValueContext) interface{} {
	return removeStringQuotes(ctx.GetText())
}

// VisitNullValue visits a parse tree produced by CsqlParser#nullValue.
func (c CsqlVisitorImpl) VisitNullValue(ctx *parser.NullValueContext) interface{} {
	return ctx.GetText()
}

// VisitLiteralValue visits a parse tree produced by CsqlParser#literalValue.
func (c CsqlVisitorImpl) VisitLiteralValue(ctx *parser.LiteralValueContext) interface{} {
	l := Literal{
		Location: SLFromToken(ctx.GetStart()),
	}
	if ctx.SignedNumber() != nil {
		v := ctx.SignedNumber().Accept(c).(string)
		l.NumericValue = &v
	} else if ctx.StringValue() != nil {
		v := removeStringQuotes(ctx.StringValue().Accept(c).(string))
		l.StringValue = &v
	}
	l.IsNull = ctx.NullValue() != nil
	return l
}

// VisitAlias visits a parse tree produced by CsqlParser#alias.
func (c CsqlVisitorImpl) VisitAlias(ctx *parser.AliasContext) interface{} {
	return parseIdentifier(ctx.GetText())
}

// VisitSourceName visits a parse tree produced by CsqlParser#sourceName.
func (c CsqlVisitorImpl) VisitSourceName(ctx *parser.SourceNameContext) interface{} {
	name := ctx.Name().Accept(c).(Identifier)
	var alias *Identifier
	if ctx.Alias() != nil {
		v := ctx.Alias().Accept(c).(Identifier)
		alias = &v
	} else {
		alias = nil
	}
	return SourceName{
		Name:  name,
		Alias: alias,
	}
}

// VisitCompoundName visits a parse tree produced by CsqlParser#compoundName.
func (c CsqlVisitorImpl) VisitCompoundName(ctx *parser.CompoundNameContext) interface{} {
	result := CompoundName{
		Location: SLFromToken(ctx.GetStart()),
	}
	if ctx.Qualifier() == nil {
		result.Qualifier = nil
	} else {
		q := ctx.Qualifier().Accept(c).(Identifier)
		result.Qualifier = &q
	}
	result.Name = Identifier(ctx.Name().GetText())
	return result
}

// VisitName visits a parse tree produced by CsqlParser#name.
func (c CsqlVisitorImpl) VisitName(ctx *parser.NameContext) interface{} {
	return parseIdentifier(ctx.GetText())
}

// VisitQualifier visits a parse tree produced by CsqlParser#qualifier.
func (c CsqlVisitorImpl) VisitQualifier(ctx *parser.QualifierContext) interface{} {
	if ctx == nil {
		return nil
	}
	return parseIdentifier(ctx.GetText())
}

func (c CsqlVisitorImpl) VisitFunctionExpr(ctx *parser.FunctionExprContext) interface{} {
	return ctx.FunCall().Accept(c)
}

func (c CsqlVisitorImpl) VisitFunCall(ctx *parser.FunCallContext) interface{} {
	funcName := strings.ToLower(ctx.Function().GetText())
	args := make([]Expression, len(ctx.AllValueExpr()))
	for idx, ve := range ctx.AllValueExpr() {
		args[idx] = ve.Accept(c).(Expression)
	}
	return FunctionCall{
		function: funcName,
		args:     args,
		Location: SLFromToken(ctx.GetStart()),
	}
}

func (c CsqlVisitorImpl) VisitGroupBy(ctx *parser.GroupByContext) interface{} {
	fields := make([]GroupByField, len(ctx.AllGroupByField()))
	for idx, f := range ctx.AllGroupByField() {
		fields[idx] = f.Accept(c).(GroupByField)
	}
	return GroupByExpression{
		Location:    SLFromToken(ctx.GetStart()),
		GroupFields: fields,
	}
}

func (c CsqlVisitorImpl) VisitGroupByField(ctx *parser.GroupByFieldContext) interface{} {
	gbf := GroupByField{
		Location: SLFromToken(ctx.GetStart()),
	}
	if ctx.CompoundName() != nil {
		fieldName := ctx.CompoundName().Accept(c).(CompoundName)
		gbf.FieldName = &fieldName
	} else {
		gbf.FieldIndex = ctx.FieldIndex().Accept(c).(int32)
	}
	return gbf

}

func (c CsqlVisitorImpl) VisitAggregateFunCall(ctx *parser.AggregateFunCallContext) interface{} {
	if ctx.CountFunc() != nil {
		return ctx.CountFunc().Accept(c).(CountFunctionCall)
	} else {
		funcName := strings.ToLower(ctx.AggregateFunc().GetText())
		args := ctx.ValueExpr().Accept(c).(Expression)
		return AggregateFunctionCall{
			function: funcName,
			arg:      args,
			Location: SLFromToken(ctx.GetStart()),
		}
	}
}

func (c CsqlVisitorImpl) VisitCountFunc(ctx *parser.CountFuncContext) interface{} {
	distinct := ctx.K_DISTINCT() != nil
	field := ctx.ProjectionFieldName().Accept(c).(CompoundName)
	return CountFunctionCall{
		arg:      field,
		distinct: distinct,
		Location: SLFromToken(ctx.GetStart()),
	}
}
