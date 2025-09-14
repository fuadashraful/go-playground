package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func sayHello(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func TestWaitGroup1(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	sayHello("Hi Fuad !!", &wg)

	wg.Wait()

	w.Close()

	result, _ := io.ReadAll(r)

	out := string(result)

	os.Stdout = stdOut
	expected := "Hi Fuad !!\n"

	if out != expected {
		t.Errorf("Expected %q but got %q", expected, out)
	}

}
