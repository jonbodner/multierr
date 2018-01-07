package main

import (
	"github.com/jonbodner/multierr"
	"fmt"
	"errors"
)

func main() {
	var e error
	var e2 multierr.Error

	e = multierr.Append(e, e2)
	checkError(e)

	e = multierr.Append(e, errors.New("This is an error"))
	checkError(e)

	e = multierr.Append(e, errors.New("I'm a second error"))
	checkError(e)

	e2 = multierr.Error{errors.New("I'm a third error"), errors.New("I'm a fourth error")}
	e = multierr.Append(e, e2)
	checkError(e)
}

func checkError(e error) {
	fmt.Println("is nil",e == nil)
	switch e := e.(type) {
	case multierr.Error:
		fmt.Println("This is a multierr of length",len(e))
	default:
		fmt.Println("This is a single error")
	}
}

