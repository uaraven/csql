package data

import (
	"fmt"
	"strconv"
	"sync"
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

func (rv stringValue) Type() DataType {
	return TypeString
}

func (rv stringValue) Value() interface{} {
	return rv.value
}

func (rv stringValue) AsString() Value {
	return &rv
}

func (rv stringValue) AsInt() Value {
	iv, err := strconv.ParseInt(rv.value, 10, 64)
	if err == nil {
		return NewIntValue(iv)
	} else {
		return nil
	}
}

func (rv stringValue) AsFloat() Value {
	fv, err := strconv.ParseFloat(rv.value, 64)
	if err == nil {
		return NewFloatValue(fv)
	} else {
		return nil
	}
}

func (rv stringValue) AsBool() Value {
	b, err := strconv.ParseBool(rv.value)
	if err == nil {
		return NewBoolValue(b)
	} else {
		return nil
	}
}

func (rv stringValue) String() string {
	return fmt.Sprintf("\"%s\"", rv.value)
}

func (rv stringValue) Evaluate() Value {
	return &rv
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

func (iv intValue) Evaluate() Value {
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

func (fv floatValue) Evaluate() Value {
	return fv
}

func (fv floatValue) String() string {
	return fmt.Sprintf("%f", fv.value)
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

func (bv boolValue) Evaluate() Value {
	return &bv
}

type rowValue struct {
	lock       sync.Mutex
	row        Row
	identifier string
	value      Value
}

func NewRowValue(row Row, identifier string) Evaluator {
	return &rowValue{
		row:        row,
		identifier: identifier,
		value:      nil,
	}
}

func (rv *rowValue) Evaluate() Value {
	rv.lock.Lock()
	defer rv.lock.Unlock()
	if rv.value == nil {
		var err error
		rv.value, err = rv.getValue()
		if err != nil {
			panic(err)
		}
	}
	return rv.value
}

func (rv *rowValue) getValue() (Value, error) {
	val := rv.row.GetByName(rv.identifier)
	if val.IsPresent() {
		switch v := decode(val.Value()).(type) {
		case string:
			return NewStringValue(v), nil
		case int64:
			return NewIntValue(v), nil
		case float64:
			return NewFloatValue(v), nil
		default:
			return nil, fmt.Errorf("unsupported data type: %v", v)

		}
	} else {
		return nil, fmt.Errorf("unknown column %s", rv.identifier)
	}
}

func decode(value string) interface{} {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return value
		}
		return f64
	}
	return i64
}

var (
	ValueFalse = NewBoolValue(false)
	ValueTrue  = NewBoolValue(true)
)
