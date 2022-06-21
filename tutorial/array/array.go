package main

import "fmt"

func main() {
	evens := [4]int{1, 2, 3, 4}
	var array [4]int

	array = evens

	fmt.Println(evens)
	fmt.Println(array)
}
