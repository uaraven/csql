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

package errors

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CsqlError struct {
	Location *SourceLocation
	Message  string
}

func (ce CsqlError) Error() string {
	return ce.Message
}

func NewError(where *SourceLocation, message string) error {
	return &CsqlError{Location: where, Message: message}
}

func UnknownColumnError(location *SourceLocation, columnName string) error {
	return &CsqlError{Location: location, Message: fmt.Sprintf("Unknown column name: %s", columnName)}
}

type SourceLocation struct {
	Line   int
	Column int
}

func (sl SourceLocation) String() string {
	return fmt.Sprintf("[%d:%d]", sl.Line, sl.Column)
}

func Loc(line, column int) *SourceLocation {
	return &SourceLocation{
		Line:   line,
		Column: column,
	}
}

func SLFromToken(t antlr.Token) *SourceLocation {
	return &SourceLocation{
		Line:   t.GetLine(),
		Column: t.GetColumn(),
	}
}
