package util

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestEqualsIgnoreCase(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(EqualsIgnoreCase("string1", "StRiNg1")).To(BeTrue())
	g.Expect(EqualsIgnoreCase("string1", "StRiNg2")).To(BeFalse())
}
