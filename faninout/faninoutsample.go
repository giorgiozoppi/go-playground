package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxGoroutines = 10
)

func evenNumbersGenerator(numbers int) <-chan int {
	generatorChannel := make(chan int)
	go func() {
		for i := 0; i < numbers; i++ {
			if i%2 == 0 {
				generatorChannel <- i
			}
			time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		}
	}()
	return generatorChannel
}
func fanOut(outputsChan ...<-chan int) <-chan int {
	// create a WaitGroup
	
	// make return channel
	merged := make(chan int, 100)
	wg.Add(len(outputsChan))
	
	for outChan:=range outputsChan {
	// to push numbers to merged channel
		go func(sc <-chan int) {
			// run until channel (square numbers sender) closes
			for sqr := range sc {
				merged <- sqr
			}
		}(outChan)
	return merged
}
func fanIn(mychannel1, mychannel2 <-chan int) chan int {
	outChannel := make(chan int)
	go func() {
		for {
			outChannel <- <-mychannel1
		}
	}()
	go func() {
		for {
			outChannel <- <-mychannel2
		}
	}()
	return outChannel
}
func fanOut(channel <- chan int) {

}
func main() {
	// fanIn example
	pos1 := evenNumbersGenerator(1000)
	pos2 := evenNumbersGenerator(2000)
	mergedChannel := fanIn(pos1, pos2)
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\n", <-mergedChannel)
	}
}
