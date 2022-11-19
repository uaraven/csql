package data

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestLoadCsv(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")

	g.Expect(err).To(gomega.BeNil())

	hdr := ds.Header()

	g.Expect(hdr.ColumnCount()).To(gomega.Equal(4))
}

func TestLoadCsvRow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")

	g.Expect(err).To(gomega.BeNil())

	row, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	g.Expect(row.Values()).To(gomega.Equal([]string{"1", "1", "John", "Snow"}))
}

func TestLoadCsvRewind(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")

	g.Expect(err).To(gomega.BeNil())

	row1, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	err = ds.Rewind()
	g.Expect(err).To(gomega.BeNil())

	row2, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	g.Expect(row1).To(gomega.Equal(row2))
}

func TestCsvGetValueByName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")
	g.Expect(err).To(gomega.BeNil())

	row, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	val := row.GetByName("first_name")

	g.Expect(val.IsPresent()).To(gomega.BeTrue())
	g.Expect(val.Value()).To(gomega.Equal("John"))
}

func TestCsvGetValueByInvalidName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	ds, err := NewCsvDataSource("../test-data/employees.csv")
	g.Expect(err).To(gomega.BeNil())

	row, err := ds.NextRow()
	g.Expect(err).To(gomega.BeNil())

	val := row.GetByName("middle_name")

	g.Expect(val.IsPresent()).To(gomega.BeFalse())
}
