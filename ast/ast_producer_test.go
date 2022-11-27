package ast

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSimpleQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table")

	g.Expect(s.Projection.Distinct).To(BeFalse())
	g.Expect(s.Projection.ProjectionFields).To(HaveLen(1))
	g.Expect(s.Projection.ProjectionFields[0].Name).To(Equal(Identifier("col")))

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

func TestSimpleWhereLikeExpression(t *testing.T) {
	g := NewGomegaWithT(t)
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT LIKE 'abc%'")
	g.Expect(s.Filter).ToNot(BeNil())
	g.Expect(s.Filter).To(BeEquivalentTo(LikeExpression{
		What:    Term{Name: &CompoundName{Name: Identifier("col2")}},
		Pattern: "abc%",
		NotLike: true,
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
			Values: []Literal{{NumericValue: &v1}, {NumericValue: &v2}, {NumericValue: &v3}},
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
