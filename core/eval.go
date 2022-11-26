package core

import (
	"csql/funky"
	"fmt"
)

type DataType int

const (
	TypeInt DataType = iota
	TypeFloat
	TypeString
	TypeBoolean
	TypeNull
)

// EvaluationContext provides a way to retrieve value by name during evaluation
type EvaluationContext interface {
	Get(string) funky.Option[Value]
	Id() int
}

type Evaluator interface {
	Evaluate(EvaluationContext) Value
}

type Condition interface {
	Evaluator
}

type Value interface {
	Evaluator
	fmt.Stringer
	Type() DataType
	Value() interface{}
	AsString() Value
	AsInt() Value
	AsFloat() Value
	AsBool() Value
	Equals(Value) bool
	Less(Value) bool
}

func IsNumeric(v Value) bool {
	return v.Type() == TypeFloat || v.Type() == TypeInt
}
