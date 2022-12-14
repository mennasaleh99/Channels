package main

import (
	"fmt"
)

func writemsgtoChannel(c chan string, msg string) {
	c <- msg //send msg to c
}

func greet(c chan string) {
	<-c
	<-c
}

func main() {
	fmt.Println()
	// channels are the pipes that connect concurrent goroutines
	// you can send values into channels from one goroutine and recieve those values
	// into another goroutines

	// a channel can transport data of only one data type

	// create a channel
	// make(chan type)

	// data read and write
	// use channel operator <- to send and recieve values
	// channel name <- data (push or write data to the channel)
	// <- channel name read some data from channel

	// channel operations are blocked by dfault
	// in goroutines we deal with sleep blocking a goroutines

	// var s chan int
	// fmt.Printf("channel1 - value:%v\t type:%T\n", s, s)
	// channel1 - value:<nil>   type:chan int
	// it is not useful because you cannot send or recieve data from a channel is nil

	// alternative
	// c1:=make(chan string)
	// fmt.Printf("channel 2 value :%v\t type :%T\n",c1,c1)
	// ______________________________________________________________________
	fmt.Println("Welcome to channels")             // 1
	message := make(chan string)                   //2
	go writemsgtoChannel(message, "hello Channel") //4

	fmt.Println("before recieve value")  //3
	fmt.Println("Message : ", <-message) //5 //stop main untill recieve from c a message
	fmt.Println("after recieve value")   //6
	// _______________________________________________________________________

}
