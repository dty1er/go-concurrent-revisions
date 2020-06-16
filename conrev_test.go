package conrev

import (
	"testing"
)

func TestForkJoin_1(t *testing.T) {
	v := NewVersionableInt()
	r := Fork(func() {
		v.Set(2)
	})
	v.Set(1)

	assertEq(t, v.Get(), 1)
	Join(r)
	assertEq(t, v.Get(), 2)
}

func TestForkJoin_2(t *testing.T) {
	x := NewVersionableInt()
	y := NewVersionableInt()

	r := Fork(func() {
		x.Set(1)
	})

	x.Set(0)
	y.Set(0)

	assertEq(t, x.Get(), 0)
	y.SetVersionable(x)
	assertEq(t, y.Get(), 0)

	Join(r)

	assertEq(t, x.Get(), 1)
	assertEq(t, y.Get(), 0)
}

func assertEq(t testing.TB, x, y int) {
	if x != y {
		t.Fatalf("fail: %v != %v", x, y)
	}
}
