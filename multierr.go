package multierr

import "bytes"

type MultiErr []error

func (me MultiErr) Error() string {
	var out bytes.Buffer
	for _, v := range me {
		out.WriteString(v.Error())
		out.WriteString("\n")
	}
	return out.String()
}

func Append(e1 error, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	switch e1 := e1.(type) {
	case MultiErr:
		return append(e1, e2)
	default:
		return MultiErr{e1,e2}
	}
}