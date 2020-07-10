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
func IsPointer(t types.Type) (*types.Pointer, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tp, ok := tu.(*types.Pointer)
	return tp, ok
}

func IsFunction(t types.Type) (*types.Signature, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tf, ok := tu.(*types.Signature)
	return tf, ok
}

func IsSlice(t types.Type) (*types.Slice, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	ts, ok := tu.(*types.Slice)
	return ts, ok
}

func IsMap(t types.Type) (*types.Map, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tm, ok := tu.(*types.Map)
	return tm, ok
}

func IsInterface(t types.Type) (*types.Interface, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	ti, ok := tu.(*types.Interface)
	return ti, ok
}

func IsChan(t types.Type) (*types.Chan, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tc, ok := tu.(*types.Chan)
	return tc, ok
}

func IsBasic(t types.Type) (*types.Basic, bool) {
	if t == nil {
		return nil, false
	}
	tu := t.Underlying()
	tb, ok := tu.(*types.Basic)
	return tb, ok
}
