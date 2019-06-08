package main

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 0)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{0, 1, 3}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, 10)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, 12)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, 13)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(i+1, "-", s, len(s), cap(s))
	}
	t.Log(s)
}

func TestSliceSharedMem(t *testing.T) {
	years := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(years, len(years), cap(years))

	s2 := years[3:6]
	t.Log(s2, len(s2), cap(s2))

	s4 := years[9:12]
	t.Log(s4, len(s4), cap(s4))

	summer := years[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "June"
	t.Log(years, s2, summer)
}

func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2, 3}
	// s2 := []int{1, 2, 3}
	// t.Log(s1 == s2)//slice之间不能比较
	t.Log(s1 == nil)
}
