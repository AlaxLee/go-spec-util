package gospec

func (s *Spec) Conversion(v, t string) bool {
	vo := s.MustGetValidTypeObject(v)
	T := s.MustGetValidType(t)

	x := &operand{typ: vo.Type()}
	if constObj, ok := ToConstObject(vo); ok {
		x.mode = constant_
		x.val = constObj.Val()
	} else {
		x.mode = value
	}

	_conversion(s.checker, x, T)
	if x.mode > 0 {
		return true
	}
	return false
}

func Conversion(code, v, t string) bool {
	s := NewSpec(code)
	return s.Conversion(v, t)
}
