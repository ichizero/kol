package kol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name    string
		list    List[int]
		targets []int
		want    List[int]
	}{
		{
			name:    "remove one",
			list:    NewList[int](1, 2, 3),
			targets: []int{2},
			want:    NewList[int](1, 3),
		},
		{
			name:    "remove all",
			list:    NewList[int](1, 2, 3),
			targets: []int{1, 2, 3},
			want:    NewList[int](),
		},
		{
			name:    "not matched",
			list:    NewList[int](1, 2, 3),
			targets: []int{4, 5},
			want:    NewList[int](1, 2, 3),
		},
		{
			name:    "duplicate entries should be removed for the num of times listed in args",
			list:    NewList[int](1, 2, 3, 3, 4, 3),
			targets: []int{3, 3},
			want:    NewList[int](1, 2, 4, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Remove(tt.targets...)
			assert.Equal(t, tt.want, tt.list)
		})
	}
}

func TestList_Retain(t *testing.T) {
	tests := []struct {
		name    string
		list    List[int]
		targets []int
		want    List[int]
	}{
		{
			name:    "retain one",
			list:    NewList[int](1, 2, 3),
			targets: []int{2},
			want:    NewList[int](2),
		},
		{
			name:    "retain all",
			list:    NewList[int](1, 2, 3),
			targets: []int{1, 2, 3},
			want:    NewList[int](1, 2, 3),
		},
		{
			name:    "not matched",
			list:    NewList[int](1, 2, 3),
			targets: []int{4, 5},
			want:    NewList[int](),
		},
		{
			name:    "the num of duplicate entries should be kept",
			list:    NewList[int](1, 2, 3, 3, 4, 3),
			targets: []int{1, 3, 3},
			want:    NewList[int](1, 3, 3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Retain(tt.targets...)
			assert.Equal(t, tt.want, tt.list)
		})
	}
}

func TestList_All(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(v int) bool
		want      bool
	}{
		{
			name:      "all elements matches the given predicate",
			list:      NewList[int](1, 2, 3),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "no element matches the given predicate",
			list:      NewList[int](101, 102, 103),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "one element does not match the given predicate",
			list:      NewList[int](999, 999, 1),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "empty list should be falsy",
			list:      NewList[int](),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.All(tt.predicate))
		})
	}
}

func TestList_Any(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(v int) bool
		want      bool
	}{
		{
			name:      "all elements matches the given predicate",
			list:      NewList[int](1, 2, 3),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "no element matches the given predicate",
			list:      NewList[int](101, 102, 103),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "one element matches the given predicate",
			list:      NewList[int](999, 999, 1),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "empty list should be falsy",
			list:      NewList[int](),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Any(tt.predicate))
		})
	}
}

func TestList_Contains(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		elem int
		want bool
	}{
		{
			name: "contains",
			list: NewList[int](1, 2, 3),
			elem: 1,
			want: true,
		},
		{
			name: "not contained",
			list: NewList[int](1, 2, 3),
			elem: 999,
			want: false,
		},
		{
			name: "empty list should be falsy",
			list: NewList[int](),
			elem: 0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Contains(tt.elem))
		})
	}
}

func TestList_Count(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(v int) bool
		want      int
	}{
		{
			name: "count",
			list: NewList[int](1, 2, 3),
			predicate: func(v int) bool {
				return v < 2
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Count(tt.predicate))
		})
	}
}

func TestList_Distinct(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		want Set[int]
	}{
		{
			name: "no duplication",
			list: NewList[int](1, 2, 3),
			want: NewList[int](1, 2, 3),
		},
		{
			name: "has duplication",
			list: NewList[int](1, 2, 3, 2, 3, 4, 5),
			want: NewList[int](1, 2, 3, 4, 5),
		},
		{
			name: "empty list",
			list: NewList[int](),
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Distinct())
		})
	}
}

func TestList_Drop(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		n    uint
		want List[int]
	}{
		{
			name: "n=0",
			list: NewList[int](1, 2, 3),
			n:    0,
			want: NewList[int](1, 2, 3),
		},
		{
			name: "n=3",
			list: NewList[int](1, 2, 3, 2, 3),
			n:    3,
			want: NewList[int](2, 3),
		},
		{
			name: "n > len",
			list: NewList[int](1, 2, 3, 2, 3),
			n:    999,
			want: NewList[int](),
		},
		{
			name: "n == len",
			list: NewList[int](1, 2, 3),
			n:    3,
			want: NewList[int](),
		},
		{
			name: "empty list",
			list: NewList[int](),
			n:    999,
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Drop(tt.n))
		})
	}
}

