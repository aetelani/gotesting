package main

import (
	"fmt"
)

const (
	BILL = iota
	TED = iota
)
type message interface {
	getMessage() string
}
type local struct {
	Me string;
}
func (m local) getMessage() (string) {
	return "foo"
}

func getMessage(i string) (string, bool) {
	return i, false;
} 

func main() {
	l := local{Me: "Hellos"}
	msg, _ := getMessage(l.Me)
	fmt.Printf(msg)
	nums := []int{2, 3, 4}
	nums[0] *= 10;

	for _, x := range []int{1,2,3} {
		fmt.Println(x*2)
	}
}