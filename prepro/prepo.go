package prepro

func Filter(source string) string {
	result := ""
	i := 0

	for i < len(source) {
		if i+1 < len(source) &&
			source[i] == '/' &&
			source[i+1] == '/' {

			for i < len(source) && source[i] != '\n' {
				i++
			}

			continue
		}

		result += string(source[i])
		i++
	}

	return result
}
