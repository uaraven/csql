package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestProjectionDataSource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	pds := NewProjectionDataSource(ds, NewProjection([]string{"first_name", "last_name"}, []string{"", "ln"}))

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(HaveLen(2))
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("last_name").IsPresent()).To(BeFalse())
	g.Expect(rows[0].Get("ln").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("ln").Value()).To(Equal(NewStringValue("Snow")))

	g.Expect(rows[3].Values()).To(HaveLen(2))
	g.Expect(rows[3].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("last_name").IsPresent()).To(BeFalse())
	g.Expect(rows[3].Get("ln").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("ln").Value()).To(Equal(NewStringValue("May")))
}
