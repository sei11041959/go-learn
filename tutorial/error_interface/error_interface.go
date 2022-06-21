package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	Code    int
	Message string
	When    time.Time
}

func (e *Myerror) Error() string {
	return fmt.Sprintf("%v[エラーコード:%d]%s", e.When, e.Code, e.Message)
}

func run() error {
	return &Myerror{
		Code:    1104,
		Message: "エラーが発生しました",
		When:    time.Now(),
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
