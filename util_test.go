package gospec

import (
	"strings"
	"testing"
)

func Test_addPackageHeadToCode(t *testing.T) {
	code := ` var a int`
	addPackageHeadToCode(&code)
	if !strings.HasPrefix(code, "package example\n") {
		t.Error(`test failed`)
	}
}

func Test_mustGetPackageNameFromCode(t *testing.T) {
	code := `
package haha
var a int`
	pn := mustGetPackageNameFromCode(code)
	if pn != "haha" {
		t.Error(`test failed`)
	}
}
