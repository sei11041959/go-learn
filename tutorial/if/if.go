package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
		max(i)
	}
}

func max(i int) {
	if i >= 10 {
		fmt.Println("10よりおおきい")
	} else {
		fmt.Println("10より小さい")
	}
}
