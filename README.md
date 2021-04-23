[![Go Report Card](https://goreportcard.com/badge/github.com/AlaxLee/go-spec-util)](https://goreportcard.com/report/github.com/AlaxLee/go-spec-util)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/569/badge)](https://bestpractices.coreinfrastructure.org/projects/569)

# go-spec-util
some tools on learning go spec


## 0. Go 原生支持的类型

```
Boolean types
Numeric types
String types
Array types
Slice types
Struct types
Pointer types
Function types
Interface types
Map types
Channel types
```

## 1. 类型本身的特性

### 1.1. [identity](https://golang.google.cn/ref/spec#Type_identity)	类型相同

Two types are either *identical* or *different*. 两个类型的关系要么相同要么不同。

A `defined type` is always different from any other type. 一个定义的类型与其他所有类型都不相同。

Otherwise, two types are identical if their `underlying type literals`（注：这里的 underlying type 与 2.1.2. 中的不是一个概念 ） are structurally equivalent; that is, they have the same literal structure and corresponding components have identical types. In detail:

此外，如果两个类型的 “基础类型文字” 在结构上相等，则它们是相同的。 也就是说，它们具有相同的文字结构，并且相应的组件具有相同的类型。具体如下：

- Two array types are identical if they have identical element types and the same array length.
- Two slice types are identical if they have identical element types.
- Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. Non-exported field names from different packages are always different.
- Two pointer types are identical if they have identical base types.
- Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match.
- Two interface types are identical if they have the same set of methods with the same names and identical function types. Non-exported method names from different packages are always different. The order of the methods is irrelevant.
- Two map types are identical if they have identical key and element types.
- Two channel types are identical if they have identical element types and the same direction.

注意：上面提到的 underlying type 与 2.1.2. 中描述的应该不一样，2.1.2. 里面的是追溯到底，例如，对于  `type B int; type A1 B` 类型 A1 的 `underlying type` 是 int，而它的 `underlying type literals` 是 A1。



综上所述，

1. 如果类型是通过 `Type A = B` 的方式声明的，那么 A 和 B 类型相同。
2. 如果类型是通过 `Type A B` 的方式声明的，那么 A 和 B 类型不同。
3. 其它情况下，如果两个类型的字面定义在结构上相等，那么他们两个类型相同。具体如下：
    1. 数组：如果 `元素类型` 和 `数组长度` 都相同，那么类型相同
    2. 切片：如果 `元素类型` 相同，那么类型相同
    3. 结构体： 如果 `属性顺序` 相同，且 `对应属性的名字、类型、标签` 都相同，那么类型相同。注意：**不同包里面的结构体的未导出的属性一定不相同**（这个见下方额外的例子，简单来说，如果在不同的包里的两个结构体满足前述条件，那么，只有当它们所有的属性都是导出属性的时候，它们俩类型才相同，只要有未导出的属性，那么它们俩类型一定不相同）。
    4. 指针：如果 `基本类型（base type）` 相同，那么类型相同
    5. 函数：如果两者具有 `相同数量的参数和返回值`，`相应的参数和返回值的类型相同`，并且要么两个函数都有可变参数，要么都没有。 参数名和结果名不需要匹配。
    6. 接口：如果两者的方法集内的 `方法的名称、类型都相同` ，那么类型相同。 **来自不同程序包的未导出方法名称始终是不同的**（这个的意思跟上述结构体中描述的类似，具体见下方额外例子）。 方法的顺序无关紧要。
    7. 字典：如果两者的 `键和值的类型都相同`，那么类型相同
    8. 管道：如果两者的 `元素类型相同`、`方向相同`，那么类型相同
    9. （待确认）basic type：诸如 int、string 这种简单类型，字面量相同，那么类型相同

代码验证如下：

```go
package main

import (
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
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
```

#### 问题：现在我已经可以静态分析源文件里的类型，并判断类型是否相等了。那么，在执行过程中，如何判断两个类型是否相等？

答：例如两个对象 a 和 b，可以考虑比较 `reflect.ValueOf(a).Type().String()` 和 `reflect.ValueOf(b).Type().String() `

#### 1.1.1. [Composite types](https://golang.google.cn/ref/spec#Types)	复合类型

Composite types —  **array, struct, pointer, function, interface, slice, map, channel** types  — may be constructed using **type literals**.

我们注意到上述 类型相同 里面比较复杂的判断都是针对复合类型的。

#### 1.1.2  Basic Type

定义来自 src/go/types/type.go 里的

```go
type Basic struct {
	kind BasicKind
	info BasicInfo
	name string
}
```

它包括如下类型：

```go
// BasicKind describes the kind of basic type.
type BasicKind int

const (
	Invalid BasicKind = iota // type is invalid

	// predeclared types
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	UnsafePointer

	// types for untyped values
	UntypedBool
	UntypedInt
	UntypedRune
	UntypedFloat
	UntypedComplex
	UntypedString
	UntypedNil

	// aliases
	Byte = Uint8
	Rune = Int32
)
```



#### 1.1.3. [defined type](https://golang.google.cn/ref/spec#Type_definitions)			定义的类型

A type definition creates a new, distinct type with the same underlying type and operations as the given type, and binds an identifier to it. 类型定义使用与给定类型相同的基础类型和操作创建一个新的独特类型，并将标识符绑定到该类型。

The new type is called a **defined type**. It is different from any other type, including the type it is created from. 新类型称为“定义的类型”。 它不同于任何其他类型，包括创建它的类型。

朴素的来讲，这个指的是使用 `type A B` 方式来定义的类型中，A 就被叫做 defined type，A 和 B 是不同的类型。与之对应的是使用 `type A = B` 方式来定义类型，A 被叫做 alias ，A 和 B 是相同的类型。

注意：实际使用过程中的 defined type 还额外包括 1.1.2. 里的 basic type，即解析为类型 \*types.Named 和 \*types.Basic 的都算是 defined type，见源代码 src/go/types/operand.go 的 (\*operand). assignableTo 方法的如下这一段：

```go
	// x's type V and T have identical underlying types
	// and at least one of V or T is not a named type
	if check.identical(Vu, Tu) && (!isNamed(V) || !isNamed(T)) {
		return true
	}
```

示例代码如下：

```go
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
```



## 2. 类型的值（对象）的三种特性

### 2.1. [Assignability](https://golang.google.cn/ref/spec#Assignability)		可赋值性

Assignability governs which pairs of types may appear on the left- and right-hand side of an assignment, including implicit assignments such as function calls, map and channel operations, and so on.

可赋值性控制着哪些类型可能出现在赋值的左侧和右侧，包括隐式分配，例如函数调用，字典和管道操作等。

这个“可赋值性”，意思是，“类型为V的一个值x，可以赋值给类型为T的变量”，注意，这里的 x 是值，即 1 或 struct{}{} 等，x 不是变量，切记。
A value x is assignable to a variable of type T ("x is assignable to T") if one of the following conditions applies: 即，只要满足如下条件之一即可。

1. x's type is identical to T.
2. x's type V and T have identical underlying types and at least one of V or T is not a defined type.
3. T is an interface type and x implements T.
4. x is a bidirectional channel value, T is a channel type, x's type V and T have identical element types, and at least one of V or T is not a defined type.
5. x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
6. x is an untyped constant representable by a value of type T.

综上所述，满足以下任意一个条件，就可以说 “值 x 对于 类型 T 是可赋值的” ：

1. x 的类型 与 T 相同
2. x 的类型 V 和 T 有相同的 underlying type 并且 V 或 T 至少有一个是未（显示）定义类型
3. T 是一个接口，x 实现了 T
4. x 是一个双向管道的值，T 是一个管道类型，x 的类型 V 和 T 有相同的元素类型，并且 V 或 T 至少有一个是未（显示）定义类型
5. x 是 nil，T 是一个 指针、函数、切片、字典、管道 或 接口
6. x 是一个未显示定义的常量，且是个可以被 T 代表的值（可代表性representable见2.1.3.）

```go
package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
	"go/types"
)

func main() {
	assignExample01()
	assignExample02()
	assignExample03()
	assignExample04()
	assignExample05()
	assignExample06()
}

func assignExample01() {
	//1. x's type is identical to T.
	// x 的类型 与 T 相同
	s := gospec.NewSpec(`
type T = int
var x = 1
`)
	if s.Identical("x", "T") {
		fmt.Println("x's type is identical to T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x's type is identical to T
	x could assignable to T
	*/
}

func assignExample02() {
	//2. x's type V and T have identical underlying types
	// and at least one of V or T is not a defined type.
	// x 的类型 V 和 T 有相同的 underlying type 并且 V 或 T 至少有一个是未（显示）定义类型
	s := gospec.NewSpec(`
type V = func()
type T func()
var x V
`)
	V := s.GetType("x")
	T := s.GetType("T")
	if gospec.Identical(V.Underlying(), T.Underlying()) {
		fmt.Printf("x's type V and T have identical underlying types: %s\n", V.Underlying())
	}
	if !gospec.IsDefinedType(V) {
		fmt.Println("V is not a defined type")
	}
	if gospec.IsDefinedType(T) {
		fmt.Println("T is a defined type")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x's type V and T have identical underlying types: func()
	V is not a defined type
	T is a defined type
	x could assignable to T
	*/
}

func assignExample03() {
	//3. T is an interface type and x implements T.
	// T 是一个接口，x 实现了 T
	s := gospec.NewSpec(`
type T interface {
	m()
}
type V struct{}
func (v V) m() {}
var x V
`)
	if _, ok := gospec.IsInterface(s.GetType("T")); ok {
		fmt.Println("T is an interface type")
	}
	if s.Implements("x", "T") {
		fmt.Println("x implements T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	T is an interface type
	x implements T
	x could assignable to T
	*/
}

func assignExample04() {
	//4. x is a bidirectional channel value, T is a channel type,
	// x's type V and T have identical element types,
	// and at least one of V or T is not a defined type.
	// x 是一个双向管道的值，T 是一个管道类型
	// x 的类型 V 和 T 有相同的元素类型，并且 V 或 T 至少有一个是未（显示）定义类型
	s := gospec.NewSpec(`
type T chan int
var x = make(chan int)
`)
	V := s.GetType("x")
	T := s.GetType("T")
	vc, ok := gospec.IsChan(V)
	if ok && vc.Dir() == types.SendRecv {
		fmt.Println("x is a bidirectional channel value")
	}
	tc, ok := gospec.IsChan(T)
	if ok {
		fmt.Println("T is a channel type")
	}
	if gospec.Identical(vc.Elem(), tc.Elem()) {
		fmt.Println("x's type V and T have identical element types")
	}
	if !gospec.IsDefinedType(V) {
		fmt.Println("V is not defined type")
	}
	if gospec.IsDefinedType(T) {
		fmt.Println("T is defined type")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x is a bidirectional channel value
	T is a channel type
	x's type V and T have identical element types
	V is not defined type
	T is defined type
	x could assignable to T
	*/
}

func assignExample05() {

	//5. x is the predeclared identifier nil and T is a pointer, function, slice, map, channel, or interface type.
	// x 是 nil，T 是一个 指针、函数、切片、字典、管道 或 接口
	type typeInfo struct {
		code string
		f    func(types.Type) bool
	}
	typeNames := []string{"pointer", "function", "slice", "map", "channel", "interface"}
	typeMap := map[string]typeInfo{
		"pointer": {"type T *int", func(t types.Type) bool {
			_, ok := gospec.IsPointer(t)
			return ok
		}},
		"function": {"type T func()", func(t types.Type) bool {
			_, ok := gospec.IsFunction(t)
			return ok
		}},
		"slice": {"type T []string", func(t types.Type) bool {
			_, ok := gospec.IsSlice(t)
			return ok
		}},
		"map": {"type T map[string]int", func(t types.Type) bool {
			_, ok := gospec.IsMap(t)
			return ok
		}},
		"channel": {"type T chan int", func(t types.Type) bool {
			_, ok := gospec.IsChan(t)
			return ok
		}},
		"interface": {"type T interface{}", func(t types.Type) bool {
			_, ok := gospec.IsInterface(t)
			return ok
		}},
	}
	for _, v := range typeNames {
		ti, ok := typeMap[v]
		if !ok {
			continue
		}
		s := gospec.NewSpec(ti.code)
		if s.IsInUniverse("nil") {
			fmt.Print("x is the predeclared identifier nil and ")
		}
		if ti.f(s.GetType("T")) {
			fmt.Println("T is a " + v)
		}
		if s.Assignment("nil", "T") {
			fmt.Println("x could assignable to T")
		}
	}
	/* the output is:
	x is the predeclared identifier nil and T is a pointer
	x could assignable to T
	x is the predeclared identifier nil and T is a function
	x could assignable to T
	x is the predeclared identifier nil and T is a slice
	x could assignable to T
	x is the predeclared identifier nil and T is a map
	x could assignable to T
	x is the predeclared identifier nil and T is a channel
	x could assignable to T
	x is the predeclared identifier nil and T is a interface
	x could assignable to T
	*/
}

func assignExample06() {

	//6. x is an untyped constant representable by a value of type T.
	// x 是一个未显示定义的常量，且是个可以被 T 代表的值
	// 具体见下方的 Representability 可被代表性
	s := gospec.NewSpec(`
type T int
const x = 1
`)
	x := s.GetType("x")
	if gospec.IsUntyped(x) && gospec.IsConstType(x) {
		fmt.Print("x is an untyped constant ")
	}
	if s.Representable("x", "T") {
		fmt.Println("representable by a value of type T")
	}
	if s.Assignment("x", "T") {
		fmt.Println("x could assignable to T")
	}
	/* the output is:
	x is an untyped constant representable by a value of type T
	x could assignable to T
	*/
}
```

#### 问题，golang 在对代码做语法检查的时候，使用的代码与typechecker检查语法的代码时一样的吗？是不是可以从 go build 的逻辑里面看出来？

注：src/go/types/operand.go 里的 operand 类型的 assignableTo 方法，清晰的描述了上述条件的判断过程，看源码很好。assignableTo 针对上述第二条，检查的是

注：src/go/types/api.go 里面的 AssignableTo(V, T)方法 与实际情况有所不同，所以建议以实际验证结果为准，而非 AssignableTo 方法结果为准。例如：

注：在使用typechecker过程中，src/go/types/assignments.go 里面 assignment 方法使用了上方 operand.go 里的 assignableTo 方法

#### 2.1.1. [base type](https://golang.google.cn/ref/spec#Pointer_types)（来自官网spec文档）		基本类型

A pointer type denotes the set of all pointers to variables of a given type, called the **base type** of the pointer.

```
PointerType = "*" BaseType .
BaseType    = Type .
```

例如： *[4]int 的 base type 是 [4]int

```go
package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
)

func main() {
	fmt.Println(gospec.GetBaseType(`var la *[4]int`, "la"))
	/* 执行结果是
	[4]int
	*/
}
```



#### 2.1.2. [underlying type](https://golang.google.cn/ref/spec#Types)（来自官网spec文档）		基础类型

Each type T has an **underlying type**: If T is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is T itself. Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.
每个类型T都有一个基础类型：如果T是预先声明的布尔类型，数字类型或字符串类型或类型字面值之一，则相应的基础类型是T本身。 否则，T的基础类型是T在其类型声明中**引用的类型的基础类型**。

根据我自己的理解，根据上述定义的描述，在寻找一个类型 V 的基础类型时，最终总是能上溯到一个基础类型 W，这个 W 的基础类型就是他本身，即 W 就是 “预先声明的布尔类型，数字类型或字符串类型或类型字面值之一”

```go
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1  // []B1 is a type literal
	B4 B3
)
```
The underlying type of string, A1, A2, B1, and B2 is string. The underlying type of []B1, B3, and B4 is []B1.

问题：B1 的基础类型是 string，那 []B1 的基础类型为何是 []B1 而不是 []string ？

答：我是这么理解的，一个基础类型必然是一个Go原生支持的类型。在这里 B1 显然不是原生类型，所以对其的基础类型要进行追溯，直到得到 string 为止。而对于 []B1 来说，它是一个 slice，是原生类型，它的基础类型就是 []B1，无需继续追溯。

```go
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
	/* 执行结果是
	A1's underlying type is string
	A2's underlying type is string
	B1's underlying type is string
	B2's underlying type is string
	B3's underlying type is []test.B1
	B4's underlying type is []test.B1
	*/
}
```



#### 2.1.3. [Representability](https://golang.google.cn/ref/spec#Representability)		可被代表性

具体实现见代码 src/go/types/expr.go 里的方法 check.representable

这个是为了用于解释 Assignability 的第6条的。

A constant x is representable by a value of type T if one of the following conditions applies:

```
1. x is in the set of values determined by T.
x 是 类型 T 集合内的值（如 true 之于 bool）

2. T is a floating-point type and x can be rounded to T's precision without overflow. Rounding uses IEEE 754 round-to-even rules but with an IEEE negative zero further simplified to an unsigned zero. Note that constant values never result in an IEEE negative zero, NaN, or infinity.
T 是 浮点数类型，x 不超过其范围

3. T is a complex type, and x's components real(x) and imag(x) are representable by values of T's component type (float32 or float64).
T 是 复数类型，x 的 实部和虚部都不超过范围
```

x is representable by a value of T because（即如下 x 都可以被 T 的值代表 ）

```
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
```

x is not representable by a value of T because（即如下 x 都不能被 T 的值代表 ）

```
x                   T           

0                   bool        0 is not in the set of boolean values
'a'                 string      'a' is a rune, it is not in the set of string values
1024                byte        1024 is not in the set of unsigned 8-bit integers
-1                  uint16      -1 is not in the set of unsigned 16-bit integers
1.1                 int         1.1 is not an integer value
42i                 float32     (0 + 42i) is not in the set of float32 values
1e1000              float64     1e1000 overflows to IEEE +Inf after rounding
```

代码如下：

```go
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
```



### 2.2. [Comparability](https://golang.google.cn/ref/spec#Comparison_operators) 	可比较性（代表的是值的相等性）

#### 注意：这一部分应该会涉及到 reflect，因为比较是一个运行时的概念，其中又涉及到类型的问题，所以举例子和做验证的时候要涉及 reflect。同时 reflect 里面的某些概念和类型的某些概念的一一映射后续也要列出来。

可比较性 确定哪些类型可能出现在比较表达式 `x == y` 里 或 switch case 里 或 是否可以用作字典的键。

比较操作符

```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

In any comparison, the first operand must be assignable to the type of the second operand, or vice versa.

在任何比较的场景下，第一个操作数对于第二个操作数的类型来说，必须是可赋值的，反之亦然（即，第二个操作数对于第一个操作数的类型来说，是可赋值的）。

The equality operators `==` and `!=` apply to operands that are ***comparable***. The ordering operators `<`, `<=`, `>`, and `>=` apply to operands that are ***ordered***. These terms and the result of the comparisons are defined as follows:

- Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
- Integer values are comparable and ordered, in the usual way.
- Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
- Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
- String values are comparable and ordered, lexically byte-wise（逐字节地）.
- Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value `nil`. Pointers to distinct `zero-size` variables may or may not be equal.
- Channel values are comparable. Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
- Interface values are comparable. Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
- A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable when values of type `X` are comparable and `X` implements `T`. They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
- Struct values are comparable if all their fields are comparable. Two struct values are equal if their corresponding non-blank fields are equal.
- Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.
- Slice, map, and function values are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.

综上所述，

**首先，第一个操作数对第二个操作数的类型是可赋值的 或 第二个操作数对第一个操作数的类型是可赋值的**，然后，按下方规则进行判断：

1. 布尔值：如果两者都是 true 或 false，那么它们相等
2. 整型：相等就相等
3. 浮点型：相等就相等
4. 复数：如果两者的 实部 和 虚部 分别都相等，那么它们相等
5. 字符串：如果两者每个字节都相等，那么它们相等
6. 指针：如果两者指向同一个变量，或 值都是 nil，那么它们相等。注意：指向不同的零值变量，可能相等也可能不等（注：这个没找到相等的例子啊）
7. 管道：如果两者被同一个 make 调用创建出来，或 值都是 nil，那么它们相等
8. 接口：如果两者的动态类型相同动态值相等，或 值都是 nil，那么它们相等（问题：如果一个接口的值，存的是另一个接口定义的变量，那么这里的动态类型指什么？要追溯到底么？）
9. 接口与非接口：如果非接口类型实现了接口，则它们可比较。如果接口的动态类型与非接口类型相同，接口的动态值和非接口类型值相等，那么它们相等。
10. 结构体：如果两者所有的属性都是可比较的，则它们可比较。如果两者对应的 non-blank 属性都相等，那么它们相等。
11. 数组：如果元素类型是可比较的，则它们可比较。如果两者对应的元素都相等，那么它们相等。
12. 切片：它们不可比较。即，两个切片之间不可比较。
13. 字典：它们不可比较。即，两个字典之间不可比较。
14. 函数：它们不可比较。即，两个函数之间不可比较。

综合来看，上述 2、3、5 是可排序的，1 ~ 8 是无条件可比较的，9、10、11 是有条件可比较的，12、13、14 是不可比较的。

另外， 指针、管道、接口、切片、字典、函数，它们6个的零值是 nil，所以还是可以单方面和 nil 作比较。

代码如下：

```go
package main

import (
	"fmt"
	gospec "github.com/AlaxLee/go-spec-util"
	"go/types"
)

func main() {
	comparableExample01()
	comparableExample02()
	comparableExample03()
	comparableExample04()
	comparableExample05()
	comparableExample06()
	comparableExample07()
	comparableExample08()
	comparableExample09()
	comparableExample10()
	comparableExample11()
	comparableExample12()
}

func comparableExample01() {
	// Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
	s := gospec.NewSpec(`
var a bool = true
var b bool = true
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsBoolean(ta) && gospec.IsBoolean(tb) && gospec.Comparable(ta) && gospec.Comparable(tb) {
		fmt.Print("Boolean values are comparable. ")
	}
	var a bool = true
	var b bool = true
	if (a && b) || (!a && !b) {
		if a == b {
			fmt.Println("Two boolean values are equal if they are either both `true` or both `false`.")
		}
	}
	/* the output is:
	Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
	*/
}

func comparableExample02() {
	// Integer values are comparable and ordered, in the usual way.
	s := gospec.NewSpec(`
var a int = 1
var b int = 1
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsInteger(ta) && gospec.IsOrdered(ta) && gospec.Comparable(ta) &&
		gospec.IsInteger(tb) && gospec.IsOrdered(tb) && gospec.Comparable(tb) {
		var a int = 1
		var b int = 1
		if a == b {
			fmt.Println("Integer values are comparable and ordered, in the usual way.")
		}
	}
	/* the output is:
	Integer values are comparable and ordered, in the usual way.
	*/
}

func comparableExample03() {
	// Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
	s := gospec.NewSpec(`
var a float64 = 1.0
var b float64 = 1.0
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsFloat(ta) && gospec.IsOrdered(ta) && gospec.Comparable(ta) &&
		gospec.IsFloat(tb) && gospec.IsOrdered(tb) && gospec.Comparable(tb) {
		var a float64 = 1.0
		var b float64 = 1.0
		if a == b {
			fmt.Println("Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.")
		}
	}
	/* the output is:
	Floating-point values are comparable and ordered, as defined by the IEEE-754 standard.
	*/
}

func comparableExample04() {
	// Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
	s := gospec.NewSpec(`
var u complex64 = 1 + 2i
var v complex64 = 1 + 2i
`)
	tu := s.GetType("u")
	tv := s.GetType("v")
	if gospec.IsComplex(tu) && gospec.Comparable(tu) &&
		gospec.IsComplex(tv) && gospec.Comparable(tv) {
		fmt.Print("Complex values are comparable. ")
		var u complex64 = 1 + 2i
		var v complex64 = 1 + 2i
		if real(u) == real(v) && imag(u) == imag(v) {
			if u == v {
				fmt.Println("Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.")
			}
		}
	}
	/* the output is:
	Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
	*/
}

func comparableExample05() {
	// String values are comparable and ordered, lexically byte-wise（逐字节地）.
	s := gospec.NewSpec(`
var a string = "haha"
var b string = "haha"
`)
	ta := s.GetType("a")
	tb := s.GetType("b")
	if gospec.IsString(ta) && gospec.Comparable(ta) &&
		gospec.IsString(tb) && gospec.Comparable(tb) {
		var a string = "haha"
		var b string = "haha"
		if a == b {
			fmt.Println("String values are comparable and ordered")
		}
	}
	/* the output is:
	String values are comparable and ordered
	*/
}

func comparableExample06() {
	// Pointer values are comparable.
	// Two pointer values are equal if they point to the same variable or if both have value `nil`.
	// Pointers to distinct `zero-size` variables may or may not be equal.
	s := gospec.NewSpec(`
var a *int
`)
	ta := s.GetType("a")
	if _, isPointer := gospec.IsPointer(ta); isPointer && gospec.Comparable(ta) {
		fmt.Println("Pointer values are comparable.")
	}
	var a *int
	var b *int
	var c int = 1
	var u *int = &c
	var v *int = &c
	if a == nil && b == nil && a == b && u == v {
		fmt.Println("Two pointer values are equal if they point to the same variable or if both have value `nil`.")
	}
	/* the output is:
	Pointer values are comparable.
	Two pointer values are equal if they point to the same variable or if both have value `nil`.
	*/
}

func comparableExample07() {
	// Channel values are comparable.
	// Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
	s := gospec.NewSpec(`
var a chan int
`)
	ta := s.GetType("a")
	if _, isChan := gospec.IsChan(ta); isChan && gospec.Comparable(ta) {
		fmt.Println("Channel values are comparable.")
	}
	var a chan int
	var b chan int
	c := make(chan int, 1)
	var u chan int = c
	var v chan int = c
	if a == nil && b == nil && a == b && u == v {
		fmt.Println("Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.")
	}
	/* the output is:
	Channel values are comparable.
	Two channel values are equal if they were created by the same call to `make` or if both have value `nil`.
	*/
}

func comparableExample08() {
	// Interface values are comparable.
	// Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
	s := gospec.NewSpec(`
	var a interface{}
`)
	ta := s.GetType("a")
	if _, isInterface := gospec.IsInterface(ta); isInterface && gospec.Comparable(ta) {
		fmt.Println("Interface values are comparable.")
	}

	var a interface{}
	var b interface{}
	var c = [...]int{1, 2, 3}
	var d = [...]int{1, 2, 3}
	var u interface{} = c
	var v interface{} = d
	if a == nil && b == nil && a == b &&
		gospec.GetDynamicTypeAtRuntime(u) == gospec.GetDynamicTypeAtRuntime(v) && c == d {
		fmt.Println("Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.")
	}
	/* the output is:
	Interface values are comparable.
	Two interface values are equal if they have **identical** `dynamic types` and equal `dynamic values` or if both have value `nil`.
	*/
}

func comparableExample09() {
	// A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
	// when values of type `X` are comparable and `X` implements `T`.
	// They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
	s := gospec.NewSpec(`
	type X struct {}
	type T interface {}
	var x X
	var t T
`)
	if s.Comparable("X") && s.Implements("X", "T") {
		fmt.Println("A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable ")
		fmt.Println("when values of type `X` are comparable and `X` implements `T`.")
	}
	type X struct{}
	type T interface{}
	var x X
	var y X
	var t T
	t = y
	if gospec.GetDynamicTypeAtRuntime(t) == gospec.GetDynamicTypeAtRuntime(x) && y == x {
		fmt.Println("They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.")
	}
	/* the output is:
	A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable
	when values of type `X` are comparable and `X` implements `T`.
	They are equal if `t`'s dynamic type is **identical** to `X` and `t`'s dynamic value is equal to `x`.
	*/
}

func comparableExample10() {
	// Struct values are comparable if all their fields are comparable.
	// Two struct values are equal if their corresponding non-blank fields are equal.
	s := gospec.NewSpec(`
	type U struct {
		a int
		b []int
	}
	type V struct {
		a complex64
		b string
	}
`)

	u, ok := s.GetType("U").Underlying().(*types.Struct)
	if !ok {
		panic("unexpect")
	}
	allUFieldComparable := true
	for i := 0; i < u.NumFields(); i++ {
		if !gospec.Comparable(u.Field(i).Type()) {
			allUFieldComparable = false
		}
	}

	v, ok := s.GetType("V").Underlying().(*types.Struct)
	if !ok {
		panic("unexpect")
	}
	allVFieldComparable := true
	for i := 0; i < v.NumFields(); i++ {
		if !gospec.Comparable(v.Field(i).Type()) {
			allVFieldComparable = false
		}
	}

	if !allUFieldComparable && !s.Comparable("U") {
		if allVFieldComparable && s.Comparable("V") {
			fmt.Println("Struct values are comparable if all their fields are comparable.")
		}
	}

	type V struct {
		a complex64
		b string
	}

	m := V{1i, "lala"}
	n := V{1i, "lala"}
	if m.a == n.a && m.b == n.b && m == n {
		fmt.Println("Two struct values are equal if their corresponding non-blank fields are equal.")
	}
	/* the output is:
	Struct values are comparable if all their fields are comparable.
	Two struct values are equal if their corresponding non-blank fields are equal.
	*/
}

func comparableExample11() {
	// Array values are comparable if values of the array element type are comparable.
	// Two array values are equal if their corresponding elements are equal.
	s := gospec.NewSpec(`
	var a [3]interface{}
	var b [3][]int
`)
	if s.Comparable("a") && !s.Comparable("b") {
		fmt.Println("Array values are comparable if values of the array element type are comparable.")
	}

	var m [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	var n [3]interface{} = [3]interface{}{1, "lala", [3]int{4, 5, 6}}
	if m[0] == n[0] && m[1] == n[1] && m[2] == n[2] && m == n {
		fmt.Println("Two array values are equal if their corresponding elements are equal.")
	}
	/* the output is:
	Array values are comparable if values of the array element type are comparable.
	Two array values are equal if their corresponding elements are equal.
	*/
}

func comparableExample12() {
	// Slice, map, and function values are not comparable.
	// However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
	s := gospec.NewSpec(`
	var a []int
	var b map[int]string
	var c func(string)int
`)
	if !s.Comparable("a") && !s.Comparable("b") && !s.Comparable("c") {
		fmt.Println("Slice, map, and function values are not comparable.")
	}
	var a []int
	var b map[int]string
	var c func(string) int
	if a == nil && b == nil && c == nil {
		fmt.Println("However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.")
	}
	/*  the output is:
	Slice, map, and function values are not comparable.
	However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`.
	*/
}
```



#### 2.2.1. [variable](https://golang.google.cn/ref/spec#Variables)			变量

A variable is a storage location for holding a value.

#### 2.2.2. [structured variables](https://golang.google.cn/ref/spec#Variables)	结构化变量

Structured variables of array, slice, and struct types have elements and fields that may be addressed individually. Each such element acts like a `variable`.

数组、切片和结构体类型的结构化变量，都有可以单独取地址的元素和属性。这些元素和属性表现的就像一个变量一样。

#### 2.2.3. [static type](https://golang.google.cn/ref/spec#Variables)		静态类型

The **static type** (or just type) of a variable is the type given in its declaration, the type provided in the new call or composite literal, or the type of an element of a `structured variable`.

变量的静态类型是：类型声明中给出的类型，新调用或复合字面量提供的类型，或结构化变量的元素的类型。

#### 2.2.4. [dynamic type](https://golang.google.cn/ref/spec#Variables)	动态类型

Variables of interface type also have a distinct **dynamic type**, which is the concrete type of the value assigned to the variable **at run time** (unless the value is the predeclared identifier nil, which has no type). The dynamic type may vary during execution but values stored in interface variables are always assignable to the `static type` of the variable.

接口类型的变量有 动态类型，它是在运行时分配给变量的值的具体类型（除了无类型的nil以外）。动态类型在执行过程中可能有不同，但接口存储的值对该值的静态类型来说始终是可赋值的。

```
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```

##### 问题：如果一个接口的值，存的是另一个接口定义的变量，那么这里的动态类型是什么？要追溯到底么？

### 2.3. [Convertibility](https://golang.org/ref/spec#Conversions)		可转换性

A conversion changes the [type](https://golang.google.cn/ref/spec#Types) of an expression to the type specified by the conversion. A conversion may appear literally in the source, or it may be *implied* by the context in which an expression appears.

转换指的是将表达式的类型更改为指定的类型。 转换可能从字面上出现在源代码中，也可能被表达式所显示的上下文“暗含”。即是所谓的“显示转换”与“隐式转换”

#### 2.3.1. 显示转换

An *explicit* conversion is an expression of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`.

显式转换是形式为T(x)的表达式，其中T是类型，而x是可以转换为类型T的表达式。

A [constant](https://golang.google.cn/ref/spec#Constants) value `x` can be converted to type `T` if `x` is [representable](https://golang.google.cn/ref/spec#Representability) by a value of `T`. As a special case, an integer constant `x` can be explicitly converted to a [string type](https://golang.google.cn/ref/spec#String_types) using the [same rule](https://golang.google.cn/ref/spec#Conversions_to_and_from_a_string_type) as for non-constant `x`.

如果x可以用T的值表示，则可以将常数x转换为T类型。在特殊情况下，可以使用与非常数x相同的规则将整数常数x显式转换为字符串类型。

```
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
MyString("foo" + "bar")  // "foobar" of type MyString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
```

A non-constant value `x` can be converted to type `T` in any of these cases:

- `x` is [assignable](https://golang.google.cn/ref/spec#Assignability) to `T`.
- ignoring struct tags (see below), `x`'s type and `T` have [identical](https://golang.google.cn/ref/spec#Type_identity) [underlying types](https://golang.google.cn/ref/spec#Types).
- ignoring struct tags (see below), `x`'s type and `T` are pointer types that are not [defined types](https://golang.google.cn/ref/spec#Type_definitions), and their pointer base types have identical underlying types.
- `x`'s type and `T` are both integer or floating point types.
- `x`'s type and `T` are both complex types.
- `x` is an integer or a slice of bytes or runes and `T` is a string type.
- `x` is a string and `T` is a slice of bytes or runes.

一个非常数值 x 可以被转化为类型 T，只要满足以下任意一个条件即可：

1. x 可以赋值给 T
2. 忽略掉 struct 的 tag，x 的类型 与 T 有相同的基础类型
3. 忽略掉 struct 的 tag，x 的类型 与 T 是指针类型，且不是定义的类型，并且他们指向的基本类型有相同的基础类型
4. x 的类型 和 T 都是 整数 或 浮点数 类型
5. x 的类型 和 T 都是 复数 类型
6. x 是一个 整数 或 是一个 rune的切片，T 是一个 字符串 类型
7. x 是一个 字符串，T 是一个 byte的切片 或 是一个 rune的切片

#### 2.3.2. 隐式转换

一般来说最常见的

## 3. 其它

常用的方式有两种：

1. 比较

   这个涉及到可比较性，其实底层也与变量具体类型相关

2. 转化

   这个要先把基本类型描述清楚

### 问题：汇编里面看到的类型是如何描述的，如何存放的

```go
package main

import "fmt"

type Lala interface{}

var i Lala

func main() {
	fmt.Println(i)
}
```

这下面存的都是啥？

```
type.*"".Lala SRODATA size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 ee 10 1a be 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.algarray+80
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*main.Lala.+0
        rel 48+8 t=1 type."".Lala+0
type."".Lala SRODATA size=96
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 aa 36 31 4c 07 08 08 14 00 00 00 00 00 00 00 00  .61L............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0050 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.algarray+144
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*main.Lala.+0
        rel 44+4 t=5 type.*"".Lala+0
        rel 48+8 t=1 type..importpath."".+0
        rel 56+8 t=1 type."".Lala+96
        rel 80+4 t=5 type..importpath."".+0
```


