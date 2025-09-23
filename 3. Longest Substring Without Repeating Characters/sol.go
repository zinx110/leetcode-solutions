package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	res, left := 0, 0
	seen := make(map[rune]struct{})
	for right, char := range s {
		for _, exists := seen[char]; exists; _, exists = seen[char] {
			delete(seen, rune(s[left]))
			left++
		}
		seen[char] = struct{}{}
		if currLength := right - left + 1; currLength > res {
			res = currLength
		}

	}

	return res

}
