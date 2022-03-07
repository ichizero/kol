package kol

type Iterable[E comparable] interface {
	All(predicate func(element E) bool) bool
	Any(predicate func(element E) bool) bool
	Contains(element E) bool
	Count(predicate func(element E) bool) int
	Distinct() List[E]
	Drop(n uint) List[E]
	DropWhile(predicate func(element E) bool) List[E]
	ElementAt(index int) (*E, error)
	ElementAtOrElse(index int, defaultValue func() E) E
	Filter(predicate func(element E) bool) List[E]
	FilterIndexed(predicate func(idx int, element E) bool) List[E]
	Find(predicate func(element E) bool) *E
	FindLast(predicate func(element E) bool) *E
	ForEach(action func(element E))
	ForEachIndexed(action func(index int, element E))
	IndexOf(element E) int
	IndexOfFirst(predicate func(element E) bool) int
	IndexOfLast(predicate func(element E) bool) int
	Intersect() Set[E]     // TODO
	Iterator() Iterator[E] // TODO
	Minus(element E) List[E]
	None(predicate func(element E) bool) bool
	Partition(predicate func(element E) bool) (List[E], List[E])
	Plus(element E) List[E]
	Reversed() List[E]
	Shuffled() List[E]
	Single(predicate func(element E) bool) *E
	Subtract() Set[E] // TODO
	Take(n uint) List[E]
	TakeWhile(predicate func(element E) bool) List[E]
	ToList() List[E]
	ToSlice() []E
	Union() Set[E] // TODO
}
