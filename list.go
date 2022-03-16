package kol

import (
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

type List[E comparable] interface {
	Collection[E]

	Drop(n uint) Collection[E]
	DropWhile(predicate func(element E) bool) Collection[E]
	ElementAt(index int) (E, bool)
	ElementAtOrElse(index int, defaultValue func() E) E
	FilterIndexed(predicate func(idx int, element E) bool) Collection[E]
	FindLast(predicate func(element E) bool) (E, bool)
	ForEachIndexed(action func(index int, element E))
	IndexOf(element E) int
	IndexOfFirst(predicate func(element E) bool) int
	IndexOfLast(predicate func(element E) bool) int
	MapIndexed(predicate func(idx int, element E) E) Collection[E]
	Partition(predicate func(element E) bool) (List[E], List[E])
	Reversed() List[E]
	Shuffled() List[E]
	Take(n uint) Collection[E]
	TakeWhile(predicate func(element E) bool) Collection[E]
}

type list[E comparable] struct {
	elements []E
}

func NewList[E comparable](elements ...E) List[E] {
	if len(elements) == 0 {
		elements = make([]E, 0)
	}
	return &list[E]{
		elements: elements,
	}
}

func (l *list[E]) clone() List[E] {
	return &list[E]{
		elements: slices.Clone(l.elements),
	}
}

var _ List[int] = (*list[int])(nil)

func (l *list[E]) Add(elements ...E) {
	l.elements = append(l.elements, elements...)
}

func (l *list[E]) Clear() {
	l.elements = []E{}
}

func (l *list[E]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *list[E]) Remove(targets ...E) {
	for _, t := range targets {
		if idx := slices.Index(l.elements, t); idx >= 0 {
			l.elements = slices.Delete(l.elements, idx, idx+1)
		}
	}
}

func (l *list[E]) Retain(targets ...E) {
	var retained = make([]E, 0, len(targets))
	for _, e := range l.elements {
		if slices.Contains(targets, e) {
			retained = append(retained, e)
		}
	}
	l.elements = retained
}

func (l *list[E]) Size() int {
	return len(l.elements)
}

var _ Iterable[int] = (*list[int])(nil)

func (l *list[E]) All(p func(e E) bool) bool {
	if l.Size() == 0 {
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
	for _, e := range l.elements {
		if p(e) {
			return true
		}
	}
	return false
}

func (l *list[E]) Contains(e E) bool {
	return slices.Contains(l.elements, e)
}

func (l *list[E]) Count(p func(e E) bool) int {
	var count = 0
	for _, e := range l.elements {
		if p(e) {
			count++
		}
	}
	return count
}

func (l *list[E]) Distinct() Collection[E] {
	size := l.Size()
	if size == 0 {
		return NewList[E]()
	}
	existence := make(map[E]interface{}, size)
	filtered := make([]E, 0, size)
	for _, e := range l.elements {
		if _, ok := existence[e]; ok {
			continue // duplicated
		}
		existence[e] = struct{}{}
		filtered = append(filtered, e)
	}
	return NewList(filtered...)
}

func (l *list[E]) Drop(n uint) Collection[E] {
	if int(n) >= l.Size() {
		return NewList[E]()
	}
	return NewList[E](l.elements[n:]...)
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
	if idx < 0 || idx >= l.Size() {
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
	return l.FilterIndexed(func(_ int, e E) bool {
		return p(e)
	})
}

func (l *list[E]) FilterIndexed(p func(idx int, e E) bool) Collection[E] {
	var filtered = make([]E, 0)
	l.ForEachIndexed(func(idx int, e E) {
		if p(idx, e) {
			filtered = append(filtered, e)
		}
	})
	return NewList(filtered...)
}

func (l *list[E]) Find(p func(e E) bool) (E, bool) {
	if idx := l.IndexOfFirst(p); idx == -1 {
		var zero E
		return zero, false
	} else {
		return l.elements[idx], true
	}
}

func (l *list[E]) FindLast(p func(e E) bool) (E, bool) {
	if idx := l.IndexOfLast(p); idx == -1 {
		var zero E
		return zero, false
	} else {
		return l.elements[idx], true
	}
}

func (l *list[E]) ForEach(a func(e E)) {
	l.ForEachIndexed(func(_ int, e E) {
		a(e)
	})
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
	for i, e := range l.elements {
		if p(e) {
			return i
		}
	}
	return -1
}

func (l *list[E]) IndexOfLast(p func(e E) bool) int {
	for i := l.Size() - 1; i >= 0; i-- {
		if p(l.elements[i]) {
			return i
		}
	}
	return -1
}

func (l *list[E]) Intersect(other Iterable[E]) Set[E] {
	return l.ToSet().Intersect(other)
}

func (l *list[E]) Iterator() Iterator[E] {
	panic("not implemented")
}

func (l *list[E]) Map(p func(e E) E) Collection[E] {
	return l.MapIndexed(func(_ int, e E) E {
		return p(e)
	})
}

func (l *list[E]) MapIndexed(p func(idx int, e E) E) Collection[E] {
	var mapped = make([]E, 0)
	l.ForEachIndexed(func(idx int, e E) {
		mapped = append(mapped, p(idx, e))
	})
	return NewList(mapped...)
}

func (l *list[E]) Minus(e ...E) Collection[E] {
	cloned := NewList(slices.Clone(l.elements)...)
	cloned.Remove(e...)
	return cloned
}

func (l *list[E]) None(p func(e E) bool) bool {
	for _, e := range l.elements {
		if p(e) {
			return false
		}
	}
	return true
}

func (l *list[E]) Partition(p func(e E) bool) (List[E], List[E]) {
	var first = make([]E, 0)
	var second = make([]E, 0)
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
	return NewList(append(slices.Clone(l.elements), e...)...)
}

func (l *list[E]) Reversed() List[E] {
	cloned := slices.Clone(l.elements)
	size := len(cloned)

	for i := 0; i < size/2; i = i + 1 {
		j := size - 1 - i
		cloned[i], cloned[j] = cloned[j], cloned[i]
	}
	return NewList(cloned...)
}

func (l *list[E]) Shuffled() List[E] {
	cloned := slices.Clone(l.elements)
	rand.Shuffle(len(cloned), func(i, j int) {
		cloned[i], cloned[j] = cloned[j], cloned[i]
	})
	return NewList(cloned...)
}

func (l *list[E]) Single(p func(e E) bool) (E, bool) {
	var found = false
	var res E
	for _, e := range l.elements {
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

func (l *list[E]) Subtract(other Iterable[E]) Set[E] {
	return l.ToSet().Intersect(other)
}

func (l *list[E]) Take(n uint) Collection[E] {
	max := uint(l.Size())
	if max < n {
		n = max
	}
	return NewList(l.elements[:n]...)
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
	return l.clone()
}

func (l *list[E]) ToSet() Set[E] {
	return NewSet(l.elements...)
}

func (l *list[E]) ToSlice() []E {
	return slices.Clone(l.elements)
}

func (l *list[E]) Union(other Iterable[E]) Set[E] {
	return l.ToSet().Plus(other.ToSlice()...)

}

func MapList[E1 comparable, E2 comparable](collection Collection[E1], predicate func(E1) E2) List[E2] {
	result := make([]E2, 0, collection.Size())

	collection.ForEach(func(e1 E1) {
		result = append(result, predicate(e1))
	})

	return NewList(result...)
}
