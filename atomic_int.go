package conrev

import "sync/atomic"

type AtomicInt struct {
	val int64
}

func (ai *AtomicInt) Get() int64 {
	return atomic.LoadInt64(&ai.val)
}

func (ai *AtomicInt) Incr() {
	atomic.AddInt64(&ai.val, 1)
}
