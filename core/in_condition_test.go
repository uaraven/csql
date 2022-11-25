package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestInCondition_EvaluateNotFound(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("a", "12")

	in := NewInList(NewIntValue(8), []Evaluator{NewIntValue(1), NewIntValue(2), NewRowValue("a")})
	g.Expect(in.Evaluate(context).Value()).To(BeFalse())
}

func TestInCondition_EvaluateFound(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("a", "12")

	in := NewInList(NewIntValue(12), []Evaluator{NewIntValue(1), NewIntValue(2), NewRowValue("a")})
	g.Expect(in.Evaluate(context).Value()).To(BeTrue())

	in = NewInList(NewIntValue(2), []Evaluator{NewIntValue(1), NewIntValue(2), NewRowValue("a")})
	g.Expect(in.Evaluate(context).Value()).To(BeTrue())
}

func TestInCondition_EvaluateFoundMixed(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("a", "12")
	context.SetValue("b", "thirteen")

	in := NewInList(NewIntValue(12), []Evaluator{NewIntValue(1), NewIntValue(2), NewRowValue("a"), NewRowValue("b")})
	g.Expect(in.Evaluate(context).Value()).To(BeTrue())

	in = NewInList(NewStringValue("b"), []Evaluator{NewIntValue(1), NewIntValue(2), NewRowValue("a"), NewRowValue("b")})
	g.Expect(in.Evaluate(context).Value()).To(BeFalse())
}
