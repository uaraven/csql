package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/uaraven/csql/parser"
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

func (e *Errors) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ee antlr.RecognitionException) {
	e.errors = append(e.errors, ParsingError{
		line:    line,
		column:  column,
		message: msg,
	})
}

func (e *Errors) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}
func (e *Errors) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}
func (e *Errors) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

// ParseSQL reads the sql statement and converts it into an AST
func ParseSQL(sql string) Select {
	fs := antlr.NewInputStream(sql)
	lex := parser.NewCsqlLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCsqlParser(tokens)
	errors := newErrors()
	p.AddErrorListener(errors)
	visitor := NewCsqlVisitor()
	tree := p.Query()
	if len(errors.errors) != 0 {
		var sb strings.Builder
		for _, err := range errors.errors {
			sb.WriteString(fmt.Sprintf("[%d, %d] %s\n", err.line, err.column, err.message))
		}
		panic(sb.String())
	}
	return visitor.Visit(tree).(Select)
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
	return ctx.SelectStatement().Accept(c).(Select)
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
	return astSelect
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
	result := make([]Literal, 0)
	for _, item := range ctx.AllLiteralValue() {
		result = append(result, item.Accept(c).(Literal))
	}
	return ListLiteral{
		Values: result,
	}
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

// VisitCondition visits a parse tree produced by CsqlParser#condition.
func (c CsqlVisitorImpl) VisitCondition(ctx *parser.ConditionContext) interface{} {
	return ComparisonExpression{
		LHS:      ctx.Expr(0).Accept(c).(Expression),
		RHS:      ctx.Expr(1).Accept(c).(Expression),
		Operator: ctx.ComparisonOperator().Accept(c).(string),
	}
}

func (c CsqlVisitorImpl) VisitEvaluation(ctx *parser.EvaluationContext) interface{} {
	return BinaryExpression{
		LHS:      ctx.Expr(0).Accept(c).(Expression),
		RHS:      ctx.Expr(1).Accept(c).(Expression),
		Operator: ctx.BinaryOperation().Accept(c).(string),
	}
}

// VisitBetweenExpr visits a parse tree produced by CsqlParser#betweenExpr.
func (c CsqlVisitorImpl) VisitBetweenExpr(ctx *parser.BetweenExprContext) interface{} {
	return BetweenExpression{
		What: ctx.Expr(0).Accept(c).(Expression),
		Low:  ctx.Expr(1).Accept(c).(Expression),
		High: ctx.Expr(2).Accept(c).(Expression),
	}
}

func (c CsqlVisitorImpl) VisitTermItem(ctx *parser.TermItemContext) interface{} {
	return ctx.Term().Accept(c).(Term)
}

// VisitNotExpr visits a parse tree produced by CsqlParser#notExpr.
func (c CsqlVisitorImpl) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return NotExpression{
		Child: ctx.Expr().Accept(c).(Expression),
	}
}

// VisitInExpr visits a parse tree produced by CsqlParser#inExpr.
func (c CsqlVisitorImpl) VisitInExpr(ctx *parser.InExprContext) interface{} {
	return InListExpression{
		What:  ctx.Expr().Accept(c).(Expression),
		List:  ctx.List().Accept(c).(ListLiteral),
		NotIn: ctx.K_NOT() != nil,
	}
}

func (c CsqlVisitorImpl) VisitIsNullExpr(ctx *parser.IsNullExprContext) interface{} {
	return &IsNullExpression{
		What: ctx.Expr().Accept(c).(Expression),
		Not:  ctx.K_NOT() != nil,
	}
}

// VisitOrExpr visits a parse tree produced by CsqlParser#orExpr.
func (c CsqlVisitorImpl) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	return OrExpression{
		LHS: ctx.Expr(0).Accept(c).(Expression),
		RHS: ctx.Expr(1).Accept(c).(Expression),
	}
}

// VisitParensExpr visits a parse tree produced by CsqlParser#parensExpr.
func (c CsqlVisitorImpl) VisitParensExpr(ctx *parser.ParensExprContext) interface{} {
	return ctx.Expr().Accept(c).(Expression)
}

// VisitAndExpr visits a parse tree produced by CsqlParser#andExpr.
func (c CsqlVisitorImpl) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	return AndExpression{
		LHS: ctx.Expr(0).Accept(c).(Expression),
		RHS: ctx.Expr(1).Accept(c).(Expression),
	}
}

