package kol

type Sequence[E comparable] interface {
	Distinct() Sequence[E]
	Filter(predicate func(element E) bool) Sequence[E]
	Map(predicate func(element E) E) Sequence[E]
	Take(n int) Sequence[E]
	Drop(n int) Sequence[E]
	ToList() List[E]
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
	var res = make([]E, 0)
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

type seq[E comparable] interface {
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

type mapSequence[E comparable] struct {
	parent    seq[E]
	predicate func(element E) E
}

var _ seq[int] = (*mapSequence[int])(nil)

func newMapSequence[E comparable](parent seq[E], predicate func(e E) E) seq[E] {
	return &mapSequence[E]{parent: parent, predicate: predicate}
}

func (s *mapSequence[E]) Next() (E, bool) {
	for {
		e, ok := s.parent.Next()
		if !ok {
			break
		}
		return s.predicate(e), true
	}
	var zero E
	return zero, false
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
