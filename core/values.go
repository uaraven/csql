package core

import (
	"fmt"
	"strconv"
)

var (
	NullValueString = "null"
	nullV           = &nullValue{}
	ValueFalse      = NewBoolValue(false)
	ValueTrue       = NewBoolValue(true)
)

type stringValue struct {
	value string
}

type intValue struct {
	value int64
}

type floatValue struct {
	value float64
}

type boolValue struct {
	value bool
}

type nullValue struct {
}

func NewStringValue(v string) Value {
	return &stringValue{value: v}
}

func NewIntValue(iv int64) Value {
	return &intValue{value: iv}
}

func NewFloatValue(fv float64) Value {
	return &floatValue{value: fv}
}

func NewBoolValue(b bool) Value {
	return &boolValue{value: b}
}

func NewNullValue() Value {
	return nullV
}

func (sv stringValue) Type() DataType {
	return TypeString
}

func (sv stringValue) Value() interface{} {
	return sv.value
}

func (sv stringValue) AsString() Value {
	return &sv
}

func (sv stringValue) AsInt() Value {
	iv, err := strconv.ParseInt(sv.value, 10, 64)
	if err == nil {
		return NewIntValue(iv)
	} else {
		return nil
	}
}

func (sv stringValue) AsFloat() Value {
	fv, err := strconv.ParseFloat(sv.value, 64)
	if err == nil {
		return NewFloatValue(fv)
	} else {
		return nil
	}
}

func (sv stringValue) AsBool() Value {
	b, err := strconv.ParseBool(sv.value)
	if err == nil {
		return NewBoolValue(b)
	} else {
		return nil
	}
}

func (sv stringValue) String() string {
	return fmt.Sprintf("\"%s\"", sv.value)
}

func (sv stringValue) Equals(other Value) bool {
	return other.Type() == sv.Type() && other.Value().(string) == sv.value
}

func (sv stringValue) Less(other Value) bool {
	return other.Type() == sv.Type() && other.Value().(string) > sv.value
}

func (sv stringValue) Evaluate(_ EvaluationContext) Value {
	return &sv
}

func (iv intValue) Type() DataType {
	return TypeInt
}

func (iv intValue) Value() interface{} {
	return iv.value
}

func (iv intValue) AsString() Value {
	return NewStringValue(strconv.FormatInt(iv.value, 10))
}

func (iv intValue) AsInt() Value {
	return &iv
}

func (iv intValue) AsFloat() Value {
	return NewFloatValue(float64(iv.value))
}

func (iv intValue) AsBool() Value {
	return NewBoolValue(iv.value != 0)
}

func (iv intValue) String() string {
	return fmt.Sprintf("%d", iv.value)
}

func (iv intValue) Equals(other Value) bool {
	return other.Type() == iv.Type() && other.Value().(int64) == iv.value
}

func (iv intValue) Less(other Value) bool {
	return other.Type() == iv.Type() && other.Value().(int64) > iv.value
}

func (iv intValue) Evaluate(_ EvaluationContext) Value {
	return &iv
}

func (fv floatValue) Type() DataType {
	return TypeFloat
}

func (fv floatValue) Value() interface{} {
	return fv.value
}

func (fv floatValue) AsString() Value {
	return NewStringValue(strconv.FormatFloat(fv.value, 'g', -1, 64))
}

func (fv floatValue) AsInt() Value {
	return NewIntValue(int64(fv.value))
}

func (fv floatValue) AsFloat() Value {
	return fv
}

func (fv floatValue) AsBool() Value {
	return NewBoolValue(fv.value != 0)
}

func (fv floatValue) Evaluate(_ EvaluationContext) Value {
	return fv
}

func (fv floatValue) String() string {
	return fmt.Sprintf("%f", fv.value)
}

func (fv floatValue) Equals(other Value) bool {
	return other.Type() == fv.Type() && other.Value().(float64) == fv.value
}

func (fv floatValue) Less(other Value) bool {
	return other.Type() == fv.Type() && other.Value().(float64) > fv.value
}

func (bv boolValue) Type() DataType {
	return TypeBoolean
}

func (bv boolValue) Value() interface{} {
	return bv.value
}

func (bv boolValue) AsString() Value {
	return NewStringValue(fmt.Sprintf("%t", bv.value))
}

func (bv boolValue) AsInt() Value {
	if bv.value {
		return NewIntValue(1)
	} else {
		return NewIntValue(0)
	}
}

func (bv boolValue) AsFloat() Value {
	if bv.value {
		return NewFloatValue(1.0)
	} else {
		return NewFloatValue(0.0)
	}
}

func (bv boolValue) AsBool() Value {
	return &bv
}

func (bv boolValue) String() string {
	return fmt.Sprintf("%t", bv.value)
}

func (bv boolValue) Evaluate(_ EvaluationContext) Value {
	return &bv
}

func (bv boolValue) Equals(other Value) bool {
	return other.Type() == bv.Type() && other.Value().(bool) == bv.value
}

func (bv boolValue) Less(other Value) bool {
	return false
}

func (n nullValue) Evaluate(_ EvaluationContext) Value {
	return &n
}

func (n nullValue) Type() DataType {
	return TypeNull
}

func (n nullValue) Value() interface{} {
	return nil
}

func (n nullValue) AsString() Value {
	return NewStringValue(NullValueString)
}

func (n nullValue) AsInt() Value {
	panic(fmt.Sprintf("cannot cast null to int"))
}

func (n nullValue) AsFloat() Value {
	panic(fmt.Sprintf("cannot cast null to float"))
}

func (n nullValue) AsBool() Value {
	panic(fmt.Sprintf("cannot cast null to bool"))
}

func (n nullValue) String() string {
	return NullValueString
}

func (n nullValue) Equals(other Value) bool {
	return false
}

func (n nullValue) Less(other Value) bool {
	return false
}

type Identifiable interface {
	Identifier() string
}

type rowValue struct {
	identifier string
}

func NewRowValue(identifier string) Evaluator {
	return &rowValue{
		identifier: identifier,
	}
}

func (rv rowValue) Evaluate(ctx EvaluationContext) Value {
	val := ctx.Get(rv.identifier)
	if val.IsPresent() {
		return val.Value()
	} else {
		panic(fmt.Errorf("unknown column %s", rv.identifier))
	}
}

func (rv rowValue) Identifier() string {
	return rv.identifier
}

func EnsureString(v Value) Value {
	if v.Type() != TypeString {
		panic(fmt.Sprintf("expected %v to have type 'string'", v))
	}
	return v
}
