package helpers

func RestoreBytes(shuffled []byte, indexes []int) []byte {
	restored := make([]byte, len(shuffled))

	for i, idx := range indexes[:len(indexes)-8] {
		restored[idx] = shuffled[i]
	}

	return restored
}
