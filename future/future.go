package future

import (
	"sync"
)

type Future[T any] struct {
	wg     sync.WaitGroup
	Result T
	Err    error
	IsDone bool
}

func Run[T any](f func() (T, error)) *Future[T] {
	future := &Future[T]{}
	future.wg.Add(1)
	go func() {
		defer future.wg.Done()
		future.Result, future.Err = f()
		future.IsDone = true
	}()
	return future
}

func RunWithParam[T any, P any](f func(P) (T, error), param P) *Future[T] {
	future := &Future[T]{}
	future.wg.Add(1)
	go func(p P) {
		defer future.wg.Done()
		future.Result, future.Err = f(p)
		future.IsDone = true
	}(param)
	return future
}

func (f *Future[T]) Wait() {
	f.wg.Wait()
}

func (f *Future[T]) Get() (T, error) {
	f.Wait()
	return f.Result, f.Err
}

func (f *Future[T]) GetResult() T {
	f.Wait()
	return f.Result
}
