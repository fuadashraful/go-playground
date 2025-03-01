package main

import (
	"fmt"
)

func main() {
	arr := []int {1,2,3,4,5}

	out1 := fun1(arr...)
	out2 := fun2(out1)
	out3 := func3(out2)

	for num := range out3 {
		fmt.Println(num)
	}
}

func fun1(nums ...int) <-chan int{
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func fun2(channel <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range channel {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func func3(channel <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range channel {
			out <- num * 2
		}

		close(out)
	}()
	return out
}