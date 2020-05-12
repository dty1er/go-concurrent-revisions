package conrev

var versionCount AtomicInt

type Segment struct {
	Parent  *Segment
	Written []Versioned
	Version int64

	refcount int
}

func NewSegmentWithParent(parent *Segment) *Segment {
	if parent != nil {
		parent.refcount++
	}

	v := versionCount.Get()
	versionCount.Incr()

	return &Segment{
		Parent:   parent,
		Version:  v,
		refcount: 1,
	}
}

func NewSegment() *Segment {
	return NewSegmentWithParent(nil)
}

func (s *Segment) Release() {
	s.refcount--
	if s.refcount != 0 {
		return
	}

	for _, w := range s.Written {
		w.Release(s)
	}
	if s.Parent != nil {
		s.Parent.Release()
	}
}

func (s *Segment) Collapse(main *Revision) {
	for s.Parent != nil && s.Parent != main.Root && s.Parent.refcount == 1 {
		for _, w := range s.Written {
			w.Collapse(main, s.Parent)
		}
		s.Parent = s.Parent.Parent
	}
}
