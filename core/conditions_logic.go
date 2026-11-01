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
)

type andCondition struct {
	left  Evaluator
	right Evaluator
}

func NewAnd(left Evaluator, right Evaluator) Evaluator {
	return &andCondition{
		left:  left,
		right: right,
	}
}

func (ac andCondition) String() string {
	return fmt.Sprintf("%s AND %s", ac.left, ac.right)
}

// truth interprets a value as a boolean. known is false for NULL (unknown)
func truth(v Value) (value bool, known bool) {
	if v.Type() == TypeNull {
		return false, false
	}
	b := v.AsBool()
	if b == nil {
		panic(errors.NewError(nil, fmt.Sprintf("cannot use %v as a boolean condition", v)))
	}
	return b.Value().(bool), true
}

func (ac andCondition) Evaluate(ctx EvaluationContext) Value {
	lv, lKnown := truth(ac.left.Evaluate(ctx))
	if lKnown && !lv {
		return ValueFalse
	}
	rv, rKnown := truth(ac.right.Evaluate(ctx))
	if rKnown && !rv {
		return ValueFalse
	}
	if lKnown && rKnown {
		return ValueTrue
	}
	return NewNullValue()
}

type orCondition struct {
	left  Evaluator
	right Evaluator
}

func NewOr(left Evaluator, right Evaluator) Evaluator {
	return &orCondition{
		left:  left,
		right: right,
	}
}

func (oc orCondition) String() string {
	return fmt.Sprintf("%s OR %s", oc.left, oc.right)
}

func (oc orCondition) Evaluate(ctx EvaluationContext) Value {
	lv, lKnown := truth(oc.left.Evaluate(ctx))
	if lKnown && lv {
		return ValueTrue
	}
	rv, rKnown := truth(oc.right.Evaluate(ctx))
	if rKnown && rv {
		return ValueTrue
	}
	if lKnown && rKnown {
		return ValueFalse
	}
	return NewNullValue()
}

type notCondition struct {
	arg Evaluator
}

func NewNot(arg Evaluator) Evaluator {
	return &notCondition{arg: arg}
}

func (nc notCondition) String() string {
	return fmt.Sprintf("NOT %s", nc.arg)
}

func (nc notCondition) Evaluate(ctx EvaluationContext) Value {
	v, known := truth(nc.arg.Evaluate(ctx))
	if !known {
		return NewNullValue()
	}
	return NewBoolValue(!v)
}
