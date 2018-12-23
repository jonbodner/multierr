package multierr

import (
	"errors"
	"testing"
)

type estruct struct{}

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

var result bool

func BenchmarkNilInterface(b *testing.B) {
	var e error
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}

func BenchmarkNonNilInterface(b *testing.B) {
	e := errors.New("foo")
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}

func BenchmarkNilStructPointer(b *testing.B) {
	var e *estruct
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}

func BenchmarkNonNilStructPointer(b *testing.B) {
	e := &estruct{}
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}

func BenchmarkNilSlice(b *testing.B) {
	var e Error
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}

func BenchmarkNonNilSlice(b *testing.B) {
	e := Error{errors.New("foo")}
	var r bool
	for n := 0; n < b.N; n++ {
		r = isNil(e)
	}
	result = r
}
