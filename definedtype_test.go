package gospec

import "testing"

// IsDefinedType(code,v string) bool
// A's type is example.A, and analysed to *types.Named
// A is a defined type
func TestIsDefinedType01(t *testing.T) {
	if !IsDefinedType(`type A int`, "A") {
		t.Error(`test rule failed`)
	}
}

// A's type is int, and analysed to *types.Basic
// A is a defined type
func TestIsDefinedType02(t *testing.T) {
	if !IsDefinedType(`type A = int`, "A") {
		t.Error(`test rule failed`)
	}
}

// A's type is example.A, and analysed to *types.Named
// A is a defined type
func TestIsDefinedType03(t *testing.T) {
	if !IsDefinedType(`type A func()`, "A") {
		t.Error(`test rule failed`)
	}
}

// A's type is func(), and analysed to *types.Signature
// A is not a defined type
func TestIsDefinedType04(t *testing.T) {
	if IsDefinedType(`type A = func()`, "A") {
		t.Error(`test rule failed`)
	}
}

// test func (s *Spec) IsDefinedType(v string) bool
// A's type is example.A, and analysed to *types.Named
// A is a defined type
func TestIsDefinedType05(t *testing.T) {
	s := NewSpec(`type A int`)
	if !s.IsDefinedType("A") {
		t.Error(`test rule failed`)
	}
}

// test IsDefinedType(t types.Type) bool
// A's type is example.A, and analysed to *types.Named
// A is a defined type
func TestIsDefinedType06(t *testing.T) {
	s := NewSpec(`type A int`)
	typ := s.GetType("A")
	if !IsDefinedType(typ) {
		t.Error(`test rule failed`)
	}
}
