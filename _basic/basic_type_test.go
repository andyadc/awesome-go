package main

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	a := 1
	var b int64

	b = int64(a) //不支持隐式转换
	var c MyInt
	// c = b
	t.Log(a, b, c)
}

func TestDefault(t *testing.T) {
	t.Log(math.MinInt64)
	t.Log(math.MaxInt64)
}

func TestPoint(t *testing.T) {
	a := 1
	ap := &a //a 的指针
	t.Log(a, ap)
	t.Logf("%T %T", a, ap) //打印类型

	// ap = ap + 1 //不支持指针计算
}

func TestString(t *testing.T) {
	var s string //默认是空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
}
