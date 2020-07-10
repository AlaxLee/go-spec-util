package gospec

import (
	"github.com/AlaxLee/go-forceexport"
	"go/ast"
	"go/constant"
	"go/types"
	"log"
	"runtime"
)

//must same as method (*Checker).assignment in src/go/types/assignments.go
var _assignment func(checker *types.Checker, x *operand, T types.Type, context string)

//must same as method (*Checker).representable in src/go/types/expr.go
var _representable func(checker *types.Checker, x *operand, typ *types.Basic)

func init() {
	if runtime.Version() != "go1.14" {
		log.Println("warning: version is not go1.14, may be panic in usring")
	}

	// 先触发一次 go/types.(*Checker).assignment 方法，以保证在 runtime.firstmoduledata 中能查到它
	// 如果查找不到，会报错：panic: Invalid function name: go/types.(*Checker).assignment
	// 先触发一次 go/types.(*Checker).representable 方法，以保证在 runtime.firstmoduledata 中能查到它
	// 如果查找不到，会报错：panic: Invalid function name: go/types.(*Checker).representable
	NewSpec("")

	// 将 _assignment 映射为 go/types.(*Checker).assignment
	err := forceexport.GetFunc(&_assignment, "go/types.(*Checker).assignment")
	if err != nil {
		panic(err)
	}

	// 将 _representable 映射为 go/types.(*Checker).representable
	err = forceexport.GetFunc(&_representable, "go/types.(*Checker).representable")
	if err != nil {
		panic(err)
	}
}

//must be kept in sync with operand in src/go/types/operand.go
type operand struct {
	mode operandMode
	expr ast.Expr
	typ  types.Type
	val  constant.Value
	id   builtinId
}
type operandMode byte
type builtinId int

const constant_ operandMode = 4
const value operandMode = 7

//must be kept in sync with isNamed in src/go/types/predicates.go
func isNamed(typ types.Type) bool {
	if _, ok := typ.(*types.Basic); ok {
		return ok
	}
	_, ok := typ.(*types.Named)
	return ok
}

func IsUntyped(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsUntyped != 0
}

func IsConstType(typ types.Type) bool {
	t, ok := typ.Underlying().(*types.Basic)
	return ok && t.Info()&types.IsConstType != 0
}
