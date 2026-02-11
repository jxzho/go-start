package vars

import "fmt"

var name string = "PI"

func Vars()  {
	domain, symbol := "mathematical", "π"

	fmt.Println("name", name)
	fmt.Println("pi value", piValue)
	fmt.Println("domain", domain)
	fmt.Println("symbol", symbol)
}
