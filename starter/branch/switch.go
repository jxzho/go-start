package branch

import "fmt"

func Switch()  {
	fmt.Println("--- exec func Switch")
	val := 1
	switch val {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
	fmt.Println("--- exec func Switch")
}
