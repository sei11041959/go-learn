package main

import "fmt"

type num struct {
	x int
	y int
	z int
}

func main() {
	var nums num = num{x: 1, y: 2, z: 3}
	fmt.Println(nums.x)
	fmt.Println(nums.y)
	fmt.Println(nums.z)
	fmt.Println(nums)
	nums = num{2, 3, 4}
	fmt.Println(nums.x)
	fmt.Println(nums.y)
	fmt.Println(nums.z)
	fmt.Println(nums)
}
