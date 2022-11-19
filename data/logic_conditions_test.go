package data

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestAndCondition_Evaluate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	g.Expect(NewAndCondition(ValueTrue, ValueFalse).Evaluate()).To(gomega.Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueTrue).Evaluate()).To(gomega.Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueFalse, ValueFalse).Evaluate()).To(gomega.Equal(ValueFalse))
	g.Expect(NewAndCondition(ValueTrue, ValueTrue).Evaluate()).To(gomega.Equal(ValueTrue))
}

func TestOrCondition_Evaluate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	g.Expect(NewOrCondition(ValueTrue, ValueFalse).Evaluate()).To(gomega.Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueTrue).Evaluate()).To(gomega.Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueTrue, ValueTrue).Evaluate()).To(gomega.Equal(ValueTrue))
	g.Expect(NewOrCondition(ValueFalse, ValueFalse).Evaluate()).To(gomega.Equal(ValueFalse))
}

func TestNotCondition_Evaluate(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	g.Expect(NewNotCondition(ValueTrue).Evaluate()).To(gomega.Equal(ValueFalse))
	g.Expect(NewNotCondition(ValueFalse).Evaluate()).To(gomega.Equal(ValueTrue))
}
