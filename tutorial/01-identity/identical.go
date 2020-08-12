package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
	"go/types"
)

func main() {
	identicalExample01()
	identicalExample02()
}

func identicalExample01() {
	var s *gospec.Spec

	// 1. type A1 = B  则 A1 与 B 类型相同
	// type A1 = B , they are identical
	s = gospec.NewSpec(`type B int; type A1 = B`)
	FormatIfIdentical(s, "A1", "B")
	/* the output is:
	A1     type is example.B
	B      type is example.B
	their type are identical
	*/

	// 2. type A2 B  则 A2 与 B 类型不同
	// type A2 B , they are different
	s = gospec.NewSpec(`type B int; type A2 B`)
	FormatIfIdentical(s, "A2", "B")
	/* the output is:
	A2     type is example.A2
	B      type is example.B
	their type are different
	*/

	// 3.1. 数组：如果 元素类型 和 数组长度 都相同，那么类型相同
	// Two array types are identical if they have identical element types and the same array length.
	s = gospec.NewSpec(`var a311, a312 [2]int; var a313 [2]int64`)
	FormatIfIdentical(s, "a311", "a312")
	/* the output is:
	a311   type is [2]int
	a312   type is [2]int
	their type are identical
	*/
	FormatIfIdentical(s, "a311", "a313")
	/* the output is:
	a311   type is [2]int
	a313   type is [2]int64
	their type are different
	*/

	// 3.2  切片：如果 元素类型 相同，那么类型相同
	// Two slice types are identical if they have identical element types.
	s = gospec.NewSpec(`var a321, a322 []int; var a323 []int64`)
	FormatIfIdentical(s, "a321", "a322")
	/* the output is:
	a321   type is []int
	a322   type is []int
	their type are identical
	*/
	FormatIfIdentical(s, "a321", "a323")
	/* the output is:
	a321   type is []int
	a323   type is []int64
	their type are different
	*/

	// 3.3  结构体： 如果 属性顺序 相同，且对应属性的名字、类型、标签都相同，那么类型相同。
	// 注意：不同包里面的结构体的未导出的属性一定不相同
	// 简单来说，如果在不同的包里的两个结构体满足前述条件，那么，只有当它们所有的属性都是导出属性的时候，它们俩类型才相同，只要有未导出属性，那么它们俩类型一定不相同
	// Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. Non-exported field names from different packages are always different.
	s = gospec.NewSpec(`
type B int
type A1 = B
var a331 struct {
    x int      "one"
    Y string
    c []A1  ` + "`B slice`" + `
}
var a332 struct {
    x int   "one"
    Y string
    c []B  ` + "`B slice`" + `
}
`)
	FormatIfIdentical(s, "a331", "a332")
	/* the output is:
	a331   type is struct{x int "one"; Y string; c []example.B "B slice"}
	a332   type is struct{x int "one"; Y string; c []example.B "B slice"}
	their type are identical
	*/

	specOfStruct1 := gospec.NewSpec(`
package PA

var a333 struct {
	X int "one"
	Y string
	C []int ` + "`int slice`" + `
}
var a335 struct {
	X int "one"
	y string
	C []int ` + "`int slice`" + `
}
`)
	specOfStruct2 := gospec.NewSpec(`
package QA

var a334 struct {
	X int "one"
	Y string
	C []int ` + "`int slice`" + `
}
var a336 struct {
	X int "one"
	y string
	C []int ` + "`int slice`" + `
}
`)
	FormatIfObjectIdentical(specOfStruct1.MustGetValidTypeObject("a333"), specOfStruct2.MustGetValidTypeObject("a334"))
	/* the output is:
	a333   type is struct{X int "one"; Y string; C []int "int slice"}
	a334   type is struct{X int "one"; Y string; C []int "int slice"}
	their type are identical
	*/

	FormatIfObjectIdentical(specOfStruct1.MustGetValidTypeObject("a335"), specOfStruct2.MustGetValidTypeObject("a336"))
	/* the output is:
	a335   type is struct{X int "one"; y string; C []int "int slice"}
	a336   type is struct{X int "one"; y string; C []int "int slice"}
	their type are different
	*/

	// 3.4  指针：如果 基本类型（base type） 相同，那么类型相同
	// Two pointer types are identical if they have identical base types.
	s = gospec.NewSpec(`
type B int
type A1 = B
var a341 *B
var a342 *A1
`)
	FormatIfIdentical(s, "a341", "a342")
	/* the output is:
	a341   type is *example.B
	a342   type is *example.B
	their type are identical
	*/

	// 3.5  函数：如果两者具有 相同数量 的参数和返回值，相应的参数和返回值的 类型相同，并且要么两个函数都有可变参数，要么都没有。
	// 参数名和结果名不需要匹配。
	// Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.
	s = gospec.NewSpec(`
var a351 func(a, b int, z float64, opt ...interface{}) (success bool)
var a352 func(x int, y int, z float64, too ...interface{}) (ok bool)
`)
	FormatIfIdentical(s, "a351", "a352")
	/* the output is:
	a351   type is func(a int, b int, z float64, opt ...interface{}) (success bool)
	a352   type is func(x int, y int, z float64, too ...interface{}) (ok bool)
	their type are identical
	*/

	// 3.6  接口：如果两者的方法集内的方法的 名称、类型 都相同，那么类型相同。
	// 来自不同程序包的未导出方法名称始终是不同的。 方法的顺序无关紧要。
	// Two interface types are identical if they have the same set of methods with the same names and identical function types. Non-exported method names from different packages are always different. The order of the methods is irrelevant.
	s = gospec.NewSpec(`
type B int
type A1 = B
var a361 interface {
    X() int
	y(string)
	c() []A1
}
var a362 interface {
	c() []B
    X() int
	y(string)
}
`)
	FormatIfIdentical(s, "a361", "a362")
	/* the output is:
	a361   type is interface{X() int; c() []example.B; y(string)}
	a362   type is interface{X() int; c() []example.B; y(string)}
	their type are identical
	*/

	specOfInterface1 := gospec.NewSpec(`
package PA

var a363 interface {
	X() int
	Y(string)
	C() []int
}
var a365 interface {
	X() int
	Y(string)
	c() []int
}
`)
	specOfInterface2 := gospec.NewSpec(`
package QA

var a364 interface {
	C() []int
	X() int
	Y(string)
}
var a366 interface {
	X() int
	c() []int
	Y(string)
}
`)
	FormatIfObjectIdentical(specOfInterface1.MustGetValidTypeObject("a363"), specOfInterface2.MustGetValidTypeObject("a364"))
	/* the output is:
	a363   type is interface{C() []int; X() int; Y(string)}
	a364   type is interface{C() []int; X() int; Y(string)}
	their type are identical
	*/

	FormatIfObjectIdentical(specOfInterface1.MustGetValidTypeObject("a365"), specOfInterface2.MustGetValidTypeObject("a366"))
	/* the output is:
	a365   type is interface{c() []int; X() int; Y(string)}
	a366   type is interface{c() []int; X() int; Y(string)}
	their type are different
	*/

	// 3.7  字典：如果两者的 键和值的类型都相同，那么类型相同
	// Two map types are identical if they have identical key and element types.
	s = gospec.NewSpec(`
type B int
type A1 = B
var a371 map[A1]string
var a372 map[B]string
`)
	FormatIfIdentical(s, "a371", "a372")
	/* the output is:
	a371   type is map[example.B]string
	a372   type is map[example.B]string
	their type are identical
	*/

	// 3.8  管道：如果两者的 元素类型相同、方向相同，那么类型相同
	// Two channel types are identical if they have identical element types and the same direction.
	s = gospec.NewSpec(`
type B int
type A1 = B
var a381 chan<- A1
var a382 chan<- B
`)
	FormatIfIdentical(s, "a381", "a382")
	/* the output is:
	a381   type is chan<- example.B
	a382   type is chan<- example.B
	their type are identical
	*/

	//3.9  basic type：诸如 int、string 这种简单类型，字面量相同，那么类型相同
	s = gospec.NewSpec(`var a391 byte; var a392 byte`)
	FormatIfIdentical(s, "a391", "a392")
	/* the output is:
	a391   type is byte
	a392   type is byte
	their type are identical
	*/
}

func identicalExample02() {
	s := gospec.NewSpec(`
type U = struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

type V = struct {
	Name    string ` + "`" + `json:"name"` + "`" + `
	Address *struct {
		Street string ` + "`" + `json:"street"` + "`" + `
		City   string ` + "`" + `json:"city"` + "`" + `
	} ` + "`" + `json:"address"` + "`" + `
}
`)

	if s.IdenticalIgnoreTags("U", "V") {
		fmt.Println("they IdenticalIgnoreTags")
	}
	/* the output is:
	they IdenticalIgnoreTags
	*/
}

func FormatIfIdentical(s *gospec.Spec, v, t string) {
	vo := s.MustGetValidTypeObject(v)
	to := s.MustGetValidTypeObject(t)
	FormatIfObjectIdentical(vo, to)
}

func FormatIfObjectIdentical(vo, to types.Object) {
	result := fmt.Sprintf("%-6s type is %-10s\n%-6s type is %-10s\ntheir type are ", vo.Name(), vo.Type(), to.Name(), to.Type())
	if gospec.Identical(vo, to) {
		result += "identical"
	} else {
		result += "different"
	}
	fmt.Println(result + "\n")
}
