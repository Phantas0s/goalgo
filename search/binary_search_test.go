package search

import "testing"

func TestBinary(t *testing.T) {
	testCases := []struct {
		name     string
		search   int
		list     []int
		expected bool
	}{
		{
			name:     "middle",
			search:   4,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "smaller",
			search:   2,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "larger",
			search:   6,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "last",
			search:   7,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "first",
			search:   1,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
		{
			name:     "big numbers",
			search:   5000,
			list:     []int{1, 2, 3, 4, 5, 6, 7, 1000, 2000, 5000},
			expected: true,
		},
		{
			name:     "Not found with 8",
			search:   8,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: false,
		},
		{
			name:     "Not found with 0",
			search:   0,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: false,
		},
		{
			name:     "Not found with 1000",
			search:   1000,
			list:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := binary(tc.search, tc.list)

			if actual != tc.expected {
				t.Errorf("Expected %v, actual %v", tc.expected, actual)
			}
		})
	}
}
