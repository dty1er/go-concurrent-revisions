package conrev

type Merger interface {
	Merge(main *Revision, joinRev *Revision, joinSeg *Segment)
}
