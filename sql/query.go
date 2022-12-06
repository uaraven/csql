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
	"github.com/uaraven/csql/funky"
	"strconv"
	"strings"
)

// Identifier is really any identifier
type Identifier string

// CompoundName is a fully qualified name (with optional qualifier) usually used to reference table columns
type CompoundName struct {
	// Qualifier is a table name or alias
	Qualifier *Identifier
	// Name is a field name
	Name Identifier
}

func (cn CompoundName) String() string {
	var sb strings.Builder
	if cn.Qualifier != nil {
		sb.WriteString(fmt.Sprintf("%s.", string(*cn.Qualifier)))
	}
	sb.WriteString(string(cn.Name))
	return sb.String()
}

// Select represents a SELECT query. Select consists of
// projection, Sources AKA "FROM" and Filter expressions,
// Limit and OrderBy.
type Select struct {
	Projection SelectProjection
	From       DataSource
	Filter     Expression
	Limit      int32 // 0 means no limit
	OrderBy    *OrderByExpression
}

func (sel Select) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("SELECT %v\nFROM %v", sel.Projection, sel.From))
	if sel.Filter != nil {
		sb.WriteString(fmt.Sprintf(" %v", sel.Filter))
	}
	return sb.String()
}

// SelectProjection is a projection of fields in resultset that will be returned from the Select
type SelectProjection struct {
	Distinct         bool
	ProjectionFields []ProjectionField
}

func (sp SelectProjection) String() string {
	var sb strings.Builder
	if sp.Distinct {
		sb.WriteString("DISTINCT ")
	}
	for i, pf := range sp.ProjectionFields {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(pf.String())
	}

	return sb.String()
}

type EvaluatedProjectionField struct {
	Expr Expression
}

func (epf EvaluatedProjectionField) String() string {
	return epf.Expr.String()
}

type ProjectionField struct {
	NamedField     *CompoundName
	EvaluatedField *EvaluatedProjectionField
	Alias          *Identifier
}

func (pf ProjectionField) String() string {
	var sb strings.Builder
	if pf.NamedField != nil {
		sb.WriteString(pf.NamedField.String())
	} else {
		sb.WriteString(pf.EvaluatedField.String())
	}
	if pf.Alias != nil {
		sb.WriteString(" AS ")
		sb.WriteString(string(*pf.Alias))
	}
	return sb.String()
}

type OrderByField struct {
	FieldName  *CompoundName
	FieldIndex int32 // 0 means no index
	Descending bool
}

func (obf OrderByField) String() string {
	var sb strings.Builder
	if obf.FieldName != nil {
		sb.WriteString(obf.FieldName.String())
	} else {
		sb.WriteString(strconv.FormatInt(int64(obf.FieldIndex), 10))
	}
	if obf.Descending {
		sb.WriteString(" DESC")
	}
	return sb.String()
}

type OrderByExpression struct {
	OrderFields []OrderByField
}

func (obe OrderByExpression) String() string {
	fieldNames := funky.Map(obe.OrderFields, func(o OrderByField) string {
		return o.String()
	})
	return "ORDER BY " + strings.Join(fieldNames, ", ")
}

// SourceName is a name of a Table/Data file with optional Alias
type SourceName struct {
	Name  Identifier
	Alias *Identifier
}

func (sn SourceName) String() string {
	var sb strings.Builder
	sb.WriteString(string(sn.Name))
	if sn.Alias != nil {
		sb.WriteString(" ")
		sb.WriteString(string(*sn.Alias))
	}
	return sb.String()
}

// DataSource is a source for data in the query. It can be either a Table represented by TableName
// or a result of a join represented as JoinedSource
type DataSource struct {
	TableName *SourceName
	Join      *JoinedSource
}

func (ds DataSource) String() string {
	if ds.TableName != nil {
		return ds.TableName.String()
	}
	return ds.Join.String()
}

// JoinType is just that
type JoinType int

const (
	// InvalidJoin represents any unsupported or erroneous JOIN type
	InvalidJoin JoinType = iota
	// InnerJoin represents inner join
	InnerJoin
	// LeftOuterJoin represents LEFT JOIN or LEFT OUTER JOIN
	LeftOuterJoin
	// RightOuterJoin represents RIGHT JOIN or RIGHT OUTER JOIN
	RightOuterJoin
	// CrossJoin represents CROSS JOIN, including implicit CROSS JOIN
	CrossJoin
	// FullJoin represents FULL OUTER JOIN
	FullJoin

	// InOperator is just that
	InOperator = "IN"
	// NotInOperator is just that
	NotInOperator = "NOT IN"
	// LikeOperator is just that
	LikeOperator = "LIKE"
	// NotLikeOperator is just that
	NotLikeOperator = "NOT LIKE"
	// MatchOperator is just that
	MatchOperator = "MATCH"
	// NotMatchOperator is just that
	NotMatchOperator = "NOT MATCH"
)

