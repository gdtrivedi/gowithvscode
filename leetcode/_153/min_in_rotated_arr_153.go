package main

func main() {
	MinInRotatedArray()
}

func MinInRotatedArray() {
	println(findMinInRotatedArray([]int{4, 5, 1, 2, 3}))
}

func findMinInRotatedArray(arr []int) int {
	res := arr[0]

	len := len(arr)
	l := 0
	r := len - 1

	for l <= r {
		if arr[l] < arr[r] {
			res = min(res, arr[l])
			break
		}

		m := (l + r) / 2
		res = min(res, arr[m])
		if arr[m] >= arr[l] {
			l = m + 1
		} else {
			r = m - 1
		}
		m = (l + r) / 2
	}

	return res
}
