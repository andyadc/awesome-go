package main

import (
	"fmt"
	"math"
	"os"
)

const helloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

// Hello func
func Hello(name string) string {
	return helloPrefix + name
}

// Hello2 func
func Hello2(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(os.Args)

	fmt.Println(Hello("andy"))

	fmt.Println("hello go")
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "abc"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	fmt.Println(helloPrefix)

	const n = 50000000
	const m = 3e20 / n
	fmt.Println(m)

	fmt.Println(int64(m))
	fmt.Println(math.Sin(n))
}
