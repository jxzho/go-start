package foreach

import "fmt"

func Foreach()  {
	for i := 0; i < 2; i++ {
		fmt.Println("i:", i)
	}

	for v := range 2 {
		fmt.Println("value:", v)
	}

	i := 1
	for i <= 2 {
		fmt.Println("i:", i)
		i = i + 1
	}
}
