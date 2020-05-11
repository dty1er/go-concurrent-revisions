package conrev

import "sync/atomic"

type AtomicInt int

func (ai AtomicInt) Get() int {
	return atomic.LoadInt64(&ai)
}

func (ai AtomicInt) Incr() {
	atomic.AddInt64(&ai, 1)
}
