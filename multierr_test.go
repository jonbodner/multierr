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
Error2
` {
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
Error3
` {
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
Error3
` {
			t.Fail()
		}
	default:
		t.Fail()
	}
}

func TestMultiErr_Append(t *testing.T) {
	var me multierr.Error
	if me != nil {
		t.Fail()
	}
	me.Append(errors.New("Error1"))
	if me == nil {
		t.Fail()
	}
	if me.Error() != `Error1
` {
		t.Fail()
	}
	me.Append(errors.New("Error2"))
	if me.Error() != `Error1
Error2
` {
		t.Fail()
	}
}
