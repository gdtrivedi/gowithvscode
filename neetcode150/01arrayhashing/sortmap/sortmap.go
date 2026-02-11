package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	sortMapByKeys()
	sortMapByValues()
	sortMapByValueAndKey()
}

func sortMapByKeys() {
	fruits := map[string]int{
		"apple":  3,
		"cherry": 2,
		"banana": 1,
		"date":   3,
	}

	keys := make([]string, 0, len(fruits))

	for k := range fruits {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, fruits[k])
	}
}

func sortMapByValues() {
	fruits := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
		"date":   3,
	}

	type kv struct {
		key   string
		value int
	}

	kvarr := make([]kv, 0, len(fruits))
	for k, v := range fruits {
		kvarr = append(kvarr, kv{key: k, value: v})
	}

	sort.Slice(kvarr, func(i, j int) bool {
		return kvarr[i].value > kvarr[j].value
	})

	for _, kva := range kvarr {
		fmt.Println(kva.key, kva.value)
	}
}

func sortMapByValueAndKey() {
	fruits := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
		"date":   3,
	}

	type kv struct {
		key   string
		value int
	}

	kvarr := make([]kv, 0, len(fruits))
	for k, v := range fruits {
		kvarr = append(kvarr, kv{key: k, value: v})
	}

	slices.SortFunc(kvarr, func(a, b kv) int {
		//1. sort by value - descending
		if a.value != b.value {
			return b.value - a.value
		}

		//2. sort by key - ascending (when values match)
		if a.key < b.key {
			return -1
		}
		if a.key > b.key {
			return 1
		}

		return 0
	})

	for _, kva := range kvarr {
		fmt.Println(kva.key, kva.value)
	}
}
