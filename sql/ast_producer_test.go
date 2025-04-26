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
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/uaraven/csql/errors"
)

func TestSimpleQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	us := ParseSQL("SELECT col FROM TTable")

	g.Expect(us.Select).ToNot(BeNil())
	g.Expect(us.Union).To(BeNil())

	g.Expect(us.DropTmp).To(BeNil())

	s := us.Select

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.From.TableName).ToNot(BeNil())
	g.Expect(s.From.TableName.Name).To(Equal(Identifier("TTable")))
	g.Expect(s.From.TableName.Alias).To(BeNil())

}

func TestSimpleQueryQuotedIdentifier(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"col\" FROM \"Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	g.Expect(s.Select.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryQuotedQualifiedIdentifier(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"t\".\"col\" FROM \"Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	g.Expect(s.Select.Projection.ProjectionFields).To(HaveLen(1))
	identifier := Identifier("t")
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField.Qualifier).To(Equal(&identifier))
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryWithQualifiedProjection(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT ATable.col FROM ATable")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	g.Expect(s.Select.Projection.ProjectionFields).To(HaveLen(1))
	tName := Identifier("ATable")
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField).To(BeEquivalentTo(&CompoundName{
		Location:  Loc(1, 7),
		Qualifier: &tName,
		Name:      Identifier("col"),
	}))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("ATable")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryWithNegativeLiteral(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col, -2, -1e-10 FROM \"Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	value := "-2"
	vf := "-1e-10"
	g.Expect(s.Select.Projection.ProjectionFields).To(BeEquivalentTo(
		[]ProjectionField{
			{NamedField: &CompoundName{Location: Loc(1, 7), Name: Identifier("col")}},
			{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{Location: Loc(1, 12), NumericValue: &value}}}},
			{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{Location: Loc(1, 16), NumericValue: &vf}}}},
		}))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())
}

func TestSimpleErrorQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(func() { ParseSQL("SELECT col FROM") }).To(Panic())
}

func TestSimpleQueryDistinct(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT distinct col FROM \"Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeTrue())
}

func TestLeftOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 left outer join Table2 ON A = B")

	g.Expect(s.Select.From.Join.Type).To(Equal(LeftOuterJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Condition).ToNot(BeNil())
}

func TestRightOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 right join Table2 ON A = B")

	g.Expect(s.Select.From.Join.Type).To(Equal(RightOuterJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Condition).ToNot(BeNil())
}

func TestImplicitCrossJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1, Table2, Table3")
	if s.Select.From.Join.Type != CrossJoin {
		t.Errorf("Expected join type to bo CROSS JOIN, Got: %v", s.Select.From.Join.Type)
	}

	g.Expect(s.Select.From.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.Select.From.Join.Target.Join.Source.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Target.Join.Target.TableName.Name).To(Equal(Identifier("Table3")))
}

func TestExplicitCrossJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 cross join Table2 cross join Table3")
	if s.Select.From.Join.Type != CrossJoin {
		t.Errorf("Expected join type to bo CROSS JOIN, Got: %v", s.Select.From.Join.Type)
	}

	g.Expect(s.Select.From.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.Select.From.Join.Target.Join.Source.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Target.Join.Target.TableName.Name).To(Equal(Identifier("Table3")))
}

func TestConditionalJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 RIGHT JOIN Table2 ON tcol1 >= tcol2")

	g.Expect(s.Select.From.Join.Type).To(Equal(RightOuterJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Condition).To(Equal(ComparisonExpression{
		LHS: Term{Name: &CompoundName{
			Qualifier: nil,
			Location:  Loc(1, 44),
			Name:      Identifier("tcol1"),
		}},
		RHS: Term{Name: &CompoundName{
			Location:  Loc(1, 53),
			Qualifier: nil,
			Name:      Identifier("tcol2"),
		}},
		Operator: ">=",
		Location: Loc(1, 44),
	}))
}

func TestInnerJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 JOIN Table2 ON tcol1 >= tcol2")

	g.Expect(s.Select.From.Join.Type).To(Equal(InnerJoin))
	g.Expect(s.Select.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.Select.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.Select.From.Join.Condition).To(Equal(ComparisonExpression{
		LHS: Term{Name: &CompoundName{
			Location:  Loc(1, 38),
			Qualifier: nil,
			Name:      Identifier("tcol1"),
		}},
		RHS: Term{Name: &CompoundName{
			Location:  Loc(1, 47),
			Qualifier: nil,
			Name:      Identifier("tcol2"),
		}},
		Operator: ">=",
		Location: Loc(1, 38),
	}))
}

func TestSimpleWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	val := "10"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(ComparisonExpression{
		LHS: Term{
			Name: &CompoundName{
				Name:     Identifier("col2"),
				Location: Loc(1, 29),
			},
		},
		RHS:      Term{Value: &Literal{Location: Loc(1, 36), NumericValue: &val}},
		Operator: ">",
		Location: Loc(1, 29),
	}))
}

func TestParensWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > (col1 - 10)")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	val := "10"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(ComparisonExpression{
		LHS: Term{
			Name: &CompoundName{
				Location: Loc(1, 29),
				Name:     Identifier("col2"),
			},
		},
		RHS: BinaryExpression{
			LHS:      Term{Name: &CompoundName{Location: Loc(1, 37), Name: Identifier("col1")}},
			RHS:      Term{Value: &Literal{Location: Loc(1, 44), NumericValue: &val}},
			Location: Loc(1, 37),
			Operator: "-",
		},
		Location: Loc(1, 29),
		Operator: ">",
	}))
}

func TestSimpleWhereLikeExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT LIKE 'abc%'")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	p := "abc%"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(LikeExpression{
		What:     Term{Name: &CompoundName{Location: Loc(1, 29), Name: Identifier("col2")}},
		Pattern:  Term{Value: &Literal{Location: Loc(1, 43), StringValue: &p}},
		NotLike:  true,
		Location: Loc(1, 29),
	}))
}

func TestLikeWithExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 LIKE 'abc%' + 'bcd'")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	v1 := "abc%"
	v2 := "bcd"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(LikeExpression{
		What: Term{Name: &CompoundName{Location: Loc(1, 29), Name: Identifier("col2")}},
		Pattern: BinaryExpression{LHS: Term{Value: &Literal{Location: Loc(1, 39), StringValue: &v1}},
			RHS:      Term{Value: &Literal{Location: Loc(1, 48), StringValue: &v2}},
			Operator: "+",
			Location: Loc(1, 39),
		},
		NotLike:  false,
		Location: Loc(1, 29),
	}))
}

func TestSimpleWhereMatchExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 MATCH 'abc.*'")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	p := "abc.*"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(MatchExpression{
		Location: Loc(1, 29),
		What:     Term{Name: &CompoundName{Location: Loc(1, 29), Name: Identifier("col2")}},
		Pattern:  Term{Value: &Literal{Location: Loc(1, 40), StringValue: &p}},
		Not:      false,
	}))
}

func TestSimpleAndOrWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10 and col2 < 20\n" +
		"OR col3 >= 0")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	v1 := "10"
	v2 := "20"
	v3 := "0"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(OrExpression{
		LHS: AndExpression{
			LHS: ComparisonExpression{
				LHS: Term{Name: &CompoundName{
					Name:     "col2",
					Location: Loc(1, 29),
				}},
				RHS: Term{Value: &Literal{
					Location:     Loc(1, 36),
					NumericValue: &v1,
				}},
				Location: Loc(1, 29),
				Operator: ">",
			},
			RHS: ComparisonExpression{
				LHS: Term{Name: &CompoundName{
					Location: Loc(1, 43),
					Name:     Identifier("col2")}},
				RHS: Term{Value: &Literal{
					Location:     Loc(1, 50),
					NumericValue: &v2}},
				Operator: "<",
				Location: Loc(1, 43),
			},
			Location: Loc(1, 29),
		},
		RHS: ComparisonExpression{
			LHS: Term{Name: &CompoundName{
				Location: Loc(2, 3), Name: Identifier("col3")}},
			RHS: Term{Value: &Literal{
				Location:     Loc(2, 11),
				NumericValue: &v3}},
			Operator: ">=",
			Location: Loc(2, 3),
		},
		Location: Loc(1, 29),
	}))
}

func TestInListExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT IN (1, 2, 3)")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	v1 := "1"
	v2 := "2"
	v3 := "3"
	g.Expect(s.Select.Filter).To(BeEquivalentTo(InListExpression{
		What:     Term{Name: &CompoundName{Location: Loc(1, 29), Name: Identifier("col2")}},
		NotIn:    true,
		Location: Loc(1, 29),
		List: ListLiteral{
			Values: []Expression{
				Term{Value: &Literal{Location: Loc(1, 42), NumericValue: &v1}},
				Term{Value: &Literal{Location: Loc(1, 45), NumericValue: &v2}},
				Term{Value: &Literal{Location: Loc(1, 48), NumericValue: &v3}},
			},
		},
	}))
}

func TestIsNullExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col1 IS NULL")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	g.Expect(s.Select.Filter).To(BeEquivalentTo(IsNullExpression{
		What: Term{Name: &CompoundName{Name: "col1", Location: Loc(1, 29)}},
		Not:  false,
	}))

	s = ParseSQL("SELECT col FROM Table1 WHERE col1 IS NOT NULL")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	g.Expect(s.Select.Filter).To(BeEquivalentTo(IsNullExpression{
		What: Term{Name: &CompoundName{Location: Loc(1, 29), Name: "col1"}},
		Not:  true,
	}))

	s = ParseSQL("SELECT col FROM Table1 WHERE NOT col1 IS NULL")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	g.Expect(s.Select.Filter).To(BeEquivalentTo(NotExpression{
		Location: Loc(1, 29),
		Child: IsNullExpression{
			What: Term{Name: &CompoundName{Location: Loc(1, 33), Name: "col1"}},
			Not:  false,
		},
	}))
}

func TestBetweenExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 BETWEEN 10 AND col1")
	g.Expect(s.Select.Filter).ToNot(BeNil())
	v1 := "10"

	g.Expect(s.Select.Filter).To(BeEquivalentTo(BetweenExpression{
		What:     Term{Name: &CompoundName{Name: "col2", Location: Loc(1, 29)}},
		Low:      Term{Value: &Literal{NumericValue: &v1, Location: Loc(1, 42)}},
		High:     Term{Name: &CompoundName{Name: "col1", Location: Loc(1, 49)}},
		Location: Loc(1, 29),
	}))

}

func TestSelectEvaluation(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT 1+2 FROM Table1")
	g.Expect(s.Select.Filter).To(BeNil())
	v1 := "1"
	v2 := "2"
	g.Expect(s.Select.Projection).To(BeEquivalentTo(SelectProjection{
		Distinct: false,
		ProjectionFields: []ProjectionField{
			{
				EvaluatedField: &EvaluatedProjectionField{
					Expr: BinaryExpression{
						LHS:      Term{Value: &Literal{NumericValue: &v1, Location: Loc(1, 7)}},
						RHS:      Term{Value: &Literal{NumericValue: &v2, Location: Loc(1, 9)}},
						Operator: "+",
						Location: Loc(1, 7),
					},
				},
			},
		},
	}))
}

func TestSelectLiteralValue(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT 1,2.0, '3' FROM Table1")
	v1 := "1"
	v2 := "2.0"
	v3 := "3"

	g.Expect(s.Select.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{NumericValue: &v1, Location: Loc(1, 7)}}}},
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{NumericValue: &v2, Location: Loc(1, 9)}}}},
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{StringValue: &v3, Location: Loc(1, 14)}}}},
	}))
}

func TestSelectParenExpr(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT (1 / 2) FROM Table1")
	v1 := "1"
	v2 := "2"

	g.Expect(s.Select.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{
			EvaluatedField: &EvaluatedProjectionField{Expr: BinaryExpression{
				Location: Loc(1, 8),
				LHS:      Term{Value: &Literal{NumericValue: &v1, Location: Loc(1, 8)}},
				RHS:      Term{Value: &Literal{NumericValue: &v2, Location: Loc(1, 12)}},
				Operator: "/",
			},
			},
		},
	}))
}

