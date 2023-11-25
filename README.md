# go-future
This library helps to use goroutines and wait for result.

## Usage example

### Single Future
#### Anonymous Function usage
```go
fut := future.Run(func() (string, error) {
	time.Sleep(1000 * time.Millisecond)
	return "name", nil
})

// do anything ...

// wait until the process is complete and getResult
result := fut.GetResult()
```
#### Handle error
```go
result, err := fut.Get()
if err != nil {
    println("An error occurred.")
}
```

#### Function Ref
```go
fut := future.Run(execute)
result := fut.GetResult()

func execute() (string, error) {
    time.Sleep(1000 * time.Millisecond)
    return "name", nil
}

```

#### With Single Param
```go
fut := future.RunWithParam(execute, "str-param")
result := fut.GetResult()

func execute(strParam string) (string, error) {
    time.Sleep(1000 * time.Millisecond)
    return strParam, nil
}

```

#### With Multi Param
```go
fut := future.RunWithParam(execute, future.Params{"str-param", true, 10})
result := fut.GetResult()

func execute(params future.Params) (string, error) {
    time.Sleep(1000 * time.Millisecond)
    param1, _ := future.GetParam[string](params, 0)
    param2, _ := future.GetParam[bool](params, 1)
    param3, _ := future.GetParam[int](params, 2)
    return fmt.Sprintf("%s_%t_%d", param1, param2, param3), nil
}
```
---
### Multiple Future
#### Parallel request example
```go
ids := []string{"A", "B", "C", "D", "E"}
futures := make([]*future.Future[string], 5)
for i := range ids {
    id := ids[i]
    f := future.Run(func() (string, error) {
    return APICall(id)
    })
    futures[i] = f
}

// wait all until the process is complete and getResult
results, err := future.GetAll(futures)
if err != nil {
    log.Println("An error occurred...")
    return
}
```

#### Parallel request different result example
```go
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
```