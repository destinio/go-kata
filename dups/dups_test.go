package dups_test

import (
	"kata/dups"
	"testing"
)

type TestCase struct {
	Input    []int
	Expected bool
}

var testCases = []TestCase{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestDups(t *testing.T) {
	for _, testCase := range testCases {
		actual := dups.ContainsDuplicate(testCase.Input)
		if actual != testCase.Expected {
			t.Errorf("Input %v, expected %v, actual %v", testCase.Input, testCase.Expected, actual)
		}
	}

}
