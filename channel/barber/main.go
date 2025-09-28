package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arriavalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barberName string) {

	shop.NumberOfBarbers++

	go func() {

		isSleeping := false

		fmt.Printf("%s going to waiting room to check for clients\n", barberName)

		for {
			if len(shop.ClientsChan) == 0 {
				fmt.Printf("No clients, so %s is taking a nap\n", barberName)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					fmt.Printf("Client %s wakes up Barber %s\n", client, barberName)
					isSleeping = false
				}

				shop.cutHair(barberName, client)
			} else {
				shop.sendBarberHome(barberName)
				return
			}
		}

	}()
}

func (shop *BarberShop) cutHair(barberName, client string) {
	fmt.Printf("Barber %s is cutting hair of client %s\n", barberName, client)
	time.Sleep(shop.HairCutDuration)
	fmt.Printf("Barber %s has finished cutting hair of client %s\n", barberName, client)
}

func (shop *BarberShop) sendBarberHome(barberName string) {
	fmt.Printf("Barber %s is going home for the day\n", barberName)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	fmt.Println("Barber shop is closing for the day")

	close(shop.ClientsChan)
	shop.Open = false

	for i := 0; i < shop.NumberOfBarbers; i++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	fmt.Println("*** Barber shop is closed for the day, enveryone goes for home ***")
}

func (shop *BarberShop) addClient(clientName string) {
	fmt.Printf("Client %s arrived at the barber shop\n", clientName)

	if shop.Open {
		select {
		case shop.ClientsChan <- clientName:
			fmt.Printf("Client %s taking a seat in the waiting room\n", clientName)
		default:
			fmt.Printf("No seats available, client %s is leaving\n", clientName)
		}
	} else {
		fmt.Printf("Barber shop is closed, client %s is leaving\n", clientName)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Barber shop is opening with %d seats in the waiting room\n", seatingCapacity)

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}

	fmt.Println("Barber shop is open for business")

	shop.addBarber("Barber 1")
	shop.addBarber("Barber 2")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			ranndomMillseconds := rand.Int() % (2 * arriavalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Duration(ranndomMillseconds) * time.Millisecond):
				clientName := fmt.Sprintf("Client %d", i)
				shop.addClient(clientName)
				i++
			}
		}
	}()

	<-closed
}
