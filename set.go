package kol

import (
	"iter"
	"maps"
)

// Set is an un-ordered collection of elements without duplicate elements.
type Set[E comparable] interface {
	Collection[E]
}

type set[E comparable] struct {
	m map[E]struct{}
}

func NewSet[E comparable](elements ...E) Set[E] {
	return &set[E]{m: toMap(elements)}
}

// toMap returns a new map whose keys are the given elements.
func toMap[E comparable](elements []E) map[E]struct{} {
	m := make(map[E]struct{}, len(elements))
	for _, e := range elements {
		m[e] = struct{}{}
	}
	return m
}

// intersect returns a set containing keys of m that are contained by other.
func intersect[E comparable](m map[E]struct{}, other Iterable[E]) Set[E] {
	if o, ok := other.(*set[E]); ok {
		// Cloning a map is cheaper than building one element by element.
		res := maps.Clone(o.m)
		maps.DeleteFunc(res, func(e E, _ struct{}) bool {
			_, ok := m[e]
			return !ok
		})
		return &set[E]{m: res}
	}

	res := make(map[E]struct{})
	other.ForEach(func(e E) {
		if _, ok := m[e]; ok {
			res[e] = struct{}{}
		}
	})
	return &set[E]{m: res}
}

// subtract removes all elements of other from m and returns m as a set.
func subtract[E comparable](m map[E]struct{}, other Iterable[E]) Set[E] {
	other.ForEach(func(e E) {
		delete(m, e)
	})
	return &set[E]{m: m}
}

// union adds all elements of other to m and returns m as a set.
func union[E comparable](m map[E]struct{}, other Iterable[E]) Set[E] {
	other.ForEach(func(e E) {
		m[e] = struct{}{}
	})
	return &set[E]{m: m}
}

// single returns the single element of seq matching the given predicate.
func single[E comparable](seq iter.Seq[E], p func(e E) bool) (E, bool) {
	var (
		res   E
		found bool
	)
	for e := range seq {
		if !p(e) {
			continue
		}
		if found {
			var zero E
			return zero, false
		}
		res, found = e, true
	}
	return res, found
}

var _ Collection[int] = (*set[int])(nil)

func (s *set[E]) Add(elements ...E) {
	for _, e := range elements {
		s.m[e] = struct{}{}
	}
}

func (s *set[E]) Clear() {
	clear(s.m)
}

func (s *set[E]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *set[E]) Remove(targets ...E) {
	for _, t := range targets {
		delete(s.m, t)
	}
}

func (s *set[E]) Retain(targets ...E) {
	keep := toMap(targets)
	maps.DeleteFunc(s.m, func(e E, _ struct{}) bool {
		_, ok := keep[e]
		return !ok
	})
}

func (s *set[E]) Size() int {
	return len(s.m)
}

var _ Iterable[int] = (*set[int])(nil)

func (s *set[E]) All(p func(e E) bool) bool {
	if len(s.m) == 0 {
		return false
	}
	for e := range s.m {
		if !p(e) {
			return false
		}
	}
	return true
}

func (s *set[E]) Any(p func(e E) bool) bool {
	for e := range s.m {
		if p(e) {
			return true
		}
	}
	return false
}

func (s *set[E]) Contains(e E) bool {
	_, ok := s.m[e]
	return ok
}

func (s *set[E]) Count(p func(e E) bool) int {
	count := 0
	for e := range s.m {
		if p(e) {
			count++
		}
	}
	return count
}

func (s *set[E]) Distinct() Collection[E] {
	return s
}

func (s *set[E]) Find(p func(e E) bool) (E, bool) {
	for e := range s.m {
		if p(e) {
			return e, true
		}
	}
	var zero E
	return zero, false
}

func (s *set[E]) Filter(p func(e E) bool) Collection[E] {
	filtered := make(map[E]struct{})
	for e := range s.m {
		if p(e) {
			filtered[e] = struct{}{}
		}
	}
	return &set[E]{m: filtered}
}

func (s *set[E]) ForEach(a func(e E)) {
	for e := range s.m {
		a(e)
	}
}

func (s *set[E]) Intersect(other Iterable[E]) Set[E] {
	return intersect(s.m, other)
}

func (s *set[E]) Map(t func(e E) E) Collection[E] {
	mapped := make(map[E]struct{}, len(s.m))
	for e := range s.m {
		mapped[t(e)] = struct{}{}
	}
	return &set[E]{m: mapped}
}

func (s *set[E]) Minus(e ...E) Collection[E] {
	cloned := s.ToSet()
	cloned.Remove(e...)
	return cloned
}

func (s *set[E]) None(p func(e E) bool) bool {
	return !s.Any(p)
}

func (s *set[E]) Plus(e ...E) Collection[E] {
	cloned := s.ToSet()
	cloned.Add(e...)
	return cloned
}

func (s *set[E]) Single(p func(e E) bool) (E, bool) {
	return single(maps.Keys(s.m), p)
}

func (s *set[E]) Subtract(other Iterable[E]) Set[E] {
	return subtract(maps.Clone(s.m), other)
}

func (s *set[E]) ToList() List[E] {
	return NewList(s.ToSlice()...)
}

func (s *set[E]) ToSet() Set[E] {
	return &set[E]{m: maps.Clone(s.m)}
}

func (s *set[E]) ToSlice() []E {
	elements := make([]E, 0, len(s.m))
	for e := range s.m {
		elements = append(elements, e)
	}
	return elements
}

func (s *set[E]) Union(other Iterable[E]) Set[E] {
	return union(maps.Clone(s.m), other)
}

func MapSet[E1 comparable, E2 comparable](collection Collection[E1], transform func(E1) E2) Set[E2] {
	result := make(map[E2]struct{}, collection.Size())

	collection.ForEach(func(e1 E1) {
		result[transform(e1)] = struct{}{}
	})

	return &set[E2]{m: result}
}
