package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 3, 3}
	fmt.Println(hasDuplicate1(a))
	fmt.Println(a)
	fmt.Println(hasDuplicate2(a))
}

func hasDuplicate2(nums []int) bool {
	numsmap := make(map[int]bool, 0)

	for _, n := range nums {
		if _, ok := numsmap[n]; ok {
			// value exist
			return true
		}
		numsmap[n] = true
	}

	return false
}

func hasDuplicate1(nums []int) bool {
	flag := false

	sort.Ints(nums)

	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			flag = true
			break
		}
		i++
	}

	return flag
}
