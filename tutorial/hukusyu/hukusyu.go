package main

import "fmt"

func main() {
	fmt.Println("復習！")
	fmt.Println("addの戻り値は", add(7, 9), "です")
	fmt.Println("multipledの戻り値は", multipled(7, 9), "です")
}
func add(x, y int) (r int) {
	fmt.Println("こんにちは!func add()です!")
	r = x + y
	fmt.Println("引数の", x, "と", y, "を足して", r, "になったよ!", "戻り値としてこれを返すね!")
	return
}
func multipled(x, y int) (r int) {
	fmt.Println("こんにちは!func multipledです!")
	r = x * y
	fmt.Println("引数の", x, "と", y, "を足して", r, "になったよ!", "戻り値としてこれを返すね!")
	return
}
