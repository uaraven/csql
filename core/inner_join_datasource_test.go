package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestInnerJoinDatasource_NextRow(t *testing.T) {
	g := NewGomegaWithT(t)
	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cj, err := NewInnerJoin(emp, dept, NewEq(NewRowValue("dept_id"), NewRowValue("dept.id")))
	g.Expect(err).ToNot(HaveOccurred())
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

func TestInnerJoinDataSource_CurrentRow(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cond := NewAnd(
		NewEq(NewRowValue("employees.dept_id"), NewRowValue("dept.id")),
		NewAnd(
			NewEq(NewRowValue("first_name"), NewStringValue("James")),
			NewEq(NewRowValue("last_name"), NewStringValue("May"))))

	fds, err := NewInnerJoin(emp, dept, cond)
	g.Expect(err).ToNot(HaveOccurred())

	nrow, err := fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(nrow.Id()).To(Equal(0))
	g.Expect(nrow.Get("first_name").Value()).To(Equal(NewStringValue("James")))
	g.Expect(nrow.Get("last_name").Value()).To(Equal(NewStringValue("May")))

	crow, err := fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).To(Equal(nrow))

	nrow, err = fds.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(nrow).To(BeNil())

	crow, err = fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).ToNot(BeNil())

	g.Expect(fds.Rewind()).To(Succeed())
	crow, err = fds.CurrentRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(crow).To(BeNil())

}

func TestInnerJoinDatasource_Header(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cond := NewEq(NewRowValue("employees.dept_id"), NewRowValue("dept.id"))

	fds, err := NewInnerJoin(emp, dept, cond)
	g.Expect(err).ToNot(HaveOccurred())
	header := fds.Header()
	g.Expect(header.ColumnCount()).To(Equal(7))
	g.Expect(header.IndexByName("first_name")).To(Equal(2))
	g.Expect(func() { header.IndexByName("id") }).To(Panic())

}

func TestInnerJoinDatasource_GetName(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	fds, err := NewCrossJoin(emp, dept)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(fds.GetName()).To(Equal("employees JOIN dept"))

}

func TestInnerJoinDatasource_Rewind(t *testing.T) {
	g := NewGomegaWithT(t)
	emp := loadDefaultTestMemDatasource()
	dept := loadTestMemDatasource("dept")

	cj, err := NewInnerJoin(emp, dept, NewEq(NewRowValue("dept_id"), NewRowValue("dept.id")))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cj).ToNot(BeNil())

	g.Expect(cj.CurrentRow()).To(BeNil())

	row, err := cj.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(0))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("John")))

	row, err = cj.NextRow()
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(row.Id()).To(Equal(1))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("James")))

	err = cj.Rewind()
	g.Expect(err).ToNot(HaveOccurred())

	row, err = cj.NextRow()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(row.Id()).To(Equal(0))
	g.Expect(row.Get("first_name").Value()).To(Equal(NewStringValue("John")))

}
