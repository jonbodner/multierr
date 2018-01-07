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
