package hangul

var (
	// "가"의 유니코드 포인트
	// `rune` 은 int32의 alias
	start = rune(44032)
	// "힣" 다음 글자의 유니코드 포인트
	end = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		// Only concerns hangul
		if start <= r && r < end {
			index := int(r-start)
			// If remainder is 0, `r` has no consonant suffix
			result = index%numEnds != 0
		}
	}
	return result
}
