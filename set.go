package kol

import (
	"golang.org/x/exp/maps" // nolint:typecheck
)

// Set is an un-ordered collection of elements without duplicate elements.
type Set[E comparable] interface {
	Collection[E]
}

type set[E comparable] struct {
	m map[E]struct{}
}

func NewSet[E comparable](elements ...E) Set[E] {
	m := make(map[E]struct{}, 0)
	for _, e := range elements {
		m[e] = struct{}{}
	}
	return &set[E]{m: m}
}

func newSet[E comparable](m map[E]struct{}) Set[E] {
	return &set[E]{m: m}
}

func (s *set[E]) clone() Set[E] {
	return newSet(maps.Clone(s.m))
}

var _ Collection[int] = (*set[int])(nil)

func (s *set[E]) Add(elements ...E) {
	for _, e := range elements {
		s.m[e] = struct{}{}
	}
}

func (s *set[E]) Clear() {
	maps.Clear(s.m)
}

func (s *set[E]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *set[E]) Remove(targets ...E) {
	for _, t := range targets {
		delete(s.m, t)
	}
}

func (s *set[E]) Retain(targets ...E) {
	tl := NewList(targets...)
	for e := range s.m {
		if !tl.Contains(e) {
			delete(s.m, e)
		}
	}
}

func (s *set[E]) Size() int {
	return len(s.m)
}

var _ Iterable[int] = (*set[int])(nil)

func (s *set[E]) All(p func(e E) bool) bool {
	if s.Size() == 0 {
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
	filtered := make(map[E]struct{}, 0)
	s.ForEach(func(e E) {
		if p(e) {
			filtered[e] = struct{}{}
		}
	})
	return newSet(filtered)
}

func (s *set[E]) ForEach(a func(e E)) {
	for e := range s.m {
		a(e)
	}
}

func (s *set[E]) Intersect(other Iterable[E]) Set[E] {
	res := other.ToSet()
	for e := range s.m {
		if _, ok := res.(*set[E]).m[e]; !ok {
			res.Remove(e)
		}
	}
	return res
}

func (s *set[E]) Map(t func(e E) E) Collection[E] {
	mapped := make(map[E]struct{}, 0)
	s.ForEach(func(e E) {
		mapped[t(e)] = struct{}{}
	})
	return newSet(mapped)
}

func (s *set[E]) Minus(e ...E) Collection[E] {
	cloned := s.clone()
	cloned.Remove(e...)
	return cloned
}

func (s *set[E]) None(p func(e E) bool) bool {
	for e := range s.m {
		if p(e) {
			return false
		}
	}
	return true
}

func (s *set[E]) Plus(e ...E) Collection[E] {
	cloned := s.clone()
	cloned.Add(e...)
	return cloned
}

func (s *set[E]) Single(p func(e E) bool) (E, bool) {
	found := false
	var res E
	for e := range s.m {
		if p(e) {
			if found {
				var zero E
				return zero, false
			}
			res = e
			found = true
		}
	}
	return res, true
}

func (s *set[E]) Subtract(other Iterable[E]) Set[E] {
	res := s.clone()
	for _, e := range other.ToSlice() {
		if _, ok := res.(*set[E]).m[e]; ok {
			res.Remove(e)
		}
	}
	return res
}

func (s *set[E]) ToList() List[E] {
	return NewList(maps.Keys(s.m)...)
}

func (s *set[E]) ToSet() Set[E] {
	return s.clone()
}

func (s *set[E]) ToSlice() []E {
	return maps.Keys(s.m)
}

func (s *set[E]) Union(other Iterable[E]) Set[E] {
	return s.clone().Plus(other.ToSlice()...)
}

func MapSet[E1 comparable, E2 comparable](collection Collection[E1], transform func(E1) E2) Set[E2] {
	result := make([]E2, 0, collection.Size())

	collection.ForEach(func(e1 E1) {
		result = append(result, transform(e1))
	})

	return NewSet(result...)
}
