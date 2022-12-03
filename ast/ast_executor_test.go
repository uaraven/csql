package ast

import (
	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/core"
	"testing"
)

func TestAstExecutor_Select(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL("select * from \"../test-data/employees.csv\"")

	ds := Execute(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(5))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
}

func TestAstExecutor_SelectWhere(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL("select * from \"../test-data/employees.csv\" where dept_id = 1")

	ds := Execute(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(5))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(2))
	g.Expect(rows[0].Get("first_name").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("width").Value()).To(Equal(core.NewFloatValue(1.0)))
	g.Expect(rows[1].Get("first_name").Value()).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[1].Get("width").Value()).To(Equal(core.NewFloatValue(1.2)))
}

func TestAstExecutor_SelectProjection(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL("select first_name, width*2 from \"../test-data/employees.csv\"")

	ds := Execute(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(2))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].GetByIndex(0)).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].GetByIndex(1)).To(Equal(core.NewFloatValue(2.0)))
	g.Expect(rows[1].GetByIndex(0)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[1].GetByIndex(1)).To(Equal(core.NewFloatValue(1.2 * 2)))
	g.Expect(rows[3].GetByIndex(0)).To(Equal(core.NewStringValue("James")))
	g.Expect(rows[3].GetByIndex(1)).To(Equal(core.NewFloatValue(1.6 * 2)))
}
