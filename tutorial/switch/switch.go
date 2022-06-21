package main

import "fmt"

func main() {
	value(1)
	value(10)
	value(15)
	value(-1)
}
func value(x int) {
	fmt.Println("値は", x, "です")
	switch {
	case x <= 10 && x >= 0:
		fmt.Println("10より小さい")
	case x >= 10:
		fmt.Println("10より大きい")
	default:
		fmt.Println("0未満")
	}
}
