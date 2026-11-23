package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 2}, {4, 5}},
			want:      [][]int{{1, 2}, {4, 5}},
		},
		{
			name:      "overlapping and touching intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {6, 8}},
			want:      [][]int{{1, 8}},
		},
		{
			name:      "nested intervals",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 5}},
			want:      [][]int{{1, 10}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := merge(test.intervals); !reflect.DeepEqual(got, test.want) {
				t.Errorf("merge() = %v, want %v", got, test.want)
			}
		})
	}
}
