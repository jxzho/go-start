package pointers

import (
	"fmt"
	"os"
)

func changeValZero(val int) {
	val = 0
}

func changeValZeroPtr(val *int) {
	*val = 0
}

func Ptr() {
	fmt.Println("--- Pointers")

	val := 1

	changeValZero(val)
	if val == 1 {
		fmt.Println("val not changed", val)
	}

	changeValZeroPtr(&val)
	if val == 0 {
		fmt.Println("val changed:", val)
	}

	// mock Println
	fmt.Fprintln(os.Stdout, "val is:", val)
}
