package kol

import "fmt"

type Iterator[E comparable] interface {
	HasNext() bool
	Next() (E, bool)
}

type iterator[E comparable] struct {
	elements []E
	cursor   int
}

var _ Iterator[int] = (*iterator[int])(nil)

func newIterator[E comparable](elements []E) *iterator[E] {
	return &iterator[E]{elements: elements}
}

func (i *iterator[E]) HasNext() bool {
	return i.cursor < len(i.elements)
}

var _ seq[int] = (*iterator[int])(nil)

func (i *iterator[E]) Next() (E, bool) {
	if i.cursor < len(i.elements) {
		e := i.elements[i.cursor]
		i.cursor++
		return e, true
	}
	var zero E
	return zero, false
}

func (i *iterator[E]) String() string {
	return fmt.Sprintf("cursor: %d, elements: %v", i.cursor, i.elements)
}
