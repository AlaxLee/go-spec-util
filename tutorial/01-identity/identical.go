package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
	identicalExample01()
	identicalExample02()
}

func identicalExample01() {
	var s *gospec.Spec

	//1. type A1 = B  则 A1 与 B 类型相同
	gospec.OutputIfIdentical(
		`type B int; type A1 = B`,
		"A1", "B")
	/* 执行结果是
	A1     的类型是 example.B
	B      的类型是 example.B
	他们类型 相同
	*/

	//2. type A2 B  则 A2 与 B 类型不同
	gospec.OutputIfIdentical(
		`type B int; type A2 B`,
		"A2", "B")
	/* 执行结果是
	A2     的类型是 example.A2
	B      的类型是 example.B
	他们类型 不同
	*/

	//3.1  数组：如果 元素类型 和 数组长度 都相同，那么类型相同
	s = gospec.NewSpec(`var a311, a312 [2]int; var a313 [2]int64`)
	s.OutputIfIdentical("a311", "a312")
	/* 执行结果是
	a311   的类型是 [2]int
	a312   的类型是 [2]int
	他们类型 相同
	*/
	s.OutputIfIdentical("a311", "a313")
	/* 执行结果是
	a311   的类型是 [2]int
	a313   的类型是 [2]int64
	他们类型 不同
	*/

	//3.2  切片：如果 元素类型 相同，那么类型相同
	s = gospec.NewSpec(`var a321, a322 []int; var a323 []int64`)
	s.OutputIfIdentical("a321", "a322")
	/* 执行结果是
	a321   的类型是 []int
	a322   的类型是 []int
	他们类型 相同
	*/
	s.OutputIfIdentical("a321", "a323")
	/* 执行结果是
	a321   的类型是 []int
	a323   的类型是 []int64
	他们类型 不同
	*/

	//3.3  结构体： 如果 属性顺序 相同，且对应属性的名字、类型、标签都相同，那么类型相同。
	//注意：不同包里面的结构体的未导出的属性一定不相同
	//简单来说，如果在不同的包里的两个结构体满足前述条件，那么，只有当它们所有的属性都是导出属性的时候，它们俩类型才相同，只要有未导出属性，那么它们俩类型一定不相同
	gospec.OutputIfIdentical(`
type B int
type A1 = B
var a331 struct {
    x int      "one"
    Y string
    c []A1  `+"`B slice`"+`
}
var a332 struct {
    x int   "one"
    Y string
    c []B  `+"`B slice`"+`
}`,
		"a331", "a332")
	/* 执行结果是
	a331   的类型是 struct{x int "one"; Y string; c []example.B "B slice"}
	a332   的类型是 struct{x int "one"; Y string; c []example.B "B slice"}
	他们类型 相同
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
	gospec.FormatIfIdentical(specOfStruct1.MustGetValidTypeObject("a333"), specOfStruct2.MustGetValidTypeObject("a334"))
	/* 执行结果是
	a333   的类型是 struct{X int "one"; Y string; C []int "int slice"}
	a334   的类型是 struct{X int "one"; Y string; C []int "int slice"}
	他们类型 相同
	*/

	gospec.FormatIfIdentical(specOfStruct1.MustGetValidTypeObject("a335"), specOfStruct2.MustGetValidTypeObject("a336"))
	/* 执行结果是
	a335   的类型是 struct{X int "one"; y string; C []int "int slice"}
	a336   的类型是 struct{X int "one"; y string; C []int "int slice"}
	他们类型 不同
	*/

	//3.4  指针：如果 基本类型（base type） 相同，那么类型相同
	gospec.OutputIfIdentical(`
type B int
type A1 = B
var a341 *B
var a342 *A1`,
		"a341", "a342")
	/* 执行结果是
	a341   的类型是 *example.B
	a342   的类型是 *example.B
	他们类型 相同
	*/

	//3.5  函数：如果两者具有 相同数量 的参数和返回值，相应的参数和返回值的 类型相同，并且要么两个函数都有可变参数，要么都没有。
	//参数名和结果名不需要匹配。
	gospec.OutputIfIdentical(`
var a351 func(a, b int, z float64, opt ...interface{}) (success bool)
var a352 func(x int, y int, z float64, too ...interface{}) (ok bool)`,
		"a351", "a352")
	/* 执行结果是
	a351   的类型是 func(a int, b int, z float64, opt ...interface{}) (success bool)
	a352   的类型是 func(x int, y int, z float64, too ...interface{}) (ok bool)
	他们类型 相同
	*/

	//3.6  接口：如果两者的方法集内的方法的 名称、类型 都相同，那么类型相同。
	//来自不同程序包的未导出方法名称始终是不同的。 方法的顺序无关紧要。
	gospec.OutputIfIdentical(`
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
}`,
		"a361", "a362")
	/* 执行结果是
	a361   的类型是 interface{X() int; c() []example.B; y(string)}
	a362   的类型是 interface{X() int; c() []example.B; y(string)}
	他们类型 相同
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
	gospec.FormatIfIdentical(specOfInterface1.MustGetValidTypeObject("a363"), specOfInterface2.MustGetValidTypeObject("a364"))
	/* 执行结果是
	a363   的类型是 interface{C() []int; X() int; Y(string)}
	a364   的类型是 interface{C() []int; X() int; Y(string)}
	他们类型 相同
	*/

	gospec.FormatIfIdentical(specOfInterface1.MustGetValidTypeObject("a365"), specOfInterface2.MustGetValidTypeObject("a366"))
	/* 执行结果是
	a365   的类型是 interface{c() []int; X() int; Y(string)}
	a366   的类型是 interface{c() []int; X() int; Y(string)}
	他们类型 不同
	*/

	//3.7  字典：如果两者的 键和值的类型都相同，那么类型相同
	gospec.OutputIfIdentical(`
type B int
type A1 = B
var a371 map[A1]string
var a372 map[B]string`,
		"a371", "a372")
	/* 执行结果是
	a371   的类型是 map[example.B]string
	a372   的类型是 map[example.B]string
	他们类型 相同
	*/

	//3.8  管道：如果两者的 元素类型相同、方向相同，那么类型相同
	gospec.OutputIfIdentical(`
type B int
type A1 = B
var a381 chan<- A1
var a382 chan<- B`,
		"a381", "a382")
	/* 执行结果是
	a381   的类型是 chan<- example.B
	a382   的类型是 chan<- example.B
	他们类型 相同
	*/

	//3.9  basic type：诸如 int、string 这种简单类型，字面量相同，那么类型相同
	gospec.OutputIfIdentical(`var a391 byte; var a392 byte`,
		"a391", "a392")
	/* 执行结果是
	a391   的类型是 byte
	a392   的类型是 byte
	他们类型 相同
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