func TestSelectTableAlias(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT T.col as tcol FROM Table1 T")

	tableAlias := Identifier("T")
	g.Expect(s.Select.From).To(BeEquivalentTo(DataSource{TableName: &SourceName{Name: "Table1", Alias: &tableAlias}}))

	renamedCol := Identifier("tcol")
	g.Expect(s.Select.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{NamedField: &CompoundName{Qualifier: &tableAlias, Name: Identifier("col"), Location: Loc(1, 7)}, Alias: &renamedCol},
	}))
}

func TestSelectNull(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT NULL FROM Table1 T")

	g.Expect(s.Select.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{
			EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{IsNull: true, Location: Loc(1, 7)}}},
		},
	}))
}

func TestSelectLimit(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT NULL FROM Table1")
	g.Expect(s.Select.Limit).To(Equal(int32(0)))

	s = ParseSQL("SELECT NULL FROM Table1 LIMIT 10")

	g.Expect(s.Select.Limit).To(Equal(int32(10)))

	g.Expect(func() {
		ParseSQL("SELECT NULL FROM Table1 LIMIT '10'")
	}).To(Panic())
}

func TestSelectOrderBy(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT a, b, c FROM Table1 ORDER BY a desc, 2, Table1.c")
	q := Identifier("Table1")
	g.Expect(s.Select.OrderBy).To(BeEquivalentTo(&OrderByExpression{
		Location: Loc(1, 27),
		OrderFields: []OrderByField{
			{Location: Loc(1, 36), FieldName: &CompoundName{Name: Identifier("a"), Location: Loc(1, 36)}, FieldIndex: 0, Descending: true},
			{Location: Loc(1, 44), FieldName: nil, FieldIndex: 2, Descending: false},
			{Location: Loc(1, 47), FieldName: &CompoundName{Qualifier: &q, Name: Identifier("c"), Location: Loc(1, 47)}, FieldIndex: 0, Descending: false},
		}}))
}

func TestSelectWithFunctions(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT len('abc') FROM Table1 WHERE TO_STRING(b) != TRUNC(13.2)")

	v1 := "abc"
	v2 := "13.2"
	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: FunctionCall{
			function: "len",
			args:     []Expression{Term{Value: &Literal{StringValue: &v1, Location: Loc(1, 11)}}},
			Location: Loc(1, 7),
		}}}))
	g.Expect(s.Select.Filter).To(BeEquivalentTo(
		ComparisonExpression{
			LHS: FunctionCall{
				function: "to_string",
				args: []Expression{Term{Name: &CompoundName{
					Name:     "b",
					Location: Loc(1, 46),
				}}},
				Location: Loc(1, 36),
			},
			RHS: FunctionCall{
				function: "trunc",
				args:     []Expression{Term{Value: &Literal{NumericValue: &v2, Location: Loc(1, 58)}}},
				Location: Loc(1, 52),
			},
			Operator: "!=",
			Location: Loc(1, 36),
		},
	))
}

func TestSelectWithCountFunctions(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT count(distinct a) FROM Table1")

	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: CountFunctionCall{
			distinct: true,
			arg: CompoundName{
				Name:     "a",
				Location: Loc(1, 22),
			},
			Location: Loc(1, 7),
		}}}))
}

func TestSelectWithCountAll(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT count(*) FROM Table1")

	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: CountFunctionCall{
			distinct: false,
			arg: CompoundName{
				Name:     "*",
				Location: Loc(1, 13),
			},
			Location: Loc(1, 7),
		}}}))
}

func TestSelectWithAvg(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT avg(abc) FROM Table1")

	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: AggregateFunctionCall{
			function: "avg",
			arg: Term{
				Name: &CompoundName{
					Name:     "abc",
					Location: Loc(1, 11),
				}},
			Location: Loc(1, 7),
		}}}))
}

