package main

func main() {
	println(longestConsecutive([]int{2, 20, 4, 10, 3, 4, 5}))
	println(longestConsecutive([]int{0, 3, 2, 5, 4, 6, 1, 1}))
	println(longestConsecutive([]int{5, 10, 15, 20, 25}))
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1

	store := make(map[int]struct{})
	for _, n := range nums {
		store[n] = struct{}{}
	}

	for _, n := range nums {
		streak := 0
		cur := n
		_, ok := store[cur]
		for ok {
			streak++
			cur++
			res = max(res, streak)
			_, ok = store[cur]
		}
	}

	return res
}
