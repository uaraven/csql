package data

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestRowValue_Evaluate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")
	g.Expect(err).To(gomega.BeNil())

	row, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	firstName := NewRowValue(row, "first_name")
	g.Expect(firstName.Evaluate().Type()).To(gomega.Equal(TypeString))
	g.Expect(firstName.Evaluate().Value()).To(gomega.Equal("John"))

	lastName := NewRowValue(row, "last_name")
	g.Expect(lastName.Evaluate().Type()).To(gomega.Equal(TypeString))
	g.Expect(lastName.Evaluate().Value()).To(gomega.Equal("Snow"))

	deptId := NewRowValue(row, "dept_id")
	g.Expect(deptId.Evaluate().Type()).To(gomega.Equal(TypeInt))
	g.Expect(deptId.Evaluate().Value()).To(gomega.Equal(int64(1)))

	width := NewRowValue(row, "width")
	g.Expect(width.Evaluate().Type()).To(gomega.Equal(TypeFloat))
	g.Expect(width.Evaluate().Value()).To(gomega.Equal(1.0))
}
