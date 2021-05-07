package main

import "fmt"

func main() {

	var (
		count int    = 0
		name  string = "test"
		// initialization can skip explicit data type declaration
		left, right, up, down = 0, 1, 2, 3
	)
	// short assignment statement := can be used instead of var
	//short assignment := can be used only inside the function
	max := 10

	fmt.Println("Hello Golangie")
	fmt.Println(count, name, left, right, up, down, max)
}
