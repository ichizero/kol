package kol

type Iterable[E comparable] interface {
	// All returns `true` if all elements match the given predicate.
	All(predicate func(element E) bool) bool
	// Any returns `true` if collection has at least one element matched the given predicate.
	Any(predicate func(element E) bool) bool
	// Contains returns `true` if the given element is found in the collection.
	Contains(element E) bool
	// Count returns the number of elements that matches the given predicate.
	Count(predicate func(element E) bool) int
	// Distinct returns a collection containing only distinct elements.
	Distinct() Collection[E]
	// Filter returns a collection only elements matching the given predicate.
	Filter(predicate func(element E) bool) Collection[E]
	// Find returns the first element matching the given predicate.
	// If there is no such element, it returns `false` as a second return value.
	Find(predicate func(element E) bool) (E, bool)
	// ForEach performs the given action on each element.
	ForEach(action func(element E))
	// Intersect returns a set containing all elements that are contained
	// by both this collection and the specified collection.
	Intersect(other Iterable[E]) Set[E]
	// Map returns a collection containing the result of applying the given transform function to each element.
	Map(transform func(element E) E) Collection[E]
	// Minus returns a collection containing all elements of the original collection except given elements.
	Minus(element ...E) Collection[E]
	// None returns `true` if no elements match the given predicate.
	None(predicate func(element E) bool) bool
	// Plus returns a list containing all elements of the original collection and given elements.
	Plus(element ...E) Collection[E]
	// Single returns the single element matching the given predicate.
	// If there is no such element, it returns `false` as a second return value.
	Single(predicate func(element E) bool) (E, bool)
	// Subtract returns a set containing all elements that are contained by this collection
	// and not contained by the specified collection.
	Subtract(other Iterable[E]) Set[E]
	// ToList converts this collection into List.
	ToList() List[E]
	// ToSet converts this collection into Set.
	ToSet() Set[E]
	// ToSlice converts this collection into slice.
	ToSlice() []E
	// Union returns a set containing all distinct elements from both collections.
	Union(other Iterable[E]) Set[E]
}
