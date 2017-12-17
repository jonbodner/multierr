package multierr_test

import (
	"testing"
	"github.com/jonbodner/multierr"
	"errors"
)

func TestIt(t *testing.T) {
	var e error
	e = multierr.Append(e, errors.New("Error1"))
	switch e := e.(type) {
	case multierr.MultiErr:
		t.Fail()
	default:
		if e.Error() != "Error1" {
			t.Fail()
		}
	}

	e = multierr.Append(e, errors.New("Error2"))
	switch e := e.(type) {
	case multierr.MultiErr:
		if e.Error() != `Error1
Error2
` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	e = multierr.Append(e, errors.New("Error3"))
	switch e := e.(type) {
	case multierr.MultiErr:
		if e.Error() != `Error1
Error2
Error3
` {
			t.Fail()
		}
	default:
		t.Fail()
	}

	e = multierr.Append(e, nil)
	switch e := e.(type) {
	case multierr.MultiErr:
		if e.Error() != `Error1
Error2
Error3
` {
			t.Fail()
		}
	default:
		t.Fail()
	}
}
