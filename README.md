# multierr
A simple multiple error holder for Go

## Motivation

Sometimes you need to return more than one error from a function. The
`multierr` package provides a simple slice of error instances called
`Error`, which meets the `error` interface.

The way to use this package is via the package-level function `Append`.

## Usage

```go
package main

import (
	"errors"
	"fmt"

	"github.com/jonbodner/multierr"
)

func CallThing(i int) error {
	if i%2 == 1 {
		return errors.New("Error")
	}
	return nil
}

func DoSomethingThatMakesErrors(numCalls int) error {
	var out error
	for i := 0; i < numCalls; i++ {
		e := CallThing(i)
		out = multierr.Append(out, e)
	}
	return out
}

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("For i == ", i)
		e := DoSomethingThatMakesErrors(i)
		if e != nil {
			switch e := e.(type) {
			case multierr.Error:
				// you can use slice-supporting built-in functions,
				// like len or range
				fmt.Println(len(e), "errors found")
			default:
				fmt.Println("This is a normal error")
			}
			fmt.Println(e.Error())
		} else {
			fmt.Println("No errors!")
		}
	}
}
```

The `Append` function will do the right thing if `e` is `nil` (it will return its first parameter),
so you don't need to wrap calls to it in `if e != nil` if you are just aggregating errors.

It is very important to notice that `out` in `DoSomethingThatMakesErrors` is declared to be of type `error`, not `multierr.Error`. This
is due to the way that Go handles interface types and `nil`. For a variable of a non-interface type, having a `nil` value makes
the variable equal to `nil`. For a variable of an interface type, _both_ the type must be undefined _and_ the value must be `nil`
in order to make the variable equal to `nil`.

Look at this example:

```go
package main

import (
	"fmt"

	"github.com/jonbodner/multierr"
)

func meErr() error {
	var out multierr.Error
	fmt.Println(out == nil)
	return out
}

func main() {
	e := meErr()
	if e != nil {
		fmt.Println("There was an error")
	}
}
```

If you run this code, it will print out:

```
true
There was an error
```

If the type of `out` is `multierr.Error` and no errors are ever appended to it, `out` is equal to `nil` in the `meErr` function. But once the
value is returned, it is assigned to a variable of type `error`, which is an interface. The returned value has a type of `multierr.Error` and a value of `nil`.
This is sufficient for it to be considered non-nil, which makes `e != nil` true.

The `multierr.Append` function uses reflection to properly handle parameters where the underlying value is `nil`. This will work correctly:

```go
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
```

If you run this code, it will print out:

```
is nil true
This is a single error
is nil false
This is a single error
is nil false
This is a multierr of length 2
is nil false
This is a multierr of length 4
```
