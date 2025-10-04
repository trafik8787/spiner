package spiner

import (
	"fmt"
	"time"
	"sync"
)


type Spiner struct {
	Status bool
	mu sync.Mutex
}


func (s *Spiner) Start () {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Status {
		return
	}

	s.Status = true
	go s.spin()
}

func (s *Spiner)  Stop () {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Status = false
}

func (s *Spiner) spin () {
	
	for s.Status {
		for _, r := range `-\|/` {
			fmt.Printf("\r\033[31m%c\033[0m", r)
			time.Sleep(time.Millisecond * 90)
		}
	}
	 fmt.Print("\r")
}


