package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFileSize(t *testing.T) {
	f, err := os.Stat("/Temp")
	if err == nil {
		fmt.Println("name:", f.Name())
		fmt.Println("size:", f.Size())
		fmt.Println("is dir:", f.IsDir())
		fmt.Println("mode::", f.Mode())
		fmt.Println("modTime:", f.ModTime())
	} else if os.IsNotExist(err) {
		fmt.Println("file not exist")
	} else {
		log.Println(err)
		fmt.Println(err)
	}
}
