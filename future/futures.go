package future

import (
	"sync"
)

// WaitAllSilently function waits for the completion of specified Futures without collecting or
// returning any error information. It performs a silent wait using a sync.WaitGroup to wait for
// all Futures to complete.
//
// Parameters:
//   - futures: List of Futures to wait for.
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

// WaitAllContinueOnError function waits for the completion of specified Futures and continues
// even if an error occurs during the process. It returns a list of errors that occurred.
//
// Parameters:
//   - futures: List of Futures to wait for.
//
// Return:
//   - A list of errors that occurred during the execution. The list may be empty if there are no errors.
func WaitAllContinueOnError[T any](futures []*Future[T]) []error {
	return waitAll(futures, false)
}

// WaitAll function waits for the completion of specified Futures. If any error occurs during
// the process, it stops waiting and returns the first error encountered.
//
// Parameters:
//   - futures: List of Futures to wait for.
//
// Return:
//   - If at least one error occurs, it returns the first error encountered.
//   - If there are no errors, it returns nil.
func WaitAll[T any](futures []*Future[T]) error {
	errors := waitAll(futures, true)
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}

// waitAll function waits for the completion of specified Futures.
//
// Parameters:
//   - futures: List of Futures to wait for.
//   - abortOnError: If true, it stops waiting and returns the first error if any occurs.
//     If false, it collects errors as they occur and waits for other Futures to complete.
//
// Return:
//
//   - A list of errors that occurred during the execution. The list may be empty if there are no errors.
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

// GetAll function waits for the completion of specified Futures using the WaitAll function.
// If any error occurs during the waiting process, it returns an error immediately.
// Otherwise, it extracts the results from completed Futures and returns them as a slice.
//
// Parameters:
//   - futures: List of Futures to wait for.
//
// Return:
//   - A slice containing the results of completed Futures.
//   - If any error occurs during the waiting process, it returns an error.
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

// WaitFor function is a convenience wrapper around the WaitAll function.
// It waits for the completion of specified Futures and returns an error if any of the Futures encounters an error.
//
// Parameters:
//   - futures: Variadic parameter representing a list of Futures to wait for.
//
// Return:
//   - If any error occurs during the waiting process, it returns an error.
func WaitFor(futures ...*Future[any]) error {
	return WaitAll(futures)
}

// WaitForContinueOnError function is a convenience wrapper around the WaitAllContinueOnError function.
// It waits for the completion of specified Futures and continues even if an error occurs during the process.
// It returns a list of errors that occurred during the execution.
//
// Parameters:
//   - futures: Variadic parameter representing a list of Futures to wait for.
//
// Return:
//   - A list of errors that occurred during the execution. The list may be empty if there are no errors.
func WaitForContinueOnError(futures ...*Future[any]) []error {
	return WaitAllContinueOnError(futures)
}

// GetResult function retrieves the result of a completed Future.
// It assumes that the Future has completed successfully.
// Note: This function may panic if the type assertion fails.
//
// Parameters:
//   - f: Future from which to retrieve the result.
//
// Return:
//   - The result of the completed Future.
func GetResult[T any](f *Future[any]) T {
	return f.Result.(T)
}
