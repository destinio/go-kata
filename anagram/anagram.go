package anagram

import "sort"

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	sChars := []rune(s)
	tChars := []rune(t)

	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] < sChars[j]
	})

	sort.Slice(tChars, func(i, j int) bool {
		return tChars[i] < tChars[j]
	})

	return string(sChars) == string(tChars)
}
