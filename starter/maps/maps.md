---
定义: Go 的 Map 是 Key-Value 结构，无序。
核心操作: Comma-ok 模式
go_by_example: https://gobyexample.com/maps
---

```go
package main

import "fmt"

func main() {
    // 1. 初始化
    // 必须用 make 初始化，否则是 nil map，无法写入！
    scores := make(map[string]int)
    scores["Alice"] = 95
    scores["Bob"] = 88

    // 字面量初始化
    // dict := map[string]string{"a": "apple"}

    // 2. 查找 (Key 是否存在？)
    // score 是值，ok 是布尔值
    score, ok := scores["Alice"]
    if ok {
        fmt.Println("Alice's score:", score)
    } else {
        fmt.Println("Alice not found")
    }

    // 3. 删除
    delete(scores, "Bob")
}

// * 避坑指南：
// * Nil Map Panic：var m map[string]int 只是声明了一个 nil map，直接 m["key"] = 1 会导致程序崩溃（Panic）。必须用 make 或字面量初始化。
// * 无序性：千万不要依赖 Map 的遍历顺序，每次跑可能都不一样。
