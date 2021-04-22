package gospec

import "testing"

// func (s *Spec) Implements(v, t string) bool
func TestImplements01(t *testing.T) {
	s := NewSpec(`
type T interface {
	m()
}
type V struct{}
func (v V) m() {}
var x V
`)
	if _, ok := ToInterface(s.GetType("T")); ok {
		if s.Implements("x", "T") {
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}
}

// Implements(v, t types.Type) bool
// or
// Implements(v, t types.Object) bool
// or
// Implements((code, v, t string) bool
func TestImplements02(t *testing.T) {
	code := `
type T interface {
	m()
}
type V struct{}
func (v V) m() {}
var x V
`
	s := NewSpec(code)
	if _, ok := ToInterface(s.GetType("T")); ok {
		if Implements(s.GetType("x"), s.GetType("T")) {
		} else {
			t.Errorf("test failed")
		}
		if Implements(s.GetTypeObject("x"), s.GetTypeObject("T")) {
		} else {
			t.Errorf("test failed")
		}
		if Implements(code, "x", "T") {
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}
}
