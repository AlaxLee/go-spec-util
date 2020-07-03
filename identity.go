package gospec

import (
	"go/types"
)

func (s *Spec) Identical(v, t string) bool {
	V := s.MustGetType(v)
	T := s.MustGetType(t)
	return types.Identical(V, T)
}

func Identical(code, v, t string) bool {
	s := NewSpec(code)
	return s.Identical(v, t)
}
