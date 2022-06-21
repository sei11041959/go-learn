package main

import "fmt"

func main() {
	fmt.Println("func add()の戻り値は", add(11, 04), "でした")
	str("Hello,func!")
	fmt.Println("func str()の戻り値はないです")
	fmt.Println("func multipled()の戻り値は", multipled(02, 25), "でした")
	divided := func(x, y int) int { return x / y }
	fmt.Println("無名関数の戻り値は", divided(10, 2), "でした")
}

//normal
func add(x, y int) int {
	var p int = x + y
	fmt.Println("func add()です")
	fmt.Println("戻り値あり！", p, "を投げるね！")
	return p
}

//return const look name
func multipled(x, y int) (m int) {
	m = x * y
	fmt.Println("func multipled()です")
	fmt.Println("戻り値あり！", m, "を投げるね！")
	return
}

//return nothing
func str(s string) {
	fmt.Println("func str()です")
	fmt.Println(s)
	fmt.Println("戻り値なし！")
	return
}
