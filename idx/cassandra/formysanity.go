package cassandra

import (
	"fmt"
	"os"
	"runtime"
)

func failed(c chan int) {
	success := false
	attempts := 0
	// set a random error
	err := os.ErrExist
	num := <-c
	fmt.Println("starting routine ", num)
	for !success {
		select {
		case <-c:
			//fmt.Println("attempt  ", attempts )
		default:

			if err != nil {
				success = false
				if (attempts % 20) == 0 {
					fmt.Println("attempt  ", attempts )
				}
				attempts++
			} else {
				success = true
				fmt.Println("succeed  ", attempts)

			}
		}
	}
}

func isthisok() {
	fmt.Println("main() started")
	c := make(chan int, 5)
	go failed(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	go failed(c)
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	c <- 9
	go failed(c)
	c <- 10
	c <- 11
	c <- 12
	c <- 13 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())

	go failed(c)

	fmt.Println("active goroutines", runtime.NumGoroutine())

	c <- 14
	c <- 15
	c <- 16
	c <- 17 // blocks here

	fmt.Println("active goroutines", runtime.NumGoroutine())
	fmt.Println("main() stopped")
}
