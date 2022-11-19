package data

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

type condMaker func(Evaluator, Evaluator) Condition

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
	matcher OmegaMatcher
}

func test(maker condMaker, left interface{}, right interface{}, matcher OmegaMatcher) conditionTest {
	return conditionTest{
		maker:   maker,
		left:    left,
		right:   right,
		matcher: matcher,
	}
}

var (
	eqTests = []conditionTest{
		test(NewEq, "s1", "s1", BeTrue()),
		test(NewEq, "s1", "s2", BeFalse()),
		test(NewEq, 10, 10, BeTrue()),
		test(NewEq, 10, -10, BeFalse()),
		test(NewEq, 10.1, 10.1, BeTrue()),
		test(NewEq, 10.1, 10.2, BeFalse()),
	}
	neqTests = []conditionTest{
		test(NewNeq, "s1", "s1", BeFalse()),
		test(NewNeq, "s1", "s2", BeTrue()),
		test(NewNeq, 10, 10, BeFalse()),
		test(NewNeq, 10, -10, BeTrue()),
		test(NewNeq, 42.42, 42.42, BeFalse()),
		test(NewNeq, 10.2, 10.1, BeTrue()),
	}
	gtTests = []conditionTest{
		test(NewGt, "s1", "s2", BeFalse()),
		test(NewGt, "s2", "s1", BeTrue()),
		test(NewGt, 10, 10, BeFalse()),
		test(NewGt, 11, 10, BeTrue()),
		test(NewGt, 0.42, 42.0, BeFalse()),
		test(NewGt, 10.2, 10.1, BeTrue()),
	}
	ltTests = []conditionTest{
		test(NewLt, "s1", "s2", BeTrue()),
		test(NewLt, "s2", "s1", BeFalse()),
		test(NewLt, 9, 10, BeTrue()),
		test(NewLt, 11, 10, BeFalse()),
		test(NewLt, 0.42, 42.0, BeTrue()),
		test(NewLt, 10.2, 10.1, BeFalse()),
	}
	gteTests = []conditionTest{
		test(NewGte, "s1", "s2", BeFalse()),
		test(NewGte, "s2", "s1", BeTrue()),
		test(NewGte, "s1", "s1", BeTrue()),
		test(NewGte, 10, 10, BeTrue()),
		test(NewGte, 11, 10, BeTrue()),
		test(NewGte, 9, 10, BeFalse()),
		test(NewGte, 0.42, 42.0, BeFalse()),
		test(NewGte, 10.2, 10.1, BeTrue()),
		test(NewGte, 10.2, 10.2, BeTrue()),
	}
	lteTests = []conditionTest{
		test(NewLte, "s1", "s2", BeTrue()),
		test(NewLte, "s1", "s1", BeTrue()),
		test(NewLte, "s2", "s1", BeFalse()),
		test(NewLte, 9, 10, BeTrue()),
		test(NewLte, 10, 10, BeTrue()),
		test(NewLte, 11, 10, BeFalse()),
		test(NewLte, 0.42, 42.0, BeTrue()),
		test(NewLte, 42.0, 42.0, BeTrue()),
		test(NewLte, 10.2, 10.1, BeFalse()),
	}
)

func testSet(g Gomega, testSet []conditionTest) {
	c := newMapContext()
	for _, test := range testSet {
		cc := cond(test.maker, test.left, test.right)
		g.Expect(cc.Evaluate(c).Value()).To(test.matcher, fmt.Sprintf("Condition %s", cc))
	}
}

func Test_EqConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, eqTests)
}

func Test_NeqConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, neqTests)
}

func Test_GtConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, gtTests)
}

func Test_LtConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, ltTests)
}

func Test_GteConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, gteTests)
}

func Test_LteConditions(t *testing.T) {
	g := NewGomegaWithT(t)

	testSet(g, lteTests)
}
