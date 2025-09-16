package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
)

var wg sync.WaitGroup
var balance int
var mutex sync.Mutex

type Employee struct {
	name   string
	salary int
}

func calculateEmployeeSalary() {
	employees := []Employee{
		{"Alice", 1000},
		{"Bob", 2000},
		{"Charlie", 3000},
		{"Diana", 4000},
		{"Eve", 1500},
		{"Frank", 2500},
		{"Grace", 3500},
		{"Heidi", 4500},
		{"Ivan", 1200},
		{"Judy", 2200},
	}
	wg.Add(len(employees))

	for _, emp := range employees {
		go func(e Employee) {
			defer wg.Done()
			mutex.Lock()
			balance += e.salary
			mutex.Unlock()
		}(emp)
	}

	wg.Wait()
	fmt.Println("Total Salary Calculated ", balance)
}

func TestRaceCondition(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	calculateEmployeeSalary()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stdout = oldStdout

	output := buf.String()
	expected := 25400
	if balance != expected {
		t.Errorf("Expected balance %d but got %d", expected, balance)
	}

	expectedLog := "Total Salary Calculated  25400\n"
	if output != expectedLog {
		t.Errorf("Expected console log %q but got %q", expectedLog, output)
	}
}
