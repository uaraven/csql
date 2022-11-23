package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLeftOuterJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	left, err := NewMemDataSource("../test-data/authors.csv")
	g.Expect(err).ToNot(HaveOccurred())
	right, err := NewMemDataSource("../test-data/books.csv")
	g.Expect(err).ToNot(HaveOccurred())

	joinCondition := NewEq(NewRowValue("authors.id"),
		NewRowValue("books.author_id"))
	jds, err := NewLeftOuterJoinDatasource(left, right, joinCondition)

	g.Expect(jds.Header().ColumnCount()).To(Equal(5))
	rows, err := ReadAllRows(jds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
}
