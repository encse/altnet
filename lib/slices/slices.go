package slices

import (
	"golang.org/x/exp/constraints"
	xslices "golang.org/x/exp/slices"
)

func Contains[E comparable](s []E, v E) bool {
	return xslices.Contains(s, v)
}

func Clone[E comparable](s []E) []E {
	return xslices.Clone(s)
}

func Sort[E constraints.Ordered](x []E) {
	xslices.Sort(x)
}

func Take[E comparable](s []E, n int) []E {
	if len(s) < n {
		return s
	}
	res := make([]E, n)
	copy(s[:n], res)
	return res
}
