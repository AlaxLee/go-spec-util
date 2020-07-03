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
	addPackageHead(&code)
	s.code = code

	var err error
	fset := token.NewFileSet()
	s.file, err = parser.ParseFile(fset, "example.go", code, 0)
	if err != nil {
		log.Panicf("parse code failed: %s", err)
	}
	c := new(types.Config)
	s.pkg = types.NewPackage("example", "")
	s.checker = types.NewChecker(c, fset, s.pkg, nil)

	// 此方法会触发一次 go/types.(*Checker).assignment 方法，以保证在 runtime.firstmoduledata 中能查到它
	err = s.checker.Files([]*ast.File{s.file})
	if err != nil {
		log.Panicf("check file failed: %s", err)
	}
	return s
}

func (s *Spec) GetTypeObject(v string) types.Object {
	return s.pkg.Scope().Lookup(v)
}

func (s *Spec) GetType(v string) types.Type {
	o := s.GetTypeObject(v)
	if o == nil {
		return nil
	} else {
		return o.Type()
	}
}

func (s *Spec) MustGetType(v string) types.Type {
	o := s.GetTypeObject(v)
	if o == nil {
		panic("find <" + v + "> in code <" + s.code + "> failed")
	}
	return o.Type()
}

func GetTypeObject(code, v string) types.Object {
	s := NewSpec(code)
	return s.GetTypeObject(v)
}

func GetType(code, v string) types.Type {
	s := NewSpec(code)
	return s.GetType(v)
}

func MustGetType(code, v string) types.Type {
	s := NewSpec(code)
	return s.MustGetType(v)
}
