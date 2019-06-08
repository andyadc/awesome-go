package main

import (
	"testing"
)

func TestExchange(t *testing.T) {
	a := 1
	b := 123
	t.Log(a, b)

	a, b = b, a
	t.Log(a, b)
}
