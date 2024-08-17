package gotest

import "testing"

func WithT(t *testing.T, fn func(gt *GoTest)) {
	fn(&GoTest{t})
}

type GoTest struct {
	t *testing.T
}

func (gt *GoTest) Expect(v any) *Expect {
	print("test");
	return &Expect{
		t: gt.t,
		v: v,
	}
}
