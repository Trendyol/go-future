package future

import (
	"sync"
)

func WaitAllSilently[T any](futures []*Future[T]) {
	var wg sync.WaitGroup
	wg.Add(len(futures))
	for i := range futures {
		go func(future *Future[T]) {
			future.Wait()
			wg.Done()
		}(futures[i])
	}
	wg.Wait()
}

func WaitAll[T any](futures []*Future[T]) error {
	ch := make(chan *Future[T], len(futures))

	for i := range futures {
		go func(future *Future[T]) {
			future.Wait()
			ch <- future
		}(futures[i])
	}

	// Wait for all. If any error occurred, don't wait and return first error.
	for i := 0; i < len(futures); i++ {
		ft := <-ch
		if ft.Err != nil {
			return ft.Err
		}
	}
	return nil
}

func GetAll[T any](futures []*Future[T]) ([]T, error) {
	err := WaitAll(futures)
	if err != nil {
		return nil, err
	}

	results := make([]T, len(futures))
	for i := range futures {
		results[i] = futures[i].Result
	}

	return results, nil
}

func WaitFor(futures ...*Future[any]) error {
	return WaitAll(futures)
}

func GetResult[T any](f *Future[any]) T {
	return f.Result.(T)
}
