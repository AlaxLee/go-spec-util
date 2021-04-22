package gospec

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func init() {
	spec = NewSpec("")
}

var spec *Spec

type SearchKind int

// src/go/types/scope.go Scope has 4 level：Universe、Package、File、Local
const (
	SearchOnlyPackage = iota
	SearchPackageAndUniverse
	SearchAll
)

type Spec struct {
	code    string
	file    *ast.File
	pkg     *types.Package
	checker *types.Checker
	SearchKind
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
	c.Error = func(err error) {}    // 防止触发 go/types.(*Checker).err 方法里的 panic
	c.Importer = importer.Default() // 增加golang包导入，使之可以识别 import 的包
	s.pkg = types.NewPackage(packageName, "")
	s.checker = types.NewChecker(c, fset, s.pkg, nil)
	s.SearchKind = SearchPackageAndUniverse // default search in universe and pkg scope

	// 此方法会触发一次 go/types.(*Checker).assignment 方法，以保证在 runtime.firstmoduledata 中能查到它
	// 此方法会触发一次 go/types.(*Checker).representable 方法，以保证在 runtime.firstmoduledata 中能查到它
	err = s.checker.Files([]*ast.File{s.file})
	if err != nil {
		log.Panicf("check file failed: %s", err)
	}
	return s
}

func (s *Spec) GetTypeObject(v string) types.Object {
	switch s.SearchKind {
	case SearchOnlyPackage:
		return s.pkg.Scope().Lookup(v)
	case SearchPackageAndUniverse:
		o := s.pkg.Scope().Lookup(v)
		if o == nil {
			o = types.Universe.Lookup(v)
		}
		return o
	case SearchAll:
		o := lookupByBFS(s.pkg.Scope(), v)
		if o == nil {
			o = types.Universe.Lookup(v)
		}
		return o
	default:
		panic("unexpect")
	}
	return nil
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

func (s *Spec) IsInUniverse(v string) bool {
	switch s.SearchKind {
	case SearchOnlyPackage, SearchPackageAndUniverse:
		if s.pkg.Scope().Lookup(v) == nil {
			if types.Universe.Lookup(v) != nil {
				return true
			}
		}
	case SearchAll:
		if lookupByBFS(s.pkg.Scope(), v) == nil {
			if types.Universe.Lookup(v) != nil {
				return true
			}
		}
	default:
		panic("unexpect")
	}
	return false
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
