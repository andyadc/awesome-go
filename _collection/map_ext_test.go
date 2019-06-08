package main

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m)
	t.Log(m[1](1))
	t.Log(m[2](3))
	t.Log(m[3](5))

}

func TestMapForSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true

	n := 3
	if myset[n] {
		t.Logf("%d is exist", n)
	} else {
		t.Logf("%d is not exist", n)
	}

	myset[2] = false
	delete(myset, 1)
	t.Log(myset)
}
