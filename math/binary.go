package math

import "github.com/SPSZerone/sps-go-zerone/generic"

func IsPowerOf2[T generic.Integer](n T) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
