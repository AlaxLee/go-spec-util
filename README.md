[![Go Report Card](https://goreportcard.com/badge/github.com/AlaxLee/go-spec-util)](https://goreportcard.com/report/github.com/AlaxLee/go-spec-util)

# go-spec-util
some tools on learning go spec

# [identity](https://golang.google.cn/ref/spec#Type_identity) 
judge two types if identical
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