package ch03

func count(s string, codeCount map[rune]int)  {
	for _, r := range s {
		codeCount[r]++
	}
}
