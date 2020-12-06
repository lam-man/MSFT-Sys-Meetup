package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	restaurant()
}

// Try Goroutine
func restaurant() {
	go chef("炒饭")
	go chef("炒时蔬")
	go chef("水煮肉片")
	fmt.Scanln()
}

// Live lock chef
func chef(dish string) {
	for i := 1; true; i++ {
		fmt.Println(i, dish)
		time.Sleep(time.Millisecond * 1000)
	}
}

// Goroutine with Synchronization sync.WaitGroup
func betterRestaurant() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		betterChef("炒饭")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		betterChef("炒时蔬")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		betterChef("水煮肉片")
		wg.Done()
	}()

	wg.Wait()
}

func betterChef(dish string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, dish)
		time.Sleep(time.Second * 1)
	}
}

// Unbuffered channel (Deadlock)
func unBufferedDeadLock() {
	c := make(chan string)
	c <- "Hello"

	msg := <-c
	fmt.Println(msg)
}

// Main --> Wait for another Goroutine to receive a message
// Main --> Cannot send anything without a receive

// Unbuffered channel is channel with 0 space
func unBufferedDeadLock0() {
	c := make(chan string, 0)
	c <- "Hello"

	msg := <-c
	fmt.Println(msg)
}

// Synchronization between threads
func unbufferedWorkedChannel() {
	c := make(chan string)

	go func() {
		msg := <-c
		fmt.Println(msg)
	}()
	c <- "Hello" // signal the above goroutine
}

// BufferedChannel
func bufferedChannel() {
	c := make(chan string, 3)
	c <- "Hello"
	c <- " System"
	c <- " Meetup"

	msg0 := <-c
	msg1 := <-c
	msg2 := <-c

	fmt.Print(msg0)
	fmt.Print(msg1)
	fmt.Println(msg2)
}

// Channel coordination
func mutipleChannelsRestaurant() {
	c0 := make(chan string, 1)
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		for {
			c0 <- "炒饭"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			c1 <- "炒时蔬"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "水煮肉片"
			time.Sleep(time.Second * 3)
		}
	}()

	// for {
	// 	fmt.Println(<-c0)
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	for {
		select {
		case msg0 := <-c0:
			fmt.Println(msg0)
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

// Restaurant with Unidirectional Channel
func bestRestaurant() {
	var accepted = []string{"炒饭", "炒时蔬", "水煮肉片", "炒饭", "炒时蔬", "水煮肉片"}
	orders := make(chan string, 6)
	defer close(orders)
	dishes := make(chan string, 6)

	go bestChef(orders, dishes)

	for i := 0; i < len(accepted); i++ {
		orders <- accepted[i]
	}

	for i := 0; i < len(accepted); i++ {
		fmt.Println(<-dishes)
	}
}

// Chef as worker
func bestChef(orders <-chan string, dishes chan<- string) {
	var processed = 0
	for order := range orders {
		time.Sleep(time.Millisecond * 1000)
		// close dishes after 5 orders
		if processed == 5 {
			close(dishes)
		}
		dishes <- order
		processed++
	}
}

// 1. receive from order
// 2. len(orders)
// 3.

func doAnything(anything string, t time.Duration) {
	for i := 0; i < 5; i++ {
		fmt.Println(anything)
		time.Sleep(time.Millisecond * t)
	}
}

func anonymousFunc() {
	nums := []int{1, 2, 3, 4}
	var wg sync.WaitGroup
	for i := range nums {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i * 5)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
