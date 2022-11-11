package data

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestNewStringValue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	v := NewStringValue("str")

	g.Expect(v.Type()).To(gomega.Equal(StringType))
	g.Expect(v.Value()).To(gomega.Equal("str"))
}

func TestStringValue_AsInt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	v := NewStringValue("12")
	vi := v.AsInt()

	g.Expect(vi.Type()).To(gomega.Equal(IntType))
	g.Expect(vi.Value()).To(gomega.Equal(int64(12)))
}

func TestStringValue_AsFloat(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	v := NewStringValue("12.44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(gomega.Equal(FloatType))
	g.Expect(vi.Value()).To(gomega.Equal(12.44))
}

func TestStringValue_AsFloatExp(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	v := NewStringValue("12e44")
	vi := v.AsFloat()

	g.Expect(vi.Type()).To(gomega.Equal(FloatType))
	g.Expect(vi.Value()).To(gomega.Equal(12e44))
}
