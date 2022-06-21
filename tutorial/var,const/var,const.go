package main

import "fmt"

//global var
var (
	global1 int = 1
	global2 int = 2
)

const global3 = "golang global const"

func main() {
	var local1, local2 = 3, 5
	const local3 = "golang local const"
	fmt.Println(global1, " + ", global2, "=", global1+global2)
	fmt.Println(global3)
	fmt.Println(local1, " + ", local2, "=", local1+local2)
	fmt.Println(local3)
}
