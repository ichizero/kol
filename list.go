package kol

import (
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

type List[E comparable] interface {
	Collection[E]
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

var _ Collection[int] = (*list[int])(nil)

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

func (l *list[E]) All(predicate func(element E) bool) bool {
	if l.Size() == 0 {
		return false
	}
	for _, s := range l.elements {
		if !predicate(s) {
			return false
		}
	}
	return true
}

func (l *list[E]) Any(predicate func(element E) bool) bool {
	for _, s := range l.elements {
		if predicate(s) {
			return true
		}
	}
	return false
}

func (l *list[E]) Contains(element E) bool {
	return slices.Contains(l.elements, element)
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

func (l *list[E]) Distinct() List[E] {
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

func (l *list[E]) Drop(n uint) List[E] {
	if int(n) >= l.Size() {
		return NewList[E]()
	}
	return NewList[E](l.elements[n:]...)
}

func (l *list[E]) DropWhile(p func(e E) bool) List[E] {
	for i, e := range l.elements {
		if !p(e) {
			return NewList(l.elements[i:]...)
		}
	}
	return NewList[E]()
}

func (l *list[E]) ElementAt(idx int) (*E, error) {
	if idx < 0 || idx >= l.Size() {
		return nil, ErrIndexOutOfRange
	}
	return &l.elements[idx], nil
}

func (l *list[E]) ElementAtOrElse(idx int, f func() E) E {
	e, err := l.ElementAt(idx)
	if err != nil {
		return f()
	}
	return *e
}

func (l *list[E]) Filter(p func(e E) bool) List[E] {
	return l.FilterIndexed(func(_ int, e E) bool {
		return p(e)
	})
}

func (l *list[E]) FilterIndexed(p func(idx int, e E) bool) List[E] {
	var filtered = make([]E, 0)
	for i, e := range l.elements {
		if p(i, e) {
			filtered = append(filtered, e)
		}
	}
	return NewList(filtered...)
}

func (l *list[E]) Find(p func(e E) bool) *E {
	if idx := l.IndexOfFirst(p); idx == -1 {
		return nil
	} else {
		return &l.elements[idx]
	}
}

func (l *list[E]) FindLast(p func(e E) bool) *E {
	if idx := l.IndexOfLast(p); idx == -1 {
		return nil
	} else {
		return &l.elements[idx]
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

func (l *list[E]) Intersect() Set[E] {
	panic("not implemented")
}

func (l *list[E]) Iterator() Iterator[E] {
	panic("not implemented")
}

func (l *list[E]) Minus(e E) List[E] {
	cloned := NewList(slices.Clone(l.elements)...)
	cloned.Remove(e)
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

func (l *list[E]) Plus(e E) List[E] {
	return NewList(append(slices.Clone(l.elements), e)...)
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

func (l *list[E]) Single(p func(e E) bool) *E {
	var found = false
	var res *E
	for _, e := range l.elements {
		e := e
		if p(e) {
			if found {
				return nil
			}
			res = &e
			found = true
		}
	}
	return res
}

func (l *list[E]) Subtract() Set[E] {
	panic("not implemented")
}

func (l *list[E]) Take(n uint) List[E] {
	max := uint(l.Size())
	if max < n {
		n = max
	}
	return NewList(l.elements[:n]...)
}

func (l *list[E]) TakeWhile(p func(e E) bool) List[E] {
	for i, e := range l.elements {
		if !p(e) {
			return NewList(l.elements[:i]...)
		}
	}
	return NewList(l.elements...)
}

func (l *list[E]) ToList() List[E] {
	return l
}

func (l *list[E]) ToSlice() []E {
	return l.elements
}

func (l *list[E]) Union() Set[E] {
	panic("not implemented")
}
