package gospec

func (s *Spec) Assignment(v, t string) bool {
	V := s.MustGetValidType(v)
	T := s.MustGetValidType(t)

	x := &operand{mode: value, typ: V}
	_assignment(s.checker, x, T, "")
	if x.mode > 0 {
		return true
	}
	return false
}

func Assignment(code, v, t string) bool {
	s := NewSpec(code)
	return s.Assignment(v, t)
}
