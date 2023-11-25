package main

import (
	"fmt"
	"github.com/Trendyol/go-future/future"
	"time"
)

func main() {
	fut := future.RunWithParam(execute, "str-param")
	println("Waiting for future result...")
	result := fut.GetResult()
	fmt.Printf("Result: %s", result)
}

func execute(strParam string) (string, error) {
	time.Sleep(1000 * time.Millisecond)
	return strParam, nil
}
