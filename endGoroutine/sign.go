package endGoroutine

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Service struct {
	ch        chan bool
	waitGroup *sync.WaitGroup
}

func NewService() (*Service) {
	return &Service{
		ch:        make(chan bool),
		waitGroup: &sync.WaitGroup{},
	}
}
func (s *Service) Stop() {
	fmt.Println("I AM DONE...")
	close(s.ch)
	s.waitGroup.Wait()
}
func (s *Service) Serve() {
	s.waitGroup.Add(1)
	defer s.waitGroup.Done()
	for i := 1; ; i++ {
		select {
		case <-s.ch:
			fmt.Println("stopping service... ")
			return
		default:
		}
		time.Sleep(1e4)
		s.waitGroup.Add(1)
		go s.AnotherService(i)
	}

}
func (s *Service) AnotherService(id int) {
	defer s.waitGroup.Done()
	for {
		select {
		case <-s.ch:
			fmt.Printf("stopping another %d... \n", id)
			return
		default:

		}
	}
}
func SignMain() {
	service := NewService()
	go service.Serve()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-ch
		switch s {
		case syscall.SIGINT:
			println("ctrl+c")
		case syscall.SIGTERM:
			println("结束程序")
		default:
			println("other")
		}
		break
	}
	service.Stop()
}
