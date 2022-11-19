package data

type andCondition struct {
	left  Evaluator
	right Evaluator
}

func NewAndCondition(left Evaluator, right Evaluator) Evaluator {
	return &andCondition{
		left:  left,
		right: right,
	}
}

func (ac andCondition) Evaluate() Value {
	if ac.left.Evaluate().AsBool().Value().(bool) {
		return ac.right.Evaluate().AsBool()
	} else {
		return ValueFalse
	}
}

type orCondition struct {
	left  Evaluator
	right Evaluator
}

func NewOrCondition(left Evaluator, right Evaluator) Evaluator {
	return &orCondition{
		left:  left,
		right: right,
	}
}

func (oc orCondition) Evaluate() Value {
	if oc.left.Evaluate().AsBool().Value().(bool) {
		return ValueTrue
	} else {
		return oc.right.Evaluate().AsBool()
	}
}

type notCondition struct {
	arg Evaluator
}

func NewNotCondition(arg Evaluator) Evaluator {
	return &notCondition{arg: arg}
}

func (nc notCondition) Evaluate() Value {
	return NewBoolValue(!nc.arg.Evaluate().AsBool().Value().(bool))
}
