package gospec

import "go/types"

func (s *Spec) Representable(v, t string) bool {
	vo := s.MustGetValidTypeObject(v)
	T := s.MustGetValidType(t)

	vc, ok := vo.(*types.Const)
	if !ok {
		return false
	}

	x := &operand{mode: constant_, typ: vo.Type(), val: vc.Val()}

	tb, ok := IsBasic(T)
	if !ok {
		return false
	}

	_representable(s.checker, x, tb)
	if x.mode > 0 {
		return true
	}
	return false
}

func Representable(code, v, t string) bool {
	s := NewSpec(code)
	return s.Representable(v, t)
}
