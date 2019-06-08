package main

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	m1["five"] = 5
	t.Log(m1, len(m1))

	m2 := map[string]string{}
	m2["four"] = "f"
	t.Log(m2, len(m2))

	m3 := make(map[string]int, 10 /*Initial Capcity*/)
	t.Log(m3, len(m3))
}

//当访问key不存在时，仍会返还零值，不能通过返回nil来判断元素是否存在
func TestGetValue(t *testing.T) {
	m1 := map[string]string{}
	m1["monday"] = "星期一"
	t.Log(m1["monday"], m1["sunday"])

	m2 := map[int]int{}
	m2[1] = 10
	t.Log(m2[1], m2[2])

	v, ok := m2[11]
	t.Log(v, ok)
}

func TestMapIterator(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}
