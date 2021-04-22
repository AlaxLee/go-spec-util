package gospec

import (
	"go/types"
	"testing"
)

// func (s *Spec) Assignment(v, t string) bool
// 1. x's type is identical to T.
// x 的类型 与 T 相同
func TestAssignment01(t *testing.T) {
	s := NewSpec(`
type T = int
var x = 1
`)
	if s.Identical("x", "T") {
		if !s.Assignment("x", "T") {
			t.Error(`test rule: 
1. x's type is identical to T.
failed`)
		}
	}
}

// 2. x's type V and T have identical underlying types
// and at least one of V or T is not a defined type.
// x 的类型 V 和 T 有相同的 underlying type 并且 V 或 T 至少有一个是未（显示）定义类型
func TestAssignment02(t *testing.T) {
	s := NewSpec(`
type V = func()
type T func()
var x V
`)
	V := s.GetType("x")
	T := s.GetType("T")
	if Identical(V.Underlying(), T.Underlying()) && !IsDefinedType(V) && IsDefinedType(T) {
		if !s.Assignment("x", "T") {
			t.Error(`test rule:
2. x's type V and T have identical underlying types and at least one of V or T is not a defined type.
failed`)
		}
	}
}

// 3. T is an interface type and x implements T.
// T 是一个接口，x 实现了 T
func TestAssignment03(t *testing.T) {
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
			if !s.Assignment("x", "T") {
				t.Error(`test rule:
3. T is an interface type and x implements T.
failed`)
			}
		}
	}
}

// 4. x is a bidirectional channel value, T is a channel type,
// x's type V and T have identical element types,
// and at least one of V or T is not a defined type.
// x 是一个双向管道的值，T 是一个管道类型
// x 的类型 V 和 T 有相同的元素类型，并且 V 或 T 至少有一个是未（显示）定义类型
func TestAssignment04(t *testing.T) {
	s := NewSpec(`
type T chan int
var x = make(chan int)
`)
	V := s.GetType("x")
	T := s.GetType("T")

	var xIsABidirectionalChanValue, TisAChanType bool

	vc, ok := ToChan(V)
	if ok && vc.Dir() == types.SendRecv {
		xIsABidirectionalChanValue = true
	}
	tc, ok := ToChan(T)
	if ok {
		TisAChanType = true
	}

	if xIsABidirectionalChanValue && TisAChanType {
		if Identical(vc.Elem(), tc.Elem()) {
			if !IsDefinedType(V) && IsDefinedType(T) {
				if !s.Assignment("x", "T") {
					t.Error(`test rule:
4. x is a bidirectional channel value, T is a channel type, x's type V and T have identical element types, and at least one of V or T is not a defined type.
failed`)
				}
			}
		}
	}
}

// 5. x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
// x 是 nil，T 是一个 指针、函数、切片、字典、管道 或 接口
func TestAssignment05(t *testing.T) {
	type typeInfo struct {
		code string
		f    func(types.Type) bool
	}
	typeNames := []string{"pointer", "function", "slice", "map", "channel", "interface"}
	typeMap := map[string]typeInfo{
		"pointer": {"type T *int", func(t types.Type) bool {
			_, ok := ToPointer(t)
			return ok
		}},
		"function": {"type T func()", func(t types.Type) bool {
			_, ok := ToFunction(t)
			return ok
		}},
		"slice": {"type T []string", func(t types.Type) bool {
			_, ok := ToSlice(t)
			return ok
		}},
		"map": {"type T map[string]int", func(t types.Type) bool {
			_, ok := ToMap(t)
			return ok
		}},
		"channel": {"type T chan int", func(t types.Type) bool {
			_, ok := ToChan(t)
			return ok
		}},
		"interface": {"type T interface{}", func(t types.Type) bool {
			_, ok := ToInterface(t)
			return ok
		}},
	}
	for _, v := range typeNames {
		ti, ok := typeMap[v]
		if !ok {
			continue
		}
		s := NewSpec(ti.code)
		if s.IsInUniverse("nil") && ti.f(s.GetType("T")) {
			if !s.Assignment("nil", "T") {
				t.Error(`test rule:
5. x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
failed`)
			}
		}
	}
}

// 6. x is an untyped constant representable by a value of type T.
// x 是一个未显示定义的常量，且是个可以被 T 代表的值
// 具体见 Representability 可被代表性
func TestAssignment06(t *testing.T) {
	s := NewSpec(`
type T int
const x = 1
`)
	x := s.GetType("x")
	if IsUntyped(x) && IsConstType(x) {
		if s.Representable("x", "T") {
			if !s.Assignment("x", "T") {
				t.Error(`test rule:
6. x is an untyped constant representable by a value of type T.failed`)
			}
		}
	}
}

// func Assignment(code, v, t string) bool
func TestAssignment07(t *testing.T) {
	code := `type T = int; var x = 1`
	if Identical(code, "x", "T") {
		if !Assignment(code, "x", "T") {
			t.Error(`test rule: 
1. x's type is identical to T.
failed`)
		}
	}
}
