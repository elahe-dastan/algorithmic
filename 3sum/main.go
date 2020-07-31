package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	result := make([][]int, 0)
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = m[num] + 1
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			v := 0 - (nums[i]+nums[j])
			m[nums[i]] = m[nums[i]] - 1
			m[nums[j]] = m[nums[j]] - 1
			if m[v] > 0 {
				ans = append(ans, []int{nums[i], nums[j], v})
			}
			m[nums[i]] = m[nums[i]] + 1
			m[nums[j]] = m[nums[j]] + 1
		}
	}

	for i := 0; i < len(ans); i++ {
		f := ans[i]
		duplicate := false
		for j := i + 1; j < len(ans); j++ {
			s := ans[j]
			if same(f, s) {
				duplicate = true
				break
			}
		}

		if !duplicate {
			result = append(result, f)
		}
	}

	return result
}

func same(f []int, s []int) bool {
	sort.Ints(f)
	sort.Ints(s)

	if f[0] == s[0] && f[1] == s[1] && f[2] == s[2] {
		return true
	}

	return false
}
