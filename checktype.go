package gospec

import "go/types"

// basic types
func IsBoolean(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsBoolean != 0
}

func IsInteger(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsInteger != 0
}

func IsUnsigned(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsUnsigned != 0
}

func IsFloat(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsFloat != 0
}

func IsComplex(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsComplex != 0
}

func IsString(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsString != 0
}

func IsUntyped(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsUntyped != 0
}

func IsTyped(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return !ok || t.Info()&types.IsUntyped == 0
}

//	IsNumeric   = IsInteger | IsFloat | IsComplex
func IsNumeric(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsNumeric != 0
}

// 	IsOrdered   = IsInteger | IsFloat | IsString
func IsOrdered(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsOrdered != 0
}

// 	IsConstType = IsBoolean | IsNumeric | IsString
func IsConstType(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsConstType != 0
}

// other type
func IsByte(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Kind() == types.Byte
}

func IsRune(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Kind() == types.Rune
}

//check types.Type if be a basic, pointer, function, slice, map, channel, or interface

func ToBasic(t types.Type) (*types.Basic, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tb, ok := tu.(*types.Basic)
	return tb, ok
}

func ToPointer(t types.Type) (*types.Pointer, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tp, ok := tu.(*types.Pointer)
	return tp, ok
}

func ToFunction(t types.Type) (*types.Signature, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tf, ok := tu.(*types.Signature)
	return tf, ok
}

func ToSlice(t types.Type) (*types.Slice, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	ts, ok := tu.(*types.Slice)
	return ts, ok
}

func ToMap(t types.Type) (*types.Map, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tm, ok := tu.(*types.Map)
	return tm, ok
}

func ToInterface(t types.Type) (*types.Interface, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	ti, ok := tu.(*types.Interface)
	return ti, ok
}

func ToChan(t types.Type) (*types.Chan, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tc, ok := tu.(*types.Chan)
	return tc, ok
}

//check types.Object if be a const
func ToConstObject(o types.Object) (*types.Const, bool) {
	if o == nil {
		return nil, false
	}
	constObj, ok := o.(*types.Const)
	return constObj, ok
}

func IsConstObject(o types.Object) bool {
	_, isConstObject := ToConstObject(o)
	return isConstObject
}
