package main

import "fmt"

func SumOfSquares(c, quit chan int) {
	for i := 1; i <= 5; i++ {
		tmp := i * i
		fmt.Println(tmp)
		c <- tmp
	}
	close(c)
	select {
	case msg1 := <-quit:
		fmt.Println(msg1)
	}

	//	close(quit)
}
func main() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	data := 0
	go func() {
		for i := 0; i < 6; i++ {
			data = <-mychannel
			fmt.Printf("value %d\n", data)
			sum += data

		}
		quitchannel <- sum
	}()
	SumOfSquares(mychannel, quitchannel)
}
