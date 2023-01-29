package main

import (
	"github.com/hamzagoc/go-future/future"
	"log"
)

func main() {
	fut1 := future.Run(func() (any, error) {
		return GetValueA()
	})

	fut2 := future.Run(func() (any, error) {
		return GetValueB()
	})

	log.Println("Waiting for future result...")
	err := future.WaitFor(fut1, fut2)
	if err != nil {
		log.Println("An error occurred...")
		return
	}
	result1 := future.GetResult[string](fut1)
	result2 := future.GetResult[int64](fut2)
	log.Printf("Result1: %s\n", result1)
	log.Printf("Result2: %d\n", result2)
}

func GetValueA() (string, error) {
	return "get string", nil
}

func GetValueB() (int64, error) {
	return 1, nil
}
