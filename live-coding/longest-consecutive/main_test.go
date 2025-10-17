package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{6, 7, 3, 1, 100, 102, 6, 12},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{5, 6, 1, 2, 8, 9, 7},
			expected: 5,
		},
		{
			name:     "Problem example",
			input:    []int{4, 3, 8, 1, 2, 6, 100, 9},
			expected: 4,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: 1,
		},
		{
			name:     "No consecutive",
			input:    []int{1, 3, 5, 7, 9},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestConsecutive(tt.input)
			if result != tt.expected {
				t.Errorf("LongestConsecutive(%v) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}