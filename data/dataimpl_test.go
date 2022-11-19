package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRowImpl_GetRowId(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row, err := ds.NextRow()
	if err != nil {
		t.Fatal(err)
	}

	val := row.GetByName(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal("0"))

	row, err = ds.NextRow()
	g.Expect(err).ShouldNot(HaveOccurred())

	val = row.GetByName(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal("1"))

	g.Expect(ds.Rewind()).ShouldNot(HaveOccurred())
	row, err = ds.NextRow()
	g.Expect(err).ShouldNot(HaveOccurred())

	val = row.GetByName(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal("0"))
}

func TestRowImpl_Satisfies(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row, err := ds.NextRow()
	g.Expect(err).ShouldNot(HaveOccurred())

	condT := NewAndCondition(
		NewEq(NewRowValue(row, "first_name"), NewStringValue("John")),
		NewLt(NewRowValue(row, "dept_id"), NewIntValue(3)))

	condF := NewEq(NewRowValue(row, "first_name"), NewStringValue("Connor"))

	g.Expect(row.Satisfies(condT)).To(BeTrue())
	g.Expect(row.Satisfies(condF)).To(BeFalse())
}
