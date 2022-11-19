package data

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
	g.Expect(row.GetByName("first_name").Value()).To(Equal("Jeremy"))
	g.Expect(row.GetByName("last_name").Value()).To(Equal("Clarkson"))

	row, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(int64(1)))
	g.Expect(row.GetByName("first_name").Value()).To(Equal("James"))
	g.Expect(row.GetByName("last_name").Value()).To(Equal("May"))
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
	g.Expect(row.GetByName("first_name").Value()).To(Equal("Jeremy"))
	g.Expect(row.GetByName("last_name").Value()).To(Equal("Clarkson"))
}
