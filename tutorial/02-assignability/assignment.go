package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
	"go/types"
)

func main() {
	assignExample01()
	assignExample02()
	assignExample03()
	assignExample04()
	assignExample05()
	assignExample06()
}

func assignExample01() {
	//1. x's type is identical to T.
	// x 的类型 与 T 相同
	s := gospec.NewSpec(`
type T = int
var x = 1
`)
	if s.Identical("x", "T") {
		fmt.Println("x's type is identical to T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x's type is identical to T
	x could assignable to T
	*/
}

func assignExample02() {
	//2. x's type V and T have identical underlying types
	// and at least one of V or T is not a defined type.
	// x 的类型 V 和 T 有相同的 underlying type 并且 V 或 T 至少有一个是未（显示）定义类型
	s := gospec.NewSpec(`
type V = func()
type T func()
var x V
`)
	V := s.GetType("x")
	T := s.GetType("T")
	if gospec.Identical(V.Underlying(), T.Underlying()) {
		fmt.Printf("x's type V and T have identical underlying types: %s\n", V.Underlying())
	}
	if !gospec.IsDefinedType(V) {
		fmt.Println("V is not a defined type")
	}
	if gospec.IsDefinedType(T) {
		fmt.Println("T is a defined type")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x's type V and T have identical underlying types: func()
	V is not a defined type
	T is a defined type
	x could assignable to T
	*/
}

func assignExample03() {
	//3. T is an interface type and x implements T.
	// T 是一个接口，x 实现了 T
	s := gospec.NewSpec(`
type T interface {
	m()
}
type V struct{}
func (v V) m() {}
var x V
`)
	if _, ok := gospec.IsInterface(s.GetType("T")); ok {
		fmt.Println("T is an interface type")
	}
	if s.Implements("x", "T") {
		fmt.Println("x implements T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	T is an interface type
	x implements T
	x could assignable to T
	*/
}

func assignExample04() {
	//4. x is a bidirectional channel value, T is a channel type,
	// x's type V and T have identical element types,
	// and at least one of V or T is not a defined type.
	// x 是一个双向管道的值，T 是一个管道类型
	// x 的类型 V 和 T 有相同的元素类型，并且 V 或 T 至少有一个是未（显示）定义类型
	s := gospec.NewSpec(`
type T chan int
var x = make(chan int)
`)
	V := s.GetType("x")
	T := s.GetType("T")
	vc, ok := gospec.IsChan(V)
	if ok && vc.Dir() == types.SendRecv {
		fmt.Println("x is a bidirectional channel value")
	}
	tc, ok := gospec.IsChan(T)
	if ok {
		fmt.Println("T is a channel type")
	}
	if gospec.Identical(vc.Elem(), tc.Elem()) {
		fmt.Println("x's type V and T have identical element types")
	}
	if !gospec.IsDefinedType(V) {
		fmt.Println("V is not defined type")
	}
	if gospec.IsDefinedType(T) {
		fmt.Println("T is defined type")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x is a bidirectional channel value
	T is a channel type
	x's type V and T have identical element types
	V is not defined type
	T is defined type
	x could assignable to T
	*/
}

func assignExample05() {

	//5. x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
	// x 是 nil，T 是一个 指针、函数、切片、字典、管道 或 接口
	type typeInfo struct {
		code string
		f    func(types.Type) bool
	}
	typeNames := []string{"pointer", "function", "slice", "map", "channel", "interface"}
	typeMap := map[string]typeInfo{
		"pointer": {"type T *int", func(t types.Type) bool {
			_, ok := gospec.IsPointer(t)
			return ok
		}},
		"function": {"type T func()", func(t types.Type) bool {
			_, ok := gospec.IsFunction(t)
			return ok
		}},
		"slice": {"type T []string", func(t types.Type) bool {
			_, ok := gospec.IsSlice(t)
			return ok
		}},
		"map": {"type T map[string]int", func(t types.Type) bool {
			_, ok := gospec.IsMap(t)
			return ok
		}},
		"channel": {"type T chan int", func(t types.Type) bool {
			_, ok := gospec.IsChan(t)
			return ok
		}},
		"interface": {"type T interface{}", func(t types.Type) bool {
			_, ok := gospec.IsInterface(t)
			return ok
		}},
	}
	for _, v := range typeNames {
		ti, ok := typeMap[v]
		if !ok {
			continue
		}
		s := gospec.NewSpec(ti.code)
		if s.IsInUniverse("nil") {
			fmt.Print("x is the predeclared identifier nil and ")
		}
		if ti.f(s.GetType("T")) {
			fmt.Println("T is a " + v)
		}
		if s.Assignment("nil", "T") {
			fmt.Println("x could assignable to T")
		}
	}
	/* the output is:
	x is the predeclared identifier nil and T is a pointer
	x could assignable to T
	x is the predeclared identifier nil and T is a function
	x could assignable to T
	x is the predeclared identifier nil and T is a slice
	x could assignable to T
	x is the predeclared identifier nil and T is a map
	x could assignable to T
	x is the predeclared identifier nil and T is a channel
	x could assignable to T
	x is the predeclared identifier nil and T is a interface
	x could assignable to T
	*/
}

func assignExample06() {

	//6. x is an untyped constant representable by a value of type T.
	// x 是一个未显示定义的常量，且是个可以被 T 代表的值
	// 具体见下方的 Representability 可被代表性
	s := gospec.NewSpec(`
type T int
const x = 1
`)
	x := s.GetType("x")
	if gospec.IsUntyped(x) && gospec.IsConstType(x) {
		fmt.Print("x is an untyped constant ")
	}
	if s.Representable("x", "T") {
		fmt.Println("representable by a value of type T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x is an untyped constant representable by a value of type T
	x could assignable to T
	*/
}
