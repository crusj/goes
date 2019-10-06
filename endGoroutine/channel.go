package endGoroutine

import (
	"fmt"
	"time"
)

func Run(done chan int) {
	for {
		select {
		case <-done:
			fmt.Println("exiting....")
			done <- 1
			break
		default:
		}
		time.Sleep(time.Second * 1)
		fmt.Println("do something")
	}
}
func Main() {
	c := make(chan int)
	go Run(c)
	fmt.Println("wait")
	time.Sleep(1e9 * 5)
	c <- 1
	<-c
	fmt.Println("main exited")
}
