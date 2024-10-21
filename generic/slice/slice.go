package slice

func SafeGet[T any](slice []T, idx int) (t T, ok bool) {
	if slice == nil {
		return
	}
	lenSlice := len(slice)
	if lenSlice == 0 {
		return
	}
	if idx < 0 || idx >= lenSlice {
		return
	}
	return slice[idx], true
}
