package funky

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSome(t *testing.T) {
	g := NewGomegaWithT(t)
	expected := 12
	v := SomeOf(expected)
	g.Expect(v.IsEmpty()).To(BeFalse(), "expected IsEmpty() to return false")
	g.Expect(v.IsPresent()).To(BeTrue(), "expected IsPresent() to return true")
	g.Expect(v.Value()).To(Equal(expected))
	g.Expect(v.OrGet(10)).To(Equal(expected))
	var executed bool = false
	v.IfPresent(func(value int) {
		g.Expect(value).To(Equal(expected))
		executed = true
	})
	g.Expect(executed).To(BeTrue())
}

func TestNone(t *testing.T) {
	g := NewGomegaWithT(t)
	v := NoneOf[int]()
	g.Expect(v.IsEmpty()).To(BeTrue(), "expected IsEmpty() to return true")
	g.Expect(v.IsPresent()).To(BeFalse(), "expected IsPresent() to return false")
	g.Expect(func() { v.Value() }).To(Panic())

	expected := 10
	g.Expect(v.OrGet(expected)).To(Equal(expected))
	v.IfPresent(func(value int) {
		t.Fatal("Should not execute")
	})
}
