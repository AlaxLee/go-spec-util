package gospec

import "strings"

func addPackageHead(codePtr *string) {
	if !strings.HasPrefix(*codePtr, "package") {
		*codePtr = "package example\n" + *codePtr
	}
}