func (jt JoinType) String() string {
	switch jt {
	case InnerJoin:
		return "INNER JOIN"
	case LeftOuterJoin:
		return "LEFT OUTER JOIN"
	case RightOuterJoin:
		return "RIGHT OUTER JOIN"
	case CrossJoin:
		return "CROSS JOIN"
	default:
		return "*Unsupported Join*"
	}
}

// JoinedSource represents a DataSource that is a result of Join of two other DataSources
type JoinedSource struct {
	Source    DataSource
	Type      JoinType
	Target    DataSource
	Condition Expression
}

func (js JoinedSource) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%v\n%v %v", js.Source, js.Type, js.Target))
	if js.Condition != nil {
		sb.WriteString(fmt.Sprintf(" ON %v", js.Condition))
	}
	return sb.String()
}

// Literal represents a literal that can be one of:
// - NULL
// - Number
// - String
// - List
type Literal struct {
	IsNull       bool
	NumericValue *string
	StringValue  *string
}

func IntLiteral(i string) Literal {
	return Literal{NumericValue: &i}
}

func (lit Literal) String() string {
	if lit.IsNull {
		return "NULL"
	}
	if lit.NumericValue != nil {
		return *lit.NumericValue
	}
	return fmt.Sprintf("'%s'", *lit.StringValue)
}

type ListLiteral struct {
	Values []Expression
}

func (ll ListLiteral) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for i, item := range ll.Values {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(item.String())
	}
	sb.WriteString(")")
	return sb.String()
}

// Term is and expression element that is either a Name (reference to a table field) or a literal
type Term struct {
	Name  *CompoundName
	Value *Literal
}

func (t Term) String() string {
	if t.Name != nil {
		return t.Name.String()
	}
	return (*t.Value).String()
}

// Expression represents any expression
type Expression interface {
	String() string
}

type ParensExpression struct {
	Child Expression
}

func (pe ParensExpression) String() string {
	return fmt.Sprintf("(%v)", pe.Child)
}

type AndExpression struct {
	LHS Expression
	RHS Expression
}

func (ae AndExpression) String() string {
	return fmt.Sprintf("%v AND %v", ae.LHS, ae.RHS)
}

type OrExpression struct {
	LHS Expression
	RHS Expression
}

func (oe OrExpression) String() string {
	return fmt.Sprintf("%v OR %v", oe.LHS, oe.RHS)
}

type NotExpression struct {
	Child Expression
}

func (ne NotExpression) String() string {
	return fmt.Sprintf("NOT %v", ne.Child)
}

type ComparisonExpression struct {
	LHS      Expression
	RHS      Expression
	Operator string
}

func (ce ComparisonExpression) String() string {
	return fmt.Sprintf("%v %s %v", ce.LHS, ce.Operator, ce.RHS)
}

type BinaryExpression struct {
	LHS      Expression
	RHS      Expression
	Operator string
}

func (be BinaryExpression) String() string {
	return fmt.Sprintf("%v %s %v", be.LHS, be.Operator, be.RHS)
}

type LikeExpression struct {
	What    Expression
	Pattern Expression
	NotLike bool
}

func (like LikeExpression) String() string {
	var op string
	if like.NotLike {
		op = NotLikeOperator
	} else {
		op = LikeOperator
	}
	return fmt.Sprintf("%v %s '%s'", like.What, op, like.Pattern)
}

type MatchExpression struct {
	What    Expression
	Pattern Expression
	Not     bool
}

func (match MatchExpression) String() string {
	var op string
	if match.Not {
		op = NotMatchOperator
	} else {
		op = MatchOperator
	}
	return fmt.Sprintf("%v %s '%s'", match.What, op, match.Pattern)
}

type BetweenExpression struct {
	What Expression
	Low  Expression
	High Expression
}

func (be BetweenExpression) String() string {
	return fmt.Sprintf("%v BETWEEN %v AND %v", be.What, be.Low, be.High)
}

type InListExpression struct {
	What  Expression
	List  ListLiteral
	NotIn bool
}

func (ile InListExpression) String() string {
	var op string
	if ile.NotIn {
		op = NotInOperator
	} else {
		op = InOperator
	}
	return fmt.Sprintf("%v %s %v", ile.What, op, ile.List)
}

type IsNullExpression struct {
	What Expression
	Not  bool
}

func (ine IsNullExpression) String() string {
	var sb strings.Builder
	sb.WriteString(ine.What.String())
	sb.WriteRune(' ')
	sb.WriteString("IS ")
	if ine.Not {
		sb.WriteString("NOT ")
	}
	sb.WriteString("NULL")
	return sb.String()
}
