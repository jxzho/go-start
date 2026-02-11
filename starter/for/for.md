---
定义: 忘掉 while，忘掉 do-while。Go 只有 for，一种关键字打天下。
语法_标准版: init; condition; post (类似 C/Java)。
语法_While 版: 只保留 condition。
语法_死循环版: 什么都不写。
go_by_example: https://gobyexample.com/for
---

```go
package main

import "fmt"

func main() {
    // 1. 标准写法
    for i := 0; i < 3; i++ {
        fmt.Println("Standard:", i)
    }

    // 2. 模拟 While (重点掌握)
    n := 0
    for n < 3 {
        fmt.Println("While-style:", n)
        n++
    }

    // 3. 死循环 (常用于服务器监听)
    // for {
    //     doWork()
    // }
}