func TestSelectWithAvgGroupBy(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT avg(abc) FROM Table1 GROUP BY bc, 1")

	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: AggregateFunctionCall{
			function: "avg",
			arg: Term{
				Name: &CompoundName{
					Name:     "abc",
					Location: Loc(1, 11),
				}},
			Location: Loc(1, 7),
		}}}))

	g.Expect(s.Select.GroupBy).To(BeEquivalentTo(&GroupByExpression{
		Location: Loc(1, 28),
		GroupFields: []GroupByField{
			{
				FieldName: &CompoundName{
					Name:     "bc",
					Location: Loc(1, 37),
				},
				FieldIndex: 0,
				Location:   Loc(1, 37),
			},
			{
				FieldName:  nil,
				FieldIndex: 1,
				Location:   Loc(1, 41),
			},
		},
	}))
}

func TestIntoClause(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"col\" FROM \"Table\" INTO \"Other-Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	g.Expect(s.Select.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())

	g.Expect(s.Into).ToNot(BeNil())
	g.Expect(s.Into.Destination).To(Equal(Identifier("Other-Table.csv")))
	g.Expect(s.Into.TempTable).To(BeFalse())
}

func TestIntoTempClause(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"col\" FROM \"Table\" INTO TEMP \"Other-Table\"")

	g.Expect(s.Select.Projection.Distinct).To(BeFalse())
	g.Expect(s.Select.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Select.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.Select.From.TableName).ToNot(BeNil())
	g.Expect(s.Select.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.Select.From.TableName.Alias).To(BeNil())

	g.Expect(s.Into).ToNot(BeNil())
	g.Expect(s.Into.Destination).To(Equal(Identifier("Other-Table")))
	g.Expect(s.Into.TempTable).To(BeTrue())
}

func TestDropTempTable(t *testing.T) {
	g := NewGomegaWithT(t)
	us := ParseSQL("DROP TEMP TABLE TTable")

	g.Expect(us.Select).To(BeNil())
	g.Expect(us.Union).To(BeNil())

	g.Expect(us.DropTmp).ToNot(BeNil())
	g.Expect(us.DropTmp.Name).To(Equal(Identifier("TTable")))
}

func TestSelectWithAvgInExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT avg(abc) / 2 FROM Table1")

	v1 := "2"
	g.Expect(s.Select.Projection.ProjectionFields[0]).To(BeEquivalentTo(ProjectionField{
		EvaluatedField: &EvaluatedProjectionField{Expr: BinaryExpression{
			LHS: AggregateFunctionCall{
				function: "avg",
				arg: Term{
					Name: &CompoundName{
						Name:     "abc",
						Location: Loc(1, 11),
					}},
				Location: Loc(1, 7),
			},
			RHS: Term{Value: &Literal{
				Location:     Loc(1, 18),
				NumericValue: &v1,
			}},
			Operator: "/",
			Location: Loc(1, 7),
		}}}))
}

func TestSelectWithHaving(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT * FROM Table1 GROUP BY bc HAVING min(a) > 10 and b = 's'")

	v1 := "10"
	v2 := "s"
	g.Expect(s.Select.Having).ToNot(BeNil())
	g.Expect(s.Select.Having).To(BeEquivalentTo(
		AndExpression{
			LHS: ComparisonExpression{
				Operator: ">",
				LHS: AggregateFunctionCall{
					function: "min",
					arg:      Term{Name: &CompoundName{Name: "a", Location: Loc(1, 44)}},
					Location: Loc(1, 40),
				},
				RHS: Term{
					Value: &Literal{NumericValue: &v1, Location: Loc(1, 49)},
				},
				Location: Loc(1, 40),
			},
			RHS: ComparisonExpression{
				LHS:      Term{Name: &CompoundName{Name: "b", Location: Loc(1, 56)}},
				RHS:      Term{Value: &Literal{StringValue: &v2, Location: Loc(1, 60)}},
				Operator: "=",
				Location: Loc(1, 56),
			},
			Location: Loc(1, 40),
		},
	))
}
