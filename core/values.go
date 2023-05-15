/*
 *    Copyright 2022 Oleksiy Voronin
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 *    SPDX-License-Identifier: Apache-2.0
 *    SPDX-FileCopyrightText: (c) 2022 Oleksiy Voronin <ovoronin@gmail.com>
 */

package core

import (
	"fmt"
	"github.com/uaraven/csql/errors"
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
	loc   *errors.SourceLocation
}

type intValue struct {
	value int64
	loc   *errors.SourceLocation
}

type floatValue struct {
	value float64
	loc   *errors.SourceLocation
}

type boolValue struct {
	value bool
}

type nullValue struct {
	loc *errors.SourceLocation
}

func NewStringValue(v string) Value {
	return &stringValue{value: v}
}

func NewStringValueL(v string, l *errors.SourceLocation) Value {
	return &stringValue{value: v, loc: l}
}

func NewIntValue(iv int64) Value {
	return &intValue{value: iv}
}

func NewIntValueL(iv int64, l *errors.SourceLocation) Value {
	return &intValue{value: iv, loc: l}
}

func NewFloatValue(fv float64) Value {
	return &floatValue{value: fv}
}

func NewFloatValueL(fv float64, l *errors.SourceLocation) Value {
	return &floatValue{value: fv, loc: l}
}

func NewBoolValue(b bool) Value {
	return &boolValue{value: b}
}

func NewNullValue() Value {
	return nullV
}

func NewNullValueL(loc *errors.SourceLocation) Value {
	return &nullValue{loc: loc}
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
	return fmt.Sprintf("'%s'", sv.value)
}

func (sv stringValue) Repr() string {
	return sv.value
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

func (iv intValue) Repr() string {
	return iv.String()
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
	return &fv
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

func (fv floatValue) Repr() string {
	return fv.String()
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

func (bv boolValue) Repr() string {
	return bv.String()
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
	panic(errors.NewError(n.loc, fmt.Sprintf("cannot cast null to int")))
}

func (n nullValue) AsFloat() Value {
	panic(errors.NewError(n.loc, fmt.Sprintf("cannot cast null to float")))
}

func (n nullValue) AsBool() Value {
	panic(errors.NewError(n.loc, fmt.Sprintf("cannot cast null to bool")))
}

func (n nullValue) String() string {
	return NullValueString
}

func (n nullValue) Repr() string {
	return n.String()
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
	location   *errors.SourceLocation
}

func (rv rowValue) String() string {
	return rv.identifier
}

func NewRowValue(identifier string) Evaluator {
	return &rowValue{
		identifier: identifier,
	}
}

func NewRowValueL(identifier string, loc *errors.SourceLocation) Evaluator {
	return &rowValue{
		identifier: identifier,
		location:   loc,
	}
}

func (rv rowValue) Evaluate(ctx EvaluationContext) Value {
	val := ctx.Get(rv.identifier)
	if val.IsPresent() {
		return val.Value()
	} else {
		panic(errors.NewError(rv.location, fmt.Sprintf("unknown column %s", rv.identifier)))
	}
}

func (rv rowValue) Identifier() string {
	return rv.identifier
}

func EnsureString(v Value) (string, error) {
	if v.Type() != TypeString {
		return "", fmt.Errorf(errors.TypeMismatchMsg(v, "string", nil))
	}
	return v.Value().(string), nil
}

func EnsureInt(v Value) (int64, error) {
	if v.Type() != TypeInt && v.Type() != TypeFloat {
		return 0, fmt.Errorf(errors.TypeMismatchMsg(v, "integer", nil))
	}
	return v.AsInt().Value().(int64), nil
}

func EnsureFloat(v Value) (float64, error) {
	if v.Type() != TypeInt && v.Type() != TypeFloat {
		return 0, fmt.Errorf(errors.TypeMismatchMsg(v, "float", nil))
	}
	return v.AsFloat().Value().(float64), nil
}

func DecodeNumericL(value string, l *errors.SourceLocation) Value {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(errors.NewTypeMismatch(l, value, "number", ""))
		}
		return NewFloatValueL(f64, l)
	}
	return NewIntValueL(i64, l)
}

// ParseStringToValue parses a string and converts it into a value of appropriate type
func ParseStringToValue(value string) Value {
	i64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		f64, err := strconv.ParseFloat(value, 64)
		if err != nil {
			if value == NullValueString {
				return NewNullValue()
			}
			return NewStringValue(value)
		}
		return NewFloatValue(f64)
	}
	return NewIntValue(i64)
}
