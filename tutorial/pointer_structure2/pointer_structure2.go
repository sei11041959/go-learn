package main

import "fmt"

type Nums struct {
	x int
	y int
}

func (nums Nums) Multi() int {
	return nums.x * nums.y
}
func main() {
	nums := Nums{x: 2, y: 4}
	fmt.Println(nums.Multi())
}
