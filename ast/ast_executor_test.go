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

	sql := ParseSQL(`select first_name, width*2 from "../test-data/employees.csv"`)

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

func TestAstExecutor_SelectProjectionAliases(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select emp.first_name as fn, emp.last_name from "../test-data/employees.csv" emp`)

	ds := Execute(&sql)
	g.Expect(ds.Header().ColumnCount()).To(Equal(2))
	rows := core.ReadAllRows(ds)
	g.Expect(rows).To(HaveLen(4))
	g.Expect(rows[0].Get("fn").Value()).To(Equal(core.NewStringValue("John")))
	g.Expect(rows[0].Get("last_name").Value()).To(Equal(core.NewStringValue("Snow")))
	g.Expect(rows[0].Get("first_name").IsPresent()).To(BeFalse())
}

func TestAstExecutor_LeftOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select * from "../test-data/authors.csv" 
    left join "../test-data/books.csv"
	on authors.id = books.author_id`)

	jds := Execute(&sql)

	rows := core.ReadAllRows(jds)
	g.Expect(rows[0].Values()).To(Equal([]core.Value{
		core.NewIntValue(1),
		core.NewStringValue("J.R.R. Tolkien"),
		core.NewIntValue(1),
		core.NewIntValue(1),
		core.NewStringValue("Lord Of The Rings")}))
	g.Expect(rows[1].Values()).To(Equal([]core.Value{
		core.NewIntValue(2),
		core.NewStringValue("Stephen King"),
		core.NewIntValue(2),
		core.NewIntValue(2),
		core.NewStringValue("Carrie")}))
	g.Expect(rows[2].Values()).To(Equal([]core.Value{
		core.NewIntValue(2),
		core.NewStringValue("Stephen King"),
		core.NewIntValue(3),
		core.NewIntValue(2),
		core.NewStringValue("The Dark Tower")}))
}

func TestAstExecutor_RightOuterJoin(t *testing.T) {
	g := NewGomegaWithT(t)

	sql := ParseSQL(`select * from "../test-data/authors.csv" 
    right join "../test-data/books.csv"
	on authors.id = books.author_id`)

	jds := Execute(&sql)

	rows := core.ReadAllRows(jds)
	g.Expect(rows).To(HaveLen(5))
	g.Expect(rows[0].Values()).To(Equal([]core.Value{
		core.NewIntValue(1),
		core.NewStringValue("J.R.R. Tolkien"),
		core.NewIntValue(1),
		core.NewIntValue(1),
		core.NewStringValue("Lord Of The Rings"),
	}))

	g.Expect(rows[4].Values()).To(Equal([]core.Value{
		core.NewNullValue(),
		core.NewNullValue(),
		core.NewIntValue(5),
		core.NewIntValue(6),
		core.NewStringValue("Unwritten Book 2"),
	}))
}
