package composite

import (
	"fmt"
	"slices"
)

func push(arr []int, val int) []int {
	return append(arr, val)
}

func pop(arr []int) ([]int, bool) {
	if len(arr) == 0 {
		return nil, false
	}

	return arr[:len(arr)-1], true
}

func shift(arr []int) ([]int, bool) {
	if len(arr) == 0 {
		return nil, false
	}

	return arr[1:], true
}

func unShift(arr []int, val int) []int {
	return append([]int{val}, arr...)
}

func slice(arr []int, start, end int) ([]int, bool) {
	if start < 0 || end < start || end > len(arr) {
		return nil, false
	}
	return arr[start:end], true
}

func sort[T any](arr []T, cmp func(a, b T) int) []T {
	rt := slices.Clone(arr)
	slices.SortFunc(rt, cmp)
	return rt
}


func Slice() {
	base := []int{1, 2}
	sli := append(base, 3, 4, 5, 6)
	fmt.Println("value", sli) // [1 2 3 4 5 6]
	fmt.Println("---")

	fmt.Println(
		"push:",
		push(sli, 7),
	)

	if result, ok := pop(sli); ok {
		fmt.Println("pop:", result)
	} else {
		fmt.Println("pop: invalid input")
	}

	if result, ok := shift(sli); ok {
		fmt.Println("shift:", result)
	} else {
		fmt.Println("shift: invalid input")
	}

	fmt.Println(
		"unShift:",
		unShift(sli, 0),
	)

	if result, ok := slice(sli, 1, len(sli)-1); ok {
		fmt.Println("slice:", result)
	} else {
		fmt.Println("slice: invalid range")
	}

	if _, ok := slice(sli, -1, 2); !ok {
		fmt.Println("slice(invalid): invalid range")
	}

	fmt.Println(
		"sorted(desc):",
		sort(sli, func(a, b int) int {
			if a < b {
				return 1
			}
			if a > b {
				return -1
			}
			return 0
		}),
	)
}
