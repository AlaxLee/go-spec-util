package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
	s := gospec.NewSpec(`
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
)`)
	for _, t := range []string{"A1", "A2", "B1", "B2", "B3", "B4"} {
		fmt.Printf("%s's underlying type is %s\n", t, s.GetUnderlyingType(t).String())
	}
	/* the output is:
	A1's underlying type is string
	A2's underlying type is string
	B1's underlying type is string
	B2's underlying type is string
	B3's underlying type is []test.B1
	B4's underlying type is []test.B1
	*/
}
