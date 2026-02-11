package main

func main() {
	SearchInRotatedArray()
}

func SearchInRotatedArray() {
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) //output 4
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) //output -1
	println(search([]int{4}, 0))                   //output -1
}

func search(nums []int, target int) int {
	len := len(nums)
	l, r := 0, len-1

	res := -1
	if len == 1 && nums[0] != target {
		return res
	}

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			// left sorted portion
			if target > nums[m] || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			// right sorted portion
			if target < nums[m] || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
