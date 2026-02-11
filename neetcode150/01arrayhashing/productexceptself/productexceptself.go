package main

import "fmt"

func main() {
	fmt.Println("Result:", productExceptSelf([]int{1, 2, 4, 6}))
	fmt.Println("Result:", productExceptSelf([]int{-1, 0, 1, 2, 3}))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	prefix := make([]int, l)

	prefix[0] = 1
	for i := 1; i < l; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix := make([]int, l)
	suffix[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	result := make([]int, l)
	for i := 0; i < l; i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}
