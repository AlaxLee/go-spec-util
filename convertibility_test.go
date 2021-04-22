package gospec

import (
	"errors"
	"fmt"
	"testing"
)

// func (s *Spec) Conversion(v, t string) bool
func TestConversion01(t *testing.T) {
	type Info struct {
		x            string
		T            string
		couldConvert bool
	}
	newSpec := func(code string) (s *Spec, e error) {
		defer func() {
			if r := recover(); r != nil {
				e = errors.New(fmt.Sprintf("%s", r))
			}
		}()
		s = NewSpec(code)
		return
	}

	/* Converting a constant yields a typed constant as result.
	uint(iota)               // iota value of type uint
	float32(2.718281828)     // 2.718281828 of type float32
	complex128(1)            // 1.0 + 0.0i of type complex128
	float32(0.49999999)      // 0.5 of type float32
	float64(-1e-1000)        // 0.0 of type float64
	string('x')              // "x" of type string
	string(0x266c)           // "♬" of type string
	MyString("foo" + "bar")  // "foobar" of type MyString
	string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
	(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
	int(1.2)                 // illegal: 1.2 cannot be represented as an int
	string(65.0)             // illegal: 65.0 is not an integer constant
	*/
	infos := []Info{
		{`iota`, `uint`, true},
		{`2.718281828`, `float32`, true},
		{`1`, `complex128`, true},
		{`0.49999999`, `float32`, true},
		{`-1e-1000`, `float64`, true},
		{`'x'`, `string`, true},
		{`0x266c`, `string`, true},
		{`"foo" + "bar"`, `MyString; type MyString string`, true},
		{`[]byte{'a'}`, `string`, false},
		{`nil`, `(*int)`, false},
		{`1.2`, `int`, false},
		{`65.0`, `string`, false},
	}

	for i, v := range infos {
		code := fmt.Sprintf("type T %s; const x = %s", v.T, v.x)
		s, err := newSpec(code)
		if i >= 0 && i <= 7 {
			if err == nil && s.Conversion("x", "T") && v.couldConvert {
			} else {
				t.Errorf("test failed")
			}
		} else if i >= 8 && i <= 9 {
			if err == nil {
				t.Errorf("test failed")
			}
		} else {
			if err == nil && !s.Conversion("x", "T") && !v.couldConvert {
			} else {
				t.Errorf("test failed")
			}
		}
	}
}

