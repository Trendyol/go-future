package future

type Group[T any] struct {
	futures []*Future[T]
}

func (g *Group[T]) Go(f func() (T, error)) *Future[T] {
	fut := Run(f)
	g.futures = append(g.futures, fut)
	return fut
}

func (g *Group[T]) GoWithParams(f func(params Params) (T, error), param Params) *Future[T] {
	fut := RunWithParam(f, param)
	g.futures = append(g.futures, fut)
	return fut
}

func (g *Group[T]) Wait() error {
	return WaitAll(g.futures)
}

func (g *Group[T]) WaitContinueOnError() []error {
	return WaitAllContinueOnError(g.futures)
}

func (g *Group[T]) Get() ([]T, error) {
	return GetAll(g.futures)
}

func (g *Group[T]) GetFutures() []*Future[T] {
	return g.futures
}
