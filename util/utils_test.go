package util

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEqualsIgnoreCase(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(EqualsIgnoreCase("string1", "StRiNg1")).To(BeTrue())
	g.Expect(EqualsIgnoreCase("string1", "StRiNg2")).To(BeFalse())
}

func TestMust(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(Must(15, nil)).To(Equal(15))

	g.Expect(func() { Must(15, fmt.Errorf("error")) }).To(PanicWith(fmt.Errorf("error")))
}
