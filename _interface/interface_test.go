package _interface

import (
	"fmt"
	"testing"
)

type Code string
type Programer interface {
	WriterHello() Code
}

type GoProgramer struct{}
type JavaProgramer struct{}

func (g *GoProgramer) WriterHello() Code {
	return "fmt.Println(\"hello golang\")"
}

func (g *JavaProgramer) WriterHello() Code {
	return "System.out.println(\"hello java\")"
}

func writeFirstProgram(p Programer) {
	fmt.Printf("%T %v\n", p, p.WriterHello())
}

func TestPolymorphism(t *testing.T) {
	p1 := new(JavaProgramer)
	p2 := new(GoProgramer)
	p3 := &GoProgramer{}

	writeFirstProgram(p1)
	writeFirstProgram(p2)
	writeFirstProgram(p3)
}

func TestClient(t *testing.T) {
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriterHello())
}
