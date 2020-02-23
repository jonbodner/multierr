package multierr

import (
	"errors"
	"reflect"
	"strings"
)

type Error []error

// Is performs two different checks. First it checks if the supplied error is a multierror.Error. If so,
// Is returns true if all the errors match. If the supplied error is not a multierr.Error, Is checks
// to see if any of the errors contained within the multierr.Error matches the supplied error.
func (me Error) Is(err error) bool {
	if me2, ok := err.(Error); ok {
		if len(me2) != len(me) {
			return false
		}
		// make sure all match
		for i, e := range me {
			if !errors.Is(e, me2[i]) {
				return false
			}
		}
		return true
	}
	// check if any match
	for _, e := range me {
		if errors.Is(e, err) {
			return true
		}
	}
	return false
}

// As checks to see if any of the errors contained within Error match the supplied type.
func (me Error) As(err interface{}) bool {
	for _, e := range me {
		if errors.As(e, err) {
			return true
		}
	}
	return false
}

// Error prints out all the Errors contained within the Error.
// Error prints each contained error on its own line (a \n is appended to the output)
func (me Error) Error() string {
	a := make([]string, len(me))
	for k, v := range me {
		a[k] = v.Error()
	}
	return strings.Join(a, "\n")
}

// Append takes in two errors and returns one error
// that combines the two. If either parameter is nil, the other parameter is returned.
// If the first parameter is an Error, then the function returns the second parameter
// appended to the first. Otherwise, the function returns an Error containing both parameters.
func Append(e1 error, e2 error) error {
	if isNil(e1) && isNil(e2) {
		return nil
	}
	if isNil(e1) {
		return e2
	}
	if isNil(e2) {
		return e1
	}
	switch e1 := e1.(type) {
	case Error:
		switch e2 := e2.(type) {
		case Error:
			return append(e1, e2...)
		default:
			return append(e1, e2)
		}
	default:
		switch e2 := e2.(type) {
		case Error:
			return append(Error{e1}, e2...)
		default:
			return Error{e1, e2}
		}
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
