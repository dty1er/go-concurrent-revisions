package conrev

import "sync"

type Versioned struct {
	Releaser
	Collapser
	Merger

	versions *sync.Map
}

func (ve *Versioned) set(r *Revision, val interface{}) {
	s := r.current
	_, ok := ve.versions.Load(s.version)
	if !ok {
		r.current.written = append(r.current.written, ve)
	}
	ve.versions.Store(r.current.version, val)
}
