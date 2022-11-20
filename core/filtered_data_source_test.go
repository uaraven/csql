package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

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
