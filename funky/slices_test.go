package funky

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestFilter(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	result := Filter(list, func(i int) bool { return (i % 2) == 0 })

	g.Expect(result).To(gomega.ContainElements(2, 4, 6))
}

func TestFind_Success(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	idx, result := Find(list, func(i int) bool { return i == 4 })

	g.Expect(idx).To(gomega.Equal(3))

	g.Expect(result.IsPresent()).To(gomega.BeTrue())
	g.Expect(result.Value()).To(gomega.Equal(4))
}

func TestFind_NotFound(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	idx, result := Find(list, func(i int) bool { return i == 10 })

	g.Expect(idx).To(gomega.Equal(-1))

	g.Expect(result.IsPresent()).To(gomega.BeFalse())
}

func TestMap(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	list := []int{1, 2, 3, 4, 5, 6}
	result := Map(list, func(i int) int { return i * 2 })

	g.Expect(result).To(gomega.ContainElements(2, 4, 6, 8, 10, 12))
}
