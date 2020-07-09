package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {

	FormatIfDefinedType(`type A int`, "A")
	/* the output is:
	A's type is example.A, and analysed to *types.Named
	A is a defined type
	*/

	FormatIfDefinedType(`type A = int`, "A")
	/* the output is:
	A's type is int, and analysed to *types.Basic
	A is a defined type
	*/

	FormatIfDefinedType(`type A func()`, "A")
	/* the output is:
	A's type is example.A, and analysed to *types.Named
	A is a defined type
	*/

	FormatIfDefinedType(`type A = func()`, "A")
	/* the output is:
	A's type is func(), and analysed to *types.Signature
	A is not a defined type
	*/
}

func FormatIfDefinedType(code, v string) {
	s := gospec.NewSpec(code)
	t := s.GetType(v)
	fmt.Printf("%s's type is %v, and analysed to %T\n", v, t, t)
	if gospec.IsDefinedType(t) {
		fmt.Printf("%s is a defined type\n", v)
	} else {
		fmt.Printf("%s is not a defined type\n", v)
	}
}
