package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel1 <- "Player 1 Buzzed"
	}()

	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
		channel2 <- "Player 2 Buzzed"
	}()
	select {
	case message1 := <-channel1:
		{
			fmt.Println(message1)
			fmt.Println("Player 2 Buzzed")
		}
	case message2 := <-channel2:
		{
			fmt.Println(message2)
			fmt.Println("Player 1 Buzzed")
		}
	}
}
