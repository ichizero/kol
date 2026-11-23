package kol

import (
	"math/rand/v2"
	"slices"
)

// List is an ordered collection of elements.
type List[E comparable] interface {
	Collection[E]

	// Drop returns a list containing all elements except first n elements.
	Drop(n uint) Collection[E]
	// DropWhile returns a list containing all elements except first elements that satisfy the given predicate.
	DropWhile(predicate func(element E) bool) Collection[E]
	// ElementAt returns an element at the given index.
	// If the given index is out of range of this collection, it returns `false` as a second return value.
	ElementAt(index int) (E, bool)
	// ElementAtOrElse returns an element at the given index or the result calling of the defaultValue function
	// if the index is out of range of this collection.
	ElementAtOrElse(index int, defaultValue func() E) E
	// FilterIndexed returns a list containing only elements matching the given predicate.
	FilterIndexed(predicate func(idx int, element E) bool) Collection[E]
	// FindLast returns the last element matching the given predicate.
	// If there is no such element, it returns `false` as a second return value.
	FindLast(predicate func(element E) bool) (E, bool)
	// ForEachIndexed performs the given action on each element.
	ForEachIndexed(action func(index int, element E))
	// IndexOf returns an index of the first element matching the given element, or -1 if not preset.
	IndexOf(element E) int
	// IndexOfFirst returns an index of the first element matching the given predicate, or -1 if not present.
	IndexOfFirst(predicate func(element E) bool) int
	// IndexOfLast returns an index of the last element matching the given predicate, or -1 if not present.
	IndexOfLast(predicate func(element E) bool) int
	// MapIndexed returns a list containing the results of applying the given transform function
	// to each element and its index.
	MapIndexed(transform func(idx int, element E) E) Collection[E]
	// Partition splits the original collection in to two lists.
	// If any element returns `true` from the given predicate,
	// it is included in the first list, otherwise it is included in the second list.
	Partition(predicate func(element E) bool) (List[E], List[E])
	// Reversed returns a list with elements in reversed order.
	Reversed() List[E]
	// Shuffled returns a list with elements in shuffled order.
	Shuffled() List[E]
	// Take returns a list containing first n elements.
	Take(n uint) Collection[E]
	// TakeWhile returns a list containing first elements satisfying the given predicate.
	TakeWhile(predicate func(element E) bool) Collection[E]
}

type list[E comparable] struct {
	elements []E
}

func NewList[E comparable](elements ...E) List[E] {
	if len(elements) == 0 {
		elements = []E{}
	}
	return &list[E]{elements: elements}
}

var _ List[int] = (*list[int])(nil)

func (l *list[E]) Add(elements ...E) {
	l.elements = append(l.elements, elements...)
}

func (l *list[E]) Clear() {
	l.elements = []E{}
}

func (l *list[E]) IsEmpty() bool {
	return len(l.elements) == 0
}

func (l *list[E]) Remove(targets ...E) {
	for _, t := range targets {
		if idx := slices.Index(l.elements, t); idx >= 0 {
			l.elements = slices.Delete(l.elements, idx, idx+1)
		}
	}
}

func (l *list[E]) Retain(targets ...E) {
	retained := make([]E, 0, min(len(l.elements), len(targets)))
	if len(targets) > 0 {
		keep := toMap(targets)
		for _, e := range l.elements {
			if _, ok := keep[e]; ok {
				retained = append(retained, e)
			}
		}
	}
	l.elements = retained
}

func (l *list[E]) Size() int {
	return len(l.elements)
}

var _ Iterable[int] = (*list[int])(nil)

func (l *list[E]) All(p func(e E) bool) bool {
	if len(l.elements) == 0 {
		return false
	}
	for _, e := range l.elements {
		if !p(e) {
			return false
		}
	}
	return true
}

func (l *list[E]) Any(p func(e E) bool) bool {
	return slices.ContainsFunc(l.elements, p)
}

func (l *list[E]) Contains(e E) bool {
	return slices.Contains(l.elements, e)
}

func (l *list[E]) Count(p func(e E) bool) int {
	count := 0
	for _, e := range l.elements {
		if p(e) {
			count++
		}
	}
	return count
}

func (l *list[E]) Distinct() Collection[E] {
	seen := make(map[E]struct{}, len(l.elements))
	distinct := make([]E, 0, len(l.elements))
	for _, e := range l.elements {
		if _, ok := seen[e]; !ok {
			seen[e] = struct{}{}
			distinct = append(distinct, e)
		}
	}
	return NewList(distinct...)
}

func (l *list[E]) Drop(n uint) Collection[E] {
	if n >= uint(len(l.elements)) {
		return NewList[E]()
	}
	return NewList(l.elements[n:]...)
}

func (l *list[E]) DropWhile(p func(e E) bool) Collection[E] {
	for i, e := range l.elements {
		if !p(e) {
			return NewList(l.elements[i:]...)
		}
	}
	return NewList[E]()
}

