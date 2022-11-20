package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRowValue_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource()

	row, err := ds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	firstName := NewRowValue("first_name")
	g.Expect(firstName.Evaluate(row).Type()).To(Equal(TypeString))
	g.Expect(firstName.Evaluate(row).Value()).To(Equal("John"))

	lastName := NewRowValue("last_name")
	g.Expect(lastName.Evaluate(row).Type()).To(Equal(TypeString))
	g.Expect(lastName.Evaluate(row).Value()).To(Equal("Snow"))

	deptId := NewRowValue("dept_id")
	g.Expect(deptId.Evaluate(row).Type()).To(Equal(TypeInt))
	g.Expect(deptId.Evaluate(row).Value()).To(Equal(int64(1)))

	width := NewRowValue("width")
	g.Expect(width.Evaluate(row).Type()).To(Equal(TypeFloat))
	g.Expect(width.Evaluate(row).Value()).To(Equal(1.0))
}

func TestRowImpl_CacheTest(t *testing.T) {
	g := NewGomegaWithT(t)

	ctx1 := newMapContextWithId(1)
	ctx1.SetValue("number", "10")

	ctx2 := newMapContextWithId(1)
	ctx2.SetValue("number", "15")

	width := NewRowValue("number")
	g.Expect(width.Evaluate(ctx1).Value()).To(Equal(int64(10)))
	g.Expect(width.Evaluate(ctx2).Value()).To(Equal(int64(10)))

	ctx2.SetId(2)
	width = NewRowValue("number")
	g.Expect(width.Evaluate(ctx1).Value()).To(Equal(int64(10)))
	g.Expect(width.Evaluate(ctx2).Value()).To(Equal(int64(15)))

}

func TestRowImpl_Get_QualifiedAccess(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource()

	row, err := ds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	v := row.Get("first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("employees.first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("employee.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

	v = row.Get("emp.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())
}

func TestRowImpl_Get_QualifiedAccessAlias(t *testing.T) {
	g := NewGomegaWithT(t)

	dsa := loadTestMemDatasourceWithAlias("emp")

	row, err := dsa.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	v := row.Get("first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("emp.first_name")
	g.Expect(v.IsPresent()).To(BeTrue())
	g.Expect(v.Value()).To(Equal(NewStringValue("John")))

	v = row.Get("emp.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

	v = row.Get("employee.middle_name")
	g.Expect(v.IsPresent()).To(BeFalse())

}
