package main

import "fmt"

type Nums struct {
	x  int
	y  int
	xb *int
	yb *int
}

func (nums *Nums) double() {
	nums.x = nums.x * 2
	nums.y = nums.y * 2
	nums.xb = &nums.x
	nums.yb = &nums.y
}
func main() {
	nums := Nums{x: 2, y: 4}
	nums.double()
	fmt.Println(nums)
}
