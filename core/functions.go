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
	"github.com/uaraven/csql/funky"
	"math"
	"strings"
)

type SqlFunction interface {
	Evaluator
	GetArgs() []Evaluator
	Location() *errors.SourceLocation
	Name() string
}

type SqlFunctionImplementor func(ctx EvaluationContext, function SqlFunction) Value

type sqlFunctionImpl struct {
	name        string
	args        []Evaluator
	loc         *errors.SourceLocation
	implementor SqlFunctionImplementor
}

func (sf sqlFunctionImpl) Name() string {
	return sf.name
}

func (sf sqlFunctionImpl) GetArgs() []Evaluator {
	return sf.args
}

func (sf sqlFunctionImpl) Location() *errors.SourceLocation {
	return sf.loc
}

func (sf sqlFunctionImpl) Evaluate(ctx EvaluationContext) Value {
	return sf.implementor(ctx, sf)
}

func (sf sqlFunctionImpl) String() string {
	argsS := funky.Map(sf.args, func(a Evaluator) string {
		return a.String()
	})
	return fmt.Sprintf("%s(%s)", sf.name, strings.Join(argsS, ", "))
}

func newSqlFunction(name string, implementor SqlFunctionImplementor, loc *errors.SourceLocation, args ...Evaluator) SqlFunction {
	return &sqlFunctionImpl{
		name:        name,
		args:        args,
		loc:         loc,
		implementor: implementor,
	}
}

func onlyArg(f SqlFunction) Evaluator {
	if len(f.GetArgs()) > 1 {
		panic(errors.NewError(f.Location(), fmt.Sprintf("%s expects one argument", f.Name())))
	}
	return f.GetArgs()[0]
}

func NewLenFunction(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("LEN", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeString {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "string", f))
		}
		return NewIntValue(int64(len(value.Value().(string))))
	}, loc, arg)
}

func NewRoundFunction(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("ROUND", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeFloat && value.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "numeric", f))
		}
		return NewIntValue(int64(math.Round(value.AsFloat().Value().(float64))))
	}, loc, arg)
}

func NewToStringFunc(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TO_STRING", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		return value.AsString()
	}, loc, arg)
}

func NewToFloatFunc(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TO_FLOAT", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		newVal := value.AsFloat()
		if newVal == nil {
			panic(errors.NewError(loc, fmt.Sprintf("Cannot convert %v to float", arg1)))
		}
		return newVal
	}, loc, arg)
}

func NewToIntFunc(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TO_INT", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		newVal := value.AsInt()
		if newVal == nil {
			panic(errors.NewError(loc, fmt.Sprintf("Cannot convert %v to int", arg1)))
		}
		return newVal
	}, loc, arg)
}

func NewTruncFunc(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TRUNC", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeFloat && value.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "numeric", f))
		}
		return NewIntValue(int64(value.AsFloat().Value().(float64)))
	}, loc, arg)
}

func NewFracFunc(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("FRAC", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeFloat && value.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "numeric", f))
		}
		fl := value.AsFloat().Value().(float64)
		return NewFloatValue(fl - math.Trunc(fl))
	}, loc, arg)
}

func NewSubStringFunc(loc *errors.SourceLocation, args ...Evaluator) SqlFunction {
	return newSqlFunction("SUBSTRING", func(ctx EvaluationContext, f SqlFunction) Value {
		args := f.GetArgs()
		if len(args) != 3 {
			panic(errors.NewError(f.Location(), "SUBSTRING function expects 3 parameters"))
		}
		st := args[0].Evaluate(ctx)
		from := args[1].Evaluate(ctx)
		to := args[2].Evaluate(ctx)
		if st.Type() != TypeString {
			panic(errors.NewTypeMismatch(f.Location(), args[0], "string", f))
		}
		if from.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), args[1], "int", f))
		}
		if to.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), args[2], "int", f))
		}
		defer func() {
			if err := recover(); err != nil {
				panic(errors.NewError(f.Location(), fmt.Sprintf("Failure in SUBSTRING function %v", err)))
			}
		}()
		s := st.Value().(string)
		f1 := int(from.Value().(int64))
		if f1 < 0 {
			f1 = 0
		}
		f2 := f1 + int(to.Value().(int64))
		if f2 > len(s) {
			f2 = len(s)
		}
		return NewStringValue(s[f1:f2])
	}, loc, args...)
}

func NewToUpperFunction(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TO_UPPER", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeString {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "string", f))
		}
		return NewStringValue(strings.ToUpper(value.Value().(string)))
	}, loc, arg)
}

func NewToLowerFunction(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("TO_LOWER", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeString {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "string", f))
		}
		return NewStringValue(strings.ToLower(value.Value().(string)))
	}, loc, arg)
}

func NewSqrtFunction(arg Evaluator, loc *errors.SourceLocation) SqlFunction {
	return newSqlFunction("SQRT", func(ctx EvaluationContext, f SqlFunction) Value {
		arg1 := onlyArg(f)
		value := arg1.Evaluate(ctx)
		if value.Type() != TypeFloat && value.Type() != TypeInt {
			panic(errors.NewTypeMismatch(f.Location(), arg1, "numeric", f))
		}
		num := value.AsFloat().Value().(float64)
		return NewFloatValue(math.Sqrt(num))
	}, loc, arg)
}

func NewPowFunction(loc *errors.SourceLocation, args ...Evaluator) SqlFunction {
	return newSqlFunction("POW", func(ctx EvaluationContext, f SqlFunction) Value {
		args := f.GetArgs()
		if len(args) != 2 {
			panic(errors.NewError(f.Location(), "POW function expects 2 parameters"))
		}
		base := args[0].Evaluate(ctx)
		pow := args[1].Evaluate(ctx)
		if base.Type() != TypeInt && base.Type() != TypeFloat {
			panic(errors.NewTypeMismatch(f.Location(), args[0], "numeric", f))
		}
		if pow.Type() != TypeInt && pow.Type() != TypeFloat {
			panic(errors.NewTypeMismatch(f.Location(), args[1], "numeric", f))
		}

		return NewFloatValue(math.Pow(base.AsFloat().Value().(float64), pow.AsFloat().Value().(float64)))
	}, loc, args...)
}

func NewCoalesceFunction(loc *errors.SourceLocation, args ...Evaluator) SqlFunction {
	return newSqlFunction("COALESCE", func(ctx EvaluationContext, f SqlFunction) Value {
		for _, arg := range f.GetArgs() {
			value := arg.Evaluate(ctx)
			if value.Type() != TypeNull {
				return value
			}
		}
		return NewNullValue()
	}, loc, args...)
}
