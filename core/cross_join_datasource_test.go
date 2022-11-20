package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestCrossJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)
	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cj := NewCrossJoin(emp, dept)
	g.Expect(cj).ToNot(BeNil())

	g.Expect(cj.CurrentRow()).To(BeNil())

	rows, err := ReadAllRows(cj)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(rows).To(HaveLen(8))

	g.Expect(rows[0].Get("employees.dept_id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[0].Get("dept.id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[0].Get("employees.first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[0].Get("dept.name").Value()).To(Equal(NewStringValue("The Wall")))

	g.Expect(rows[1].Get("employees.dept_id").Value()).To(Equal(NewIntValue(1)))
	g.Expect(rows[1].Get("dept.id").Value()).To(Equal(NewIntValue(2)))
	g.Expect(rows[1].Get("employees.first_name").Value()).To(Equal(NewStringValue("John")))
	g.Expect(rows[1].Get("dept.name").Value()).To(Equal(NewStringValue("The Grand Tour")))

}
