package main

func main() {
	println([]int{1, 2, 3, 4}, 3)
}

func twoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1

	res := make([]int, 2)
	for lp < rp {
		sum := numbers[lp] + numbers[rp]
		if sum == target {
			res = []int{lp + 1, rp + 1}
			break
		}

		if sum < target {
			lp++
		}

		if sum > target {
			rp--
		}
	}
	return res
}
