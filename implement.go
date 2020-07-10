package gospec

import "go/types"

func (s *Spec) Implements(v, t string) bool {
	V := s.MustGetValidType(v)
	T := s.MustGetValidType(t)

	return implements(V, T)
}

// Implements(v, t types.Type) bool
// or
// Implements(v, t types.Object) bool
// or
// Implements((code, v, t string) bool
func Implements(a ...interface{}) bool {
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
			return implements(v1, t1)
		} else if isAllObject {
			return implements(v2.Type(), t2.Type())
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
		return s.Implements(v, t)
	default:
		panic("unexpect")
	}

	return true
}

// logic like src/go/types/operand.go line 254
func implements(v, t types.Type) bool {
	tu := t.Underlying()
	ti, ok := tu.(*types.Interface)
	if !ok {
		return false
	}
	return types.Implements(v, ti)
}
