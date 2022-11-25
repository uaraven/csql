package core

import (
	. "github.com/onsi/gomega"
	"math"
	"testing"
)

func TestArithExpression_EvaluateAdd(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("ten", "10")
	context.SetValue("ten_four", "10.4")

	expr := NewExpression(addition, NewIntValue(10), NewIntValue(2))
	result := expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(12)))

	expr = NewExpression(addition, NewRowValue("ten"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(14)))

	expr = NewExpression(addition, NewRowValue("ten_four"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(14.4))
}

func TestArithExpression_EvaluateSub(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("ten", "10")
	context.SetValue("ten_four", "10.4")

	expr := NewExpression(subtraction, NewIntValue(10), NewIntValue(2))
	result := expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(8)))

	expr = NewExpression(subtraction, NewRowValue("ten"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(6)))

	expr = NewExpression(subtraction, NewRowValue("ten_four"), NewIntValue(-2))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(12.4))
}

func TestArithExpression_EvaluateMul(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("ten", "10")
	context.SetValue("ten_four", "10.4")

	expr := NewExpression(multiplication, NewIntValue(10), NewIntValue(2))
	result := expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(20)))

	expr = NewExpression(multiplication, NewRowValue("ten"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(40)))

	expr = NewExpression(multiplication, NewRowValue("ten_four"), NewIntValue(1))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(10.4))

	expr = NewExpression(multiplication, NewRowValue("ten_four"), NewFloatValue(0.5))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(5.2))
}

func TestArithExpression_EvaluateDiv(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("ten", "10")
	context.SetValue("ten_four", "10.4")

	expr := NewExpression(division, NewIntValue(10), NewIntValue(2))
	result := expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(5)))

	expr = NewExpression(division, NewRowValue("ten"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(2)))

	expr = NewExpression(division, NewRowValue("ten_four"), NewIntValue(2))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(5.2))

	expr = NewExpression(division, NewRowValue("ten_four"), NewFloatValue(0.5))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(20.8))

	expr = NewExpression(division, NewRowValue("ten_four"), NewFloatValue(0))
	result = expr.Evaluate(context)
	g.Expect(result.Type()).To(Equal(TypeFloat))
	g.Expect(result.Value()).To(Equal(math.Inf(1)))
}

func TestArithExpression_EvaluateMod(t *testing.T) {
	g := NewGomegaWithT(t)
	context := newMapContext()
	context.SetValue("ten", "10")
	context.SetValue("ten_four", "10.4")

	expr := NewExpression(modulo, NewIntValue(10), NewIntValue(2))
	result := expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(0)))

	expr = NewExpression(modulo, NewRowValue("ten"), NewIntValue(4))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(int64(2)))

	expr = NewExpression(modulo, NewRowValue("ten_four"), NewIntValue(6))
	result = expr.Evaluate(context)
	g.Expect(result.Value()).To(Equal(4.0))

}
