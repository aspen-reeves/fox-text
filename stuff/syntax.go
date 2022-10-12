package stuff

// parse string for syntax highlighting
func parseLine(src Bruh) []string {
	var words []string
	var word string
	for i := 0; i < len(src.Lines[src.YOffset]); i++ {
		if src.Lines[src.YOffset][i] == ' ' {
			words = append(words, word)
			word = ""
		} else {
			word += string(src.Lines[src.YOffset][i])
		}
	}
	return words
}
