// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Csql

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by CsqlParser.
type CsqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CsqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by CsqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by CsqlParser#binaryOperation.
	VisitBinaryOperation(ctx *BinaryOperationContext) interface{}

	// Visit a parse tree produced by CsqlParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by CsqlParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by CsqlParser#valueTerm.
	VisitValueTerm(ctx *ValueTermContext) interface{}

	// Visit a parse tree produced by CsqlParser#valueBinaryExpr.
	VisitValueBinaryExpr(ctx *ValueBinaryExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#valueParensExpr.
	VisitValueParensExpr(ctx *ValueParensExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#functionExpr.
	VisitFunctionExpr(ctx *FunctionExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#aggregateFuncExpr.
	VisitAggregateFuncExpr(ctx *AggregateFuncExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by CsqlParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#termItem.
	VisitTermItem(ctx *TermItemContext) interface{}

	// Visit a parse tree produced by CsqlParser#inExpr.
	VisitInExpr(ctx *InExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#isNullExpr.
	VisitIsNullExpr(ctx *IsNullExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#likeExpr.
	VisitLikeExpr(ctx *LikeExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#matchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#parensExpr.
	VisitParensExpr(ctx *ParensExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#valueExprItem.
	VisitValueExprItem(ctx *ValueExprItemContext) interface{}

	// Visit a parse tree produced by CsqlParser#betweenExpr.
	VisitBetweenExpr(ctx *BetweenExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by CsqlParser#where.
	VisitWhere(ctx *WhereContext) interface{}

	// Visit a parse tree produced by CsqlParser#having.
	VisitHaving(ctx *HavingContext) interface{}

	// Visit a parse tree produced by CsqlParser#distinct.
	VisitDistinct(ctx *DistinctContext) interface{}

	// Visit a parse tree produced by CsqlParser#funCall.
	VisitFunCall(ctx *FunCallContext) interface{}

	// Visit a parse tree produced by CsqlParser#aggregateFunCall.
	VisitAggregateFunCall(ctx *AggregateFunCallContext) interface{}

	// Visit a parse tree produced by CsqlParser#countFunc.
	VisitCountFunc(ctx *CountFuncContext) interface{}

	// Visit a parse tree produced by CsqlParser#aggregateFunc.
	VisitAggregateFunc(ctx *AggregateFuncContext) interface{}

	// Visit a parse tree produced by CsqlParser#projection.
	VisitProjection(ctx *ProjectionContext) interface{}

	// Visit a parse tree produced by CsqlParser#projectionField.
	VisitProjectionField(ctx *ProjectionFieldContext) interface{}

	// Visit a parse tree produced by CsqlParser#projectionFieldName.
	VisitProjectionFieldName(ctx *ProjectionFieldNameContext) interface{}

	// Visit a parse tree produced by CsqlParser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by CsqlParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#leftJoin.
	VisitLeftJoin(ctx *LeftJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#rightJoin.
	VisitRightJoin(ctx *RightJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#fullJoin.
	VisitFullJoin(ctx *FullJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#crossJoin.
	VisitCrossJoin(ctx *CrossJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#conditionalJoinType.
	VisitConditionalJoinType(ctx *ConditionalJoinTypeContext) interface{}

	// Visit a parse tree produced by CsqlParser#conditionalJoinTarget.
	VisitConditionalJoinTarget(ctx *ConditionalJoinTargetContext) interface{}

	// Visit a parse tree produced by CsqlParser#unconditionalJoinTarget.
	VisitUnconditionalJoinTarget(ctx *UnconditionalJoinTargetContext) interface{}

	// Visit a parse tree produced by CsqlParser#dataSourceConditionalJoin.
	VisitDataSourceConditionalJoin(ctx *DataSourceConditionalJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#dataSourceItem.
	VisitDataSourceItem(ctx *DataSourceItemContext) interface{}

	// Visit a parse tree produced by CsqlParser#dataSourceCrossJoin.
	VisitDataSourceCrossJoin(ctx *DataSourceCrossJoinContext) interface{}

	// Visit a parse tree produced by CsqlParser#dataSourceNamed.
	VisitDataSourceNamed(ctx *DataSourceNamedContext) interface{}

	// Visit a parse tree produced by CsqlParser#sources.
	VisitSources(ctx *SourcesContext) interface{}

	// Visit a parse tree produced by CsqlParser#unionSelects.
	VisitUnionSelects(ctx *UnionSelectsContext) interface{}

	// Visit a parse tree produced by CsqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by CsqlParser#signedNumber.
	VisitSignedNumber(ctx *SignedNumberContext) interface{}

	// Visit a parse tree produced by CsqlParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by CsqlParser#nullValue.
	VisitNullValue(ctx *NullValueContext) interface{}

	// Visit a parse tree produced by CsqlParser#literalValue.
	VisitLiteralValue(ctx *LiteralValueContext) interface{}

	// Visit a parse tree produced by CsqlParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by CsqlParser#sourceName.
	VisitSourceName(ctx *SourceNameContext) interface{}

	// Visit a parse tree produced by CsqlParser#compoundName.
	VisitCompoundName(ctx *CompoundNameContext) interface{}

	// Visit a parse tree produced by CsqlParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by CsqlParser#qualifier.
	VisitQualifier(ctx *QualifierContext) interface{}

	// Visit a parse tree produced by CsqlParser#limit.
	VisitLimit(ctx *LimitContext) interface{}

	// Visit a parse tree produced by CsqlParser#limitValue.
	VisitLimitValue(ctx *LimitValueContext) interface{}

	// Visit a parse tree produced by CsqlParser#orderBy.
	VisitOrderBy(ctx *OrderByContext) interface{}

	// Visit a parse tree produced by CsqlParser#orderByField.
	VisitOrderByField(ctx *OrderByFieldContext) interface{}

	// Visit a parse tree produced by CsqlParser#fieldIndex.
	VisitFieldIndex(ctx *FieldIndexContext) interface{}

	// Visit a parse tree produced by CsqlParser#groupBy.
	VisitGroupBy(ctx *GroupByContext) interface{}

	// Visit a parse tree produced by CsqlParser#groupByField.
	VisitGroupByField(ctx *GroupByFieldContext) interface{}

	// Visit a parse tree produced by CsqlParser#function.
	VisitFunction(ctx *FunctionContext) interface{}
}
