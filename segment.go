package conrev

var vercount int

type Segment struct {
	Parent  *Segment
	Written []Releaser
	Version int

	refcount int
}

func NewSegmentWithParent(parent *Segment) *Segment {
	if parent != nil {
		parent.refcount++
	}

	vercount++

	return &Segment{
		Parent:   parent,
		Version:  vercount,
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
