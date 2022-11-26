package core

type andCondition struct {
	left  Evaluator
	right Evaluator
}

func NewAnd(left Evaluator, right Evaluator) Evaluator {
	return &andCondition{
		left:  left,
		right: right,
	}
}

func (ac andCondition) Evaluate(ctx EvaluationContext) Value {
	if ac.left.Evaluate(ctx).AsBool().Value().(bool) {
		return ac.right.Evaluate(ctx).AsBool()
	} else {
		return ValueFalse
	}
}

type orCondition struct {
	left  Evaluator
	right Evaluator
}

func NewOr(left Evaluator, right Evaluator) Evaluator {
	return &orCondition{
		left:  left,
		right: right,
	}
}

func (oc orCondition) Evaluate(ctx EvaluationContext) Value {
	if oc.left.Evaluate(ctx).AsBool().Value().(bool) {
		return ValueTrue
	} else {
		return oc.right.Evaluate(ctx).AsBool()
	}
}

type notCondition struct {
	arg Evaluator
}

func NewNot(arg Evaluator) Evaluator {
	return &notCondition{arg: arg}
}

func (nc notCondition) Evaluate(ctx EvaluationContext) Value {
	return NewBoolValue(!nc.arg.Evaluate(ctx).AsBool().Value().(bool))
}
