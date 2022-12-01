package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestOrderedDatasource_OrderBy(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{
		NewOrderBy("width", false),
	}, 0)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(3)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(2)))
	g.Expect(rows[2].GetByIndex(0).Value()).To(Equal(int64(1)))
	g.Expect(rows[3].GetByIndex(0).Value()).To(Equal(int64(4)))
}

func TestOrderedDatasource_OrderByDesc(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{
		NewOrderBy("width", true),
	}, 0)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(4)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(1)))
	g.Expect(rows[2].GetByIndex(0).Value()).To(Equal(int64(2)))
	g.Expect(rows[3].GetByIndex(0).Value()).To(Equal(int64(3)))
}

func TestOrderedDatasource_OrderByTwoFields(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{
		NewOrderBy("last_name", false),
		NewOrderBy("first_name", true),
	}, 0)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(3)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(4)))
	g.Expect(rows[2].GetByIndex(0).Value()).To(Equal(int64(1)))
	g.Expect(rows[3].GetByIndex(0).Value()).To(Equal(int64(2)))
}

func TestOrderedDatasource_OrderByIndex(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{
		NewOrderByIndex(3, false),
		NewOrderByIndex(2, true),
	}, 0)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(3)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(4)))
	g.Expect(rows[2].GetByIndex(0).Value()).To(Equal(int64(1)))
	g.Expect(rows[3].GetByIndex(0).Value()).To(Equal(int64(2)))
}

func TestOrderedDatasource_Limit(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{}, 2)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(1)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(2)))
}

func TestOrderedDatasource_OrderByIndexLimit(t *testing.T) {
	g := NewGomegaWithT(t)

	ds := loadTestMemDatasource("people")

	ods := NewOrderedDatasource(ds, []OrderByField{
		NewOrderByIndex(3, false),
		NewOrderByIndex(2, false),
	}, 3)

	rows := ReadAllRows(ods)

	g.Expect(rows).To(HaveLen(3))
	g.Expect(rows[0].GetByIndex(0).Value()).To(Equal(int64(3)))
	g.Expect(rows[1].GetByIndex(0).Value()).To(Equal(int64(4)))
	g.Expect(rows[2].GetByIndex(0).Value()).To(Equal(int64(2)))
}
