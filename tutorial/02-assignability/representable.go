package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
	representableExample01()
	representableExample02()
}

type Info struct {
	x string
	T string
}

func representableExample01() {
	/* x is representable by a value of T because
	x                   T

	'a'                 byte        97 is in the set of byte values
	97                  rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
	"foo"               string      "foo" is in the set of string values
	1024                int16       1024 is in the set of 16-bit integers
	42.0                byte        42 is in the set of unsigned 8-bit integers
	1e10                uint64      10000000000 is in the set of unsigned 64-bit integers
	2.718281828459045   float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
	-1e-1000            float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
	0i                  int         0 is an integer value
	(42 + 0i)           float32     42.0 (with zero imaginary part) is in the set of float32 values
	*/
	infos := []Info{
		{`'a'`, `byte`},
		{`97`, `rune`},
		{`"foo"`, `string`},
		{`1024`, `int16`},
		{`42.0`, `byte`},
		{`1e10`, `uint64`},
		{`2.718281828459045`, `float32`},
		{`-1e-1000`, `float64`},
		{`0i`, `int`},
		{`(42 + 0i)`, `float32`},
	}

	for _, v := range infos {
		code := fmt.Sprintf("type T %s; const x = %s", v.T, v.x)
		s := gospec.NewSpec(code)
		if s.Representable("x", "T") {
			fmt.Printf("%20s is representable by a value of %s\n", v.x, v.T)
		}
	}
	/* the output is:
	                 'a' is representable by a value of byte
	                  97 is representable by a value of rune
	               "foo" is representable by a value of string
	                1024 is representable by a value of int16
	                42.0 is representable by a value of byte
	                1e10 is representable by a value of uint64
	   2.718281828459045 is representable by a value of float32
	            -1e-1000 is representable by a value of float64
	                  0i is representable by a value of int
	           (42 + 0i) is representable by a value of float32
	*/

	for _, v := range infos {
		code := fmt.Sprintf("const x = %s", v.x)
		s := gospec.NewSpec(code)
		if s.Representable("x", v.T) {
			fmt.Printf("%20s is representable by a value of %s\n", v.x, v.T)
		}
	}
	/* the output is:
	                 'a' is representable by a value of byte
	                  97 is representable by a value of rune
	               "foo" is representable by a value of string
	                1024 is representable by a value of int16
	                42.0 is representable by a value of byte
	                1e10 is representable by a value of uint64
	   2.718281828459045 is representable by a value of float32
	            -1e-1000 is representable by a value of float64
	                  0i is representable by a value of int
	           (42 + 0i) is representable by a value of float32
	*/
}

func representableExample02() {
	/* x is not representable by a value of T because
	x                   T

	0                   bool        0 is not in the set of boolean values
	'a'                 string      'a' is a rune, it is not in the set of string values
	1024                byte        1024 is not in the set of unsigned 8-bit integers
	-1                  uint16      -1 is not in the set of unsigned 16-bit integers
	1.1                 int         1.1 is not an integer value
	42i                 float32     (0 + 42i) is not in the set of float32 values
	1e1000              float64     1e1000 overflows to IEEE +Inf after rounding
	*/
	infos := []Info{
		{`0`, `bool`},
		{`'a'`, `string`},
		{`1024`, `byte`},
		{`-1`, `uint16`},
		{`1.1`, `int`},
		{`42i`, `float32`},
		{`1e1000`, `float64`},
	}

	for _, v := range infos {
		code := fmt.Sprintf("type T %s; const x = %s", v.T, v.x)
		s := gospec.NewSpec(code)
		if !s.Representable("x", "T") {
			fmt.Printf("%10s is not representable by a value of %s\n", v.x, v.T)
		}
	}
	/* the output is:
	        0 is not representable by a value of bool
	      'a' is not representable by a value of string
	     1024 is not representable by a value of byte
	       -1 is not representable by a value of uint16
	      1.1 is not representable by a value of int
	      42i is not representable by a value of float32
	   1e1000 is not representable by a value of float64
	*/

	for _, v := range infos {
		code := fmt.Sprintf("const x = %s", v.x)
		s := gospec.NewSpec(code)
		if !s.Representable("x", v.T) {
			fmt.Printf("%10s is not representable by a value of %s\n", v.x, v.T)
		}
	}
	/* the output is:
	        0 is not representable by a value of bool
	      'a' is not representable by a value of string
	     1024 is not representable by a value of byte
	       -1 is not representable by a value of uint16
	      1.1 is not representable by a value of int
	      42i is not representable by a value of float32
	   1e1000 is not representable by a value of float64
	*/
}
