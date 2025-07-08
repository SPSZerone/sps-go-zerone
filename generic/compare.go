package generic

type Comparable[T any] interface {
	Compare(T) int
}
