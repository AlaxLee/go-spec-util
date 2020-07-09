package gospec

import "go/types"

func (s *Spec) IsDefinedType(v string) bool {
	t := s.GetType(v)
	return isNamed(t)
}

//IsDefinedType(t types.Type) bool
//or
//IsDefinedType(code,v string) bool
func IsDefinedType(a ...interface{}) bool {
	switch len(a) {
	case 1:
		t, ok := a[0].(types.Type)
		if !ok {
			panic("arg must be a types.Type")
		}
		return isNamed(t)
	case 2:
		code, ok1 := a[0].(string)
		v, ok2 := a[1].(string)
		if !ok1 || !ok2 {
			panic("args must all string")

		}
		s := NewSpec(code)
		return s.IsDefinedType(v)
	default:
		panic("unexpect")
	}
	return true
}
