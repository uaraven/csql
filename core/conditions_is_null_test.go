package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestIsNullCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("a", "null")

	isN := NewIsNull(NewRowValue("a"), false)
	g.Expect(isN.Evaluate(context).Value()).To(BeTrue())

	isNotN := NewIsNull(NewRowValue("a"), true)
	g.Expect(isNotN.Evaluate(context).Value()).To(BeFalse())

	isNotN = NewIsNull(NewIntValue(10), false)
	g.Expect(isNotN.Evaluate(context).Value()).To(BeFalse())
}
