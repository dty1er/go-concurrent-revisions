package conrev

import (
	"fmt"
	"sync"
)

type Versionable interface {
	SetVersionable(interface{})
	Set(interface{})
	Get() interface{}
}

type VersionableInt struct {
	vals *VersionedVal
}

func NewVersionableInt() *VersionableInt {
	return &VersionableInt{
		vals: &VersionedVal{
			versions: &sync.Map{},
		},
	}
}

func (vi *VersionableInt) SetVersionable(v interface{}) {
	casted, ok := v.(*VersionableInt)
	if !ok {
		panic("invalid argument: must be *VersionableInt")
	}

	vi.vals.set(currentRev, casted.Get())
}

func (vi *VersionableInt) Set(v interface{}) {
	casted, ok := v.(int)
	if !ok {
		panic("invalid argument: must be int")
	}

	vi.vals.set(currentRev, casted)
}

func (vi *VersionableInt) Get() int {
	v := vi.vals.get(currentRev.current)
	iv, ok := v.(int)
	if !ok {
		panic(fmt.Sprintf("not int: %v", iv))
	}

	return iv
}
