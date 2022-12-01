package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestLoadMemCsv(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := NewMemDataSourceFromCsv("../test-data/employees.csv")

	hdr := ds.Header()

	g.Expect(hdr.ColumnCount()).To(Equal(5))

	rows := ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
}

func TestLoadMemCsvRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(1), NewIntValue(1),
		NewStringValue("John"),
		NewStringValue("Snow"),
		NewFloatValue(1.0)}))
	g.Expect(row.Id()).To(Equal(0))

	row = ds.NextRow()
	g.Expect(row.Values()).To(Equal([]Value{NewIntValue(2), NewIntValue(1),
		NewStringValue("James"),
		NewStringValue("Snow"),
		NewFloatValue(1.2)}))
	g.Expect(row.Id()).To(Equal(1))
}

func TestLoadMemCsvRewind(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row1 := ds.NextRow()

	g.Expect(func() { ds.Rewind() }).ToNot(Panic())

	g.Expect(ds.NextRow()).To(Equal(row1))
}

func TestMemCsvGetValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	val := row.Get("first_name")

	g.Expect(val.IsPresent()).To(BeTrue())
	g.Expect(val.Value().Value()).To(Equal("John"))
}

func TestMemCsvGetValueByInvalidName(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	row := ds.NextRow()

	val := row.Get("middle_name")

	g.Expect(val.IsPresent()).To(BeFalse())
}
