package slices

import (
	"golang.org/x/exp/constraints"
	xslices "golang.org/x/exp/slices"
)

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return xslices.Contains(s, v)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	return xslices.Clone(s)
}

// Sort sorts a slice of any ordered type in ascending order.
// Sort may fail to sort correctly when sorting slices of floating-point
// numbers containing Not-a-number (NaN) values.
// Use slices.SortFunc(x, func(a, b float64) bool {return a < b || (math.IsNaN(a) && !math.IsNaN(b))})
// instead if the input may contain NaNs.
func Sort[E constraints.Ordered](x []E) {
	xslices.Sort(x)
}

// SortFunc sorts the slice x in ascending order as determined by the less function.
// This sort is not guaranteed to be stable.
//
// SortFunc requires that less is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[E any](x []E, less func(a, b E) bool) {
	xslices.SortFunc(x, less)
}

func Take[E comparable](s []E, n int) []E {
	if len(s) < n {
		return s
	}
	res := make([]E, n)
	copy(s[:n], res)
	return res
}
