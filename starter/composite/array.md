---
定义: 是值类型。长度固定，不可变。传参时会拷贝整个数组（性能差）。
go_by_example: https://gobyexample.com/arrays
---

```go
package main

import "fmt"

func main() {
    // 1. 创建 Slice
    // 方式 A: 直接字面量 (最常用)
    nums := []int{1, 2, 3} 
    
    // 方式 B: 使用 make(type, len, cap)
    // 创建一个长度为3，容量为5的切片
    data := make([]int, 3, 5) 

    // 2. Append (动态扩容)
    // 注意：append 可能会返回一个新的 slice (如果原容量不够，Go 会换一个更大的底层数组)
    // 所以必须把结果重新赋值回去： nums = append(nums, ...)
    nums = append(nums, 4)
    
    // 3. 切片操作 (左含右不含)
    // 获取 nums 的索引 1 到 2 的元素
    sub := nums[1:3] // 结果 [2, 3]
    
    fmt.Println(nums, sub)
}
