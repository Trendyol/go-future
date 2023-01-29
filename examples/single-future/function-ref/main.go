package main

import (
	"fmt"
	"github.com/hamzagoc/go-future/future"
	"time"
)

func main() {
	fut := future.Run(ExampleFunc)
	println("Waiting for future result...")
	result := fut.GetResult()
	fmt.Printf("Result: %s", result)
}

func ExampleFunc() (string, error) {
	time.Sleep(1000 * time.Millisecond)
	return "name", nil
}
