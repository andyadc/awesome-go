package _error

import (
	"errors"
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
