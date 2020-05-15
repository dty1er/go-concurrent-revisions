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
	r.done <- struct{}{}
}
