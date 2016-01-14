package q01

import "github.com/jvillasante/interviews/str"

func IsUnique1(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i != j && s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func IsUnique2(s string) bool {
	sorted := str.SortString(s)

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			return false
		}
	}

	return true
}

func IsUnique3(s string) bool {
	if len(s) > 128 {
		return false
	}

	charset := [128]bool{}
	for _, r := range s {
		if charset[r] {
			return false
		}
		charset[r] = true
	}

	return true
}

func IsUnique4(s string) bool {
	seen := make(map[rune]struct{})

	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false
		}
		seen[r] = struct{}{}
	}

	return true
}
