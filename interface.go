package conrev

func Fork(action Action) *Revision {
	if currentRev == nil {
		root := newSegment()
		currentRev = newRevision(root, root)
	}

	return currentRev.fork(action)
}

func Join(rev *Revision) {

}