func (l *list[E]) ElementAt(idx int) (E, bool) {
	if idx < 0 || idx >= len(l.elements) {
		var zero E
		return zero, false
	}
	return l.elements[idx], true
}

func (l *list[E]) ElementAtOrElse(idx int, f func() E) E {
	e, ok := l.ElementAt(idx)
	if !ok {
		return f()
	}
	return e
}

func (l *list[E]) Filter(p func(e E) bool) Collection[E] {
	filtered := []E{}
	for _, e := range l.elements {
		if p(e) {
			filtered = append(filtered, e)
		}
	}
	return NewList(filtered...)
}

func (l *list[E]) FilterIndexed(p func(idx int, e E) bool) Collection[E] {
	filtered := []E{}
	for i, e := range l.elements {
		if p(i, e) {
			filtered = append(filtered, e)
		}
	}
	return NewList(filtered...)
}

func (l *list[E]) Find(p func(e E) bool) (E, bool) {
	return l.ElementAt(slices.IndexFunc(l.elements, p))
}

func (l *list[E]) FindLast(p func(e E) bool) (E, bool) {
	return l.ElementAt(l.IndexOfLast(p))
}

func (l *list[E]) ForEach(a func(e E)) {
	for _, e := range l.elements {
		a(e)
	}
}

func (l *list[E]) ForEachIndexed(a func(idx int, e E)) {
	for i, e := range l.elements {
		a(i, e)
	}
}

func (l *list[E]) IndexOf(e E) int {
	return slices.Index(l.elements, e)
}

func (l *list[E]) IndexOfFirst(p func(e E) bool) int {
	return slices.IndexFunc(l.elements, p)
}

func (l *list[E]) IndexOfLast(p func(e E) bool) int {
	for i, e := range slices.Backward(l.elements) {
		if p(e) {
			return i
		}
	}
	return -1
}

func (l *list[E]) Intersect(other Iterable[E]) Set[E] {
	return intersect(toMap(l.elements), other)
}

func (l *list[E]) Map(t func(e E) E) Collection[E] {
	mapped := make([]E, len(l.elements))
	for i, e := range l.elements {
		mapped[i] = t(e)
	}
	return NewList(mapped...)
}

func (l *list[E]) MapIndexed(t func(idx int, e E) E) Collection[E] {
	mapped := make([]E, len(l.elements))
	for i, e := range l.elements {
		mapped[i] = t(i, e)
	}
	return NewList(mapped...)
}

func (l *list[E]) Minus(e ...E) Collection[E] {
	cloned := l.ToList()
	cloned.Remove(e...)
	return cloned
}

func (l *list[E]) None(p func(e E) bool) bool {
	return !slices.ContainsFunc(l.elements, p)
}

func (l *list[E]) Partition(p func(e E) bool) (List[E], List[E]) {
	first := []E{}
	second := []E{}
	for _, e := range l.elements {
		if p(e) {
			first = append(first, e)
		} else {
			second = append(second, e)
		}
	}
	return NewList(first...), NewList(second...)
}

func (l *list[E]) Plus(e ...E) Collection[E] {
	return NewList(slices.Concat(l.elements, e)...)
}

func (l *list[E]) Reversed() List[E] {
	cloned := slices.Clone(l.elements)
	slices.Reverse(cloned)
	return NewList(cloned...)
}

func (l *list[E]) Shuffled() List[E] {
	cloned := slices.Clone(l.elements)
	rand.Shuffle(len(cloned), func(i, j int) { //nolint:gosec // Shuffling needs no cryptographic randomness.
		cloned[i], cloned[j] = cloned[j], cloned[i]
	})
	return NewList(cloned...)
}

func (l *list[E]) Single(p func(e E) bool) (E, bool) {
	return single(slices.Values(l.elements), p)
}

func (l *list[E]) Subtract(other Iterable[E]) Set[E] {
	return subtract(toMap(l.elements), other)
}

func (l *list[E]) Take(n uint) Collection[E] {
	return NewList(l.elements[:min(n, uint(len(l.elements)))]...)
}

func (l *list[E]) TakeWhile(p func(e E) bool) Collection[E] {
	for i, e := range l.elements {
		if !p(e) {
			return NewList(l.elements[:i]...)
		}
	}
	return NewList(l.elements...)
}

func (l *list[E]) ToList() List[E] {
	return NewList(slices.Clone(l.elements)...)
}

func (l *list[E]) ToSet() Set[E] {
	return NewSet(l.elements...)
}

func (l *list[E]) ToSlice() []E {
	return slices.Clone(l.elements)
}

func (l *list[E]) Union(other Iterable[E]) Set[E] {
	return union(toMap(l.elements), other)
}

func MapList[E1 comparable, E2 comparable](collection Collection[E1], transform func(E1) E2) List[E2] {
	result := make([]E2, 0, collection.Size())

	collection.ForEach(func(e1 E1) {
		result = append(result, transform(e1))
	})

	return NewList(result...)
}
