package conrev

func Fork(action func(), opts ...Option) *Revision {
	for _, opt := range opts {
		opt(config)
	}

	if currentRev == nil {
		root := newSegment()
		currentRev = newRevision(root, root)
	}

	return currentRev.fork(action)
}

func Join(join *Revision) {
	currentRev.join(join)
}
