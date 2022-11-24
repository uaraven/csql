package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRightOuterJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left, err := NewMemDataSourceFromCsv("../test-data/authors.csv")
	g.Expect(err).ToNot(HaveOccurred())
	right, err := NewMemDataSourceFromCsv("../test-data/books.csv")
	g.Expect(err).ToNot(HaveOccurred())

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds, err := NewRightOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.Header().ColumnCount()).To(Equal(5))
	rows, err := ReadAllRows(jds)
	g.Expect(err).ToNot(HaveOccurred())

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

	left, err := NewMemDataSourceFromCsv("../test-data/authors.csv")
	g.Expect(err).ToNot(HaveOccurred())
	right, err := NewMemDataSourceFromCsv("../test-data/books.csv")
	g.Expect(err).ToNot(HaveOccurred())

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds, err := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.CurrentRow()).To(BeNil())

	row, err := jds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(jds.CurrentRow()).To(Equal(row))
}

func TestRightOuterJoinDatasource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	left, err := NewMemDataSourceFromCsv("../test-data/authors.csv")
	g.Expect(err).ToNot(HaveOccurred())
	right, err := NewMemDataSourceFromCsv("../test-data/books.csv")
	g.Expect(err).ToNot(HaveOccurred())

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds, err := NewLeftOuterJoinDatasource(left, right, joinCondition)

	row, err := jds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))

	g.Expect(jds.Rewind()).ToNot(HaveOccurred())

	row, err = jds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewStringValue("J.R.R. Tolkien"), NewIntValue(1), NewIntValue(1), NewStringValue("Lord Of The Rings")}))
}

func TestRightOuterJoinDatasource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	left, err := NewMemDataSourceFromCsv("../test-data/authors.csv")
	g.Expect(err).ToNot(HaveOccurred())
	right, err := NewMemDataSourceFromCsv("../test-data/books.csv")
	g.Expect(err).ToNot(HaveOccurred())

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds, err := NewRightOuterJoinDatasource(left, right, joinCondition)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(jds.GetName()).To(Equal("(authors RIGHT JOIN books)"))
}
