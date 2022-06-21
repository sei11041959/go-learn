package main

import "fmt"

type MyInt int

func (myint MyInt) Multi() int {
	return int(myint) * 2
}
func main() {
	fmt.Println(MyInt(10).Multi())
}
