package piscine

func ConcatParams(args []string) string {
	var empty string
	for i, a := range args {
		if i < len(args)-1 {
			empty = empty + a + "\n"
		} else {
			empty = empty + a
		}
	}
	return empty
}
