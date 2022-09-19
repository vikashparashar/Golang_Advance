// Date : 19/09/2022 00:55 IST
// Channels In Golan

// Go program to illustrate
// how to create a channel
package main

import "fmt"

func main() {

	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

	// note : this is an example of unbuffered channel
	// Go program to illustrate send
	// and receive operation

	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
