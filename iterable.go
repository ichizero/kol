package kol

type Iterable[E comparable] interface {
	All(predicate func(element E) bool) bool
	Any(predicate func(element E) bool) bool
	Contains(element E) bool
	Count(predicate func(element E) bool) int
	Distinct() Collection[E]
	Filter(predicate func(element E) bool) Collection[E]
	Find(predicate func(element E) bool) (E, bool)
	ForEach(action func(element E))
	Intersect(other Iterable[E]) Set[E]
	Iterator() Iterator[E] // TODO
	Map(predicate func(element E) E) Collection[E]
	Minus(element ...E) Collection[E]
	None(predicate func(element E) bool) bool
	Plus(element ...E) Collection[E]
	Single(predicate func(element E) bool) (E, bool)
	Subtract(other Iterable[E]) Set[E]
	ToList() List[E]
	ToSet() Set[E]
	ToSlice() []E
	Union(other Iterable[E]) Set[E]
}
