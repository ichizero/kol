package kol

// Collection is a generic collection of elements.
type Collection[E comparable] interface {
	Iterable[E]

	// Add adds specified elements.
	Add(elements ...E)
	// Clear removes all elements.
	Clear()
	// IsEmpty returns `true` if the collection is empty, `false` otherwise.
	IsEmpty() bool
	// Remove removes specified elements.
	Remove(elements ...E)
	// Retain retains only elements in this collection that are contained in specified elements.
	Retain(elements ...E)
	// Size is size of this collection.
	Size() int
}
