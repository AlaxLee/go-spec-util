package gospec

import (
	"go/types"
	"testing"
)

func TestGetDynamicTypeAtRuntime(test *testing.T) {
	var a = [...]int{1, 2, 3}
	var b interface{} = a
	if GetDynamicTypeAtRuntime(b).String() != "[3]int" {
		test.Error("test failed")
	}
}

// Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
func TestComparable01(test *testing.T) {
	s := NewSpec(`
var a bool = true
var b bool = true
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if IsBoolean(ta) && IsBoolean(tb) && Comparable(ta) && Comparable(tb) {
		//fmt.Print("Boolean values are comparable. ")
	} else {
		test.Error("test failed")
	}
	var a bool = true
	var b bool = true
	if (a && b) || (!a && !b) {
		if a == b {
			//fmt.Println("Two boolean values are equal if they are either both `true` or both `false`.")
		} else {
			test.Error("test failed")
		}
	}
	/* the output is:
	Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
	*/
}

// Integer values are comparable and ordered, in the usual way.
func TestComparable02(test *testing.T) {
	s := NewSpec(`
var a int = 1
var b int = 1
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if IsInteger(ta) && IsOrdered(ta) && Comparable(ta) &&
		IsInteger(tb) && IsOrdered(tb) && Comparable(tb) {
		var a int = 1
		var b int = 1
		if a == b {
			//fmt.Println("Integer values are comparable and ordered, in the usual way.")
		} else {
			test.Error("test failed")
		}
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Integer values are comparable and ordered, in the usual way.
	*/
}

// Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
func TestComparable03(test *testing.T) {
	s := NewSpec(`
var a float64 = 1.0
var b float64 = 1.0
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if IsFloat(ta) && IsOrdered(ta) && Comparable(ta) &&
		IsFloat(tb) && IsOrdered(tb) && Comparable(tb) {
		var a float64 = 1.0
		var b float64 = 1.0
		if a == b {
			//fmt.Println("Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.")
		} else {
			test.Error("test failed")
		}
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
	*/
}

// Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
func TestComparable04(test *testing.T) {
	s := NewSpec(`
var u complex64 = 1 + 2i
var v complex64 = 1 + 2i
`)
	tu := s.GetType("u")
	tv := s.GetType("v")
	if IsComplex(tu) && Comparable(tu) &&
		IsComplex(tv) && Comparable(tv) {
		//fmt.Print("Complex values are comparable. ")
		var u complex64 = 1 + 2i
		var v complex64 = 1 + 2i
		if real(u) == real(v) && imag(u) == imag(v) {
			if u == v {
				//fmt.Println("Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.")
			} else {
				test.Error("test failed")
			}
		}
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
	*/
}

// String values are comparable and ordered, lexically byte-wise（逐字节地）.
func TestComparable05(test *testing.T) {
	s := NewSpec(`
var a string = "haha"
var b string = "haha"
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if IsString(ta) && Comparable(ta) &&
		IsString(tb) && Comparable(tb) {
		var a string = "haha"
		var b string = "haha"
		if a == b {
			//fmt.Println("String values are comparable and ordered")
		} else {
			test.Error("test failed")
		}
	} else {
		test.Error("test failed")
	}
	/* the output is:
	String values are comparable and ordered
	*/
}

// Pointer values are comparable.
// Two pointer values are equal if they point to the same variable or if both have value `nil`.
// Pointers to distinct `zero-size` variables may or may not be equal.
func TestComparable06(test *testing.T) {
	s := NewSpec(`
var a *int
`)
	ta := s.GetType("a")
	if _, isPointer := ToPointer(ta); isPointer && Comparable(ta) {
		//fmt.Println("Pointer values are comparable.")
	} else {
		test.Error("test failed")
	}
	var a *int
	var b *int
	var c int = 1
	var u *int = &c
	var v *int = &c
	if a == nil && b == nil && a == b && u == v {
		//fmt.Println("Two pointer values are equal if they point to the same variable or if both have value `nil`.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Pointer values are comparable.
	Two pointer values are equal if they point to the same variable or if both have value `nil`.
	*/
}

// Channel values are comparable.
// Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
func TestComparable07(test *testing.T) {
	s := NewSpec(`
var a chan int
`)
	ta := s.GetType("a")
	if _, isChan := ToChan(ta); isChan && Comparable(ta) {
		//fmt.Println("Channel values are comparable.")
	} else {
		test.Error("test failed")
	}
	var a chan int
	var b chan int
	c := make(chan int, 1)
	var u chan int = c
	var v chan int = c
	if a == nil && b == nil && a == b && u == v {
		//fmt.Println("Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Channel values are comparable.
	Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
	*/
}

// Interface values are comparable.
// Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
func TestComparable08(test *testing.T) {
	s := NewSpec(`
	var a interface{}
`)
	ta := s.GetType("a")
	if _, isInterface := ToInterface(ta); isInterface && Comparable(ta) {
		//fmt.Println("Interface values are comparable.")
	} else {
		test.Error("test failed")
	}

	var a interface{}
	var b interface{}
	var c = [...]int{1, 2, 3}
	var d = [...]int{1, 2, 3}
	var u interface{} = c
	var v interface{} = d
	if a == nil && b == nil && a == b &&
		GetDynamicTypeAtRuntime(u) == GetDynamicTypeAtRuntime(v) && c == d {
		//fmt.Println("Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Interface values are comparable.
	Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
	*/
}

// A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
// when values of type `X` are comparable and `X` implements `T`.
// They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
func TestComparable09(test *testing.T) {
	s := NewSpec(`
	type X struct {}
	type T interface {}
	var x X
	var t T
`)
	if s.Comparable("X") && s.Implements("X", "T") {
		//fmt.Println("A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable ")
		//fmt.Println("when values of type `X` are comparable and `X` implements `T`.")
	} else {
		test.Error("test failed")
	}
	type X struct{}
	type T interface{}
	var x X
	var y X
	var t T
	t = y
	if GetDynamicTypeAtRuntime(t) == GetDynamicTypeAtRuntime(x) && y == x {
		//fmt.Println("They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
	when values of type `X` are comparable and `X` implements `T`.
	They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
	*/
}

// Struct values are comparable if all their fields are comparable.
// Two struct values are equal if their corresponding non-blank fields are equal.
func TestComparable10(test *testing.T) {
	s := NewSpec(`
	type U struct {
		a int
		b []int
	}
	type V struct {
		a complex64
		b string
	}
`)

	u, ok := s.GetType("U").Underlying().(*types.Struct)
	if !ok {
		panic("unexpect")
	}
	allUFieldComparable := true
	for i := 0; i < u.NumFields(); i++ {
		if !Comparable(u.Field(i).Type()) {
			allUFieldComparable = false
		}
	}

	v, ok := s.GetType("V").Underlying().(*types.Struct)
	if !ok {
		panic("unexpect")
	}
	allVFieldComparable := true
	for i := 0; i < v.NumFields(); i++ {
		if !Comparable(v.Field(i).Type()) {
			allVFieldComparable = false
		}
	}

	if !allUFieldComparable && !s.Comparable("U") {
		if allVFieldComparable && s.Comparable("V") {
			//fmt.Println("Struct values are comparable if all their fields are comparable.")
		} else {
			test.Error("test failed")
		}
	} else {
		test.Error("test failed")
	}

	type V struct {
		a complex64
		b string
	}

	m := V{1i, "lala"}
	n := V{1i, "lala"}
	if m.a == n.a && m.b == n.b && m == n {
		//fmt.Println("Two struct values are equal if their corresponding non-blank fields are equal.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Struct values are comparable if all their fields are comparable.
	Two struct values are equal if their corresponding non-blank fields are equal.
	*/
}

// Array values are comparable if values of the array element type are comparable.
// Two array values are equal if their corresponding elements are equal.
func TestComparable11(test *testing.T) {
	s := NewSpec(`
	var a [3]interface{}
	var b [3][]int
`)
	if s.Comparable("a") && !s.Comparable("b") {
		//fmt.Println("Array values are comparable if values of the array element type are comparable.")
	} else {
		test.Error("test failed")
	}

	var m [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	var n [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	if m[0] == n[0] && m[1] == n[1] && m[2] == n[2] && m == n {
		//fmt.Println("Two array values are equal if their corresponding elements are equal.")
	} else {
		test.Error("test failed")
	}
	/* the output is:
	Array values are comparable if values of the array element type are comparable.
	Two array values are equal if their corresponding elements are equal.
	*/
}

// Slice, map, and function values are not comparable.
// However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
func TestComparable12(test *testing.T) {
	s := NewSpec(`
	var a []int
	var b map[int]string
	var c func(string)int
`)
	if !s.Comparable("a") && !s.Comparable("b") && !s.Comparable("c") {
		//fmt.Println("Slice, map, and function values are not comparable.")
	} else {
		test.Error("test failed")
	}
	var a []int
	var b map[int]string
	var c func(string) int
	if a == nil && b == nil && c == nil {
		//fmt.Println("However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.")
	} else {
		test.Error("test failed")
	}
	/*  the output is:
	Slice, map, and function values are not comparable.
	However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
	*/
}

// func (s *Spec) Comparable(v string) bool
// Comparable(t types.Type) bool
// or
// Comparable(code, v string) bool
func TestComparable13(test *testing.T) {
	code := `var a bool`
	s := NewSpec(code)
	if !s.Comparable("a") {
		test.Error("test failed")
	}
	if !Comparable(s.GetType("a")) {
		test.Error("test failed")
	}
	if !Comparable(code, "a") {
		test.Error("test failed")
	}
}
