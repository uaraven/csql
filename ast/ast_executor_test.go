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
