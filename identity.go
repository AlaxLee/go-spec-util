package gospec

import (
	"fmt"
	"go/types"
)

func (s *Spec) Identical(v, t string) bool {
	V := s.MustGetValidType(v)
	T := s.MustGetValidType(t)
	return types.Identical(V, T)
}

// Identical(v, t types.Type) bool
// or
// Identical(v, t types.Object) bool
// or
// Identical((code, v, t string) bool
func Identical(a ...interface{}) bool {
	switch len(a) {
	case 2:
		//v, t types.Type
		v1, okV1 := a[0].(types.Type)
		t1, okT1 := a[1].(types.Type)
		isAllType := okV1 && okT1

		//v, t types.Object
		v2, okV2 := a[0].(types.Object)
		t2, okT2 := a[1].(types.Object)
		isAllObject := okV2 && okT2

		if isAllType {
			return types.Identical(v1, t1)
		} else if isAllObject {
			return types.Identical(v2.Type(), t2.Type())
		} else {
			panic("args must all types.Type or all types.Object")
		}
	case 3:
		//code, v, t string
		code, ok1 := a[0].(string)
		v, ok2 := a[1].(string)
		t, ok3 := a[2].(string)
		if !ok1 || !ok2 || !ok3 {
			panic("args must all string")
		}
		s := NewSpec(code)
		return s.Identical(v, t)
	default:
		panic("unexpect")
	}

	return true
}

func (s *Spec) OutputIfIdentical(v, t string) {
	vo := s.MustGetValidTypeObject(v)
	to := s.MustGetValidTypeObject(t)
	FormatIfIdentical(vo, to)
}

func OutputIfIdentical(code, v, t string) {
	s := NewSpec(code)
	s.OutputIfIdentical(v, t)
}

func FormatIfIdentical(v, t types.Object) {
	result := fmt.Sprintf("%-6s 的类型是 %-10s\n%-6s 的类型是 %-10s\n他们类型 ", v.Name(), v.Type(), t.Name(), t.Type())
	if Identical(v, t) {
		result += "相同"
	} else {
		result += "不同"
	}
	fmt.Println(result)
}
