---
list: ["全局变量", "常量", "短变量声明", "多变量赋值"]
---

```go
package main

import "fmt"

// * 全局变量 -> 即包级变量
// * 大写跨包共享，小写同包内共享

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
