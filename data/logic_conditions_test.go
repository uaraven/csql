package data

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestAndCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewAndCondition(ValueTrue, ValueFalse).Evaluate()).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueTrue).Evaluate()).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueFalse).Evaluate()).To(Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueTrue, ValueTrue).Evaluate()).To(Equal(ValueTrue))
}

func TestOrCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewOrCondition(ValueTrue, ValueFalse).Evaluate()).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueTrue).Evaluate()).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueTrue, ValueTrue).Evaluate()).To(Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueFalse).Evaluate()).To(Equal(ValueFalse))
}

func TestNotCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(NewNotCondition(ValueTrue).Evaluate()).To(Equal(ValueFalse))
	g.Expect(NewNotCondition(ValueFalse).Evaluate()).To(Equal(ValueTrue))
}
