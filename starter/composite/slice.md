---
定义: 引用类型。长度动态。传参时只拷贝“切片头”（性能极高）。99% 的场景都用 Slice。
title: "深入理解 Go Slice"
tags: [Go, Slice, 数据结构]

summary: >
  可以把 Slice 想象成一个“窗口”，
  透过它访问底层数组。

description: "Slice 是对底层数组的一层视图（view），而不是数组本身。"

structure:
  - name: Pointer
    meaning: "指向底层数组的起始位置"
  - name: Length (len)
    meaning: "窗口中当前包含的元素个数"
  - name: Capacity (cap)
    meaning: "从起始位置到底层数组末尾的可用容量"

key_points:
  - "Slice 本身不存储数据，只是对数组的描述"
  - "多个 Slice 可以共享同一个底层数组"
  - "len 表示可访问范围"
  - "cap 表示在不扩容前的最大增长范围"

go_by_example: https://gobyexample.com/slices
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

// * 避坑指南：
// * 修改切片会影响底层数组：如果你从大数组切出一个小切片，修改小切片的值，原数组也会变！
// * 数组 vs 切片定义：
// * - [3]int -> 有数字，是数组（Array）。
// * - []int -> 没数字，是切片（Slice）。
