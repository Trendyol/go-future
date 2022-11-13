package main

import (
	"fmt"
	"go-future/future"
	"time"
)

func main() {
	fut := future.Run(func() (string, error) {
		time.Sleep(1000 * time.Millisecond)
		return "name", nil
	})
	println("Waiting for future result...")
	result, err := fut.Get()
	if err != nil {
		println("An error occurred.")
	} else {
		fmt.Printf("Result: %s", result)
	}
}
