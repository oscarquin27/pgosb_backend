package results

type GeneralErrorType string

const (
	TimeOutErr GeneralErrorType = "TimeOutErr"
	OtherErr   GeneralErrorType = "OtherErr"
)

type IError interface {
	Code() string
	Message() string
	AssociateException() error
}

type GeneralError struct {
	code    string
	message string
	err     error
}

func NewError(message string, err error) *GeneralError {
	return &GeneralError{
		message: message,
		err:     err,
	}
}

func NewErrorWithCode(code string, message string, err error) *GeneralError {
	return &GeneralError{
		message: message,
		err:     err,
	}
}

func (ge *GeneralError) Message() string {
	return ge.message
}
func (ge *GeneralError) AssociateException() error {
	return ge.err
}
func (ge *GeneralError) Code() string {
	return ge.code
}

func Zero[T comparable]() T {
	var zero T
	return zero
}

type Result struct {
	StepIdentifier string
	IsSuccessful   bool
	Err            IError
}

func NewResult(stepIdentifier string, isSuccessful bool, err IError) *Result {
	return &Result{
		StepIdentifier: stepIdentifier,
		IsSuccessful:   isSuccessful,
		Err:            err,
	}
}

func (r *Result) HasError() (IError, bool) {

	if r.Err == nil {
		return r.Err, false
	}
	return r.Err, true
}

func (r *Result) Success() *Result {
	r.IsSuccessful = true
	return r
}

func (r *Result) Failure() *Result {
	r.IsSuccessful = false
	return r
}

func (r *Result) FailureWithError(err error) *Result {
	r.IsSuccessful = false
	r.Err = NewError(err.Error(), err)
	return r
}

func (r *Result) WithCustomError(err IError) *Result {
	r.Err = err
	return r
}

type ResultWithValue[V comparable] struct {
	StepIdentifier string
	IsSuccessful   bool
	Value          V
	Err            IError
}

func NewResultWithValue[V comparable](stepIdentifier string, isSuccessful bool, value V, err IError) *ResultWithValue[V] {
	return &ResultWithValue[V]{
		StepIdentifier: stepIdentifier,
		IsSuccessful:   isSuccessful,
		Value:          value,
		Err:            err,
	}
}

func (rwv *ResultWithValue[V]) HasError() (IError, bool) {
	if rwv.Err == nil {
		return rwv.Err, false
	}
	return rwv.Err, true
}

func (rwv *ResultWithValue[V]) HasValue() (V, bool) {
	if rwv.Value == Zero[V]() {
		return rwv.Value, false
	}
	return rwv.Value, true
}

func (rwv *ResultWithValue[V]) Success() *ResultWithValue[V] {
	rwv.IsSuccessful = true
	return rwv
}

func (rwv *ResultWithValue[V]) Failure() *ResultWithValue[V] {
	rwv.IsSuccessful = false
	return rwv
}

func (rwv *ResultWithValue[V]) WithValue(value V) *ResultWithValue[V] {
	rwv.Value = value
	return rwv
}

func (rwv *ResultWithValue[V]) WithError(err IError) *ResultWithValue[V] {
	rwv.Err = err
	return rwv
}
