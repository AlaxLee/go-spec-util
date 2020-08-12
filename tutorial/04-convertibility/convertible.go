package main

import (
	"errors"
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
	convertibleExample01()
	convertibleExample02()
}

func convertibleExample01() {
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
	type Info struct {
		x string
		T string
	}

	infos := []Info{
		{`iota`, `uint`},
		{`2.718281828`, `float32`},
		{`1`, `complex128`},
		{`0.49999999`, `float32`},
		{`-1e-1000`, `float64`},
		{`'x'`, `string`},
		{`0x266c`, `string`},
		{`"foo" + "bar"`, `MyString; type MyString string`},
		{`[]byte{'a'}`, `string`},
		{`nil`, `(*int)`},
		{`1.2`, `int`},
		{`65.0`, `string`},
	}

	newSpec := func(code string) (s *gospec.Spec, e error) {
		defer func() {
			if r := recover(); r != nil {
				e = errors.New(fmt.Sprintf("%s", r))
			}
		}()
		s = gospec.NewSpec(code)
		return
	}

	for _, v := range infos {
		code := fmt.Sprintf("type T %s; const x = %s", v.T, v.x)
		s, err := newSpec(code)
		if err != nil {
			//fmt.Printf("%12s could not convert to type %s, because %s\n", v.x, v.T, err)
			fmt.Printf("%20s could not convert to type %s\n", v.x, v.T)
		} else if s.Conversion("x", "T") {
			fmt.Printf("%20s could convert to type %s\n", v.x, v.T)
		} else {
			fmt.Printf("%20s could not convert to type %s\n", v.x, v.T)
		}
	}
	/* the output is:
	            iota could convert to type uint
	     2.718281828 could convert to type float32
	               1 could convert to type complex128
	      0.49999999 could convert to type float32
	        -1e-1000 could convert to type float64
	             'x' could convert to type string
	          0x266c could convert to type string
	   "foo" + "bar" could convert to type MyString; type MyString string
	     []byte{'a'} could not convert to type string
	             nil could not convert to type (*int)
	             1.2 could not convert to type int
	            65.0 could not convert to type string
	*/
}

func convertibleExample02() {
	// A non-constant value x can be converted to type T in any of these cases:
	fmt.Println("A non-constant value x can be converted to type T in any of these cases:")
	var s *gospec.Spec

	// 1. x is assignable to T.
	s = gospec.NewSpec(`var x = func(){}; type T func()`)
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if s.Assignment("x", "T") &&
			s.Conversion("x", "T") {
			fmt.Println("1. x is assignable to T.")
		}
	}

	// 2. ignoring struct tags (see below), x's type and T have identical underlying types.
	s = gospec.NewSpec(`
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
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if gospec.IdenticalIgnoreTags(s.MustGetValidType("x").Underlying(), s.MustGetValidType("T").Underlying()) &&
			s.Conversion("x", "T") {
			fmt.Println("2. ignoring struct tags (see below), x's type and T have identical underlying types.")
		}
	}

	// 3. ignoring struct tags (see below), x's type and T are pointer types that are not defined types,
	// and their pointer base types have identical underlying types.
	s = gospec.NewSpec(`
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
	xp, xIsPointer := gospec.ToPointer(xt)
	tp, tIsPointer := gospec.ToPointer(tt)
	if !gospec.IsConstObject(s.MustGetValidTypeObject("data")) {
		if xIsPointer && tIsPointer &&
			!gospec.IsDefinedType(xt) && !gospec.IsDefinedType(tt) &&
			gospec.IdenticalIgnoreTags(xp.Elem().Underlying(), tp.Elem().Underlying()) &&
			s.Conversion("data", "T") {
			fmt.Println("3. ignoring struct tags (see below), x's type and T are pointer types that are not defined types, and their pointer base types have identical underlying types.")
		}
	}

	// 4. x's type and T are both integer or floating point types.
	s = gospec.NewSpec(`
type T uint
var x = 1
type U float32
var y = 1.0
`)
	var bothIntegerOK, bothFloatOK bool
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if gospec.IsInteger(s.MustGetValidType("x")) && gospec.IsInteger(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			bothIntegerOK = true
		}
	}
	if !gospec.IsConstObject(s.MustGetValidTypeObject("y")) {
		if gospec.IsFloat(s.MustGetValidType("y")) && gospec.IsFloat(s.MustGetValidType("U")) &&
			s.Conversion("y", "U") {
			bothFloatOK = true
		}
	}
	if bothIntegerOK && bothFloatOK {
		fmt.Println("4. x's type and T are both integer or floating point types.")
	}

	// 5. x's type and T are both complex types.
	s = gospec.NewSpec(`
type T complex64
var x = 1+2i
`)
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if gospec.IsComplex(s.MustGetValidType("x")) && gospec.IsComplex(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			fmt.Println("5. x's type and T are both complex types.")
		}
	}

	// 6. x is an integer or a slice of bytes or runes and T is a string type.
	s = gospec.NewSpec(`
type T string
var x = 1
var y = []byte{}
var z = []rune{}
`)
	var fromIntegerOK, fromByteSliceOK, fromRuneSliceOK bool
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if gospec.IsInteger(s.MustGetValidType("x")) && gospec.IsString(s.MustGetValidType("T")) &&
			s.Conversion("x", "T") {
			fromIntegerOK = true
		}
	}

	if !gospec.IsConstObject(s.MustGetValidTypeObject("y")) {
		if ys, yIsSlice := gospec.ToSlice(s.MustGetValidType("y")); yIsSlice && gospec.IsByte(ys.Elem()) && gospec.IsString(s.MustGetValidType("T")) &&
			s.Conversion("y", "T") {
			fromByteSliceOK = true
		}
	}

	if !gospec.IsConstObject(s.MustGetValidTypeObject("z")) {
		if zs, zIsSlice := gospec.ToSlice(s.MustGetValidType("z")); zIsSlice && gospec.IsRune(zs.Elem()) && gospec.IsString(s.MustGetValidType("T")) &&
			s.Conversion("z", "T") {
			fromRuneSliceOK = true
		}
	}
	if fromIntegerOK && fromByteSliceOK && fromRuneSliceOK {
		fmt.Println("6. x is an integer or a slice of bytes or runes and T is a string type.")
	}

	// 7. x is a string and T is a slice of bytes or runes.
	s = gospec.NewSpec(`
var x string = "lala"
type T []byte
type U []rune
`)
	var toByteSliceOK, toRuneSliceOK bool
	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if ts, tIsSlice := gospec.ToSlice(s.MustGetValidType("T")); tIsSlice && gospec.IsByte(ts.Elem()) &&
			s.Conversion("x", "T") {
			toByteSliceOK = true
		}
	}

	if !gospec.IsConstObject(s.MustGetValidTypeObject("x")) {
		if ts, tIsSlice := gospec.ToSlice(s.MustGetValidType("U")); tIsSlice && gospec.IsRune(ts.Elem()) &&
			s.Conversion("x", "U") {
			toRuneSliceOK = true
		}
	}

	if toByteSliceOK && toRuneSliceOK {
		fmt.Println("7. x is a string and T is a slice of bytes or runes.")
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
