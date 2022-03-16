package kol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_iterator_HasNext(t *testing.T) {
	tests := []struct {
		name string
		iter *iterator[int]
		want bool
	}{
		{
			name: "empty slice",
			iter: &iterator[int]{
				elements: []int{},
				cursor:   0,
			},
			want: false,
		},
		{
			name: "has next",
			iter: &iterator[int]{
				elements: []int{1, 2, 3},
				cursor:   1,
			},
			want: true,
		},
		{
			name: "end of slice",
			iter: &iterator[int]{
				elements: []int{1, 2, 3},
				cursor:   3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.iter.HasNext())
		})
	}
}

func Test_iterator_Next(t *testing.T) {
	tests := []struct {
		name   string
		iter   *iterator[int]
		want   int
		wantOK bool
	}{
		{
			name: "empty slice",
			iter: &iterator[int]{
				elements: []int{},
				cursor:   0,
			},
			want:   0,
			wantOK: false,
		},
		{
			name: "empty slice",
			iter: &iterator[int]{
				elements: []int{1},
				cursor:   0,
			},
			want:   1,
			wantOK: true,
		},
		{
			name: "end of slice",
			iter: &iterator[int]{
				elements: []int{1, 2, 3},
				cursor:   3,
			},
			want:   0,
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.iter.Next()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantOK, ok)
		})
	}
}
