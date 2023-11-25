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

func WaitAllContinueOnError[T any](futures []*Future[T]) []error {
	return waitAll(futures, false)
}

func WaitAll[T any](futures []*Future[T]) error {
	errors := waitAll(futures, true)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

func waitAll[T any](futures []*Future[T], abortOnError bool) []error {
	ch := make(chan *Future[T], len(futures))
	errs := make([]error, 0, len(futures))

	for i := range futures {
		go func(future *Future[T]) {
			future.Wait()
			ch <- future
		}(futures[i])
	}

	for i := 0; i < len(futures); i++ {
		ft := <-ch
		if ft.Err != nil {
			errs = append(errs, ft.Err)
			if abortOnError {
				break
			}
		}
	}
	return errs
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

func WaitForContinueOnError(futures ...*Future[any]) []error {
	return WaitAllContinueOnError(futures)
}

func GetResult[T any](f *Future[any]) T {
	return f.Result.(T)
}
