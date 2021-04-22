package gospec

import (
	"testing"
)

func TestIsBoolean(t *testing.T) {
	s := NewSpec(`var a bool`)
	if !IsBoolean(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsInteger(t *testing.T) {
	s := NewSpec(`var a int`)
	if !IsInteger(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsUnsigned(t *testing.T) {
	s := NewSpec(`var a uint`)
	if !IsUnsigned(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsFloat(t *testing.T) {
	s := NewSpec(`var a float32`)
	if !IsFloat(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsComplex(t *testing.T) {
	s := NewSpec(`var a complex64`)
	if !IsComplex(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsString(t *testing.T) {
	s := NewSpec(`var a string`)
	if !IsString(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsUntyped(t *testing.T) {
	s := NewSpec(`const a = 1`) // a is UnTyped
	if !IsUntyped(s.GetType("a")) {
		t.Error("*******test failed")
	}
}

func TestIsTyped(t *testing.T) {
	s := NewSpec(`const a = 1`) // a is UnTyped
	if IsTyped(s.GetType("a")) {
		t.Error("*******test failed")
	}
}

//	IsNumeric   = IsInteger | IsFloat | IsComplex
func TestIsNumeric(t *testing.T) {
	s := NewSpec(`var a int; var b float32; var c complex64`)
	if IsNumeric(s.GetType("a")) && IsNumeric(s.GetType("b")) && IsNumeric(s.GetType("c")) {
	} else {
		t.Error("test failed")
	}
}

// 	IsOrdered   = IsInteger | IsFloat | IsString
func TestIsOrdered(t *testing.T) {
	s := NewSpec(`var a int; var b float32; var c string`)
	if IsOrdered(s.GetType("a")) && IsOrdered(s.GetType("b")) && IsOrdered(s.GetType("c")) {
	} else {
		t.Error("test failed")
	}
}

// 	IsConstType = IsBoolean | IsNumeric | IsString
func TestIsConstType(t *testing.T) {
	s := NewSpec(`var a bool; var b int; var c float32; var d complex64; var e string`)
	if IsConstType(s.GetType("a")) &&
		IsConstType(s.GetType("b")) &&
		IsConstType(s.GetType("c")) &&
		IsConstType(s.GetType("d")) &&
		IsConstType(s.GetType("e")) {
	} else {
		t.Error("test failed")
	}
}

func TestIsByte(t *testing.T) {
	s := NewSpec(`var a byte`)
	if !IsByte(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestIsRune(t *testing.T) {
	s := NewSpec(`var a rune`)
	if !IsRune(s.GetType("a")) {
		t.Error("test failed")
	}
}

func TestToBasic(t *testing.T) {
	s := NewSpec(`var a int`)
	_, ok := ToBasic(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToPointer(t *testing.T) {
	s := NewSpec(`var a *int`)
	_, ok := ToPointer(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToFunction(t *testing.T) {
	s := NewSpec(`var a func()`)
	_, ok := ToFunction(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToSlice(t *testing.T) {
	s := NewSpec(`var a []int`)
	_, ok := ToSlice(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToMap(t *testing.T) {
	s := NewSpec(`var a map[int]int`)
	_, ok := ToMap(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToInterface(t *testing.T) {
	s := NewSpec(`var a interface{ lala() }`)
	_, ok := ToInterface(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToChan(t *testing.T) {
	s := NewSpec(`var a chan int`)
	_, ok := ToChan(s.GetType("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestToConstObject(t *testing.T) {
	s := NewSpec(`const a string = "lala"`)
	_, ok := ToConstObject(s.GetTypeObject("a"))
	if !ok {
		t.Error("test failed")
	}
}

func TestIsConstObject(t *testing.T) {
	s := NewSpec(`const a string = "lala"`)
	if !IsConstObject(s.GetTypeObject("a")) {
		t.Error("test failed")
	}
}
