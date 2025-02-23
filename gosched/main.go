package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func doWork(message string, enableGosched bool, wg *sync.WaitGroup) {
	
	for i := 0; i < 10; i++ {
		if enableGosched {
			/* 
				change current context (yield) - stops and saves the current goroutine,
				switches execution control to another goroutine.
			*/
			runtime.Gosched()
		}

		fmt.Printf("%s -- %d\n", message, i)
	}

	wg.Done()
}


func main() {
	fmt.Println("*** With switching context between goroutine ***")
	wg.Add(2)
	
	go doWork("Tic", true, &wg)
	doWork("Toc", true, &wg)
	
	wg.Wait()

	fmt.Println("*** without switching context between goroutine **")

	wg.Add(2)

	go doWork("Bae", false, &wg)
	doWork("Boe", false, &wg)

	wg.Wait()
}