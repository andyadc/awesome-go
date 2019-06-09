package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func (e Employee) string() string {
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//推荐使用这种方式，减少内存使用
func (e *Employee) string2() string {
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructMethod(t *testing.T) {
	e := &Employee{1, "adc", 18}
	fmt.Printf("address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.string())
	t.Log(e.string2())
}

func TestCreateObj(t *testing.T) {
	e := Employee{1, "adc", 18}
	e1 := Employee{Name: "zkx", Age: 19}
	e2 := new(Employee) //返回指针
	e2.Id = 3
	e2.Age = 23
	e2.Name = "sdd"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}
