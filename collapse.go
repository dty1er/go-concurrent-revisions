package conrev

type Collapser interface {
	Collapse(main *Revision, parent *Segment)
}
