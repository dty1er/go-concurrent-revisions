package go-concurrent-revisions

type Version int

type Map struct {
	data map[int64]Version
}
