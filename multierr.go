package multierr

import (
	"reflect"
	"strings"
)

type Error []error

// Error prints out all of the Errors contained within the Error.
// Each Error is printed on its own line (a \n is appended to the output)
func (me Error) Error() string {
	a := make([]string, len(me))
	for k, v := range me {
		a[k] = v.Error()
	}
	return strings.Join(a, "\n")
}

// The Append package-level function takes in two errors and returns back one error
// that combines the two. If either parameter is nil, the other parameter is returned.
// If the first parameter is a Error, then the function returns the second parameter
// appended to the first. Otherwise, the function returns a Error containing both parameters.
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
