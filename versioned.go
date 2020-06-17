package conrev

type Versioned interface {
	Release(rel *Segment)
	Collapse(main *Revision, parent *Segment)
	Merge(main *Revision, joinRev *Revision, joinSeg *Segment)
}
