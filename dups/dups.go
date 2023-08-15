package dups

func ContainsDuplicate(nums []int) bool {
	var m = make(map[int]int)

	for _, v := range nums {

		if m[v] == 1 {
			return true
		}

		m[v] = 1
	}

	return false
}
