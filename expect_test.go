package gotest_test

import (
	"testing"

	gotest "github.com/dunstack/go-test"
)

func TestExpect(t *testing.T) {
	gotest.WithT(t, func(gt *gotest.GoTest) {
		gt.Expect(1).ToBe(1)
		gt.Expect("a").Not().ToBe("b")

		gt.Expect(nil).ToBeNil()
		gt.Expect(1).Not().ToBeNil()

		gt.Expect([]string{"a"}).ToEqual([]string{"a"})
		gt.Expect([]string{"a"}).Not().ToEqual([]int{1})
	})
}
