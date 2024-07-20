package queue

import (
	"testing"
	"time"
)

func TestQueue001(t *testing.T) {
	queue := NewQueue(2*time.Second, func(s string) {
		println(s)
	})
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	time.Sleep(1900 * time.Millisecond)
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("a")
	queue.Add("b")
	time.Sleep(1900 * time.Millisecond)
	queue.Add("a")
	time.Sleep(2100 * time.Millisecond)
	queue.Add("a")
	<-make(chan struct{})
}
