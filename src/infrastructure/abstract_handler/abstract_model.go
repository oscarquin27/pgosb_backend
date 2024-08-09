package abstract_handler

type AbstactModel[T any, F any] interface {
	ToModel() T
}
