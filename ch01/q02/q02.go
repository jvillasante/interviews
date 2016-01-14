package q02

import "github.com/jvillasante/interviews/str"

func Permutation1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	sorted1 := str.SortString(s1)
	sorted2 := str.SortString(s2)
	if sorted1 != sorted2 {
		return false
	}

	return true
}

func Permutation2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := make(map[rune]int)
	for _, r := range s1 {
		charCount[r]++
	}
	for _, r := range s2 {
		charCount[r]--
		if charCount[r] < 0 {
			return false
		}
	}

	return true
}
