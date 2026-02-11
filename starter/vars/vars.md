---
定义: 在 Go 中，变量定义讲究“显式”和“类型安全”。

core:
    - 标准声明 (var)：适用于包级别（全局）或需要显式定义零值时。
    - 短变量声明 (:=)：最常用！ 只能在函数内部使用。它让编译器自动推导类型。
    - 常量 (const)：编译期确定的值，不可修改。

tips:
    - 全局变量 -> 即包级变量
    - 大写跨包共享，小写同包内共享

go_by_example: https://gobyexample.com/variables
---

```go
package main

import "fmt"

// 全局变量必须用 var
var version string = "1.0.0" 

const Pi = 3.14159

func main() {
    // 1. 标准写法 (如果你不赋值，int 默认为 0，string 默认为 "")
    var x int 
    x = 10

    // 2. 短变量声明 (推荐)
    // 相当于 var y string = "hello"
    y := "hello, go" 
    
    // 多变量同时赋值
    name, age := "Alice", 25

    fmt.Println(x, y, name, age)
}

// * 避坑指南
// * 未使用报错：Go 极其严格，如果函数内声明了变量却没使用，编译报错。
// * := 左边必须至少有一个新变量。
// * := 不能用于函数外部（全局）。
