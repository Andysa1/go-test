package gotest

import (
	"reflect"
	"testing"

	"github.com/fatih/color"
)

type Expect struct {
	t   *testing.T
	v   any
	not bool
}

func (e *Expect) ToEqual(v any) {
	if e.check(reflect.DeepEqual(e.v, v)) {
		return
	}
	e.want(v)
}
func (e *Expect) ToBeNil() {
	e.ToBe(nil)
}

func (e *Expect) ToBe(v any) {
	if e.check(e.v == v) {
		return
	}
	e.want(v)
}

func (e *Expect) Not() *Expect {
	ne := new(Expect)
	*ne = *e
	ne.not = true
	return ne
}

func (e *Expect) check(ok bool) bool {
	if e.not {
		return !ok
	}
	return ok
}

func (e *Expect) want(v any) {
	e.t.Errorf("\nwant: %s\ngot: %s",
		color.GreenString("%#v", v),
		color.RedString("%#v", e.v),
	)
}
