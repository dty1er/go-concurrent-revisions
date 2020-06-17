package conrev

import (
	"sync"
)

type VersionedVal struct {
	versions *sync.Map
}

func (ve *VersionedVal) Release(s *Segment) {
	ve.versions.Delete(s.version)
}

func (ve *VersionedVal) Collapse(main *Revision, parent *Segment) {
	_, ok := ve.versions.Load(main.current.version)
	if !ok {
		// if !ok, want to store nil
		p, _ := ve.versions.Load(parent.version)
		ve.set(main, p)
	}
	ve.versions.Delete(parent.version)
}

func (ve *VersionedVal) Merge(main *Revision, joinRev *Revision, joinSeg *Segment) {
	s := joinRev.current
	for {
		_, ok := ve.versions.Load(s.version)
		if ok {
			break
		}

		s = s.parent
	}

	if s == joinSeg {
		// ve.set(main, config.cumulativeFunc(ve.get(currentRev.current), ve.get(joinSeg), ve.get(joinRev.root)))
		v, _ := ve.versions.Load(joinSeg.version)
		ve.set(main, v)
	}
}

func (ve *VersionedVal) get(s *Segment) interface{} {
	sc := s
	for {
		v, ok := ve.versions.Load(sc.version)
		if ok {
			return v
		}

		sc = s.parent
	}
}

func (ve *VersionedVal) set(r *Revision, val interface{}) {
	_, ok := ve.versions.Load(r.current.version)
	if !ok {
		r.current.written = append(r.current.written, ve)
	}
	ve.versions.Store(r.current.version, val)
}
