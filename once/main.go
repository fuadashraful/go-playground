package main

import (
	"fmt"
	"sync"
)

var callOnce sync.Once

var lambdaFunc = func() {
	fmt.Println("This will be called once only!!")
}

func doSomeWork() {
	callOnce.Do(lambdaFunc)

	fmt.Println("Doing some work !!")
}

func main() {
	doSomeWork()
	doSomeWork()

	for i:=0; i < 10; i++ {
		callOnce.Do(lambdaFunc)
	}
}