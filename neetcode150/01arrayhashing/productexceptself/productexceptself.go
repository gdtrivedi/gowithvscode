package main

import "fmt"

func main() {
	fmt.Println("Result:", productExceptSelf([]int{1, 2, 4, 6}))
	fmt.Println("Result:", productExceptSelf([]int{-1, 0, 1, 2, 3}))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)
	suffix := make([]int, l)

	prd := 1
	for i, _ := range nums {
		if i == 0 {
			prefix[i] = 1
		} else {
			prefix[i] = prd * nums[i-1]
		}
	}
}
