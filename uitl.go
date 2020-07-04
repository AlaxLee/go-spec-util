package gospec

import (
	easyregexp "github.com/AlaxLee/easyregexp"
)

func addPackageHeadToCode(codePtr *string) {
	if !easyregexp.Match(`^\s*package\s+`, *codePtr) {
		*codePtr = easyregexp.ReplaceAll(`^\s*`, *codePtr, "package example\n")
	}
}

func mustGetPackageNameFromCode(code string) string {
	catches := easyregexp.Catch(`package\s+(\w+)`, code)
	if len(catches) == 0 {
		panic("get packageName failed")
	}
	return catches[0]
}
