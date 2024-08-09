package models

type AbstactModel[T any, F any] interface {
	ToModel() T
	FromModel(*T) *F
}
