package main

import (
	"fmt"
	"time"
)

const (
	BILL = iota
	TED = iota
)
type message interface {
	getMessage() string
}
type local struct {
	Me string
}

// Implementing interface method in struct
func (m local) getMessage() (string) {
	return "foo"
}

// Variable paramaters
func get42(params ... int) int {
	fmt.Printf("Parameters given: %d\n", len(params))

	// making list of interfaces amount of parameters
	//is := make([]interface{}, len(params))
	sum := 0;

	// range give index to variable parameters
	for i := range params {
		fmt.Printf("%d,", params[i])
		sum += params[i];
	}
	return sum;
}

// Callback
func fun(foo func(...int) int, x, y int) int {
	meaning := foo(x, y)
	return meaning
}

func getMessage(i string) (string, bool) {
	return i, false;
} 

func doChannelStuff(pinger chan string, done chan bool) {
loop:
	select {
		case m := <-pinger:
			if m == "quit" {
				fmt.Println("Server goroutine exiting")
				done <- true
				return
			}
			fmt.Println("server:" + m)
		}
	goto loop
}

func main() {
	l := local{Me: "Hellos"}
	yep := fun(get42, 11, 22)
	_ = yep
	fmt.Printf("Meaning %d\n ", yep)
	msg, _ := getMessage(l.Me)
	fmt.Printf(msg)
	nums := []int{2, 3, 4}
	nums[0] *= 10;

	for _, x := range []int{1,2,3} {
		fmt.Println(x*2)
	}
	// Buffered channel for two messages
	pinger := make(chan string)
	exit := make(chan bool)
	
	go doChannelStuff(pinger, exit);

	// Goroutine client
	go func() {
		roundsLeft := 3

		for true {
			// Delay client
			time.Sleep(time.Second)
			
			if roundsLeft > 0 {
				pinger <- "ping"
				roundsLeft -= 1
			} else { 
				pinger <- "quit" 
			}
	}
	}()
	
	<- exit
	fmt.Println("exited")
}