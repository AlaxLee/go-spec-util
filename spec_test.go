package gospec

import (
	"fmt"
	"testing"
)

// func (s *Spec) GetTypeObject(v string) types.Object
func TestSpec_GetTypeObject(t *testing.T) {
	code := `
package haha
import "fmt"
const a = 1
func main() {
    var b = 2.0
    type c struct {
        d string
    }
    fmt.Println(a, b)
}`
	/*
		|          | Universe | Package | File | Local |
		| :------: | :------: | :-----: | :--: | :---: |
		| Builtin  |    √     |         |      |       |
		|   Nil    |    √     |         |      |       |
		|  Const   |    √     |    √    |      |   √   |
		| TypeName |    √     |    √    |      |   √   |
		|   Func   |          |    √    |      |       |
		|   Var    |          |    √    |      |   √   |
		| PkgName  |          |         |  √   |       |
		|  Label   |          |         |      |   √   |

		The Universe scope contains all predeclared objects of Go. For example: type/int/bool ...
		Package: a
		File: fmt
		Local: b, c
	*/
	s := NewSpec(code)
	s.SearchKind = SearchOnlyPackage
	if s.GetTypeObject("bool") == nil &&
		s.GetTypeObject("fmt") == nil &&
		s.GetTypeObject("a").String() == "const haha.a untyped int" &&
		s.GetTypeObject("b") == nil {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchPackageAndUniverse
	if s.GetTypeObject("bool").String() == "type bool" &&
		s.GetTypeObject("fmt") == nil &&
		s.GetTypeObject("a").String() == "const haha.a untyped int" &&
		s.GetTypeObject("b") == nil {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchAll
	if s.GetTypeObject("bool").String() == "type bool" &&
		s.GetTypeObject("fmt").String() == "package fmt" &&
		s.GetTypeObject("a").String() == "const haha.a untyped int" &&
		s.GetTypeObject("b").String() == "var b float64" {
	} else {
		t.Error(`test failed`)
	}

}

//func (s *Spec) MustGetValidTypeObject(v string) types.Object
func TestSpec_MustGetValidTypeObject(t *testing.T) {
	code := `
package haha
import "fmt"
const a = 1
func main() {
    var b = 2.0
    type c struct {
        d string
    }
    fmt.Println(a, b)
}`
	testGet := func(s *Spec, v string) (r string) {
		defer func() {
			if err := recover(); err != nil {
				r = fmt.Sprintf("%s", err)
			}
		}()
		s.MustGetValidTypeObject(v)
		return
	}

	s := NewSpec(code)
	s.SearchKind = SearchOnlyPackage
	if testGet(s, "bool") == "find <bool> in code <"+code+"> failed" &&
		testGet(s, "fmt") == "find <fmt> in code <"+code+"> failed" &&
		s.MustGetValidTypeObject("a").String() == "const haha.a untyped int" &&
		testGet(s, "b") == "find <b> in code <"+code+"> failed" {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchPackageAndUniverse
	if s.MustGetValidTypeObject("bool").String() == "type bool" &&
		testGet(s, "fmt") == "find <fmt> in code <"+code+"> failed" &&
		s.MustGetValidTypeObject("a").String() == "const haha.a untyped int" &&
		testGet(s, "b") == "find <b> in code <"+code+"> failed" {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchAll
	if s.MustGetValidTypeObject("bool").String() == "type bool" &&
		s.MustGetValidTypeObject("fmt").String() == "package fmt" &&
		s.MustGetValidTypeObject("a").String() == "const haha.a untyped int" &&
		s.MustGetValidTypeObject("b").String() == "var b float64" {
	} else {
		t.Error(`test failed`)
	}
}

//func (s *Spec) GetType(v string) types.Type
func TestSpec_GetType(t *testing.T) {
	code := `
package haha
import "fmt"
const a = 1
func main() {
    var b = 2.0
    type c struct {
        d string
    }
    fmt.Println(a, b)
}`
	s := NewSpec(code)
	s.SearchKind = SearchOnlyPackage
	if s.GetType("bool") == nil &&
		s.GetType("fmt") == nil &&
		s.GetType("a").String() == "untyped int" &&
		s.GetType("b") == nil {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchPackageAndUniverse
	if s.GetType("bool").String() == "bool" &&
		s.GetType("fmt") == nil &&
		s.GetType("a").String() == "untyped int" &&
		s.GetType("b") == nil {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchAll
	if s.GetType("bool").String() == "bool" &&
		s.GetType("fmt").String() == "invalid type" &&
		s.GetType("a").String() == "untyped int" &&
		s.GetType("b").String() == "float64" {
	} else {
		t.Error(`test failed`)
	}
}

//func (s *Spec) MustGetValidType(v string) types.Type
func TestSpec_MustGetValidType(t *testing.T) {
	code := `
package haha
import "fmt"
const a = 1
func main() {
    var b = 2.0
    type c struct {
        d string
    }
    fmt.Println(a, b)
}`
	testGet := func(s *Spec, v string) (r string) {
		defer func() {
			if err := recover(); err != nil {
				r = fmt.Sprintf("%s", err)
			}
		}()
		s.MustGetValidType(v)
		return
	}

	s := NewSpec(code)
	s.SearchKind = SearchOnlyPackage
	if testGet(s, "bool") == "find <bool> in code <"+code+"> failed" &&
		testGet(s, "fmt") == "find <fmt> in code <"+code+"> failed" &&
		s.MustGetValidType("a").String() == "untyped int" &&
		testGet(s, "b") == "find <b> in code <"+code+"> failed" {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchPackageAndUniverse
	if s.MustGetValidType("bool").String() == "bool" &&
		testGet(s, "fmt") == "find <fmt> in code <"+code+"> failed" &&
		s.MustGetValidType("a").String() == "untyped int" &&
		testGet(s, "b") == "find <b> in code <"+code+"> failed" {
	} else {
		t.Error(`test failed`)
	}
	s.SearchKind = SearchAll
	if s.MustGetValidType("bool").String() == "bool" &&
		s.MustGetValidType("fmt").String() == "invalid type" &&
		s.MustGetValidType("a").String() == "untyped int" &&
		s.MustGetValidType("b").String() == "float64" {
	} else {
		t.Error(`test failed`)
	}
}

//func (s *Spec) GetUnderlyingType(v string) types.Type
func TestSpec_GetUnderlyingType(t *testing.T) {
	code := `
package test
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1  // []B1 is a type literal
	B4 B3
)`
	s := NewSpec(code)
	for k, v := range map[string]string{"A1": "string", "A2": "string", "B1": "string", "B2": "string", "B3": "[]test.B1", "B4": "[]test.B1"} {
		if s.GetUnderlyingType(k).String() != v {
			t.Error(`test failed`)
		}
	}
}

//func (s *Spec) GetBaseType(v string) types.Type
func TestSpec_GetBaseType(t *testing.T) {
	code := `
package test
func main() {
    var a int
	b := &a
	_ = b
}`
	s := NewSpec(code)
	s.SearchKind = SearchAll
	fmt.Println(s.GetBaseType("b"))
	if s.GetBaseType("a") == nil && s.GetBaseType("b").String() == "int" {
	} else {
		t.Error(`test failed`)
	}
}

//func (s *Spec) IsInUniverse(v string) bool
func TestSpec_IsInUniverse(t *testing.T) {
	code := `
package haha
import "fmt"
const a = 1
func main() {
    var b = 2.0
    type c struct {
        d string
    }
    fmt.Println(a, b)
}`
	s := NewSpec(code)
	if s.IsInUniverse("bool") && !s.IsInUniverse("a") && !s.IsInUniverse("kaka") {
	} else {
		t.Error(`test failed`)
	}
}
