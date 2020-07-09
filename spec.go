package gospec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func init() {
	spec = NewSpec("")
}

var spec *Spec

type Spec struct {
	code    string
	file    *ast.File
	pkg     *types.Package
	checker *types.Checker
}

func NewSpec(code string) *Spec {
	s := new(Spec)
	addPackageHeadToCode(&code)
	packageName := mustGetPackageNameFromCode(code)
	s.code = code

	var err error
	fset := token.NewFileSet()
	s.file, err = parser.ParseFile(fset, packageName+".go", code, 0)
	if err != nil {
		log.Panicf("parse code failed: %s", err)
	}
	c := new(types.Config)
	c.Error = func(err error) {} // 防止触发 go/types.(*Checker).err 方法里的 panic
	s.pkg = types.NewPackage(packageName, "")
	s.checker = types.NewChecker(c, fset, s.pkg, nil)

	// 此方法会触发一次 go/types.(*Checker).assignment 方法，以保证在 runtime.firstmoduledata 中能查到它
	err = s.checker.Files([]*ast.File{s.file})
	if err != nil {
		log.Panicf("check file failed: %s", err)
	}
	return s
}

func (s *Spec) GetTypeObject(v string) types.Object {
	//return lookupByBFS(s.pkg.Scope(), v)
	return s.pkg.Scope().Lookup(v)
}

func (s *Spec) MustGetValidTypeObject(v string) types.Object {
	o := s.GetTypeObject(v)
	if o == nil {
		panic("find <" + v + "> in code <" + s.code + "> failed")
	}
	return o
}

func (s *Spec) GetType(v string) types.Type {
	o := s.GetTypeObject(v)
	if o == nil {
		return nil
	} else {
		return o.Type()
	}
}

func (s *Spec) MustGetValidType(v string) types.Type {
	o := s.GetTypeObject(v)
	if o == nil {
		panic("find <" + v + "> in code <" + s.code + "> failed")
	}
	return o.Type()
}

func (s *Spec) GetUnderlyingType(v string) types.Type {
	t := s.GetType(v)
	if t == nil {
		return nil
	}
	return t.Underlying()
}

func (s *Spec) GetBaseType(v string) types.Type {
	t := s.GetType(v)
	if t == nil {
		return nil
	}
	p, isPointer := t.(*types.Pointer)
	if !isPointer {
		return nil
	}
	return p.Elem()
}

func (s *Spec) IsDefinedType(v string) bool {
	t := s.GetType(v)
	return isNamed(t)
}

func GetTypeObject(code, v string) types.Object {
	s := NewSpec(code)
	return s.GetTypeObject(v)
}

func MustGetValidTypeObject(code, v string) types.Object {
	s := NewSpec(code)
	return s.MustGetValidTypeObject(v)
}

func GetType(code, v string) types.Type {
	s := NewSpec(code)
	return s.GetType(v)
}

func MustGetType(code, v string) types.Type {
	s := NewSpec(code)
	return s.MustGetValidType(v)
}

func GetUnderlyingType(code, v string) types.Type {
	s := NewSpec(code)
	return s.GetUnderlyingType(v)
}

func GetBaseType(code, v string) types.Type {
	s := NewSpec(code)
	return s.GetBaseType(v)
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

func lookupByBFS(scope *types.Scope, v string) types.Object {
	o := scope.Lookup(v)
	if o != nil {
		return o
	}
	for i := 0; i < scope.NumChildren(); i++ {
		o = lookupByBFS(scope.Child(i), v)
		if o != nil {
			return o
		}
	}
	return nil
}

// isNamed must be kept in sync with isNamed in src/go/types/predicates.go
func isNamed(typ types.Type) bool {
	if _, ok := typ.(*types.Basic); ok {
		return ok
	}
	_, ok := typ.(*types.Named)
	return ok
}
