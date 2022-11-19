package data

type DataType int

const (
	IntType DataType = iota
	FloatType
	StringType
	BooleanType
)

type Evaluator interface {
	Evaluate() Value
}

type Condition interface {
	Evaluator
}

type Value interface {
	Evaluator
	Type() DataType
	Value() interface{}
	AsString() Value
	AsInt() Value
	AsFloat() Value
	AsBool() Value
}
