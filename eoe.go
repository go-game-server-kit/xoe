package xoe

import "errors"

type EoeError struct {
	error
	arg *Arg
}

func (e *EoeError) Unwrap() error {
	return e.error
}

func (e *EoeError) Is(target error) bool {
	return errors.Is(e.error, target)
}

func (e *EoeError) As(target any) bool {
	return errors.As(e.error, target)
}

func Eoe(err error, args ...Args) error {
	if err != nil {
		var err1 *EoeError
		if errors.As(err, &err1) {
			return err1
		}
		args = append(args, WithStackSkip(-1))
		err1 = &EoeError{err, NewArg("Eoe", err, args)}
		return err1
	}
	return nil
}
