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

func RemoveByKeepOrder[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func RemoveFast[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
