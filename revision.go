package conrev

var currentRev *Revision

type Revision struct {
	Root    *Segment
	Current *Segment

	done chan struct{}
}

func NewRevision(root, current *Segment) *Revision {
	return &Revision{Root: root, Current: current}
}

func (r *Revision) Fork(a Action) {
	nr := NewRevision(r.Root, NewSegmentWithParent(r.Current))
	r.Current.Release()
	r.Current = NewSegmentWithParent(r.Current)

	go r.runAction(a, nr)
}

func (r *Revision) runAction(a Action, newRevision *Revision) {
	prev := currentRev
	currentRev = newRevision
	a.Do()
	currentRev = prev
	close(r.done)
}

func (r *Revision) Join(join *Revision) {
	<-r.done // TODO: timeout should be set?

	s := join.Current
	for s != join.Root {
		for _, w := range s.Written {
			w.Merge(r, join, s)
		}
		s = s.Parent
	}

	join.Current.Release()
	r.Current.Collapse(r)
}
