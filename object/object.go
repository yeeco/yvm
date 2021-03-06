/*
 * // Copyright (C) 2017 gyee authors
 * //
 * // This file is part of the gyee library.
 * //
 * // the gyee library is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU General Public License as published by
 * // the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // the gyee library is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * // GNU General Public License for more details.
 * //
 * // You should have received a copy of the GNU General Public License
 * // along with the gyee library.  If not, see <http://www.gnu.org/licenses/>.
 *
 *
 */

package object

import (
	"bytes"
	"fmt"
	"github.com/yeeco/yvm/ast"
	"github.com/yeeco/yvm/code"
	"strings"
)

//const (
//	NULL_OBJ         = "NULL"
//	ERROR_OBJ        = "ERROR"
//	INTEGER_OBJ      = "INTEGER"
//	BOOLEAN_OBJ      = "BOOLEAN"
//	STRING_OBJ       = "STRING"
//	RETURN_VALUE_OBJ = "RETURN_VALUE"
//	FUNCTION_OBJ     = "FUNCTION"
//	BUILTIN_OBJ      = "BUILTIN"
//	ARRAY_OBJ        = "ARRAY"
//
//	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"
//	CLOSURE_OBJ           = "CLOSURE"
//)
//
//type ObjectType string

const (
	NULL_OBJ ObjectType = iota
	ERROR_OBJ
	INTEGER_OBJ
	BOOLEAN_OBJ
	STRING_OBJ
	RETURN_VALUE_OBJ
	FUNCTION_OBJ
	BUILTIN_OBJ
	ARRAY_OBJ
	COMPILED_FUNCTION_OBJ
	CLOSURE_OBJ
)

var objectTypeString = [...]string{
	NULL_OBJ:              "NULL",
	ERROR_OBJ:             "ERROR",
	INTEGER_OBJ:           "INTEGER",
	BOOLEAN_OBJ:           "BOOLEAN",
	STRING_OBJ:            "STRING",
	RETURN_VALUE_OBJ:      "RETURN_VALUE",
	FUNCTION_OBJ:          "FUNCTION",
	BUILTIN_OBJ:           "BUILTIN",
	ARRAY_OBJ:             "ARRAY",
	COMPILED_FUNCTION_OBJ: "COMPILED_FUNCTION_OBJ",
	CLOSURE_OBJ:           "CLOSURE",
}

type ObjectType int

type Object interface {
	Type() ObjectType
	TypeString() string
	Inspect() string
}

type Error struct {
	Message string
}

//TODO: 需要增加stack trace， line number，column number，lexer中增加后这儿也要增加

func (e *Error) Type() ObjectType   { return ERROR_OBJ }
func (e *Error) TypeString() string { return objectTypeString[ERROR_OBJ] }
func (e *Error) Inspect() string    { return "ERROR:" + e.Message }

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType   { return INTEGER_OBJ }
func (i *Integer) TypeString() string { return objectTypeString[INTEGER_OBJ] }
func (i *Integer) Inspect() string    { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType   { return BOOLEAN_OBJ }
func (b *Boolean) TypeString() string { return objectTypeString[BOOLEAN_OBJ] }
func (b *Boolean) Inspect() string    { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType   { return NULL_OBJ }
func (n *Null) TypeString() string { return objectTypeString[NULL_OBJ] }
func (n *Null) Inspect() string    { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType   { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) TypeString() string { return objectTypeString[RETURN_VALUE_OBJ] }
func (rv *ReturnValue) Inspect() string    { return rv.Value.Inspect() }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType   { return FUNCTION_OBJ }
func (f *Function) TypeString() string { return objectTypeString[FUNCTION_OBJ] }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("){\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType   { return STRING_OBJ }
func (s *String) TypeString() string { return objectTypeString[STRING_OBJ] }
func (s *String) Inspect() string    { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType   { return BUILTIN_OBJ }
func (b *Builtin) TypeString() string { return objectTypeString[BUILTIN_OBJ] }
func (b *Builtin) Inspect() string    { return "builtin function" }

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType   { return ARRAY_OBJ }
func (a *Array) TypeString() string { return objectTypeString[ARRAY_OBJ] }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type CompiledFunction struct {
	Instructions  code.Instructions
	NumLocals     int
	NumParameters int
}

func (cf *CompiledFunction) Type() ObjectType   { return COMPILED_FUNCTION_OBJ }
func (cf *CompiledFunction) TypeString() string { return objectTypeString[COMPILED_FUNCTION_OBJ] }
func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() ObjectType   { return CLOSURE_OBJ }
func (c *Closure) TypeString() string { return objectTypeString[CLOSURE_OBJ] }
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}

//TODO: go里的primitive类型也可以实现接口的，可以用来提高性能？
