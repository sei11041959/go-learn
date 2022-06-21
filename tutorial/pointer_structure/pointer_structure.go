package main

import "fmt"

type Nums struct {
	x int
	y int
	z int
}

func double(nums Nums) {
	nums.x = nums.x * 2
	nums.y = nums.y * 2
	nums.z = nums.z * 2

	fmt.Println("double内における構造体の値", nums)
}

func triple(nums *Nums) {
	nums.x = nums.x * 3
	nums.y = nums.y * 3
	nums.z = nums.z * 3

	fmt.Println("triple内における構造体の値", *nums)
}
func main() {
	nums := Nums{
		x: 1,
		y: 2,
		z: 3,
	}
	double(nums)
	fmt.Println("double実行後の構造体の値", nums)

	triple(&nums)
	fmt.Println("triple実行後の構造体の値", nums)

}
