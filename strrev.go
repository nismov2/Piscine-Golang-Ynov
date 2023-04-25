package piscine

func StrRev(a string) string {
	reverse := ""
	for _, b := range a {
		reverse = string(b) + reverse
	}
	return reverse
}
