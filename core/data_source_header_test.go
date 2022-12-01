package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRowImpl_GetRowId(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row := ds.NextRow()

	val := row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(0)))

	row = ds.NextRow()

	val = row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(1)))

	g.Expect(func() { ds.Rewind() }).ToNot(Panic())
	row = ds.NextRow()

	val = row.Get(RowId)

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal(NewIntValue(0)))
}

func TestRowImpl_Satisfies(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row := ds.NextRow()

	condT := NewAnd(
		NewEq(NewRowValue("first_name"), NewStringValue("John")),
		NewLt(NewRowValue("dept_id"), NewIntValue(3)))

	condF := NewEq(NewRowValue("first_name"), NewStringValue("Connor"))

	g.Expect(row.Satisfies(condT)).To(BeTrue())
	g.Expect(row.Satisfies(condF)).To(BeFalse())
}
