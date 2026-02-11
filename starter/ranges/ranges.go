package ranges

import "fmt"

func Ranges() {
	fmt.Println("--- Ranges")

	// slice
	sli := []int{1, 2}
	for k, v := range sli {
		fmt.Println("--- Range slice sli", k, v)
	}
	
	// map
	o := map[string]int{"index0": 0, "index1": 1}
	for k, v := range o {
		fmt.Println("--- Range map o", k, v)
	}

	// 数值
	for range 2 {
		fmt.Println("--- Range num 2")
	}

	// 字符串
	for k, v := range "kv" {
		fmt.Println("--- Range string/range", k, v)
		// k -> 6b[hsex] -> 107[dec]
		// v -> 76[hsex] -> 118[dec]
	}
}
