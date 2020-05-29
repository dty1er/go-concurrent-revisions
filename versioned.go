package conrev

import "sync"

type Versioned struct {
	Releaser
	Collapser
	Merger

	versions *sync.Map
}

func (ve *Versioned) Release(s *Segment) {
	ve.versions.Delete(s.version)
}

func (ve *Versioned) Collapse(main *Revision, parent *Segment) {
	_, ok := ve.versions.Load(main.current.version)
	if !ok {
		// if !ok, want to store nil so ok to ignore ok
		p, _ := ve.versions.Load(parent.version)
		ve.set(main, p)
	}
	ve.versions.Delete(parent.version)
}

func (ve *Versioned) get(r *Revision) interface{} {
	s := r.current
	for {
		v, ok := ve.versions.Load(s.version)
		if ok {
			return v
		}
		s = s.parent
	}
}

func (ve *Versioned) set(r *Revision, val interface{}) {
	s := r.current
	_, ok := ve.versions.Load(s.version)
	if !ok {
		r.current.written = append(s.written, ve)
	}
	ve.versions.Store(s.version, val)
}
