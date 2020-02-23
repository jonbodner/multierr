package multierr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jonbodner/multierr"
)

func TestIt(t *testing.T) {
	var e error
	e = multierr.Append(nil, nil)
	if e != nil {
		t.Fatal("should be nil")
	}

	var nilErr error
	e = multierr.Append(nilErr, nil)
	if e != nil {
		t.Fatal("should be nil")
	}

	e = multierr.Append(e, errors.New("error1"))
	switch e := e.(type) {
	case multierr.Error:
		t.Fatal()
	default:
		if e.Error() != "error1" {
			t.Fatal()
		}
	}

	e = multierr.Append(e, errors.New("error2"))
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `error1
error2` {
			t.Fatal()
		}
	default:
		t.Fatal()
	}

	e = multierr.Append(e, errors.New("error3"))
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `error1
error2
error3` {
			t.Fatal()
		}
	default:
		t.Fatal()
	}

	e = multierr.Append(e, nil)
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `error1
error2
error3` {
			t.Fatal()
		}
	default:
		t.Fatal()
	}

	var e2 error = multierr.Error{errors.New("error 0")}
	e2 = multierr.Append(e2, e)
	switch e2 := e2.(type) {
	case multierr.Error:
		if e2.Error() != `error 0
error1
error2
error3` {
			t.Fatal()
		}
	default:
		t.Fatal()
	}

	e2 = multierr.Append(errors.New("error -1"), e2)
	switch e2 := e2.(type) {
	case multierr.Error:
		if e2.Error() != `error -1
error 0
error1
error2
error3` {
			t.Fatal()
		}
	default:
		t.Fatal()
	}
}

func TestIs(t *testing.T) {
	anotherErr := errors.New("first error")
	notPresentErr := errors.New("not present")
	e := errors.New("look for me")
	// contains error
	me := multierr.Append(anotherErr, e)
	if !errors.Is(me, e) {
		t.Error("doesn't contain e")
	}
	// both are identical multierrs
	me2 := multierr.Append(anotherErr, e)
	if !errors.Is(me, me2) {
		t.Error("not the same")
	}

	// doesn't contain error
	if errors.Is(me, notPresentErr) {
		t.Error("shouldn't contain notPresentErr")
	}

	// different lengths
	me3 := multierr.Append(multierr.Append(anotherErr, e), notPresentErr)
	if errors.Is(me, me3) {
		t.Error("should be different lengths")
	}

	// different contents
	me4 := multierr.Append(anotherErr, notPresentErr)
	if errors.Is(me, me4) {
		t.Error("should be different content")
	}

	// different order
	me5 := multierr.Append(e, anotherErr)
	if errors.Is(me, me5) {
		t.Error("should be different order")
	}
}

type testCodeErr int

func (te testCodeErr) Error() string {
	return fmt.Sprintf("code: %d", te)
}

func TestAs(t *testing.T) {
	anotherErr := errors.New("first error")
	// contains error
	me := multierr.Append(anotherErr, testCodeErr(10))

	var te testCodeErr
	if !errors.As(me, &te) {
		t.Fatal("should have contained it")
	}
	if te != 10 {
		t.Error("should be 10")
	}

	type wrong interface {
		notpresent()
	}

	// pass wrong type
	var w wrong
	if errors.As(me, &w) {
		t.Fatal("should be wrong type")
	}

	// pass multierr
	var me2 multierr.Error
	if !errors.As(me, &me2) {
		t.Fatal("should work")
	}
	if len(me2) != 2 {
		t.Fatal("should have 2 elements")
	}
	if !errors.Is(me2[0], anotherErr) {
		t.Error("should be anotherErr")
	}
	if !errors.As(me2[1], &te) {
		t.Error("should be testCodeErr")
	}

	// pass interface
	var me3 interface {
		Error() string
	}
	if !errors.As(me, &me3) {
		t.Fatal("should work")
	}
}