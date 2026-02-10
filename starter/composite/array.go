package composite

import "fmt"

var arr = [2]int{1, 2}

func Arr() {
	fmt.Println("--- Arr")

	fmt.Println("array", arr)
	fmt.Println("cap", cap(arr))
	fmt.Println("len", len(arr))

	fmt.Println("--- Arr")

	Slice()
}
