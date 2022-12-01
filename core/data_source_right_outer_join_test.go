package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRightOuterJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewRightOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.Header().ColumnCount()).To(Equal(5))
	rows := ReadAllRows(jds)

	g.Expect(rows).To(HaveLen(5))
	g.Expect(rows[0].Values()).To(Equal([]Value{
		NewIntValue(1),
		NewStringValue("J.R.R. Tolkien"),
		NewIntValue(1),
		NewIntValue(1),
		NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[4].Values()).To(Equal([]Value{
		NewNullValue(),
		NewNullValue(),
		NewIntValue(5),
		NewIntValue(6),
		NewStringValue("Unwritten Book 2"),
	}))
}

func TestRightOuterJoinDatasource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.CurrentRow()).To(BeNil())

	row := jds.NextRow()
	g.Expect(jds.CurrentRow()).To(Equal(row))
}

func TestRightOuterJoinDatasource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewLeftOuterJoinDatasource(left, right, joinCondition)

	row := jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))

	g.Expect(func() { jds.Rewind() }).ToNot(Panic())

	row = jds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))
}

func TestRightOuterJoinDatasource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	left := NewMemDataSourceFromCsv("../test-data/authors.csv")
	right := NewMemDataSourceFromCsv("../test-data/books.csv")

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds := NewRightOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.GetName()).To(Equal("(authors RIGHT JOIN books)"))
}
