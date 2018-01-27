package poloniex

import (
	"fmt"
	"sync"
)

type Logger struct {
	isOpen bool
	mu     *sync.Mutex
}

func (l *Logger) LogRoutine(bus <-chan string) {
	if l.isOpen {
		for {
			message := <-bus
			l.mu.Lock()
			fmt.Println(message)
			l.mu.Unlock()
		}
	}
}