// VisitLikeExpr visits a parse tree produced by CsqlParser#likeExpr.
func (c CsqlVisitorImpl) VisitLikeExpr(ctx *parser.LikeExprContext) interface{} {
	return LikeExpression{
		What:    ctx.Expr().Accept(c).(Expression),
		Pattern: removeStringQuotes(ctx.StringValue().GetText()),
		NotLike: ctx.K_NOT() != nil,
	}
}

// VisitWhere visits a parse tree produced by CsqlParser#where.
func (c CsqlVisitorImpl) VisitWhere(ctx *parser.WhereContext) interface{} {
	return ctx.Expr().Accept(c).(Expression)
}

// VisitProjectionField visits a parse tree produced by CsqlParser#projectionField.
func (c CsqlVisitorImpl) VisitProjectionField(ctx *parser.ProjectionFieldContext) interface{} {
	var projectionField ProjectionField
	if ctx.ProjectionFieldName() != nil {
		field := ctx.ProjectionFieldName().Accept(c).(NamedProjectionField)
		if ctx.Alias() != nil {
			alias := ctx.Alias().Accept(c).(Identifier)
			field.Alias = &alias
		} else {
			field.Alias = nil
		}
		projectionField.NamedField = &field
	} else {
		expr := ctx.EvaluatedExpression().Accept(c).(Expression)
		projectionField.EvaluatedField = &EvaluatedProjectionField{
			Expr: expr,
		}
	}
	return projectionField
}

// VisitProjectionFieldName visits a parse tree produced by CsqlParser#projectionFieldName.
func (c CsqlVisitorImpl) VisitProjectionFieldName(ctx *parser.ProjectionFieldNameContext) interface{} {
	field := NamedProjectionField{}
	if ctx.Qualifier() != nil {
		tableName := ctx.Qualifier().Accept(c).(Identifier)
		field.TableName = &tableName
	} else {
		field.TableName = nil
	}
	field.Name = ctx.FieldName().Accept(c).(Identifier)
	return field
}

// VisitFieldName visits a parse tree produced by CsqlParser#fieldName.
func (c CsqlVisitorImpl) VisitFieldName(ctx *parser.FieldNameContext) interface{} {
	return parseIdentifier(ctx.GetText())
}

func (c CsqlVisitorImpl) VisitEvalTerm(ctx *parser.EvalTermContext) interface{} {
	return ctx.Term().Accept(c).(Term)
}

func (c CsqlVisitorImpl) VisitEvalBinaryExpression(ctx *parser.EvalBinaryExpressionContext) interface{} {
	return BinaryExpression{
		LHS:      ctx.Expr(0).Accept(c).(Expression),
		RHS:      ctx.Expr(1).Accept(c).(Expression),
		Operator: ctx.BinaryOperation().GetText(),
	}
}

func (c CsqlVisitorImpl) VisitEvalParens(ctx *parser.EvalParensContext) interface{} {
	return ParensExpression{
		Child: ctx.EvaluatedExpression().Accept(c).(Expression),
	}
}

func (c CsqlVisitorImpl) VisitMatchExpr(ctx *parser.MatchExprContext) interface{} {
	return MatchExpression{
		What:    ctx.Expr().Accept(c).(Expression),
		Pattern: removeStringQuotes(ctx.StringValue().GetText()),
		Not:     ctx.K_NOT() != nil,
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

	if ctx.Expr() != nil {
		e := ctx.Expr().Accept(c).(Expression)
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
	l := Literal{}
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
	result := CompoundName{}
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

// VisitInnerJoin visits a parse tree produced by CsqlParser#innerJoin.
func (c CsqlVisitorImpl) VisitInnerJoin(ctx *parser.InnerJoinContext) interface{} {
	return InnerJoin
}

// VisitLeftJoin visits a parse tree produced by CsqlParser#leftJoin.
func (c CsqlVisitorImpl) VisitLeftJoin(ctx *parser.LeftJoinContext) interface{} {
	return LeftOuterJoin
}

// VisitRightJoin visits a parse tree produced by CsqlParser#rightJoin.
func (c CsqlVisitorImpl) VisitRightJoin(ctx *parser.RightJoinContext) interface{} {
	return RightOuterJoin
}

// VisitCrossJoin visits a parse tree produced by CsqlParser#crossJoin.
func (c CsqlVisitorImpl) VisitCrossJoin(ctx *parser.CrossJoinContext) interface{} {
	return CrossJoin
}
