---
定义: Range 是 Go 里的“万能迭代器”。
Tips: [
    "遍历 Slice",
    "遍历 Map",
    "小技巧：如果你只需要值，不需要索引/Key，用下划线 _ 丢弃",
]
"go_by_example": https://gobyexample.com/range-over-built-in-types
---

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    
    // 遍历 Slice
    // i 是索引，v 是值的拷贝
    for i, v := range nums {
        fmt.Printf("Index: %d, Value: %d\n", i, v)
    }

    // 遍历 Map
    m := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range m {
        fmt.Println(k, v)
    }

    // 小技巧：如果你只需要值，不需要索引/Key，用下划线 _ 丢弃
    for _, v := range nums {
        fmt.Println("Only value:", v)
    }
}

// * 避坑指南：
// * range 返回的 v 是元素的副本。在循环里修改 v 不会改变原切片里的值。如果要修改，请使用 nums[i] = new_val。
