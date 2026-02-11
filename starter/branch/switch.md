---
title: switch 比其它语言更智能
key_points:
  - 默认不需要 break，命中即停止，不会穿透（fallthrough）
  - 支持任意类型匹配（不仅仅是整数）
go_by_example: https://gobyexample.com/switch
---

```go
day := "Mon"
switch day {
case "Mon":
    fmt.Println("Work")
case "Sat", "Sun": // 多个条件
    fmt.Println("Party")
default:
    fmt.Println("Sleep")
}
