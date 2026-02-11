package main

import "fmt"

func main() {
	a := []int{3, 4, 5, 6}
	t := 7
	fmt.Println(twoSum(a, t))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i, n := range nums {
		sub := target - n
		if val, ok := m[sub]; ok {
			return []int{val, i}
		}
		m[n] = i
	}

	return []int{-1, -1}
}
