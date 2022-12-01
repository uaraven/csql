package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFilteredDataSource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	g.Expect(fds.GetName()).To(Equal(ds.GetName()))
}

func TestFilteredDataSource_Header(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	g.Expect(fds.Header().ColumnsMetadata()).To(Equal(ds.Header().ColumnsMetadata()))
}

func TestFilteredDataSource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("first_name"), NewStringValue("James"))

	fds := NewFilteredDataSource(ds, cond)

	nrow := fds.NextRow()

	g.Expect(nrow.Id()).To(Equal(0))
	g.Expect(nrow.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(nrow.Get("last_name").Value()).To(Equal(NewStringValue("Snow")))

	crow := fds.CurrentRow()
	g.Expect(crow).To(Equal(nrow))

	nrow = fds.NextRow()

	g.Expect(nrow.Id()).To(Equal(1))
	g.Expect(nrow.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(nrow.Get("last_name").Value()).To(Equal(NewStringValue("May")))

	crow = fds.CurrentRow()
	g.Expect(crow).To(Equal(nrow))

	nrow = fds.NextRow()
	g.Expect(nrow).To(BeNil())

	crow = fds.CurrentRow()
	g.Expect(crow).ToNot(BeNil())

}

func TestFilteredDataSource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	row := fds.NextRow()

	g.Expect(row.Id()).To(Equal(0))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("Jeremy")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("Clarkson")))

	row = fds.NextRow()

	g.Expect(row.Id()).To(Equal(1))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("May")))
}

func TestFilteredDataSource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	row := fds.NextRow()
	row = fds.NextRow()

	g.Expect(func() { fds.Rewind() }).ToNot(Panic())

	row = fds.NextRow()

	g.Expect(row.Id()).To(Equal(0))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("Jeremy")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("Clarkson")))
}
