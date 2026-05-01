package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"sort"
)

func main() {
	a := []int{1, 2, 2, 3, 3, 3, 7, 9, 8, 5, 5, 7, 8, 9, 7, 8, 9, 9}
	k := 3
	fmt.Println(topKFrequent(a, k))
	fmt.Println(topKFrequent1(a, k))
}

func topKFrequent1(nums []int, k int) []int {
	tm := treemap.NewWithIntComparator() // empty (keys are of type int)

	m := make(map[int]int, 0)

	// iterate through nums and put in map with counter
	for _, n := range nums {
		m[n] = m[n] + 1
	}

	for k, v := range m {
		tm.Put(v, k)
	}

	newk := min(tm.Size(), k)
	if newk == 0 {
		return []int{}
	}

	res := []int{}
	it := tm.Iterator()
	for it.End(); it.Prev(); {
		res = append(res, it.Value().(int))
		if newk == len(res) {
			break
		}
	}
	return res
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
