package data

import (
	"fmt"
	"strconv"
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

func NewStringValue(v string) Value {
	return &stringValue{value: v}
}

func NewIntValue(iv int64) Value {
	return &intValue{value: iv}
}

func NewFloatValue(fv float64) Value {
	return &floatValue{value: fv}
}

func (rv stringValue) Type() DataType {
	return StringType
}

func (rv stringValue) Value() interface{} {
	return rv.value
}

func (rv stringValue) AsString() Value {
	return rv
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

func (rv stringValue) String() string {
	return fmt.Sprintf("\"%s\"", rv.value)
}

func (iv intValue) Type() DataType {
	return IntType
}

func (iv intValue) Value() interface{} {
	return iv.value
}

func (iv intValue) AsString() Value {
	return NewStringValue(strconv.FormatInt(iv.value, 10))
}

func (iv intValue) AsInt() Value {
	return iv
}

func (iv intValue) AsFloat() Value {
	return NewFloatValue(float64(iv.value))
}

func (iv intValue) String() string {
	return fmt.Sprintf("%d", iv.value)
}

func (fv floatValue) Type() DataType {
	return FloatType
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

func (fv floatValue) String() string {
	return fmt.Sprintf("%f", fv.value)
}

func NewRowValue(row Row, identifier string) (Value, error) {
	val := row.GetByName(identifier)
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
		return nil, fmt.Errorf("unknown column %s", identifier)
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
