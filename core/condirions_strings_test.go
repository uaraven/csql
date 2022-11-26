package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestMatchCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()

	match := NewMatchCondition(NewStringValue("John Connor"), NewStringValue("^J[oh]+n .o.*r$"))

	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))

	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("^J[oh]+n .o.*r$"))

	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
}

func TestLikeCondition_Evaluate(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()

	match := NewLikeCondition(NewStringValue("John Connor"), NewStringValue("J__n%nor"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))
	match = NewLikeCondition(NewStringValue("Jean Tenor"), NewStringValue("J__n%nor"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueTrue))

	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("Joh%Con%"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
	match = NewMatchCondition(NewStringValue("John Mellenkamp"), NewStringValue("John%"))
	g.Expect(match.Evaluate(context)).To(Equal(ValueFalse))
}
