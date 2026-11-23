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
		{
			name:      "malformed interval",
			intervals: [][]int{{1, 2}, {3}},
			want:      [][]int{},
		},
		{
			name:      "reversed interval",
			intervals: [][]int{{3, 1}},
			want:      [][]int{},
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

func TestInputValidation(t *testing.T) {
	t.Run("findOnetimenumber", func(t *testing.T) {
		if got := findOnetimenumber(nil); got != -1 {
			t.Errorf("empty input = %d, want -1", got)
		}
		if got := findOnetimenumber([]int{1, 1, 2, 3}); got != -1 {
			t.Errorf("multiple unique values = %d, want -1", got)
		}
		if got := findOnetimenumber([]int{1, 1, 2, 2}); got != -1 {
			t.Errorf("no unique value = %d, want -1", got)
		}
	})

	t.Run("isPalindrome", func(t *testing.T) {
		if !isPalindrome("") || !isPalindrome("上海海上") || isPalindrome("你好") {
			t.Error("empty and rune-based palindrome cases are incorrect")
		}
	})

	t.Run("kuohaoMatch", func(t *testing.T) {
		if !kuohaoMatch("") || kuohaoMatch("(a)") || kuohaoMatch("(") {
			t.Error("empty or invalid bracket cases are incorrect")
		}
	})

	t.Run("findstrcommonprestr", func(t *testing.T) {
		if got := findstrcommonprestr([]string{"上海人", "上海菜"}); got != "上海" {
			t.Errorf("Unicode prefix = %q, want %q", got, "上海")
		}
		if got := findstrcommonprestr([]string{}); got != "" {
			t.Errorf("empty input = %q, want empty string", got)
		}
	})

	t.Run("plusOne", func(t *testing.T) {
		if got := plusOne(nil); got != -1 {
			t.Errorf("empty input = %d, want -1", got)
		}
		if got := plusOne([]int{1, 10}); got != -1 {
			t.Errorf("invalid digit = %d, want -1", got)
		}
		if got := plusOne([]int{int(^uint(0) >> 1)}); got != -1 {
			t.Errorf("overflowing input = %d, want -1", got)
		}
	})

	t.Run("removeDuplicates", func(t *testing.T) {
		if got := removeDuplicates(nil); got != 0 {
			t.Errorf("empty input length = %d, want 0", got)
		}
		if got := removeDuplicates([]int{1, 3, 2}); got != 0 {
			t.Errorf("unsorted input length = %d, want 0", got)
		}
	})

	t.Run("twoSum", func(t *testing.T) {
		if got := twoSum(nil, 0); len(got) != 0 {
			t.Errorf("empty input result = %v, want empty", got)
		}
		maxInt := int(^uint(0) >> 1)
		if got := twoSum([]int{maxInt, 1}, maxInt); len(got) != 0 {
			t.Errorf("overflowing sum result = %v, want empty", got)
		}
	})
}
