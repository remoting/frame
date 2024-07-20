package queue

import (
	"github.com/remoting/frame/pkg/conv"
	"github.com/remoting/frame/pkg/goroutine"
	"sync"
	"time"
)

type Queue struct {
	mu     sync.Mutex
	delay  time.Duration
	timer  *time.Timer
	buffer map[string]time.Time
	stop   chan struct{}
	hook   func(string)
}

func NewQueue(delay time.Duration, hook func(string)) *Queue {
	dq := &Queue{
		delay:  delay,
		hook:   hook,
		buffer: make(map[string]time.Time),
		timer:  time.NewTimer(200 * time.Microsecond),
		stop:   make(chan struct{}),
	}
	go dq.start()
	return dq
}

func (dq *Queue) Add(item string, expires ...time.Duration) {
	dq.mu.Lock()
	defer dq.mu.Unlock()
	if len(expires) > 0 {
		dq.buffer[item] = time.Now().Add(expires[0])
	} else {
		dq.buffer[item] = time.Now().Add(dq.delay)
	}
}

// start processes items from the queue with the given delay.
func (dq *Queue) start() {
	for {
		select {
		case <-dq.stop:
			return
		case <-dq.timer.C:
			dq.timer.Reset(time.Millisecond * 20)
			dq.next()
		}
	}
}
func (dq *Queue) next() {
	dq.mu.Lock()
	defer dq.mu.Unlock()
	for key, value := range dq.buffer {
		if time.Now().After(value) {
			if dq.hook != nil {
				goroutine.SafeGo(func(args ...any) {
					dq.hook(conv.String(args[0]))
				}, key)
			}
			delete(dq.buffer, key)
		}
	}
}

func (dq *Queue) Stop() {
	close(dq.stop)
}
