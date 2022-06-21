package main

import (
	"fmt"
)

func main() {
	//型アサーション
	var x interface{} = 1
	i, isInt := x.(int)
	fmt.Println(i, isInt)
	s, isStr := x.(string)
	fmt.Println(s, isStr)
	//switch文 アサーション
	var b interface{} = true
	var ui interface{} = 1
	var str interface{} = "str"
	var f interface{} = 3.14
	typeSwitch(b)
	typeSwitch(ui)
	typeSwitch(str)
	typeSwitch(f)
}

func typeSwitch(x interface{}) {
	fmt.Println("引数は", x, "です")
	switch x.(type) {
	case bool:
		fmt.Println("論理型")
	case int, uint:
		fmt.Println("整数型")
	case string:
		fmt.Println("文字列")
	default:
		fmt.Println("その他")
	}
}
