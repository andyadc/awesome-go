package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	password := []byte("adc")

	// Hashing the password with the default cost of 10.
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

    t.Log("---")
	t.Log(string(hashedPassword))
}
