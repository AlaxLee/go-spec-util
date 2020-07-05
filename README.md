[![Go Report Card](https://goreportcard.com/badge/github.com/AlaxLee/go-spec-util)](https://goreportcard.com/report/github.com/AlaxLee/go-spec-util)

# go-spec-util
学习 go spec 过程中的一些工具

# identity
判断两个类型是否相等
```go
// type A1 = B  则 A1 与 B 类型相同

if gospec.Identical(`type B int; type A1 = B`, "A1", "B") {
    fmt.Println("Identical")
} else {
    fmt.Println("not Identical")
}
/* 执行结果是
Identical
*/

// OutputIfIdentical 方法会以一种友善的方式print出来
gospec.OutputIfIdentical(
    `type B int; type A1 = B`,
    "A1", "B")

/* 执行结果是
A1     的类型是 example.B
B      的类型是 example.B
他们类型 相同
*/
```