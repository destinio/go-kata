package anagram_test

import (
	"kata/anagram"
	"testing"
)

type TestCase struct {
	StringOne string
	StringTwo string
	Expected  bool
}

var TestCases = []TestCase{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
	{"a", "ab", false},
}

func TestIsAnagram(t *testing.T) {
	for _, testCase := range TestCases {
		actual := anagram.IsAnagram(testCase.StringOne, testCase.StringTwo)
		if actual != testCase.Expected {
			t.Errorf("Expected %v, but got %v", testCase.Expected, actual)
		}
	}

}
