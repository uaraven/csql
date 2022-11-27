package ast

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSimpleQuery(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table")
	if s.Projection.Distinct {
		t.Error("Projection should not have DISTINCT flag")
	}
	if len(s.Projection.ProjectionFields) != 1 &&
		s.Projection.ProjectionFields[0].Name != "col" {
		t.Error("Failed to parse columns")
	}
	if s.From.TableName == nil {
		t.Errorf("Expected FROM clause to contain table")
	}
	if s.From.TableName.Name != "Table" {
		t.Errorf("Expected FROM clause to contain table named 'Table'")
	}
}

func TestSimpleErrorQuery(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(func() { ParseSQL("SELECT col FROM") }).To(Panic())
}

func TestSimpleQueryDistinct(t *testing.T) {
	s := ParseSQL("SELECT distinct col FROM Table")
	if !s.Projection.Distinct {
		t.Error("Projection should have DISTINCT flag")
	}
}

func TestInnerJoin(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 left outer join Table2 ON A = B")
	if s.From.Join.Type != LeftOuterJoin {
		t.Errorf("Expected join type to be 'LEFT OUTER'")
	}
	if s.From.Join.Source.TableName.Name != "Table1" {
		t.Errorf("Expected table name Table1, Got: %v", s.From.Join.Source.TableName.Name)
	}
	if s.From.Join.Target.TableName.Name != "Table2" {
		t.Errorf("Expected joining table name Table2, Got: %v", s.From.Join.Target.TableName.Name)
	}
	if s.From.Join.Condition == nil {
		t.Errorf("Expected condition on the join")
	}
}

func TestImplicitCrossJoin(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1, Table2, Table3")
	if s.From.Join.Type != CrossJoin {
		t.Errorf("Expected join type to bo CROSS JOIN, Got: %v", s.From.Join.Type)
	}
	if s.From.Join.Source.TableName.Name != "Table1" {
		t.Errorf("Expected table name Table1, Got: %v", s.From.Join.Source.TableName.Name)
	}
	if s.From.Join.Target.Join.Type != CrossJoin {
		t.Errorf("Expected second join type to bo CROSS JOIN, Got: %v", s.From.Join.Target.Join.Type)
	}
	if s.From.Join.Target.Join.Source.TableName.Name != "Table2" {
		t.Errorf("Expected joining table name Table2, Got: %v", s.From.Join.Target.Join.Source.TableName.Name)
	}
	if s.From.Join.Target.Join.Target.TableName.Name != "Table3" {
		t.Errorf("Expected joining table name Table3, Got: %v", s.From.Join.Target.Join.Target.TableName.Name)
	}
}

func TestConditionalJoin(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 RIGHT JOIN Table2 ON tcol1 >= tcol2")
	if s.From.Join.Type != RightOuterJoin {
		t.Errorf("Expected join type to bo RIGHT JOIN, Got: %v", s.From.Join.Type)
	}
	if s.From.Join.Source.TableName.Name != "Table1" {
		t.Errorf("Expected table name Table1, Got: %v", s.From.Join.Source.TableName.Name)
	}
	if s.From.Join.Target.TableName.Name != "Table2" {
		t.Errorf("Expected joining table name Table2, Got: %v", s.From.Join.Target.Join.Source.TableName.Name)
	}
	if s.From.Join.Condition == nil {
		t.Errorf("Expected condition on the join")
	}
	cc, ok := s.From.Join.Condition.(ComparisonExpression)
	if ok != true {
		t.Errorf("Expected comparision condition")
	}
	if cc.Operator != ">=" {
		t.Errorf("Expected condition operator to be '>=', but got %s", cc.Operator)
	}
	if cc.LHS.Name.Name != Identifier("tcol1") {
		t.Errorf("Expected left operand to be 'tcol1', but got %s", cc.LHS.Name)
	}
	if cc.RHS.Name.Name != "tcol2" {
		t.Errorf("Expected right operand to be 'tcol2', but got %s", cc.RHS.Name)
	}
}

func TestSimpleWhereExpression(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10")
	if s.Filter == nil {
		t.Fatal("Expected Filter, but got nil")
	}
	switch v := (s.Filter).(type) {
	case ComparisonExpression:
		if v.LHS.Name == nil {
			t.Errorf("Expected LHS to be a compound name, but get %v", v.LHS)
		}
		if v.RHS.Value.NumericValue == nil {
			t.Errorf("Expected RHS to be a numeric value, but get %v", v.RHS)
		}
		if v.Operator != ">" {
			t.Errorf("Expected operator to be '>' but got %v", v.Operator)
		}
	default:
		t.Errorf("Expected Comparision expression, but got '%v'", s.Filter)
	}
}

func TestSimpleWhereLikeExpression(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT LIKE 'abc%'")
	if s.Filter == nil {
		t.Fatal("Expected Filter, but got nil")
	}
	switch v := (s.Filter).(type) {
	case LikeExpression:
		if !v.NotLike {
			t.Error("Expected NOT LIKE")
		}
		if v.What.String() != "col2" {
			t.Errorf("Expected expr to be a compound name, but get %v", v.What)
		}
		if v.Pattern != "abc%" {
			t.Errorf("Expected pattern to be 'abc%%', but get %v", v.Pattern)
		}

	default:
		t.Errorf("Expected LIKE expression, but got '%v'", s.Filter)
	}
}

func TestSimpleAndOrWhereExpression(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 > 10 and col2 < 20 OR col3 >= 0")
	if s.Filter == nil {
		t.Fatal("Expected Filter, but got nil")
	}
	if (s.Filter).String() != "col2 > 10 AND col2 < 20 OR col3 >= 0" {
		t.Errorf("Incorrectly parsed Filter expression: %v", s.Filter)
	}
}

func TestInListExpression(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 NOT IN (1, 2, 3)")
	if s.Filter == nil {
		t.Fatal("Expected Filter, but got nil")
	}
	switch v := (s.Filter).(type) {
	case InListExpression:
		if !v.NotIn {
			t.Error("Expected NOT LIKE")
		}
		if v.What.String() != "col2" {
			t.Errorf("Expected expr to be a compound name, but get %v", v.What)
		}
		if v.List.String() != "(1, 2, 3)" {
			t.Errorf("Expected list to be '(1, 2, 3)', but get %v", v.List)
		}
	default:
		t.Errorf("Expected IN LIST expression, but got '%v'", s.Filter)
	}
}

func TestBetweenExpression(t *testing.T) {
	s := ParseSQL("SELECT col FROM Table1 WHERE col2 BETWEEN 10 AND col1")
	if s.Filter == nil {
		t.Fatal("Expected Filter, but got nil")
	}
	switch v := (s.Filter).(type) {
	case BetweenExpression:
		high, ok := v.High.(Term)
		if !ok {
			t.Errorf("High boundary should be term")
		}
		if high.Name.String() != "col1" {
			t.Errorf("Expected high boundary to be equal col1, got: %s", high)
		}
		if high.Value != nil {
			t.Errorf("Expected high boundary to not be a value")
		}
		low, ok := v.Low.(Term)
		if !ok {
			t.Errorf("Low boundary should be term")
		}
		if *low.Value.NumericValue != "10" {
			t.Errorf("Expected low boundary to be equal 10, got: %s", low.Value)
		}
		if low.Name != nil {
			t.Errorf("Expected low boundary to not be a name")
		}
		if v.What.String() != "col2" {
			t.Errorf("Expected expr to be a compound name, but get %v", v.What)
		}
	default:
		t.Errorf("Expected BETWEEN expression, but got '%v'", s.Filter)
	}
}
