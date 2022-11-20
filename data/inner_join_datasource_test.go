package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestInnerJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)
	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cj := NewInnerJoin(emp, dept, NewEq(NewRowValue("dept_id"), NewRowValue("dept.id")))
	g.Expect(cj).ToNot(BeNil())

	g.Expect(cj.CurrentRow()).To(BeNil())

	rows, err := ReadAllRows(cj)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rows).To(HaveLen(4))

	g.Expect(rows[0].Get("employees.dept_id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[0].Get("dept.id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[0].Get("employees.first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("dept.name").Value()).To(Equal(NewStringValue("The Wall")))

	g.Expect(rows[3].Get("employees.dept_id").Value()).To(Equal(NewIntValue(2)))
	g.Expect(rows[3].Get("dept.id").Value()).To(Equal(NewIntValue(2)))
	g.Expect(rows[3].Get("employees.first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(rows[3].Get("dept.name").Value()).To(Equal(NewStringValue("The Grand Tour")))
}
