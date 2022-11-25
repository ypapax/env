package env

func SecurePrint(s string) string {
	var result string
	const printEveryChar = 2
	for i, c := range s {
		if i % printEveryChar == 0{
			result += string(c)
			continue
		}
		result += "*"
	}
	return result
}

