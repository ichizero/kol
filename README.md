# kol

[![Test](https://github.com/ichizero/kol/actions/workflows/test.yml/badge.svg)](https://github.com/ichizero/kol/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ichizero/kol.svg)](https://pkg.go.dev/github.com/ichizero/kol)
[![Codecov](https://codecov.io/gh/ichizero/kol/branch/main/graph/badge.svg)](https://codecov.io/gh/ichizero/kol)
[![Go Report Card](https://goreportcard.com/badge/github.com/ichizero/kol)](https://goreportcard.com/report/github.com/ichizero/kol)

`kol` is a collection package based on Go Generics.
It provides List, Set and Sequence like Kotlin collections package.
You can operate collections with method-chaining 🔗.

- [kotlin.collections](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.collections/)
- [kotlin.sequences](https://kotlinlang.org/api/latest/jvm/stdlib/kotlin.sequences/)

**Out of scope**

`kol` is not a complete reimplementation of the kotlin.collections and kotlin.sequences packages.
It is implemented by utilizing Go language features.

## 🚀 Installation

```bash
go get github.com/ichizero/kol
```

## 🧐 Example

```go
package main

import (
	"fmt"

	"github.com/ichizero/kol"
)

func main() {
	seq := kol.NewSequence(1, 2, 3, 4, 5).
		Filter(func(e int) bool {
			return e < 3
		}).
		Map(func(e int) int {
			return e * 2
		}).
		Distinct().
		Drop(1).
		Take(2)
	fmt.Println(seq.ToSlice())
}
```

## ✨ Features

GoDoc: [![Go Reference](https://pkg.go.dev/badge/github.com/ichizero/kol.svg)](https://pkg.go.dev/github.com/ichizero/kol)

### List & Set

List backend is a slice of Go.

Set backend is a map of Go.
It seems like Kotlin & Java's HashMap, but it cannot be iterated in a sorted order.

|                | List | Set |
| -------------- | ---- | --- |
| All            | ✅   | ✅  |
| Any            | ✅   | ✅  |
| Contains       | ✅   | ✅  |
| Count          | ✅   | ✅  |
| Distinct       | ✅   | ✅  |
| Filter         | ✅   | ✅  |
| FilterIndexed  | ✅   | 🚫  |
| Find           | ✅   | ✅  |
| ForEach        | ✅   | ✅  |
| ForEachIndexed | ✅   | 🚫  |
| Intersect      | ✅   | ✅  |
| Iterator       | ✅   | ✅  |
| Map            | ✅   | ✅  |
| MapIndexed     | ✅   | 🚫  |
| Minus          | ✅   | ✅  |
| None           | ✅   | ✅  |
| Plus           | ✅   | ✅  |
| Single         | ✅   | ✅  |
| Subtract       | ✅   | ✅  |
| Union          | ✅   | ✅  |

### Sequence

Sequence enables us to lazy evaluation of a collection.
In some cases, the process is faster than list operations.

For more details, please refer to the following documents.

- [Kotlin docs: Sequence processing example](https://kotlinlang.org/docs/sequences.html#sequence-processing-example)

|          | Sequence |
| -------- | -------- |
| Distinct | ✅       |
| Filter   | ✅       |
| Map      | ✅       |
| Take     | ✅       |
| Drop     | ✅       |

### ⚠️ Limitation

#### Map

Because of the Go Generics specification, Map methods in each interface cannot convert an element type.
You can use MapList, MapSet and MapSequence instead.

```go
MapList(
    NewList(1, 2, 3).
        Map(func(e int) int {
            return e * 2
        }),
    func(e int) string {
        return strconv.Itoa(e)
    })
```
