package _interface

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	//断言
	if v, ok := p.(int); ok {
		fmt.Printf("%d is integer\n", v)
		return
	}
	if v, ok := p.(string); ok {
		fmt.Printf("%s is string\n", v)
		return
	}
	fmt.Println("unknown")
}

func doAnything(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("integer ", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("unknow")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	doSomething(1)
	doSomething(1.1)

	doAnything("s")
}
