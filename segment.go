package conrev

var versionCount AtomicInt

type Segment struct {
	parent  *Segment
	written []*VersionedVal
	version int64

	refcount int
}

func newSegmentWithParent(parent *Segment) *Segment {
	if parent != nil {
		parent.refcount++
	}

	v := versionCount.Get()
	versionCount.Incr()

	return &Segment{
		parent:   parent,
		version:  v,
		refcount: 1,
	}
}

func newSegment() *Segment {
	return newSegmentWithParent(nil)
}

func (s *Segment) Release() {
	s.refcount--
	if s.refcount != 0 {
		return
	}

	for _, w := range s.written {
		w.Release(s)
	}
	if s.parent != nil {
		s.parent.Release()
	}
}

func (s *Segment) Collapse(main *Revision) {
	for s.parent != nil && (s.parent != main.root && s.parent.refcount == 1) {
		for _, w := range s.written {
			w.Collapse(main, s.parent)
		}
		s.parent = s.parent.parent
	}
}
