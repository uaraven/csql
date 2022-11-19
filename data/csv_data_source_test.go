package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLoadCsv(t *testing.T) {
	g := NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")

	g.Expect(err).To(BeNil())

	hdr := ds.Header()

	g.Expect(hdr.ColumnCount()).To(Equal(5))
}

func TestLoadCsvRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row, err := ds.NextRow()
	g.Expect(err).To(BeNil())

	g.Expect(row.Values()).To(Equal([]string{"1", "1", "John", "Snow", "1.0"}))
	g.Expect(row.Id()).To(Equal(int64(0)))

	row, err = ds.NextRow()
	g.Expect(err).To(BeNil())

	g.Expect(row.Values()).To(Equal([]string{"2", "1", "James", "Snow", "1.2"}))
	g.Expect(row.Id()).To(Equal(int64(1)))
}

func TestLoadCsvRewind(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row1, err := ds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(ds.Rewind()).To(Succeed())
	g.Expect(err).ShouldNot(HaveOccurred())

	g.Expect(ds.NextRow()).To(Equal(row1))
}

func TestCsvGetValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row, err := ds.NextRow()
	g.Expect(err).ShouldNot(HaveOccurred())

	val := row.GetByName("first_name")

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value()).To(Equal("John"))
}

func TestCsvGetValueByInvalidName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestDatasource()

	row, err := ds.NextRow()
	g.Expect(err).ShouldNot(HaveOccurred())

	val := row.GetByName("middle_name")

	g.Expect(val.IsPresent()).To(BeFalse())
}
