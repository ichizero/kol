package kol

import (
	"testing"
)

const benchSize = 1000

// benchInts returns n ints starting at offset, where every value appears twice.
func benchInts(n, offset int) []int {
	elems := make([]int, n)
	for i := range elems {
		elems[i] = offset + i/2
	}
	return elems
}

// benchUniqueInts returns n distinct ints starting at offset.
func benchUniqueInts(n, offset int) []int {
	elems := make([]int, n)
	for i := range elems {
		elems[i] = offset + i
	}
	return elems
}

func isEven(e int) bool { return e%2 == 0 }

func double(e int) int { return e * 2 }

func BenchmarkList(b *testing.B) {
	list := NewList(benchInts(benchSize, 0)...)
	other := NewList(benchInts(benchSize, benchSize/4)...)
	targets := benchInts(benchSize/10, 0)

	b.Run("Filter", func(b *testing.B) {
		for b.Loop() {
			list.Filter(isEven)
		}
	})
	b.Run("Map", func(b *testing.B) {
		for b.Loop() {
			list.Map(double)
		}
	})
	b.Run("Distinct", func(b *testing.B) {
		for b.Loop() {
			list.Distinct()
		}
	})
	b.Run("Partition", func(b *testing.B) {
		for b.Loop() {
			list.Partition(isEven)
		}
	})
	b.Run("Reversed", func(b *testing.B) {
		for b.Loop() {
			list.Reversed()
		}
	})
	b.Run("Retain", func(b *testing.B) {
		for b.Loop() {
			l := list.ToList()
			l.Retain(targets...)
		}
	})
	b.Run("Intersect", func(b *testing.B) {
		for b.Loop() {
			list.Intersect(other)
		}
	})
	b.Run("Union", func(b *testing.B) {
		for b.Loop() {
			list.Union(other)
		}
	})
	b.Run("MapList", func(b *testing.B) {
		for b.Loop() {
			MapList[int, int](list, double)
		}
	})
}

func BenchmarkSet(b *testing.B) {
	elems := benchUniqueInts(benchSize, 0)
	set := NewSet(elems...)
	other := NewSet(benchUniqueInts(benchSize, benchSize/2)...)
	targets := benchUniqueInts(benchSize/10, 0)

	b.Run("NewSet", func(b *testing.B) {
		for b.Loop() {
			NewSet(elems...)
		}
	})
	b.Run("Filter", func(b *testing.B) {
		for b.Loop() {
			set.Filter(isEven)
		}
	})
	b.Run("Map", func(b *testing.B) {
		for b.Loop() {
			set.Map(double)
		}
	})
	b.Run("Retain", func(b *testing.B) {
		for b.Loop() {
			s := set.ToSet()
			s.Retain(targets...)
		}
	})
	b.Run("Intersect", func(b *testing.B) {
		for b.Loop() {
			set.Intersect(other)
		}
	})
	b.Run("Subtract", func(b *testing.B) {
		for b.Loop() {
			set.Subtract(other)
		}
	})
	b.Run("Union", func(b *testing.B) {
		for b.Loop() {
			set.Union(other)
		}
	})
	b.Run("ToSlice", func(b *testing.B) {
		for b.Loop() {
			set.ToSlice()
		}
	})
	b.Run("MapSet", func(b *testing.B) {
		for b.Loop() {
			MapSet[int, int](set, double)
		}
	})
}

func BenchmarkSequence(b *testing.B) {
	elems := benchInts(benchSize, 0)

	b.Run("FilterMap", func(b *testing.B) {
		for b.Loop() {
			NewSequence(elems...).
				Filter(isEven).
				Map(double).
				ToSlice()
		}
	})
	b.Run("Chain", func(b *testing.B) {
		for b.Loop() {
			NewSequence(elems...).
				Filter(isEven).
				Map(double).
				Distinct().
				Drop(10).
				Take(100).
				ToSlice()
		}
	})
	b.Run("MapSequence", func(b *testing.B) {
		for b.Loop() {
			MapSequence(NewSequence(elems...), double).ToSlice()
		}
	})
}
