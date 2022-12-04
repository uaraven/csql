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
	. "github.com/onsi/gomega"
	"testing"
)

func TestSimpleQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table")

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.From.TableName).ToNot(BeNil())
	g.Expect(s.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryQuotedIdentifier(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"col\" FROM \"Table\"")

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.From.TableName).ToNot(BeNil())
	g.Expect(s.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryQuotedQualifiedIdentifier(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT \"t\".\"col\" FROM \"Table\"")

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	identifier := Identifier("t")
	g.Expect(s.Projection.ProjectionFields[0].NamedField.TableName).To(Equal(&identifier))
	g.Expect(s.Projection.ProjectionFields[0].NamedField.Name).To(Equal(Identifier("col")))

	g.Expect(s.From.TableName).ToNot(BeNil())
	g.Expect(s.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.From.TableName.Alias).To(BeNil())
}

func TestSimpleQueryWithQualifiedProjection(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT Table.col FROM Table")

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	tName := Identifier("Table")
	g.Expect(s.Projection.ProjectionFields[0].NamedField).To(BeEquivalentTo(&NamedProjectionField{
		TableName: &tName,
		Name:      Identifier("col"),
	}))

	g.Expect(s.From.TableName).ToNot(BeNil())
	g.Expect(s.From.TableName.Name).To(Equal(Identifier("Table")))
	g.Expect(s.From.TableName.Alias).To(BeNil())
}

func TestSimpleErrorQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(func() { ParseSQL("SELECT col FROM") }).To(Panic())
}

func TestSimpleQueryDistinct(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT distinct col FROM Table")

	g.Expect(s.Projection.Distinct).To(BeTrue())
}

func TestLeftOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 left outer join Table2 ON A = B")

	g.Expect(s.From.Join.Type).To(Equal(LeftOuterJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Condition).ToNot(BeNil())
}

func TestRightOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 right join Table2 ON A = B")

	g.Expect(s.From.Join.Type).To(Equal(RightOuterJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Condition).ToNot(BeNil())
}

func TestImplicitCrossJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1, Table2, Table3")
	if s.From.Join.Type != CrossJoin {
		t.Errorf("Expected join type to bo CROSS JOIN, Got: %v", s.From.Join.Type)
	}

	g.Expect(s.From.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.From.Join.Target.Join.Source.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Target.Join.Target.TableName.Name).To(Equal(Identifier("Table3")))
}

func TestExplicitCrossJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 cross join Table2 cross join Table3")
	if s.From.Join.Type != CrossJoin {
		t.Errorf("Expected join type to bo CROSS JOIN, Got: %v", s.From.Join.Type)
	}

	g.Expect(s.From.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.Join.Type).To(Equal(CrossJoin))
	g.Expect(s.From.Join.Target.Join.Source.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Target.Join.Target.TableName.Name).To(Equal(Identifier("Table3")))
}

func TestConditionalJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 RIGHT JOIN Table2 ON tcol1 >= tcol2")

	g.Expect(s.From.Join.Type).To(Equal(RightOuterJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Condition).To(Equal(ComparisonExpression{
		LHS: Term{Name: &CompoundName{
			Qualifier: nil,
			Name:      Identifier("tcol1"),
		}},
		RHS: Term{Name: &CompoundName{
			Qualifier: nil,
			Name:      Identifier("tcol2"),
		}},
		Operator: ">=",
	}))
}

func TestInnerJoin(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 JOIN Table2 ON tcol1 >= tcol2")

	g.Expect(s.From.Join.Type).To(Equal(InnerJoin))
	g.Expect(s.From.Join.Source.TableName.Name).To(Equal(Identifier("Table1")))
	g.Expect(s.From.Join.Target.TableName.Name).To(Equal(Identifier("Table2")))
	g.Expect(s.From.Join.Condition).To(Equal(ComparisonExpression{
		LHS: Term{Name: &CompoundName{
			Qualifier: nil,
			Name:      Identifier("tcol1"),
		}},
		RHS: Term{Name: &CompoundName{
			Qualifier: nil,
			Name:      Identifier("tcol2"),
		}},
		Operator: ">=",
	}))
}

func TestSimpleWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10")
	g.Expect(s.Filter).ToNot(BeNil())
	val := "10"
	g.Expect(s.Filter).To(BeEquivalentTo(ComparisonExpression{
		LHS: Term{
			Name: &CompoundName{
				Name: Identifier("col2"),
			},
		},
		RHS:      Term{Value: &Literal{NumericValue: &val}},
		Operator: ">",
	}))
}

func TestParensWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > (col1 - 10)")
	g.Expect(s.Filter).ToNot(BeNil())
	val := "10"
	g.Expect(s.Filter).To(BeEquivalentTo(ComparisonExpression{
		LHS: Term{
			Name: &CompoundName{
				Name: Identifier("col2"),
			},
		},
		RHS: BinaryExpression{
			LHS:      Term{Name: &CompoundName{Name: Identifier("col1")}},
			RHS:      Term{Value: &Literal{NumericValue: &val}},
			Operator: "-",
		},
		Operator: ">",
	}))
}

func TestSimpleWhereLikeExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT LIKE 'abc%'")
	g.Expect(s.Filter).ToNot(BeNil())
	p := "abc%"
	g.Expect(s.Filter).To(BeEquivalentTo(LikeExpression{
		What:    Term{Name: &CompoundName{Name: Identifier("col2")}},
		Pattern: Term{Value: &Literal{StringValue: &p}},
		NotLike: true,
	}))
}

func TestLikeWithExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 LIKE 'abc%' + 'bcd'")
	g.Expect(s.Filter).ToNot(BeNil())
	v1 := "abc%"
	v2 := "bcd"
	g.Expect(s.Filter).To(BeEquivalentTo(LikeExpression{
		What: Term{Name: &CompoundName{Name: Identifier("col2")}},
		Pattern: BinaryExpression{LHS: Term{Value: &Literal{StringValue: &v1}},
			RHS:      Term{Value: &Literal{StringValue: &v2}},
			Operator: "+",
		},
		NotLike: false,
	}))
}

func TestSimpleWhereMatchExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 MATCH 'abc.*'")
	g.Expect(s.Filter).ToNot(BeNil())
	p := "abc.*"
	g.Expect(s.Filter).To(BeEquivalentTo(MatchExpression{
		What:    Term{Name: &CompoundName{Name: Identifier("col2")}},
		Pattern: Term{Value: &Literal{StringValue: &p}},
		Not:     false,
	}))
}

func TestSimpleAndOrWhereExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10 and col2 < 20 OR col3 >= 0")
	g.Expect(s.Filter).ToNot(BeNil())
	v1 := "10"
	v2 := "20"
	v3 := "0"
	g.Expect(s.Filter).To(BeEquivalentTo(OrExpression{
		LHS: AndExpression{
			LHS: ComparisonExpression{
				LHS: Term{Name: &CompoundName{
					Name: "col2",
				}},
				RHS: Term{Value: &Literal{
					NumericValue: &v1,
				}},
				Operator: ">",
			},
			RHS: ComparisonExpression{
				LHS:      Term{Name: &CompoundName{Name: Identifier("col2")}},
				RHS:      Term{Value: &Literal{NumericValue: &v2}},
				Operator: "<",
			},
		},
		RHS: ComparisonExpression{
			LHS:      Term{Name: &CompoundName{Name: Identifier("col3")}},
			RHS:      Term{Value: &Literal{NumericValue: &v3}},
			Operator: ">=",
		},
	}))
}

func TestInListExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT IN (1, 2, 3)")
	g.Expect(s.Filter).ToNot(BeNil())
	v1 := "1"
	v2 := "2"
	v3 := "3"
	g.Expect(s.Filter).To(BeEquivalentTo(InListExpression{
		What:  Term{Name: &CompoundName{Name: Identifier("col2")}},
		NotIn: true,
		List: ListLiteral{
			Values: []Expression{
				Term{Value: &Literal{NumericValue: &v1}},
				Term{Value: &Literal{NumericValue: &v2}},
				Term{Value: &Literal{NumericValue: &v3}},
			},
		},
	}))
}

func TestIsNullExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col1 IS NULL")
	g.Expect(s.Filter).ToNot(BeNil())
	g.Expect(s.Filter).To(BeEquivalentTo(IsNullExpression{
		What: Term{Name: &CompoundName{Name: "col1"}},
		Not:  false,
	}))

	s = ParseSQL("SELECT col FROM Table1 WHERE col1 IS NOT NULL")
	g.Expect(s.Filter).ToNot(BeNil())
	g.Expect(s.Filter).To(BeEquivalentTo(IsNullExpression{
		What: Term{Name: &CompoundName{Name: "col1"}},
		Not:  true,
	}))

	s = ParseSQL("SELECT col FROM Table1 WHERE NOT col1 IS NULL")
	g.Expect(s.Filter).ToNot(BeNil())
	g.Expect(s.Filter).To(BeEquivalentTo(NotExpression{
		Child: IsNullExpression{
			What: Term{Name: &CompoundName{Name: "col1"}},
			Not:  false,
		},
	}))
}

func TestBetweenExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 BETWEEN 10 AND col1")
	g.Expect(s.Filter).ToNot(BeNil())
	v1 := "10"

	g.Expect(s.Filter).To(BeEquivalentTo(BetweenExpression{
		What: Term{Name: &CompoundName{Name: "col2"}},
		Low:  Term{Value: &Literal{NumericValue: &v1}},
		High: Term{Name: &CompoundName{Name: "col1"}},
	}))

}

func TestSelectEvaluation(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT 1+2 FROM Table1")
	g.Expect(s.Filter).To(BeNil())
	v1 := "1"
	v2 := "2"
	g.Expect(s.Projection).To(BeEquivalentTo(SelectProjection{
		Distinct: false,
		ProjectionFields: []ProjectionField{
			{
				EvaluatedField: &EvaluatedProjectionField{
					Expr: BinaryExpression{
						LHS:      Term{Value: &Literal{NumericValue: &v1}},
						RHS:      Term{Value: &Literal{NumericValue: &v2}},
						Operator: "+",
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

	g.Expect(s.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{NumericValue: &v1}}}},
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{NumericValue: &v2}}}},
		{EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{StringValue: &v3}}}},
	}))
}

func TestSelectParenExpr(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT (1 / 2) FROM Table1")
	v1 := "1"
	v2 := "2"

	g.Expect(s.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{
			EvaluatedField: &EvaluatedProjectionField{Expr: BinaryExpression{
				LHS:      Term{Value: &Literal{NumericValue: &v1}},
				RHS:      Term{Value: &Literal{NumericValue: &v2}},
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
	g.Expect(s.From).To(BeEquivalentTo(DataSource{TableName: &SourceName{Name: "Table1", Alias: &tableAlias}}))

	renamedCol := Identifier("tcol")
	g.Expect(s.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{NamedField: &NamedProjectionField{TableName: &tableAlias, Name: Identifier("col")}, Alias: &renamedCol},
	}))
}

func TestSelectNull(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT NULL FROM Table1 T")

	g.Expect(s.Projection.ProjectionFields).To(BeEquivalentTo([]ProjectionField{
		{
			EvaluatedField: &EvaluatedProjectionField{Expr: Term{Value: &Literal{IsNull: true}}},
		},
	}))
}
