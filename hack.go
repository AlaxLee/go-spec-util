package gospec

import (
	"go/ast"
	"go/constant"
	"go/types"
	_ "unsafe"
)

//go:linkname _assignment go/types.(*Checker).assignment
func _assignment(checker *types.Checker, x *operand, T types.Type, context string)

//go:linkname _representable go/types.(*Checker).representable
func _representable(checker *types.Checker, x *operand, typ *types.Basic)

//go:linkname _conversion go/types.(*Checker).conversion
func _conversion(checker *types.Checker, x *operand, T types.Type)

//go:linkname isNamed go/types.isNamed
func isNamed(typ types.Type) bool

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
