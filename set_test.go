package kol

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	assert.Equal(t,
		map[int]struct{}{1: {}, 2: {}},
		NewSet(1, 2, 1, 2).(*set[int]).m)
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		name     string
		set      Set[int]
		elements []int
		want     Set[int]
	}{
		{
			name:     "add some elements",
			set:      NewSet[int](1, 2),
			elements: []int{3, 4},
			want:     NewSet[int](1, 2, 3, 4),
		},
		{
			name:     "duplicate",
			set:      NewSet[int](1, 2, 3),
			elements: []int{1, 2},
			want:     NewSet[int](1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Add(tt.elements...)
			assert.Equal(t, tt.want, tt.set)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct {
		name    string
		set     Set[int]
		targets []int
		want    Set[int]
	}{
		{
			name:    "remove one",
			set:     NewSet[int](1, 2, 3),
			targets: []int{2},
			want:    NewSet[int](1, 3),
		},
		{
			name:    "remove all",
			set:     NewSet[int](1, 2, 3),
			targets: []int{1, 2, 3},
			want:    NewSet[int](),
		},
		{
			name:    "not matched",
			set:     NewSet[int](1, 2, 3),
			targets: []int{4, 5},
			want:    NewSet[int](1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Remove(tt.targets...)
			assert.Equal(t, tt.want, tt.set)
		})
	}
}

func TestSet_Retain(t *testing.T) {
	tests := []struct {
		name    string
		set     Set[int]
		targets []int
		want    Set[int]
	}{
		{
			name:    "retain one",
			set:     NewSet[int](1, 2, 3),
			targets: []int{2},
			want:    NewSet[int](2),
		},
		{
			name:    "retain all",
			set:     NewSet[int](1, 2, 3),
			targets: []int{1, 2, 3},
			want:    NewSet[int](1, 2, 3),
		},
		{
			name:    "not matched",
			set:     NewSet[int](1, 2, 3),
			targets: []int{4, 5},
			want:    NewSet[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.Retain(tt.targets...)
			assert.Equal(t, tt.want, tt.set)
		})
	}
}

func TestSet_All(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(v int) bool
		want      bool
	}{
		{
			name:      "all elements matches the given predicate",
			set:       NewSet[int](1, 2, 3),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "no element matches the given predicate",
			set:       NewSet[int](101, 102, 103),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "one element does not match the given predicate",
			set:       NewSet[int](999, 999, 1),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "empty set should be falsy",
			set:       NewSet[int](),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.All(tt.predicate))
		})
	}
}

func TestSet_Any(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(v int) bool
		want      bool
	}{
		{
			name:      "all elements matches the given predicate",
			set:       NewSet[int](1, 2, 3),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "no element matches the given predicate",
			set:       NewSet[int](101, 102, 103),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
		{
			name:      "one element matches the given predicate",
			set:       NewSet[int](999, 999, 1),
			predicate: func(v int) bool { return v < 100 },
			want:      true,
		},
		{
			name:      "empty set should be falsy",
			set:       NewSet[int](),
			predicate: func(v int) bool { return v < 100 },
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Any(tt.predicate))
		})
	}
}

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		elem int
		want bool
	}{
		{
			name: "contains",
			set:  NewSet[int](1, 2, 3),
			elem: 1,
			want: true,
		},
		{
			name: "not contained",
			set:  NewSet[int](1, 2, 3),
			elem: 999,
			want: false,
		},
		{
			name: "empty set should be falsy",
			set:  NewSet[int](),
			elem: 0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Contains(tt.elem))
		})
	}
}

func TestSet_Count(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(v int) bool
		want      int
	}{
		{
			name: "count",
			set:  NewSet[int](1, 2, 3),
			predicate: func(v int) bool {
				return v < 2
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Count(tt.predicate))
		})
	}
}

func TestSet_Distinct(t *testing.T) {
	tests := []struct {
		name string
		set  Set[int]
		want List[int]
	}{
		{
			name: "no duplication",
			set:  NewSet[int](1, 2, 3),
			want: NewList[int](1, 2, 3),
		},
		{
			name: "has duplication",
			set:  NewSet[int](1, 2, 3, 2, 3, 4, 5),
			want: NewList[int](1, 2, 3, 4, 5),
		},
		{
			name: "empty hash set",
			set:  NewSet[int](),
			want: NewList[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want.ToSlice(), tt.set.Distinct().ToSlice())
		})
	}
}

