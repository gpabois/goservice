package utils

type Configurator[T any] func(s *T)

func Configure[T any](s *T, options []Configurator[T]) {
	for _, c := range options {
		c(s)
	}
}
