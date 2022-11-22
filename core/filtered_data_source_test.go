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

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	g.Expect(fds.Header()).To(Equal(ds.Header()))
}

func TestFilteredDataSource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("first_name"), NewStringValue("James"))

	fds := NewFilteredDataSource(ds, cond)

	nrow, err := fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(nrow.Id()).To(Equal(int64(0)))
	g.Expect(nrow.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(nrow.Get("last_name").Value()).To(Equal(NewStringValue("Snow")))

	crow, err := fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).To(Equal(nrow))

	nrow, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(nrow.Id()).To(Equal(int64(1)))
	g.Expect(nrow.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(nrow.Get("last_name").Value()).To(Equal(NewStringValue("May")))

	crow, err = fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).To(Equal(nrow))

	nrow, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(nrow).To(BeNil())

	crow, err = fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).To(BeNil())

}

func TestFilteredDataSource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	row, err := fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(int64(0)))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("Jeremy")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("Clarkson")))

	row, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(int64(1)))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("May")))
}

func TestFilteredDataSource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	cond := NewEq(NewRowValue("dept_id"), NewIntValue(2))

	fds := NewFilteredDataSource(ds, cond)

	row, err := fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	row, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(fds.Rewind()).To(Succeed())

	row, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(int64(0)))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("Jeremy")))
	g.Expect(row.Get("last_name").Value()).To(Equal(NewStringValue("Clarkson")))
}
