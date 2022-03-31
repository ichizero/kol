package kol

import "fmt"

// Sequence returns lazily evaluated values.
type Sequence[E comparable] interface {
	// Distinct returns a sequence containing only distinct elements.
	Distinct() Sequence[E]
	// Filter returns a sequence containing only elements matching the given predicate.
	Filter(predicate func(element E) bool) Sequence[E]
	// Map returns a sequence containing the results of applying the given transform function to each element.
	Map(predicate func(element E) E) Sequence[E]
	// Take returns a sequence containing first n elements.
	Take(n int) Sequence[E]
	// Drop returns a sequence containing all elements except first n elements.
	Drop(n int) Sequence[E]
	// ToList evaluate each element and returns it as a List.
	ToList() List[E]
	// ToSlice evaluate each element and returns it as a slice.
	ToSlice() []E
}

type sequence[E comparable] struct {
	seq seq[E]
}

var _ Sequence[int] = (*sequence[int])(nil)

func NewSequence[E comparable](elements ...E) Sequence[E] {
	return &sequence[E]{seq: newIterator(elements)}
}

func newSequence[E comparable](seq seq[E]) *sequence[E] {
	return &sequence[E]{seq: seq}
}

func (s *sequence[E]) Distinct() Sequence[E] {
	return newSequence[E](newDistinctSequence[E](s.seq))
}

func (s *sequence[E]) Filter(predicate func(element E) bool) Sequence[E] {
	return newSequence[E](newFilterSequence[E](s.seq, predicate))
}

func (s *sequence[E]) Map(predicate func(element E) E) Sequence[E] {
	return newSequence[E](newMapSequence[E](s.seq, predicate))
}

func (s *sequence[E]) Take(n int) Sequence[E] {
	return newSequence[E](newTakeSequence[E](s.seq, n))
}

func (s *sequence[E]) Drop(n int) Sequence[E] {
	return newSequence[E](newDropSequence[E](s.seq, n))
}

func (s *sequence[E]) ToSlice() []E {
	res := make([]E, 0)
	for {
		e, ok := s.seq.Next()
		if !ok {
			break
		}
		res = append(res, e)
	}
	return res
}

func (s *sequence[E]) ToList() List[E] {
	return NewList[E](s.ToSlice()...)
}

var _ fmt.Stringer = (*sequence[int])(nil)

func (s *sequence[E]) String() string {
	return s.seq.String()
}

type seq[E comparable] interface {
	fmt.Stringer
	Next() (E, bool)
}

type distinctSequence[E comparable] struct {
	parent seq[E]
	m      map[E]struct{}
}

var _ seq[int] = (*distinctSequence[int])(nil)

func newDistinctSequence[E comparable](parent seq[E]) seq[E] {
	return &distinctSequence[E]{parent: parent, m: map[E]struct{}{}}
}

func (s *distinctSequence[E]) Next() (E, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		if _, exists := s.m[e]; exists {
			continue
		}
		s.m[e] = struct{}{}
		return e, true
	}
	var zero E
	return zero, false
}

func (s *distinctSequence[E]) String() string {
	return fmt.Sprintf("%s > distinct", s.parent)
}

type filterSequence[E comparable] struct {
	parent    seq[E]
	predicate func(element E) bool
}

var _ seq[int] = (*filterSequence[int])(nil)

func newFilterSequence[E comparable](parent seq[E], predicate func(e E) bool) seq[E] {
	return &filterSequence[E]{parent: parent, predicate: predicate}
}

func (s *filterSequence[E]) Next() (E, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		if s.predicate(e) {
			return e, true
		}
	}
	var zero E
	return zero, false
}

func (s *filterSequence[E]) String() string {
	return fmt.Sprintf("%s > filter", s.parent)
}

type mapSequence[E comparable] struct {
	parent    seq[E]
	transform func(element E) E
}

var _ seq[int] = (*mapSequence[int])(nil)

func newMapSequence[E comparable](parent seq[E], transform func(e E) E) seq[E] {
	return &mapSequence[E]{parent: parent, transform: transform}
}

func (s *mapSequence[E]) Next() (E, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		return s.transform(e), true
	}
	var zero E
	return zero, false
}

func (s *mapSequence[E]) String() string {
	return fmt.Sprintf("%s > map", s.parent)
}

type takeSequence[E comparable] struct {
	parent seq[E]
	limit  int
	taken  int
}

var _ seq[int] = (*takeSequence[int])(nil)

func newTakeSequence[E comparable](parent seq[E], n int) seq[E] {
	return &takeSequence[E]{parent: parent, limit: n}
}

func (s *takeSequence[E]) Next() (E, bool) {
	for {
		if s.taken >= s.limit {
			break
		}
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		s.taken++
		return e, true
	}
	var zero E
	return zero, false
}

func (s *takeSequence[E]) String() string {
	return fmt.Sprintf("%s > take %d", s.parent, s.limit)
}

type dropSequence[E comparable] struct {
	parent  seq[E]
	limit   int
	dropped int
}

var _ seq[int] = (*dropSequence[int])(nil)

func newDropSequence[E comparable](parent seq[E], n int) seq[E] {
	return &dropSequence[E]{parent: parent, limit: n}
}

func (s *dropSequence[E]) Next() (E, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		if s.dropped < s.limit {
			s.dropped++
			continue
		}
		return e, true
	}
	var zero E
	return zero, false
}

func (s *dropSequence[E]) String() string {
	return fmt.Sprintf("%s > drop %d", s.parent, s.limit)
}

func MapSequence[E1 comparable, E2 comparable](seq Sequence[E1], predicate func(E1) E2) Sequence[E2] {
	return newSequence[E2](newMapSequenceWithTypeConversion[E1, E2](seq.(*sequence[E1]).seq, predicate))
}

type mapSequenceWithTypeConversion[E1 comparable, E2 comparable] struct {
	parent    seq[E1]
	transform func(element E1) E2
}

var _ seq[int] = (*mapSequenceWithTypeConversion[string, int])(nil)

func newMapSequenceWithTypeConversion[E1 comparable, E2 comparable](parent seq[E1], transform func(e E1) E2) seq[E2] {
	return &mapSequenceWithTypeConversion[E1, E2]{parent: parent, transform: transform}
}

func (s *mapSequenceWithTypeConversion[E1, E2]) Next() (E2, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		return s.transform(e), true
	}
	var zero E2
	return zero, false
}

func (s *mapSequenceWithTypeConversion[E1, E2]) String() string {
	return fmt.Sprintf("%s > map", s.parent)
}
