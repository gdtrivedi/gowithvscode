package main

// leetcode 152

import (
	"slices"
)

func main() {
	MaxProductSubArr()
}

func MaxProductSubArr() {
	println(findMaxProdSubArr([]int{1, 2, -3, 2, 2}))
}

func findMaxProdSubArr(nums []int) int {
	res := slices.Max(nums)
	curMin, curMax := 1, 1

	for _, n := range nums {
		if n == 0 {
			curMin, curMax = 1, 1
			continue
		}

		tmp := n * curMax
		curMax = max(n*curMax, n*curMin, n)
		curMin = min(tmp, n*curMin, n)

		res = max(res, curMax)
	}

	return res
}
