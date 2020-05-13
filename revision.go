package conrev

type Revision struct {
	Root       *Segment
	CurrentRev *Segment

	done chan struct{}
}

func NewRevision(root, current *Segment) *Revision {
	return &Revision{Root: root, CurrentRev: current}
}

func (r *Revision) Fork(a Action) {
	nr := NewRevision(r.Root, NewSegmentWithParent(r.CurrentRev))
	r.CurrentRev.Release()
	r.CurrentRev = NewSegmentWithParent(r.CurrentRev)
}

// run action, wait it
