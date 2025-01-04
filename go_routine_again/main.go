package main

import (
	"fmt"
)

func SendData(ch chan string, sendChan chan string) {
	for i := 0; i < 5; i++ {
		sendString := fmt.Sprintf("This is a %v string\n", i)
		ch <- sendString
	}
	close(ch)
	sendChan <- "Done"
}

func RecieveData(ch chan string, recieverChan chan string) {
	for i := range ch {
		fmt.Println(i)
	}
	recieverChan <- "Done"
}

func main() {
	ch := make(chan string)
	senderChan := make(chan string)
	recieverChan := make(chan string)
	go SendData(ch, senderChan)
	go RecieveData(ch, recieverChan)
	<-senderChan
	// <-recieverChan
}
