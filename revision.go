package conrev

var currentRev *Revision

type Revision struct {
	root    *Segment
	current *Segment

	done chan struct{}
}

func newRevision(root, current *Segment) *Revision {
	return &Revision{root: root, current: current}
}

func (r *Revision) fork(a Action) {
	if currentRev == nil {
		root := newSegment()
		currentRev = newRevision(root, root)
	}

	nr := newRevision(r.root, newSegmentWithParent(r.current))
	r.current.Release()
	r.current = newSegmentWithParent(r.current)

	go r.runAction(a, nr)
}

func (r *Revision) runAction(a Action, newRevision *Revision) {
	prev := currentRev
	currentRev = newRevision
	a.Do()
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
