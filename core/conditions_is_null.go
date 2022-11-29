package core

type isNullCondition struct {
	expr   Evaluator
	negate bool
}

func NewIsNull(expr Evaluator, negate bool) Evaluator {
	return &isNullCondition{
		expr:   expr,
		negate: negate,
	}
}

func (inc isNullCondition) Evaluate(context EvaluationContext) Value {
	result := inc.expr.Evaluate(context)
	if inc.negate {
		return NewBoolValue(result != nullV)
	} else {
		return NewBoolValue(result == nullV)
	}
}