func TestConversion02(t *testing.T) {
	// A non-constant value x can be converted to type T in any of these cases:
	//fmt.Println("A non-constant value x can be converted to type T in any of these cases:")
	var s *Spec

	// 1. x is assignable to T.
	s = NewSpec(`var x = func(){}; type T func()`)
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if s.Assignment("x", "T") &&
			s.Conversion("x", "T") {
			//fmt.Println("1. x is assignable to T.")
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}

	// 2. ignoring struct tags (see below), x's type and T have identical underlying types.
	s = NewSpec(`
type T struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var x struct {
	Name    string ` + "`" + `json:"name"` + "`" + `
	Address *struct {
		Street string ` + "`" + `json:"street"` + "`" + `
		City   string ` + "`" + `json:"city"` + "`" + `
	} ` + "`" + `json:"address"` + "`" + `
}
`)
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if IdenticalIgnoreTags(s.MustGetValidType("x").Underlying(), s.MustGetValidType("T").Underlying()) &&
			s.Conversion("x", "T") {
			//fmt.Println("2. ignoring struct tags (see below), x's type and T have identical underlying types.")
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}

	// 3. ignoring struct tags (see below), x's type and T are pointer types that are not defined types,
	// and their pointer base types have identical underlying types.
	s = NewSpec(`
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string ` + "`" + `json:"name"` + "`" + `
	Address *struct {
		Street string ` + "`" + `json:"street"` + "`" + `
		City   string ` + "`" + `json:"city"` + "`" + `
	} ` + "`" + `json:"address"` + "`" + `
}

type T = *Person
var person = (*Person)(data)
`)
	xt := s.MustGetValidType("data")
	tt := s.MustGetValidType("T")
	xp, xIsPointer := ToPointer(xt)
	tp, tIsPointer := ToPointer(tt)
	if !IsConstObject(s.MustGetValidTypeObject("data")) {
		if xIsPointer && tIsPointer &&
			!IsDefinedType(xt) && !IsDefinedType(tt) &&
			IdenticalIgnoreTags(xp.Elem().Underlying(), tp.Elem().Underlying()) &&
			s.Conversion("data", "T") {
			//fmt.Println("3. ignoring struct tags (see below), x's type and T are pointer types that are not defined types, and their pointer base types have identical underlying types.")
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}

	// 4. x's type and T are both integer or floating point types.
	s = NewSpec(`
type T uint
var x = 1
type U float32
var y = 1.0
`)
	var bothIntegerOK, bothFloatOK bool
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if IsInteger(s.MustGetValidType("x")) && IsInteger(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			bothIntegerOK = true
		}
	}
	if !IsConstObject(s.MustGetValidTypeObject("y")) {
		if IsFloat(s.MustGetValidType("y")) && IsFloat(s.MustGetValidType("U")) &&
			s.Conversion("y", "U") {
			bothFloatOK = true
		}
	}
	if bothIntegerOK && bothFloatOK {
		//fmt.Println("4. x's type and T are both integer or floating point types.")
	} else {
		t.Errorf("test failed")
	}

	// 5. x's type and T are both complex types.
	s = NewSpec(`
type T complex64
var x = 1+2i
`)
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if IsComplex(s.MustGetValidType("x")) && IsComplex(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			//fmt.Println("5. x's type and T are both complex types.")
		} else {
			t.Errorf("test failed")
		}
	} else {
		t.Errorf("test failed")
	}

	// 6. x is an integer or a slice of bytes or runes and T is a string type.
	s = NewSpec(`
type T string
var x = 1
var y = []byte{}
var z = []rune{}
`)
	var fromIntegerOK, fromByteSliceOK, fromRuneSliceOK bool
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if IsInteger(s.MustGetValidType("x")) && IsString(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			fromIntegerOK = true
		}
	}

	if !IsConstObject(s.MustGetValidTypeObject("y")) {
		if ys, yIsSlice := ToSlice(s.MustGetValidType("y")); yIsSlice && IsByte(ys.Elem()) && IsString(s.MustGetValidType("T")) &&
			s.Conversion("y", "T") {
			fromByteSliceOK = true
		}
	}

	if !IsConstObject(s.MustGetValidTypeObject("z")) {
		if zs, zIsSlice := ToSlice(s.MustGetValidType("z")); zIsSlice && IsRune(zs.Elem()) && IsString(s.MustGetValidType("T")) &&
			s.Conversion("z", "T") {
			fromRuneSliceOK = true
		}
	}
	if fromIntegerOK && fromByteSliceOK && fromRuneSliceOK {
		//fmt.Println("6. x is an integer or a slice of bytes or runes and T is a string type.")
	} else {
		t.Errorf("test failed")
	}

	// 7. x is a string and T is a slice of bytes or runes.
	s = NewSpec(`
var x string = "lala"
type T []byte
type U []rune
`)
	var toByteSliceOK, toRuneSliceOK bool
	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if ts, tIsSlice := ToSlice(s.MustGetValidType("T")); tIsSlice && IsByte(ts.Elem()) &&
			s.Conversion("x", "T") {
			toByteSliceOK = true
		}
	}

	if !IsConstObject(s.MustGetValidTypeObject("x")) {
		if ts, tIsSlice := ToSlice(s.MustGetValidType("U")); tIsSlice && IsRune(ts.Elem()) &&
			s.Conversion("x", "U") {
			toRuneSliceOK = true
		}
	}

	if toByteSliceOK && toRuneSliceOK {
		//fmt.Println("7. x is a string and T is a slice of bytes or runes.")
	} else {
		t.Errorf("test failed")
	}

	/* the output is:
	A non-constant value x can be converted to type T in any of these cases:
	1. x is assignable to T.
	2. ignoring struct tags (see below), x's type and T have identical underlying types.
	3. ignoring struct tags (see below), x's type and T are pointer types that are not defined types, and their pointer base types have identical underlying types.
	4. x's type and T are both integer or floating point types.
	5. x's type and T are both complex types.
	6. x is an integer or a slice of bytes or runes and T is a string type.
	7. x is a string and T is a slice of bytes or runes.
	*/
}

// func Conversion(code, v, t string) bool
func TestConversion03(t *testing.T) {
	code := "type T uint; const x = iota"
	if Conversion(code, "x", "T") {
	} else {
		t.Errorf("test failed")
	}
}
