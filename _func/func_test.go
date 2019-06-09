package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	a, b := returnMultiVal()
	t.Log(a, b)

	ts := timeSpent(slowFunc)
	t.Log(ts(999))
}

func returnMultiVal() (int, int) {
	rand.Seed(time.Now().UnixNano()) //伪随机
	return rand.Intn(100), rand.Intn(1000)
}

//计算func执行耗时
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Printf("time spent: %f", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(n int) int {
	time.Sleep(time.Second * 3)
	return n
}

func TestVariableParams(t *testing.T) {
	t.Log(sum(1, 2, 5, 7, 9, 0))
}

func sum(n ...int) int {
	sum := 0
	for _, v := range n {
		sum = sum + v
	}
	return sum
}
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("I'm late")
	}()
	defer clear()
	// t.Log("hello")
	fmt.Println("hello")

	// panic("Fatal error")
}

func clear() {
	fmt.Printf("clearing")
}
