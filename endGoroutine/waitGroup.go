package endGoroutine

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("id %d start\n", id)
	time.Sleep(1e9)
	fmt.Printf("id %d end\n", id)
}

func WaitGroupMain() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		fmt.Printf("Main start working %d\n", i+1)
		wg.Add(1)
		go WaitGroup(&wg, i+1)

	}
	fmt.Println("Main wait to workers finish")
	wg.Wait()
	fmt.Println("Main all worker finish")
}
