package multierr_test

import (
	"errors"
	"github.com/jonbodner/multierr"
	"testing"
)

func TestIt(t *testing.T) {
	var e error
	e = multierr.Append(e, errors.New("Error1"))
	switch e := e.(type) {
	case multierr.Error:
		t.Fail()
	default:
		if e.Error() != "Error1" {
			t.Fail()
		}
	}

	e = multierr.Append(e, errors.New("Error2"))
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `Error1
Error2` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	e = multierr.Append(e, errors.New("Error3"))
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `Error1
Error2
Error3` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	e = multierr.Append(e, nil)
	switch e := e.(type) {
	case multierr.Error:
		if e.Error() != `Error1
Error2
Error3` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	var e2 error = multierr.Error{errors.New("Error 0")}
	e2 = multierr.Append(e2, e)
	switch e2:= e2.(type) {
	case multierr.Error:
		if e2.Error() != `Error 0
Error1
Error2
Error3` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	e2 = multierr.Append(errors.New("Error -1"), e2)
	switch e2:= e2.(type) {
	case multierr.Error:
		if e2.Error() != `Error -1
Error 0
Error1
Error2
Error3` {
			t.Fail()
		}
	default:
		t.Fail()
	}
}
