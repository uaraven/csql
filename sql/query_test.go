package sql

import (
	. "github.com/onsi/gomega"
	"github.com/uaraven/csql/core"
	"testing"
)

func TestOrderByCalculatedColumn(t *testing.T) {
	g := NewGomegaWithT(t)
	rows := ExecuteSql(` select city || ',' || country as city, population/10000 as bucket from "../test-data/cities" order by 2 desc,1`).
		GetRows()

	columns := []string{"city", "bucket"}
	g.Expect(core.RowsToSlice(rows, columns...)).To(ContainElements(
		[]interface{}{"London,UK", int64(879)},
		[]interface{}{"Warsaw,Poland", int64(305)},
		[]interface{}{"Odesa,Ukraine", int64(101)},
		[]interface{}{"London,Canada", int64(42)},
		[]interface{}{"Waterloo,Canada", int64(12)},
		[]interface{}{"Dubrovnik,Croatia", int64(4)},
		[]interface{}{"Waterloo,Belgium", int64(3)},
		[]interface{}{"London,USA", int64(1)},
		[]interface{}{"Odessa,Canada", int64(1)},
		[]interface{}{"Waterloo,Australia", int64(1)},
		[]interface{}{"London,USA", int64(0)},
		[]interface{}{"London,USA", int64(0)},
		[]interface{}{"Odessa,USA", int64(0)},
		[]interface{}{"Warsaw,USA", int64(0)},
		[]interface{}{"Waterloo,USA", int64(0)},
	))
}

func TestGroupByCalculatedColumn(t *testing.T) {
	g := NewGomegaWithT(t)
	rows := ExecuteSql(`select count(*) as cnt, population/10000 as bucket from "../test-data/cities" order by 2, 1 group by 2`).
		GetRows()

	columns := []string{"cnt", "bucket"}
	g.Expect(core.RowsToSlice(rows, columns...)).To(ContainElements(
		[]interface{}{int64(5), int64(0)},
		[]interface{}{int64(3), int64(1)},
		[]interface{}{int64(1), int64(3)},
		[]interface{}{int64(1), int64(4)},
		[]interface{}{int64(1), int64(12)},
		[]interface{}{int64(1), int64(42)},
		[]interface{}{int64(1), int64(101)},
		[]interface{}{int64(1), int64(305)},
		[]interface{}{int64(1), int64(879)},
	))
}
