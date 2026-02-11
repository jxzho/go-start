package maps

import (
	"fmt"
	"maps"
)

func Maps() {
	o := make(map[string]string)
	o["name"] = "maps"
	o["same1"] = "hashes"
	o["same2"] = "dicts"

	// clear(o)

	o1 := map[string]string{"name": "maps", "same1": "hashes", "same2": "dicts"}

	fmt.Println("--- Maps o1 sizes:", len(o1))
	fmt.Println("--- Maps \no equal o1:", o, maps.Equal(o, o1))

	// delete(o, "same2")
	if same3, ok := o["same2"]; ok {
		fmt.Println("--- Maps \nsame2:", same3)
	} else {
		fmt.Println("--- Maps same2 is empty")
	}

	clear(o)
	fmt.Println("--- Maps \no sizes:", len(o))
}
