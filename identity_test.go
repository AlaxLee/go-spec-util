package gospec

import "testing"

// func (s *Spec) Identical(v, t string) bool
// 1. type A1 = B  则 A1 与 B 类型相同
// type A1 = B , they are identical
func TestIdentical01(t *testing.T) {
	s := NewSpec(`type B int; type A1 = B`)
	if !s.Identical("A1", "B") {
		t.Error(`test rule failed`)
	}
}

// 2. type A2 B  则 A2 与 B 类型不同
// type A2 B , they are different
func TestIdentical02(t *testing.T) {
	s := NewSpec(`type B int; type A2 B`)
	if s.Identical("A2", "B") {
		t.Error(`test rule failed`)
	}
}

// 3.1. 数组：如果 元素类型 和 数组长度 都相同，那么类型相同
// Two array types are identical if they have identical element types and the same array length.
func TestIdentical03(t *testing.T) {
	s := NewSpec(`var a311, a312 [2]int; var a313 [2]int64`)
	if !s.Identical("a311", "a312") {
		t.Error(`test rule failed`)
	}
	if s.Identical("a311", "a313") {
		t.Error(`test rule failed`)
	}
}

// 3.2  切片：如果 元素类型 相同，那么类型相同
// Two slice types are identical if they have identical element types.
func TestIdentical04(t *testing.T) {
	s := NewSpec(`var a321, a322 []int; var a323 []int64`)
	if !s.Identical("a321", "a322") {
		t.Error(`test rule failed`)
	}
	if s.Identical("a321", "a323") {
		t.Error(`test rule failed`)
	}
}

// 3.3  结构体： 如果 属性顺序 相同，且对应属性的名字、类型、标签都相同，那么类型相同。
// 注意：不同包里面的结构体的未导出的属性一定不相同
// 简单来说，如果在不同的包里的两个结构体满足前述条件，那么，只有当它们所有的属性都是导出属性的时候，它们俩类型才相同，只要有未导出属性，那么它们俩类型一定不相同
// Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. Non-exported field names from different packages are always different.
func TestIdentical05(t *testing.T) {
	s := NewSpec(`
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
	if !s.Identical("a331", "a332") {
		t.Error(`test rule failed`)
	}

	specOfStruct1 := NewSpec(`
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
	specOfStruct2 := NewSpec(`
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
	// Identical(v, t types.Object) bool
	if !Identical(specOfStruct1.MustGetValidTypeObject("a333"), specOfStruct2.MustGetValidTypeObject("a334")) {
		t.Error(`test rule failed`)
	}
	if Identical(specOfStruct1.MustGetValidTypeObject("a335"), specOfStruct2.MustGetValidTypeObject("a336")) {
		t.Error(`test rule failed`)
	}
}

// 3.4  指针：如果 基本类型（base type） 相同，那么类型相同
// Two pointer types are identical if they have identical base types.
func TestIdentical06(t *testing.T) {
	s := NewSpec(`
type B int
type A1 = B
var a341 *B
var a342 *A1
`)
	if !s.Identical("a341", "a342") {
		t.Error(`test rule failed`)
	}
}

// 3.5  函数：如果两者具有 相同数量 的参数和返回值，相应的参数和返回值的 类型相同，并且要么两个函数都有可变参数，要么都没有。
// 参数名和结果名不需要匹配。
// Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.
func TestIdentical07(t *testing.T) {
	s := NewSpec(`
var a351 func(a, b int, z float64, opt ...interface{}) (success bool)
var a352 func(x int, y int, z float64, too ...interface{}) (ok bool)
`)
	if !s.Identical("a351", "a352") {
		t.Error(`test rule failed`)
	}
}

// 3.6  接口：如果两者的方法集内的方法的 名称、类型 都相同，那么类型相同。
// 来自不同程序包的未导出方法名称始终是不同的。 方法的顺序无关紧要。
// Two interface types are identical if they have the same set of methods with the same names and identical function types. Non-exported method names from different packages are always different. The order of the methods is irrelevant.
func TestIdentical08(t *testing.T) {
	s := NewSpec(`
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
	if !s.Identical("a361", "a362") {
		t.Error(`test rule failed`)
	}

	specOfInterface1 := NewSpec(`
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
	specOfInterface2 := NewSpec(`
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
	// Identical(v, t types.Object) bool
	if !Identical(specOfInterface1.MustGetValidTypeObject("a363"), specOfInterface2.MustGetValidTypeObject("a364")) {
		t.Error(`test rule failed`)
	}
	// Identical(v, t types.Object) bool
	if Identical(specOfInterface1.MustGetValidTypeObject("a365"), specOfInterface2.MustGetValidTypeObject("a366")) {
		t.Error(`test rule failed`)
	}
}

// 3.7  字典：如果两者的 键和值的类型都相同，那么类型相同
// Two map types are identical if they have identical key and element types.
func TestIdentical09(t *testing.T) {
	s := NewSpec(`
type B int
type A1 = B
var a371 map[A1]string
var a372 map[B]string
`)
	if !s.Identical("a371", "a372") {
		t.Error(`test rule failed`)
	}
}

// 3.8  管道：如果两者的 元素类型相同、方向相同，那么类型相同
// Two channel types are identical if they have identical element types and the same direction.
func TestIdentical10(t *testing.T) {
	s := NewSpec(`
type B int
type A1 = B
var a381 chan<- A1
var a382 chan<- B
`)
	if !s.Identical("a381", "a382") {
		t.Error(`test rule failed`)
	}
}

//3.9  basic type：诸如 int、string 这种简单类型，字面量相同，那么类型相同
func TestIdentical11(t *testing.T) {
	s := NewSpec(`var a391 byte; var a392 byte`)
	if !s.Identical("a391", "a392") {
		t.Error(`test rule failed`)
	}
}

// Identical(v, t types.Type) bool
// or
// Identical(v, t types.Object) bool
// or
// Identical((code, v, t string) bool
func TestIdentical12(t *testing.T) {
	s := NewSpec(`type B int; type A1 = B`)
	if !Identical(s.MustGetValidType("A1"), s.MustGetValidType("B")) {
		t.Error(`test rule failed`)
	}
	if !Identical(s.MustGetValidTypeObject("A1"), s.MustGetValidTypeObject("B")) {
		t.Error(`test rule failed`)
	}
	if !Identical(`type B int; type A1 = B`, "A1", "B") {
		t.Error(`test rule failed`)
	}
}

// func (s *Spec) IdenticalIgnoreTags(v, t string) bool
// IdenticalIgnoreTags(v, t types.Type) bool
// or
// IdenticalIgnoreTags(v, t types.Object) bool
// or
// IdenticalIgnoreTags((code, v, t string) bool
// 忽略标签，然后判断结构体类型是否相等
func TestIdentical13(t *testing.T) {
	code := `
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
`
	s := NewSpec(code)
	if !s.IdenticalIgnoreTags("U", "V") {
		t.Error(`test rule failed`)
	}
	if !IdenticalIgnoreTags(s.MustGetValidType("U"), s.MustGetValidType("V")) {
		t.Error(`test rule failed`)
	}
	if !IdenticalIgnoreTags(s.MustGetValidTypeObject("U"), s.MustGetValidTypeObject("V")) {
		t.Error(`test rule failed`)
	}
	if !IdenticalIgnoreTags(code, "U", "V") {
		t.Error(`test rule failed`)
	}
}
