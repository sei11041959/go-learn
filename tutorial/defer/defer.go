package main

import "fmt"

func main() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")

	fmt.Println("deferは後入れ先出しで実行される")
}
