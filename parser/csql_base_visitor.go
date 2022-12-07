// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Csql

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseCsqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCsqlVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitBinaryOperation(ctx *BinaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitValueTerm(ctx *ValueTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitValueBinaryExpr(ctx *ValueBinaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitValueParensExpr(ctx *ValueParensExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitTermItem(ctx *TermItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitInExpr(ctx *InExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitIsNullExpr(ctx *IsNullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitLikeExpr(ctx *LikeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitParensExpr(ctx *ParensExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitValueExprItem(ctx *ValueExprItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitBetweenExpr(ctx *BetweenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitWhere(ctx *WhereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitDistinct(ctx *DistinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitProjection(ctx *ProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitProjectionField(ctx *ProjectionFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitProjectionFieldName(ctx *ProjectionFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitLeftJoin(ctx *LeftJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitRightJoin(ctx *RightJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitFullJoin(ctx *FullJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitCrossJoin(ctx *CrossJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitConditionalJoinType(ctx *ConditionalJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitConditionalJoinTarget(ctx *ConditionalJoinTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitUnconditionalJoinTarget(ctx *UnconditionalJoinTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitDataSourceConditionalJoin(ctx *DataSourceConditionalJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitDataSourceItem(ctx *DataSourceItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitDataSourceCrossJoin(ctx *DataSourceCrossJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitDataSourceNamed(ctx *DataSourceNamedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitSources(ctx *SourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitUnionSelects(ctx *UnionSelectsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitSignedNumber(ctx *SignedNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitNullValue(ctx *NullValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitLiteralValue(ctx *LiteralValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitSourceName(ctx *SourceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitCompoundName(ctx *CompoundNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitQualifier(ctx *QualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitLimit(ctx *LimitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitLimitValue(ctx *LimitValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitOrderBy(ctx *OrderByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitOrderByField(ctx *OrderByFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCsqlVisitor) VisitFieldIndex(ctx *FieldIndexContext) interface{} {
	return v.VisitChildren(ctx)
}
