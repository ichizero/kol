package kol

type Collection[E comparable] interface {
	Iterable[E]

	Add(elements ...E)
	Clear()
	IsEmpty() bool
	Remove(elements ...E)
	Retain(elements ...E)
	Size() int
}
