package core

import (
	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/funky"
	"testing"
)

func TestProjectionDataSource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	pds, err := NewProjectionDataSource(ds, NewSimpleProjection([]string{"first_name", "last_name"}, []string{"", "ln"}), false)
	g.Expect(err).ToNot(HaveOccurred())

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

func TestProjectionDataSource_NextRowDistinct(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	pds, err := NewProjectionDataSource(ds, NewSimpleProjection([]string{"first_name"}, []string{""}), true)
	g.Expect(err).ToNot(HaveOccurred())

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(3))
	firstNames := funky.Map(rows, func(r Row) string {
		return r.GetByIndex(0).Value().(string)
	})
	g.Expect(firstNames).To(ContainElements("James", "John", "Jeremy"))
}

func TestProjectionDataSource_WithExpressions(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	projection := []ProjectionColumn{
		NewColumn("first_name"),
		NewExpressionColumn(NewExpression(addition, NewRowValue("id"), NewIntValue(10)), "sum"),
	}
	pds, err := NewProjectionDataSource(ds, projection, false)
	g.Expect(err).ToNot(HaveOccurred())

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(HaveLen(2))
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("sum").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("sum").Value().Value()).To(Equal(int64(11)))

	g.Expect(rows[3].Values()).To(HaveLen(2))
	g.Expect(rows[3].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("sum").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("sum").Value().Value()).To(Equal(int64(14)))
}

func TestProjectionDataSource_WithUnnamedExpressions(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	projection := []ProjectionColumn{
		NewColumn("first_name"),
		NewExpressionColumn(NewExpression(addition, NewRowValue("id"), NewIntValue(10)), ""),
	}
	pds, err := NewProjectionDataSource(ds, projection, false)
	g.Expect(err).ToNot(HaveOccurred())

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(HaveLen(2))
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("sum").IsPresent()).To(BeFalse())
	g.Expect(rows[0].GetByIndex(1).Value()).To(Equal(int64(11)))

	g.Expect(rows[3].Values()).To(HaveLen(2))
	g.Expect(rows[3].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("sum").IsPresent()).To(BeFalse())
	g.Expect(rows[3].GetByIndex(1).Value()).To(Equal(int64(14)))
}

func TestProjectionDataSource_WithStar(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadDefaultTestMemDatasource()

	projection := []ProjectionColumn{
		NewColumn("*"),
		NewExpressionColumn(NewExpression(addition, NewRowValue("id"), NewIntValue(10)), ""),
	}
	pds, err := NewProjectionDataSource(ds, projection, false)
	g.Expect(err).ToNot(HaveOccurred())

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(HaveLen(6))
	g.Expect(rows[0].Get("id").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("dept_id").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("last_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("width").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("sum").IsPresent()).To(BeFalse())
	g.Expect(rows[0].GetByIndex(5).Value()).To(Equal(int64(11)))

	g.Expect(rows[3].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("sum").IsPresent()).To(BeFalse())
	g.Expect(rows[3].GetByIndex(5).Value()).To(Equal(int64(14)))
}

func TestProjectionDataSource_WithTableStar(t *testing.T) {
	g := NewGomegaWithT(t)
	emps := loadDsWithAlias("employees", "e")
	depts := loadDsWithAlias("dept", "d")
	jds, err := NewInnerJoin(emps, depts, NewEq(NewRowValue("e.dept_id"), NewRowValue("d.id")))
	g.Expect(err).ToNot(HaveOccurred())

	projection := []ProjectionColumn{
		NewColumn("e.first_name"),
		NewColumn("e.last_name"),
		NewColumn("d.*"),
	}
	pds, err := NewProjectionDataSource(jds, projection, false)
	g.Expect(err).ToNot(HaveOccurred())

	rows, err := ReadAllRows(pds)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Values()).To(HaveLen(4))
	g.Expect(rows[0].Get("id").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("last_name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("name").IsPresent()).To(BeTrue())
	g.Expect(rows[0].Get("name").Value()).To(Equal(NewStringValue("The Wall")))

	g.Expect(rows[3].Values()).To(HaveLen(4))
	g.Expect(rows[3].Get("id").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("last_name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("name").IsPresent()).To(BeTrue())
	g.Expect(rows[3].Get("name").Value()).To(Equal(NewStringValue("The Grand Tour")))
}
