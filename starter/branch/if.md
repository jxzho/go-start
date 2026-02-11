---
核心技巧: "If Init Statement"
语法定义: "在判断条件前先执行语句"
标准范式: "错误处理 (if err := fn(); err != nil)"
作用域: "变量仅在 if 代码块内有效 (Scope Isolation)"
go_by_example: https://gobyexample.com/if-else
---

```go
// 传统写法
result := calculate()
if result > 10 { ... }

// Go 惯用写法 (重点！)
// 变量 n 的作用域仅限于这个 if/else 块内部，更安全
if n := calculate(); n > 10 {
    fmt.Println("Result is big:", n)
} else {
    fmt.Println("Result is small:", n)
}

func makeTea(val int) error {
    if val <= 0 {
        return ErrOutOfTea
    }
    return nil
}
if err := makeTea(i); err != nil {
    if errors.Is(err, ErrOutOfTea) {
        fmt.Println("We should buy new tea!")
    } else if errors.Is(err, ErrPower) {
        fmt.Println("Now it is dark.")
    } else {
        fmt.Printf("unknown error: %s\n", err)
    }
    continue
}
