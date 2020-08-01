package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
	"go/types"
)

func main() {
	comparableExample01()
	comparableExample02()
	comparableExample03()
	comparableExample04()
	comparableExample05()
	comparableExample06()
	comparableExample07()
	comparableExample08()
	comparableExample09()
	comparableExample10()
	comparableExample11()
	comparableExample12()
}

func comparableExample01() {
	// Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
	s := gospec.NewSpec(`
var a bool = true
var b bool = true
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsBoolean(ta) && gospec.IsBoolean(tb) && gospec.Comparable(ta) && gospec.Comparable(tb) {
		fmt.Print("Boolean values are comparable. ")
	}
	var a bool = true
	var b bool = true
	if (a && b) || (!a && !b) {
		if a == b {
			fmt.Println("Two boolean values are equal if they are either both `true` or both `false`.")
		}
	}
	/* the output is:
	Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
	*/
}

func comparableExample02() {
	// Integer values are comparable and ordered, in the usual way.
	s := gospec.NewSpec(`
var a int = 1
var b int = 1
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsInteger(ta) && gospec.IsOrdered(ta) && gospec.Comparable(ta) &&
		gospec.IsInteger(tb) && gospec.IsOrdered(tb) && gospec.Comparable(tb) {
		var a int = 1
		var b int = 1
		if a == b {
			fmt.Println("Integer values are comparable and ordered, in the usual way.")
		}
	}
	/* the output is:
	Integer values are comparable and ordered, in the usual way.
	*/
}

func comparableExample03() {
	// Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
	s := gospec.NewSpec(`
var a float64 = 1.0
var b float64 = 1.0
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsFloat(ta) && gospec.IsOrdered(ta) && gospec.Comparable(ta) &&
		gospec.IsFloat(tb) && gospec.IsOrdered(tb) && gospec.Comparable(tb) {
		var a float64 = 1.0
		var b float64 = 1.0
		if a == b {
			fmt.Println("Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.")
		}
	}
	/* the output is:
	Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
	*/
}

func comparableExample04() {
	// Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
	s := gospec.NewSpec(`
var u complex64 = 1 + 2i
var v complex64 = 1 + 2i
`)
	tu := s.GetType("u")
	tv := s.GetType("v")
	if gospec.IsComplex(tu) && gospec.Comparable(tu) &&
		gospec.IsComplex(tv) && gospec.Comparable(tv) {
		fmt.Print("Complex values are comparable. ")
		var u complex64 = 1 + 2i
		var v complex64 = 1 + 2i
		if real(u) == real(v) && imag(u) == imag(v) {
			if u == v {
				fmt.Println("Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.")
			}
		}
	}
	/* the output is:
	Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
	*/
}

func comparableExample05() {
	// String values are comparable and ordered, lexically byte-wise（逐字节地）.
	s := gospec.NewSpec(`
var a string = "haha"
var b string = "haha"
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsString(ta) && gospec.Comparable(ta) &&
		gospec.IsString(tb) && gospec.Comparable(tb) {
		var a string = "haha"
		var b string = "haha"
		if a == b {
			fmt.Println("String values are comparable and ordered")
		}
	}
	/* the output is:
	String values are comparable and ordered
	*/
}

func comparableExample06() {
	// Pointer values are comparable.
	// Two pointer values are equal if they point to the same variable or if both have value `nil`.
	// Pointers to distinct `zero-size` variables may or may not be equal.
	s := gospec.NewSpec(`
var a *int
`)
	ta := s.GetType("a")
	if _, isPointer := gospec.IsPointer(ta); isPointer && gospec.Comparable(ta) {
		fmt.Println("Pointer values are comparable.")
	}
	var a *int
	var b *int
	var c int = 1
	var u *int = &c
	var v *int = &c
	if a == nil && b == nil && a == b && u == v {
		fmt.Println("Two pointer values are equal if they point to the same variable or if both have value `nil`.")
	}
	/* the output is:
	Pointer values are comparable.
	Two pointer values are equal if they point to the same variable or if both have value `nil`.
	*/
}

func comparableExample07() {
	// Channel values are comparable.
	// Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
	s := gospec.NewSpec(`
var a chan int
`)
	ta := s.GetType("a")
	if _, isChan := gospec.IsChan(ta); isChan && gospec.Comparable(ta) {
		fmt.Println("Channel values are comparable.")
	}
	var a chan int
	var b chan int
	c := make(chan int, 1)
	var u chan int = c
	var v chan int = c
	if a == nil && b == nil && a == b && u == v {
		fmt.Println("Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.")
	}
	/* the output is:
	Channel values are comparable.
	Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
	*/
}

func comparableExample08() {
	// Interface values are comparable.
	// Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
	s := gospec.NewSpec(`
	var a interface{}
`)
	ta := s.GetType("a")
	if _, isInterface := gospec.IsInterface(ta); isInterface && gospec.Comparable(ta) {
		fmt.Println("Interface values are comparable.")
	}

	var a interface{}
	var b interface{}
	var c = [...]int{1, 2, 3}
	var d = [...]int{1, 2, 3}
	var u interface{} = c
	var v interface{} = d
	if a == nil && b == nil && a == b &&
		gospec.GetDynamicTypeAtRuntime(u) == gospec.GetDynamicTypeAtRuntime(v) && c == d {
		fmt.Println("Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.")
	}
	/* the output is:
	Interface values are comparable.
	Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
	*/
}

func comparableExample09() {
	// A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
	// when values of type `X` are comparable and `X` implements `T`.
	// They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
	s := gospec.NewSpec(`
	type X struct {}
	type T interface {}
	var x X
	var t T
`)
	if s.Comparable("X") && s.Implements("X", "T") {
		fmt.Println("A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable ")
		fmt.Println("when values of type `X` are comparable and `X` implements `T`.")
	}
	type X struct{}
	type T interface{}
	var x X
	var y X
	var t T
	t = y
	if gospec.GetDynamicTypeAtRuntime(t) == gospec.GetDynamicTypeAtRuntime(x) && y == x {
		fmt.Println("They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.")
	}
	/* the output is:
	A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
	when values of type `X` are comparable and `X` implements `T`.
	They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
	*/
}

func comparableExample10() {
	// Struct values are comparable if all their fields are comparable.
	// Two struct values are equal if their corresponding non-blank fields are equal.
	s := gospec.NewSpec(`
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
		if !gospec.Comparable(u.Field(i).Type()) {
			allUFieldComparable = false
		}
	}

	v, ok := s.GetType("V").Underlying().(*types.Struct)
	if !ok {
		panic("unexpect")
	}
	allVFieldComparable := true
	for i := 0; i < v.NumFields(); i++ {
		if !gospec.Comparable(v.Field(i).Type()) {
			allVFieldComparable = false
		}
	}

	if !allUFieldComparable && !s.Comparable("U") {
		if allVFieldComparable && s.Comparable("V") {
			fmt.Println("Struct values are comparable if all their fields are comparable.")
		}
	}

	type V struct {
		a complex64
		b string
	}

	m := V{1i, "lala"}
	n := V{1i, "lala"}
	if m.a == n.a && m.b == n.b && m == n {
		fmt.Println("Two struct values are equal if their corresponding non-blank fields are equal.")
	}
	/* the output is:
	Struct values are comparable if all their fields are comparable.
	Two struct values are equal if their corresponding non-blank fields are equal.
	*/
}

func comparableExample11() {
	// Array values are comparable if values of the array element type are comparable.
	// Two array values are equal if their corresponding elements are equal.
	s := gospec.NewSpec(`
	var a [3]interface{}
	var b [3][]int
`)
	if s.Comparable("a") && !s.Comparable("b") {
		fmt.Println("Array values are comparable if values of the array element type are comparable.")
	}

	var m [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	var n [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	if m[0] == n[0] && m[1] == n[1] && m[2] == n[2] && m == n {
		fmt.Println("Two array values are equal if their corresponding elements are equal.")
	}
	/* the output is:
	Array values are comparable if values of the array element type are comparable.
	Two array values are equal if their corresponding elements are equal.
	*/
}

func comparableExample12() {
	// Slice, map, and function values are not comparable.
	// However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
	s := gospec.NewSpec(`
	var a []int
	var b map[int]string
	var c func(string)int
`)
	if !s.Comparable("a") && !s.Comparable("b") && !s.Comparable("c") {
		fmt.Println("Slice, map, and function values are not comparable.")
	}
	var a []int
	var b map[int]string
	var c func(string) int
	if a == nil && b == nil && c == nil {
		fmt.Println("However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.")
	}
	/*  the output is:
	Slice, map, and function values are not comparable.
	However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
	*/
}
