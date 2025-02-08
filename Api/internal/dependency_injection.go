package internal

import (
	"github.com/samber/do/v2"
)

type Di struct {
	i do.Injector
}

func NewDi() *Di {
	return &Di{
		i: do.New(),
	}
}

func Provide[T any](d Di, fn func(d Di) (T, error)) {
	do.Provide(d.i, func(i do.Injector) (T, error) {
		return fn(d)
	})
}

func Invoke[T any](d Di) (T, error) {
	return do.Invoke[T](d.i)
}
