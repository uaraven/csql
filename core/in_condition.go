package core

import "csql/funky"

type inCondition struct {
	left  Evaluator
	right []Evaluator
}

func NewInList(left Evaluator, right []Evaluator) Evaluator {
	return &inCondition{
		left:  left,
		right: right,
	}
}

func (ie inCondition) Evaluate(context EvaluationContext) Value {
	return NewBoolValue(funky.AnyMatches(ie.right, func(v Evaluator) bool {
		return ie.left.Evaluate(context).Equals(v.Evaluate(context))
	}))
}
