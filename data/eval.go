package data

type DataType int

const (
	TypeInt DataType = iota
	TypeFloat
	TypeString
	TypeBoolean
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
