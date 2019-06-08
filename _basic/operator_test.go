package main

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4, 6}

	t.Log(a == b)
	t.Log(a == c)
}
