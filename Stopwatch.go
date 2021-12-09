package main

import (
	"sync"
	"time"
)

type Stopwatch struct {
	watch time.Time
	lock  sync.RWMutex
}

func (s *Stopwatch) Start() {
	s.lock.Lock()
	s.watch = time.Now()
	s.lock.Unlock()
}

func (s *Stopwatch) ElapsedMilliseconds() int64 {

	s.lock.RLock()
	t := time.Now().UnixMilli() - s.watch.UnixMilli()
	s.lock.RUnlock()

	return t
}
