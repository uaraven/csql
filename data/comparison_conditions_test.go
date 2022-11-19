package data

import (
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

type condMaker func(Value, Value) Condition

func cond(maker condMaker, left interface{}, right interface{}) Condition {
	switch l := left.(type) {
	case string:
		return maker(NewStringValue(l), NewStringValue(right.(string)))
	case int:
		return maker(NewIntValue(int64(l)), NewIntValue(int64(right.(int))))
	case int64:
		return maker(NewIntValue(l), NewIntValue(right.(int64)))
	case float64:
		return maker(NewFloatValue(l), NewFloatValue(right.(float64)))
	default:
		panic(fmt.Errorf("unsupported type: %v", left))
	}
}

type conditionTest struct {
	maker   condMaker
	left    interface{}
	right   interface{}
	matcher gomega.OmegaMatcher
}

func test(maker condMaker, left interface{}, right interface{}, matcher gomega.OmegaMatcher) conditionTest {
	return conditionTest{
		maker:   maker,
		left:    left,
		right:   right,
		matcher: matcher,
	}
}

var (
	eqTests = []conditionTest{
		test(NewEq, "s1", "s1", gomega.BeTrue()),
		test(NewEq, "s1", "s2", gomega.BeFalse()),
		test(NewEq, 10, 10, gomega.BeTrue()),
		test(NewEq, 10, -10, gomega.BeFalse()),
		test(NewEq, 10.1, 10.1, gomega.BeTrue()),
		test(NewEq, 10.1, 10.2, gomega.BeFalse()),
	}
	neqTests = []conditionTest{
		test(NewNeq, "s1", "s1", gomega.BeFalse()),
		test(NewNeq, "s1", "s2", gomega.BeTrue()),
		test(NewNeq, 10, 10, gomega.BeFalse()),
		test(NewNeq, 10, -10, gomega.BeTrue()),
		test(NewNeq, 42.42, 42.42, gomega.BeFalse()),
		test(NewNeq, 10.2, 10.1, gomega.BeTrue()),
	}
	gtTests = []conditionTest{
		test(NewGt, "s1", "s2", gomega.BeFalse()),
		test(NewGt, "s2", "s1", gomega.BeTrue()),
		test(NewGt, 10, 10, gomega.BeFalse()),
		test(NewGt, 11, 10, gomega.BeTrue()),
		test(NewGt, 0.42, 42.0, gomega.BeFalse()),
		test(NewGt, 10.2, 10.1, gomega.BeTrue()),
	}
	ltTests = []conditionTest{
		test(NewLt, "s1", "s2", gomega.BeTrue()),
		test(NewLt, "s2", "s1", gomega.BeFalse()),
		test(NewLt, 9, 10, gomega.BeTrue()),
		test(NewLt, 11, 10, gomega.BeFalse()),
		test(NewLt, 0.42, 42.0, gomega.BeTrue()),
		test(NewLt, 10.2, 10.1, gomega.BeFalse()),
	}
	gteTests = []conditionTest{
		test(NewGte, "s1", "s2", gomega.BeFalse()),
		test(NewGte, "s2", "s1", gomega.BeTrue()),
		test(NewGte, "s1", "s1", gomega.BeTrue()),
		test(NewGte, 10, 10, gomega.BeTrue()),
		test(NewGte, 11, 10, gomega.BeTrue()),
		test(NewGte, 9, 10, gomega.BeFalse()),
		test(NewGte, 0.42, 42.0, gomega.BeFalse()),
		test(NewGte, 10.2, 10.1, gomega.BeTrue()),
		test(NewGte, 10.2, 10.2, gomega.BeTrue()),
	}
	lteTests = []conditionTest{
		test(NewLte, "s1", "s2", gomega.BeTrue()),
		test(NewLte, "s1", "s1", gomega.BeTrue()),
		test(NewLte, "s2", "s1", gomega.BeFalse()),
		test(NewLte, 9, 10, gomega.BeTrue()),
		test(NewLte, 10, 10, gomega.BeTrue()),
		test(NewLte, 11, 10, gomega.BeFalse()),
		test(NewLte, 0.42, 42.0, gomega.BeTrue()),
		test(NewLte, 42.0, 42.0, gomega.BeTrue()),
		test(NewLte, 10.2, 10.1, gomega.BeFalse()),
	}
)

func testSet(g gomega.Gomega, testSet []conditionTest) {
	for _, test := range testSet {
		cc := cond(test.maker, test.left, test.right)
		g.Expect(cc.Evaluate().Value()).To(test.matcher, fmt.Sprintf("Condition %s", cc))
	}
}

func Test_EqConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, eqTests)
}

func Test_NeqConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, neqTests)
}

func Test_GtConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, gtTests)
}

func Test_LtConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, ltTests)
}

func Test_GteConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, gteTests)
}

func Test_LteConditions(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	testSet(g, lteTests)
}
