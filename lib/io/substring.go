package io

func Substring(s string, n int) string {
	if len(s) < n {
		n = len(s)
	}
	return s[:n]
}
