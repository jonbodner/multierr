package multierr

import "testing"

type estruct struct {}

func (e estruct) Error() string {
	return "I'm an error!"
}

func TestIsNil(t *testing.T) {
	var i int
	if isNil(i) {
		t.Fail()
	}
	var e error
	if !isNil(e) {
		t.Fail()
	}
	var es estruct
	if isNil(es) {
		t.Fail()
	}

	var esp *estruct
	if !isNil(esp) {
		t.Fail()
	}

	e = es
	if isNil(e) {
		t.Fail()
	}

	e = esp
	if !isNil(e) {
		t.Fail()
	}
}