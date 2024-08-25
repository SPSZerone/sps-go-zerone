package generic

type SInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Integer interface {
	SInt | UInt
}

type Float interface {
	~float32 | ~float64
}

type List[T any] []T
type Map[K Integer | string, V any] map[K]V
type ListMap[K Integer | string, V any] List[Map[K, V]]

type ListS List[string]
type MapSS Map[string, string]
type MapSA Map[string, any]
type ListMapSS ListMap[string, string]
type ListMapSA List[MapSA]

func ConvertT[T any](a any, defaultA T) T {
	if a == nil {
		return defaultA
	}
	if aT, ok := a.(T); ok {
		return aT
	}
	return defaultA
}

func ConvertF[T any](a any, defaultA func() T) T {
	if a == nil {
		return defaultA()
	}
	if aT, ok := a.(T); ok {
		return aT
	}
	return defaultA()
}
