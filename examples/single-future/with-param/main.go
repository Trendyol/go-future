package main

import (
	"fmt"
	"github.com/Trendyol/go-future/future"
	"time"
)

func main() {
	fut := future.RunWithParam(execute, future.Params{"str-param", true, 10})
	println("Waiting for future result...")
	result := fut.GetResult()
	fmt.Printf("Result: %s", result)
}

func execute(params future.Params) (string, error) {
	time.Sleep(1000 * time.Millisecond)
	param1, _ := future.GetParam[string](params, 0)
	param2, _ := future.GetParam[bool](params, 1)
	param3, _ := future.GetParam[int](params, 2)
	return fmt.Sprintf("%s_%t_%d", param1, param2, param3), nil
}
