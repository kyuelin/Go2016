package main

import "fmt"

func main() {
	presAge := make(map[string] int)
	presAge["Ken"] = 42
	presAge["Obama"] = 50
	fmt.Println(presAge)
	presAge["Obama"]++
	fmt.Println(len(presAge))
	presAge["Bush"] = 71
	fmt.Println(len(presAge))
	delete(presAge, "Obama")
	fmt.Println(presAge)

}
