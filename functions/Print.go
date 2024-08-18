package ascii

func Print(s string, letters map[int][]string) string {
	var art string
	for row := 0; row < 8; row++ {
		for _, char := range s {
			if int(char) < 32 || int(char) > 126 {
				continue
			}
			art += letters[int(char)][row]
		}
		if row < 7 {
			art += "\n"
		}
	}
	return art
}
