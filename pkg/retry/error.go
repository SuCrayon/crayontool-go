package retry

type Errors []error

func (e *Errors) HasError() bool {
	return len(*e) != 0
}

func (e *Errors) LastError() error {
	if e.HasError() {
		return (*e)[len(*e)-1]
	}
	return nil
}

func (e *Errors) Error() error {
	if e.HasError() {
		return (*e)[0]
	}
	return nil
}

func NewError(errs ...error) Errors {
	var err Errors
	err = append(err, errs...)
	return err
}
