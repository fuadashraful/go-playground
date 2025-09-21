package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numOfOrders = 10

var orderMade, orderFailed, total int

type Order struct {
	orderNo int
	message string
	success bool
}

type Producer struct {
	data chan Order
	quit chan bool
}

func (p *Producer) Close() {
	p.quit <- true
}

func newOrder(orderNo int) *Order {

	if orderNo > numOfOrders {
		return &Order{
			orderNo: orderNo,
		}
	}

	fmt.Printf("Recieved order no: %d\n", orderNo)

	delay := rand.Intn(5) + 1

	fmt.Printf("Making order no: %d. It will take %d seconds\n", orderNo, delay)
	time.Sleep(time.Duration(delay) * time.Second)

	msg := ""
	success := false

	rnd := rand.Intn(10) + 1

	fmt.Printf("Processing order no: %d. It will take %d seconds\n", orderNo, delay)

	if rnd < 5 {
		orderFailed++
	} else {
		orderMade++
	}

	total++

	if rnd <= 2 {
		msg = fmt.Sprintf("Order no: %d failed due to insufficient stock", orderNo)
	} else if rnd <= 4 {
		msg = fmt.Sprintf("Order no: %d failed due to payment issue", orderNo)
	} else {
		msg = fmt.Sprintf("Order no: %d completed successfully", orderNo)
		success = true
	}

	return &Order{
		orderNo: orderNo,
		message: msg,
		success: success,
	}
}

func processOrders(p *Producer) {
	orderNo := 1

	for {
		order := newOrder(orderNo)

		if order != nil {
			orderNo++

			select {
			case p.data <- *order:
			case <-p.quit:
				close(p.data)
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting order processing")

	orderJob := &Producer{
		data: make(chan Order),
		quit: make(chan bool),
	}

	go processOrders(orderJob)

	for order := range orderJob.data {
		if order.orderNo <= numOfOrders {
			if order.success {
				fmt.Println(order.message)
				fmt.Printf("Order %d is ready for shipping\n", order.orderNo)
			} else {
				fmt.Printf("Order processing %d failed: ", order.orderNo)
			}
		} else {
			fmt.Println("Done processing orders")
			orderJob.Close()
		}
	}

	fmt.Println("Done processing orders for the day")
}
