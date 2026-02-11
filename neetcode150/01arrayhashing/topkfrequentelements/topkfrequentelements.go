package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 2, 3, 3, 3}
	k := 3
	fmt.Println(topKFrequent(a, k))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)

	// iterate through nums and put in map with counter
	for _, n := range nums {
		m[n] = m[n] + 1
	}

	// sort map based on the value i.e. counter
	// get keys into an array
	ka := make([]int, 0)
	for k, _ := range m {
		ka = append(ka, k)
	}

	// sort keys array based on the value inside the map
	sort.Slice(ka, func(i, j int) bool {
		return m[ka[i]] > m[ka[j]]
	})

	// prepare result based after sorting
	result := make([]int, 0, k)
	cnt := 0
	for cnt < k {
		result = append(result, ka[cnt])
		cnt++
	}

	return result
}
