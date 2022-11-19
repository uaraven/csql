package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

var c = newMapContext()

func TestAndCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewAndCondition(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
}

func TestOrCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewOrCondition(ValueTrue, ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueTrue, ValueTrue).Evaluate(c)).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueFalse).Evaluate(c)).To(Equal(ValueFalse))
}

func TestNotCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewNotCondition(ValueTrue).Evaluate(c)).To(Equal(ValueFalse))
	g.Expect(NewNotCondition(ValueFalse).Evaluate(c)).To(Equal(ValueTrue))
}
