package main

import "fmt"

func main() {
	fmt.Println("Hello from Go WebAssembly!")
}

//export add
func add(x int, y int) int {
	return x + y
}
