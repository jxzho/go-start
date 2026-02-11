---
定义: |
  在 Go 里，指针就是内存地址。
  你只需要记住一个核心原则：想修改原件，就传指针；只读不改，就传值。
使用:
  - "& (取地址)：我想知道变量 a 在内存哪儿？ -> &a"
  - "* (解引用)：拿到地址 p，我想看这里面存的值 -> *p"

go_by_example: https://gobyexample.com/pointers
---

```go
package main

import "fmt"

// 1. 值传递 (Value)
// 这里的 age 只是原件的一份复印件
func changeAgeVal(age int) {
    age = 18 // 修改复印件，原件不受影响
}

// 2. 指针传递 (Pointer)
// 这里的 age 是原件的地址 (钥匙)
func changeAgePtr(age *int) {
    *age = 18 // 通过地址找到原件，直接修改！
}

func main() {
    x := 10
    
    changeAgeVal(x)
    fmt.Println("Val:", x) // 还是 10

    changeAgePtr(&x) // 传地址进去
    fmt.Println("Ptr:", x) // 变成 18 了！
}
