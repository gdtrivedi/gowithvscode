package main

import (
	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	ts := treeset.NewWithIntComparator()
	ts.Add(1)

	for _, v := range ts.Values() {
		println(v.(int))
	}
}
