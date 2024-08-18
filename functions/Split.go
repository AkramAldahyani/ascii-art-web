package ascii

func Split(s, sep string) []string {
	list := []string{}
	pointer := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if sep == s[i:i+len(sep)] {
			if pointer != i {
				list = append(list, s[pointer:i])
			}
			list = append(list, "\n")
			pointer = i + len(sep)
			i += len(sep) - 1
		}
	}
	if pointer < len(s) {
		end := s[pointer:]
		list = append(list, end)
	}
	return list
}
