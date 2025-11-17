package main

func longestConsecutive(nums []int) int {
	s := make(map[int]struct{})
	for _, n := range nums {
		s[n] = struct{}{}
	}

	starts := make([]int, 0)
	for key, _ := range s {
		if _, exists := s[key-1]; !exists {
			starts = append(starts, key)
		}
	}

	res := 0
	for _, st := range starts {
		curr := st
		currCount := 0
		exists := true
		for exists {
			currCount++
			curr = curr + 1
			_, exists = s[curr]
		}
		if currCount > res {
			res = currCount
		}

	}

	return res

}
