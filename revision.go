package conrev

var currentRev *Revision

type Revision struct {
	root    *Segment
	current *Segment

	done chan struct{}
}

func newRevision(root, current *Segment) *Revision {
	return &Revision{root: root, current: current, done: make(chan struct{}, 1)}
}

func (r *Revision) fork(action func()) *Revision {
	nr := newRevision(r.root, newSegmentWithParent(r.current))
	r.current.Release()
	r.current = newSegmentWithParent(r.current)

	go r.runAction(action, nr)
	return nr
}

func (r *Revision) runAction(action func(), newRevision *Revision) {
	prev := currentRev
	currentRev = newRevision
	action()
	currentRev = prev
	close(r.done)
}

func (r *Revision) join(join *Revision) {
	<-r.done // TODO: timeout should be set?

	s := join.current
	for s != join.root {
		for _, w := range s.written {
			w.Merge(r, join, s)
		}
		s = s.parent
	}

	join.current.Release()
	r.current.Collapse(r)
}
