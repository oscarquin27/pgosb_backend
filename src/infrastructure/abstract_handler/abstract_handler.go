package abstract_handler

type AbstractCRUDService[T any] interface {
	Get(id int) (T, error)
	GetAll() ([]T, error)
	Create(value T) (T, error)
	Update(value T) error
	Delete(id int) error
}

type AbstractHandler[T any] struct {
	service AbstractCRUDService[T]
}

func NewAbstractHandler[T any](s AbstractCRUDService[T]) *AbstractHandler[T] {

	return &AbstractHandler[T]{
		service: s,
	}
}
