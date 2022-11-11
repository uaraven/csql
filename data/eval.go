package data

type DataType int

const (
	IntType DataType = iota
	FloatType
	StringType
	BooleanType
)

type Value interface {
	Type() DataType
	Value() interface{}
	AsString() Value
	AsInt() Value
	AsFloat() Value
}

type Condition interface {
	Evaluate() bool
}
