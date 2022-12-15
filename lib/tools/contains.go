package tools

func Contains[T comparable](rgt []T, t T) bool {
	for _, a := range rgt {
		if a == t {
			return true
		}
	}
	return false
}
