package data

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestRowImpl_Satisfies(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")
	g.Expect(err).To(gomega.BeNil())

	row, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	condT := NewAndCondition(
		NewEq(NewRowValue(row, "first_name"), NewStringValue("John")),
		NewLt(NewRowValue(row, "dept_id"), NewIntValue(3)))

	condF := NewEq(NewRowValue(row, "first_name"), NewStringValue("Connor"))

	g.Expect(row.Satisfies(condT)).To(gomega.BeTrue())
	g.Expect(row.Satisfies(condF)).To(gomega.BeFalse())
}
