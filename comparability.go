package gospec

import (
	"go/types"
	"reflect"
)

func GetDynamicTypeAtRuntime(i interface{}) reflect.Type {
	return reflect.ValueOf(i).Type()
}

func (s *Spec) Comparable(v string) bool {
	V := s.MustGetValidType(v)
	return types.Comparable(V)
}

// Comparable(t types.Type) bool
// or
// Comparable(code, v string) bool
func Comparable(a ...interface{}) bool {
	switch len(a) {
	case 1:
		//t types.Type
		if t, ok := a[0].(types.Type); ok {
			return types.Comparable(t)
		} else {
			panic("args must types.Type")
		}
	case 2:
		//code, v string
		code, ok1 := a[0].(string)
		v, ok2 := a[1].(string)
		if !ok1 || !ok2 {
			panic("args must all string")
		}
		s := NewSpec(code)
		return s.Comparable(v)
	default:
		panic("unexpect")
	}
	return false
}