func TestSet_Filter(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(e int) bool
		want      Set[int]
	}{
		{
			name: "use element",
			set:  NewSet[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 1
			},
			want: NewSet[int](2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Filter(tt.predicate))
		})
	}
}

func TestSet_Map(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(e int) int
		want      Set[int]
	}{
		{
			name: "map",
			set:  NewSet[int](1, 2, 3),
			predicate: func(e int) int {
				return e * 2
			},
			want: NewSet[int](2, 4, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.Map(tt.predicate))
		})
	}
}

func TestSet_ForEach(t *testing.T) {
	tests := []struct {
		name    string
		set     Set[int]
		prepare func() (*int, func(e int))
		want    int
	}{
		{
			name: "use element",
			set:  NewSet[int](1, 2, 3),
			prepare: func() (*int, func(e int)) {
				var res int
				return &res, func(e int) {
					res += e
				}
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, action := tt.prepare()
			tt.set.ForEach(action)
			assert.Equal(t, tt.want, *res)
		})
	}
}

func TestSet_Minus(t *testing.T) {
	tests := []struct {
		name  string
		set   Set[int]
		elems []int
		want  Set[int]
	}{
		{
			name:  "minus",
			set:   NewSet[int](1, 2, 3, 4),
			elems: []int{3, 4},
			want:  NewSet[int](1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := tt.set.ToSlice()
			assert.ElementsMatch(t, tt.want.ToSlice(), tt.set.Minus(tt.elems...).ToSlice())
			assert.ElementsMatch(t, prev, tt.set.ToSlice())
		})
	}
}

func TestSet_None(t *testing.T) {
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(e int) bool
		want      bool
	}{
		{
			name: "none",
			set:  NewSet[int](1, 2, 3, 3),
			predicate: func(e int) bool {
				return e > 999
			},
			want: true,
		},
		{
			name: "matched",
			set:  NewSet[int](1, 2, 3, 3),
			predicate: func(e int) bool {
				return e < 999
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.set.None(tt.predicate))
		})
	}
}

func TestSet_Plus(t *testing.T) {
	tests := []struct {
		name  string
		set   Set[int]
		elems []int
		want  Set[int]
	}{
		{
			name:  "plus",
			set:   NewSet[int](1, 2, 3),
			elems: []int{4, 5},
			want:  NewSet[int](1, 2, 3, 4, 5),
		},
		{
			name:  "duplicated",
			set:   NewSet[int](1, 2, 3),
			elems: []int{3, 5},
			want:  NewSet[int](1, 2, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prev := tt.set.ToSlice()
			assert.Equal(t, tt.want, tt.set.Plus(tt.elems...))
			assert.ElementsMatch(t, prev, tt.set.ToSlice())
		})
	}
}

func TestSet_Single(t *testing.T) {
	type result struct {
		e  int
		ok bool
	}
	tests := []struct {
		name      string
		set       Set[int]
		predicate func(e int) bool
		want      result
	}{
		{
			name: "matched once",
			set:  NewSet[int](1, 2, 3),
			predicate: func(e int) bool {
				return e == 1
			},
			want: result{e: 1, ok: true},
		},
		{
			name: "matched twice",
			set:  NewSet[int](1, 2, 3),
			predicate: func(e int) bool {
				return e > 1
			},
			want: result{e: 0, ok: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.set.Single(tt.predicate)
			assert.Equal(t, tt.want.e, got)
			assert.Equal(t, tt.want.ok, ok)
		})
	}
}

func TestMapSet(t *testing.T) {
	tests := []struct {
		name      string
		elems     []int
		predicate func(e int) string
		want      []string
	}{
		{
			name:  "map",
			elems: []int{1, 2, 3},
			predicate: func(e int) string {
				return strconv.Itoa(e)
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, MapSet[int, string](NewSet(tt.elems...), tt.predicate).ToSlice())
		})
	}
}
