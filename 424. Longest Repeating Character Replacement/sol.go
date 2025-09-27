package characterReplacement

func characterReplacement(s string, k int) int {
	chars := make(map[byte]int)
	res, maxf := 0, 0
	l := 0
	for r := range s {
		chars[s[r]]++
		if chars[s[r]] > maxf {
			maxf = chars[s[r]]
		}

		for (r-l+1)-maxf > k {
			chars[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res

}
