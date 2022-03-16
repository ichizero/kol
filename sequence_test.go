package kol

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence_MethodChain(t *testing.T) {
	type user struct {
		ID    int
		Name  string
		Admin bool
	}
	users := []*user{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol"},
		{ID: 4, Name: "Dave"},
		{ID: 5, Name: "Ellen"},
		{ID: 6, Name: "Frank"},
	}

	// Make a user with name length 5 admin and get a list of admin users.
	result := NewSequence(users...).
		Filter(func(u *user) bool {
			return len(u.Name) == 5
		}).
		Map(func(u *user) *user {
			u.Admin = true
			return u
		}).
		Take(2).
		Drop(1).
		ToSlice()

	assert.Equal(t, []*user{
		{ID: 3, Name: "Carol", Admin: true},
	}, result)

	assert.Equal(t, []*user{
		{ID: 1, Name: "Alice", Admin: true},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Carol", Admin: true},
		{ID: 4, Name: "Dave"},
		{ID: 5, Name: "Ellen"}, // should not make an admin
		{ID: 6, Name: "Frank"}, // should not make an admin
	}, users)
}

func TestSequence_String(t *testing.T) {
	seq := NewSequence([]int{1, 2, 3}...).
		Filter(func(e int) bool { return e < 2 }).
		Map(func(e int) int { return e * 2 }).
		Distinct().
		Take(2).
		Drop(1)
	assert.Equal(t, "cursor: 0, elements: [1 2 3] > filter > map > distinct > take 2 > drop 1", fmt.Sprint(seq))
}

func TestSequence_Distinct(t *testing.T) {
	tests := []struct {
		name  string
		elems []int
		want  []int
	}{
		{
			name:  "returns only distinct elements",
			elems: []int{1, 2, 3, 3, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSequence(tt.elems...).Distinct().ToSlice())
		})
	}
}

func TestSequence_Filter(t *testing.T) {
	tests := []struct {
		name      string
		elems     []int
		predicate func(e int) bool
		want      []int
	}{
		{
			name:  "filter",
			elems: []int{5, 6, 1, 5, 4, 2, 1},
			predicate: func(e int) bool {
				return e > 3
			},
			want: []int{5, 6, 5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSequence(tt.elems...).Filter(tt.predicate).ToSlice())
		})
	}
}

func TestSequence_Map(t *testing.T) {
	tests := []struct {
		name      string
		elems     []int
		transform func(e int) int
		want      []int
	}{
		{
			name:  "map",
			elems: []int{1, 2, 3, 3},
			transform: func(e int) int {
				return e * 2
			},
			want: []int{2, 4, 6, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSequence(tt.elems...).Map(tt.transform).ToSlice())
		})
	}
}

func TestSequence_Take(t *testing.T) {
	tests := []struct {
		name  string
		elems []int
		n     int
		want  []int
	}{
		{
			name:  "take 2 elements",
			elems: []int{1, 2, 3},
			n:     2,
			want:  []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSequence(tt.elems...).Take(tt.n).ToSlice())
		})
	}

	t.Run("Take should call previous sequences times just given n", func(t *testing.T) {
		const n = 3
		var countMap1, countMap2 int
		NewSequence([]int{1, 2, 3, 4, 5}...).
			Map(func(e int) int {
				countMap1++
				return e
			}).
			Map(func(e int) int {
				countMap2++
				return e
			}).
			Take(n).
			ToSlice()
		assert.Equal(t, n, countMap1)
		assert.Equal(t, n, countMap2)
	})
}

func TestSequence_Drop(t *testing.T) {
	tests := []struct {
		name  string
		elems []int
		n     int
		want  []int
	}{
		{
			name:  "drop 2 elements",
			elems: []int{1, 2, 3},
			n:     2,
			want:  []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSequence(tt.elems...).Drop(tt.n).ToSlice())
		})
	}
}

func TestMapSequence(t *testing.T) {
	tests := []struct {
		name      string
		elems     []int
		transform func(e int) string
		want      []string
	}{
		{
			name:  "map",
			elems: []int{1, 2, 3, 3},
			transform: func(e int) string {
				return strconv.Itoa(e)
			},
			want: []string{"1", "2", "3", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MapSequence(NewSequence(tt.elems...), tt.transform).ToSlice())
		})
	}
}
