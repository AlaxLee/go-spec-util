package gospec

import (
	easyregexp "github.com/AlaxLee/easyregexp"
	"go/types"
)

func addPackageHeadToCode(codePtr *string) {
	if !easyregexp.Match(`^\s*package\s+`, *codePtr) {
		*codePtr = easyregexp.ReplaceAll(`^\s*`, *codePtr, "package example\n")
	}
}

func mustGetPackageNameFromCode(code string) string {
	catches := easyregexp.Catch(`package\s+(\w+)`, code)
	if len(catches) == 0 {
		panic("get packageName failed")
	}
	return catches[0]
}

//check types.Type if be a basic, pointer, function, slice, map, channel, or interface
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

func ToBasic(t types.Type) (*types.Basic, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tb, ok := tu.(*types.Basic)
	return tb, ok
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

func IsByte(t types.Type) bool {
	if t == nil {
		return false
	}
	tu := t.Underlying()
	tb, ok := tu.(*types.Basic)
	if !ok {
		return false
	}
	if tb.Kind() != types.Byte {
		return false
	}
	return true
}

func IsRune(t types.Type) bool {
	if t == nil {
		return false
	}
	tu := t.Underlying()
	tb, ok := tu.(*types.Basic)
	if !ok {
		return false
	}
	if tb.Kind() != types.Rune {
		return false
	}
	return true
}
