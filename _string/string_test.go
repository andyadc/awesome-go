package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s, len(s))

	c := []rune("golang")
	t.Log(len(c))
	t.Log(c)

	t.Logf("%x", "中")

}

func TestStringFunc(t *testing.T) {
	s := "a,v,b,c,x"
	t.Log(s)
	ss := strings.Split(s, ",")
	t.Log(ss[1])

	t.Log(strings.Join(ss, "-"))
}

func TestStrConv(t *testing.T) {
	s := strconv.Itoa(90)
	t.Log("str" + s)
	if n, err := strconv.Atoi("9"); err == nil {
		t.Log(10 + n)
	}
}
