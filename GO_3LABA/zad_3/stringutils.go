package stringutils

func Flipping(n string) string {

	var result string = ""
	for i := len(n) - 1; i >= 0; i-- {
		result += string(n[i])
	}
	return result
}