func TestList_DropWhile(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(v int) bool
		want      List[int]
	}{
		{
			name: "not matched",
			list: NewList[int](1, 2, 3),
			predicate: func(v int) bool {
				return v < 0
			},
			want: NewList[int](1, 2, 3),
		},
		{
			name: "some elements matched",
			list: NewList[int](1, 2, 3, 4, 5),
			predicate: func(v int) bool {
				return v < 3
			},
			want: NewList[int](3, 4, 5),
		},
		{
			name: "except last one",
			list: NewList[int](1, 2, 3),
			predicate: func(v int) bool {
				return v < 3
			},
			want: NewList[int](3),
		},
		{
			name: "drop all",
			list: NewList[int](1, 2, 3),
			predicate: func(v int) bool {
				return v < 5
			},
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.DropWhile(tt.predicate))
		})
	}
}

func TestList_ElementAt(t *testing.T) {
	type result struct {
		e  int
		ok bool
	}
	tests := []struct {
		name  string
		list  List[int]
		index int
		want  result
	}{
		{
			name:  "exists",
			list:  NewList[int](1, 2, 3),
			index: 1,
			want:  result{e: 2, ok: true},
		},
		{
			name:  "index out of range",
			list:  NewList[int](1, 2, 3),
			index: -1,
			want:  result{e: 0, ok: false},
		},
		{
			name:  "empty list",
			list:  NewList[int](),
			index: 0,
			want:  result{e: 0, ok: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.list.ElementAt(tt.index)
			assert.Equal(t, tt.want.e, got)
			assert.Equal(t, tt.want.ok, ok)
		})
	}
}

func TestList_ElementAtOrElse(t *testing.T) {
	tests := []struct {
		name         string
		list         List[int]
		index        int
		defaultValue func() int
		want         int
	}{
		{
			name:  "exists",
			list:  NewList[int](1, 2, 3),
			index: 1,
			defaultValue: func() int {
				return 999
			},
			want: 2,
		},
		{
			name:  "index out of range",
			list:  NewList[int](1, 2, 3),
			index: -1,
			defaultValue: func() int {
				return 999
			},
			want: 999,
		},
		{
			name:  "empty list",
			list:  NewList[int](),
			index: 0,
			defaultValue: func() int {
				return 999
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.ElementAtOrElse(tt.index, tt.defaultValue))
		})
	}
}

func TestList_FilterIndexed(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(idx int, e int) bool
		want      List[int]
	}{
		{
			name: "use index",
			list: NewList[int](1, 2, 3),
			predicate: func(idx int, e int) bool {
				return idx > 1
			},
			want: NewList[int](3),
		},
		{
			name: "use element",
			list: NewList[int](1, 2, 3, 2, 3),
			predicate: func(idx int, e int) bool {
				return e == 2
			},
			want: NewList[int](2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.FilterIndexed(tt.predicate))
		})
	}
}

func TestList_ForEachIndexed(t *testing.T) {
	tests := []struct {
		name    string
		list    List[int]
		prepare func() (*int, func(idx int, e int))
		want    int
	}{
		{
			name: "use both index and element",
			list: NewList[int](1, 2, 3, 4, 5),
			prepare: func() (*int, func(idx int, e int)) {
				var res int
				return &res, func(idx int, e int) {
					res += idx * e
				}
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, action := tt.prepare()
			tt.list.ForEachIndexed(action)
			assert.Equal(t, tt.want, *res)
		})
	}
}

func TestList_IndexOfFirst(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(e int) bool
		want      int
	}{
		{
			name: "found",
			list: NewList[int](1, 2, 3, 3, 3),
			predicate: func(e int) bool {
				return e == 3
			},
			want: 2,
		},
		{
			name: "not found",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 3
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.IndexOfFirst(tt.predicate))
		})
	}
}

func TestList_IndexOfLast(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(e int) bool
		want      int
	}{
		{
			name: "found",
			list: NewList[int](1, 1, 1, 2, 3),
			predicate: func(e int) bool {
				return e == 1
			},
			want: 2,
		},
		{
			name: "not found",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 3
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.IndexOfLast(tt.predicate))
		})
	}
}

