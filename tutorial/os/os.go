package main

import (
	"os"
)

func main() {
	file, err := os.Create("ばか.txt")
	if err != nil {
		return
	}
	defer file.Close()
	for i := 0; i < 100; i++ {
		file.Write([]byte("ばか"))
		if i%2 == 0 {
			file.Write([]byte("\n"))
		}
	}
}
