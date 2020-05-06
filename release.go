package conrev

type Releaser interface {
	Release(rel *Segment)
}
