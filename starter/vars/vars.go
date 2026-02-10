package vars

import "fmt"

var name string = "PI"

func Vars()  {
	val := Value()
	domain, symbol := "mathematical", "π"

	fmt.Println("name", name)
	fmt.Println("value", val)
	fmt.Println("domain", domain)
	fmt.Println("symbol", symbol)
}
