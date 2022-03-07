package kol

type Iterator[E comparable] interface {
	HasNext() bool
	Next() E
	Remove()
}
