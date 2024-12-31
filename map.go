package fuse

func mapSlice[T, U any](slc []T, fn func(T) U) []U {
	ret := make([]U, len(slc))
	for idx, elt := range slc {
		ret[idx] = fn(elt)
	}
	return ret
}
