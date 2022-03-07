package kol

import (
	"fmt"
)

func ExampleList() {
	fmt.Println(
		NewList("alice", "bob", "carol", "dave", "carol", "eve", "alice").
			Distinct().
			Drop(1).
			Take(1))
}
