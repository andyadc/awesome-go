package _error

import (
	"errors"
	"fmt"
	"testing"
)

func check(n int) (string, error) {
	if n < 0 {
		return "", errors.New("n<0")
	}

	return "success", nil
}

func TestErrorCheck(t *testing.T) {
	if v, err := check(-1); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestPanicVxExit(t *testing.T) {

	defer func() {
		fmt.Println("Finally!")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from ", err)
		}
	}()

	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
