package gospec

import (
	forceexport "github.com/AlaxLee/go-forceexport"
	"go/ast"
	"go/constant"
	"go/types"
	"log"
	"runtime"
)

//same as method (*Checker).assignment in src/go/types/assignments.go
var _assignment func(checker *types.Checker, x *operand, T types.Type, context string)

func init() {
	if runtime.Version() != "go1.14" {
		log.Println("warning: version is not go1.14, may be panic in usring")
	}

	// 先触发一次 go/types.(*Checker).assignment 方法，以保证在 runtime.firstmoduledata 中能查到它
	// 如果查找不到，会报错：panic: Invalid function name: go/types.(*Checker).assignment
	NewSpec("")

	// 将 _assignment 映射为 go/types.(*Checker).assignment
	err := forceexport.GetFunc(&_assignment, "go/types.(*Checker).assignment")
	if err != nil {
		panic(err)
	}
}

func (s *Spec) Assignment(v, t string) bool {
	V := s.MustGetType(v)
	T := s.MustGetType(t)

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

//same as type operand in src/go/types/operand.go
type operand struct {
	mode operandMode
	expr ast.Expr
	typ  types.Type
	val  constant.Value
	id   builtinId
}
type operandMode byte
type builtinId int

const value operandMode = 7
