package multierr

import (
	"bytes"
)

type Error []error

// Error prints out all of the Errors contained within the Error.
// Each Error is printed on its own line (a \n is appended to the output)
func (me Error) Error() string {
	var out bytes.Buffer
	for _, v := range me {
		out.WriteString(v.Error())
		out.WriteString("\n")
	}
	return out.String()
}

// Append adds an error to a Error. If the Error is nil, a
// new Error is instantiated
func (me *Error) Append(e error) {
	if *me == nil {
		*me = Error{}
	}
	*me = append(*me, e)
}

// The Append package-level function takes in two errors and returns back one error
// that combines the two. If either parameter is nil, the other parameter is returned.
// If the first parameter is a Error, then the function returns the second parameter
// appended to the first. Otherwise, the function returns a Error containing both parameters.
func Append(e1 error, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	switch e1 := e1.(type) {
	case Error:
		return append(e1, e2)
	default:
		return Error{e1, e2}
	}
}
