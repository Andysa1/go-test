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

	// интересное сообщение в тесте не показывается в консоли в режиме разработки и в режиме тестирования по умолчанию в терминале в режиме разработки
	return &Expect{
		t: gt.t,
		v: v,
	}
}
