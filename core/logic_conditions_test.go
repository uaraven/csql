package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

var c = newMapContext()

func TestAndCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewAnd(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAnd(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
}

func TestOrCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewOr(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOr(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
}

func TestNotCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewNot(ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewNot(ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
}
