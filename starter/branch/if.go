package branch

import "fmt"

func useIf() {
	if true {
		fmt.Println("print if")
	} else {
		fmt.Println("else")
	}
}

func Branch()  {
	Switch()
	useIf()
}
