package main

import "fmt"

func main() {
	i := 100
	s := "ポインタ"

	var pi *int = &i
	ps := &s
	fmt.Println(pi)
	fmt.Println(ps)
	fmt.Println(*pi)
	fmt.Println(*ps)
}
