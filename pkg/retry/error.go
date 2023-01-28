package retry

type Error []error

func (e *Error) HasError() bool {
	return len(*e) != 0
}

func (e *Error) LastError() error {
	if e.HasError() {
		return (*e)[len(*e)-1]
	}
	return nil
}

func (e *Error) Error() error {
	if e.HasError() {
		return (*e)[0]
	}
	return nil
}

func NewError(errs ...error) Error {
	var err Error
	err = append(err, errs...)
	return err
}
