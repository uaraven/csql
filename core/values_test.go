package core

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestNewStringValue(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("str")

	g.Expect(v.Type()).To(Equal(TypeString))
	g.Expect(v.Value()).To(Equal("str"))
}

func TestStringValue_AsInt(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12")
	vi := v.AsInt()

	g.Expect(vi.Type()).To(Equal(TypeInt))
	g.Expect(vi.Value()).To(Equal(int64(12)))
}

func TestStringValue_AsFloat(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12.44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(Equal(TypeFloat))
	g.Expect(vi.Value()).To(Equal(12.44))
}

func TestStringValue_AsFloatExp(t *testing.T) {
	g := NewGomegaWithT(t)

	v := NewStringValue("12e44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(Equal(TypeFloat))
	g.Expect(vi.Value()).To(Equal(12e44))
}
