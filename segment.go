package conrev

var vercount int

type Segment struct {
	Parent  *Segment
	Written []string
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