func TestList_Minus(t *testing.T) {
	tests := []struct {
		name  string
		list  List[int]
		elems []int
		want  List[int]
	}{
		{
			name:  "minus",
			list:  NewList[int](1, 2, 3, 3, 3),
			elems: []int{3, 3},
			want:  NewList[int](1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := tt.list.ToSlice()
			assert.Equal(t, tt.want, tt.list.Minus(tt.elems...))
			assert.Equal(t, prev, tt.list.ToSlice())
		})
	}
}

func TestList_None(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(e int) bool
		want      bool
	}{
		{
			name: "none",
			list: NewList[int](1, 2, 3, 3),
			predicate: func(e int) bool {
				return e > 999
			},
			want: true,
		},
		{
			name: "matched",
			list: NewList[int](1, 2, 3, 3),
			predicate: func(e int) bool {
				return e < 999
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.None(tt.predicate))
		})
	}
}

func TestList_Partition(t *testing.T) {
	tests := []struct {
		name       string
		list       List[int]
		predicate  func(e int) bool
		wantFirst  List[int]
		wantSecond List[int]
	}{
		{
			name: "partitioned",
			list: NewList[int](1, 2, 3, 4, 5),
			predicate: func(e int) bool {
				return e > 3
			},
			wantFirst:  NewList(4, 5),
			wantSecond: NewList(1, 2, 3),
		},
		{
			name: "all elements matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e < 999
			},
			wantFirst:  NewList(1, 2, 3),
			wantSecond: NewList[int](),
		},
		{
			name: "no element matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 999
			},
			wantFirst:  NewList[int](),
			wantSecond: NewList(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first, second := tt.list.Partition(tt.predicate)
			assert.Equal(t, tt.wantFirst, first)
			assert.Equal(t, tt.wantSecond, second)
		})
	}
}

func TestList_Plus(t *testing.T) {
	tests := []struct {
		name  string
		list  List[int]
		elems []int
		want  List[int]
	}{
		{
			name:  "plus",
			list:  NewList[int](1, 2, 3),
			elems: []int{3, 4},
			want:  NewList[int](1, 2, 3, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := tt.list.ToSlice()
			assert.Equal(t, tt.want, tt.list.Plus(tt.elems...))
			assert.Equal(t, prev, tt.list.ToSlice())
		})
	}
}

func TestList_Reversed(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		want List[int]
	}{
		{
			name: "reversed",
			list: NewList[int](1, 2, 3, 3),
			want: NewList[int](3, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := tt.list.ToSlice()
			assert.Equal(t, tt.want, tt.list.Reversed())
			assert.Equal(t, prev, tt.list.ToSlice())
		})
	}
}

func TestList_Single(t *testing.T) {
	type result struct {
		e  int
		ok bool
	}
	tests := []struct {
		name      string
		list      List[int]
		predicate func(e int) bool
		want      result
	}{
		{
			name: "matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e == 1
			},
			want: result{e: 1, ok: true},
		},
		{
			name: "duplicated",
			list: NewList[int](1, 2, 3, 3),
			predicate: func(e int) bool {
				return e == 3
			},
			want: result{e: 0, ok: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.list.Single(tt.predicate)
			assert.Equal(t, tt.want.e, got)
			assert.Equal(t, tt.want.ok, ok)
		})
	}
}

func TestList_Take(t *testing.T) {
	tests := []struct {
		name string
		list List[int]
		n    uint
		want List[int]
	}{
		{
			name: "n < len",
			list: NewList[int](1, 2, 3),
			n:    2,
			want: NewList[int](1, 2),
		},
		{
			name: "n == len",
			list: NewList[int](1, 2, 3),
			n:    3,
			want: NewList[int](1, 2, 3),
		},
		{
			name: "n > len",
			list: NewList[int](1, 2, 3),
			n:    10,
			want: NewList[int](1, 2, 3),
		},
		{
			name: "n == 0",
			list: NewList[int](1, 2, 3),
			n:    0,
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.Take(tt.n))
		})
	}
}

func TestList_TakeWhile(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate func(e int) bool
		want      List[int]
	}{
		{
			name: "some elements matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e < 3
			},
			want: NewList[int](1, 2),
		},
		{
			name: "all elements matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e < 999
			},
			want: NewList[int](1, 2, 3),
		},
		{
			name: "no element matched",
			list: NewList[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 999
			},
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.list.TakeWhile(tt.predicate))
		})
	}
}
