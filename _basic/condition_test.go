package main

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1; a > 0 {
		t.Log(a, "> 0")
	}

	if a := 1 == 1; a {
		t.Log("1==1")
	}

	if v, err := someFun(); err {
		t.Log(v)
	}
}

func someFun() (s string, err bool) {
	return "test", true
}

func TestSwitch(t *testing.T) {
	n := 1
	switch {
	case n >= 0 && n <= 10:
		t.Log("0~10")
	case n > 11 && n <= 20:
		t.Log("11~20")
	case n > 0 && n <= 30:
		t.Log("21~30")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1:
			t.Log(0, 1)
		case 2, 3:
			t.Log(2, 3)
		default:
			t.Log("ddd")
		}
	}
}

func TestSwitchMultiCaseCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log("even", i)
		case 1%2 == 1:
			t.Log("odd", i)
		default:
			t.Log(i)
		}
	}
}
