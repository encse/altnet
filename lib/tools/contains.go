package tools

func Contains(rgst []string, st string) bool {
	for _, a := range rgst {
		if a == st {
			return true
		}
	}
	return false
}
