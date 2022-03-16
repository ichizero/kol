package kol

import (
	"fmt"
)

func ExampleList() {
	l1 := NewList(1, 2, 3)
	l2 := NewList(2, 3, 4).
		MapIndexed(func(idx int, e int) int {
			return idx + e
		})
	fmt.Println(l1.Subtract(l2).ToSlice())
}

func ExampleSet() {
	s1 := NewSet(1, 2, 3, 3, 3)
	s2 := NewSet(2, 3, 4).
		Map(func(e int) int {
			return e + 1
		})
	fmt.Println(s1.Union(s2).ToSlice())
}

func ExampleSequence() {
	seq := NewSequence(1, 2, 3, 4, 5).
		Filter(func(e int) bool {
			return e < 3
		}).
		Map(func(e int) int {
			return e * 2
		}).
		Distinct().
		Drop(1).
		Take(2)
	fmt.Println(seq.ToSlice())
}
