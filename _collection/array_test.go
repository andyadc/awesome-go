package main

import "testing"

func TestArrayInt(t *testing.T) {
	var arr [3]int
	arr[1] = 100

	arr1 := [5]int{1, 3, 5, 7, 9}
	arr2 := [...]int{2, 4, 6, 8} //不指定元素个数

	t.Log(arr[1], len(arr))
	t.Log(arr1, len(arr1))
	t.Log(arr2, len(arr2))

	t.Log(arr1[1:2])
	t.Log(arr1[1:len(arr1)])
	t.Log(arr2[1:])
	t.Log(arr2[:len(arr2)])
	t.Log(arr[:])
}

func TestArrayIterator(t *testing.T) {
	arr := [...]int{2, 4, 6, 8}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx /*索引*/, v /*元素*/ := range arr {
		t.Log(idx, v)
	}
}
