package main

import (
	"context"
	"slices"
	"testing"
	"time"
)

func TestProcessJobs(t *testing.T) {
	tests := []struct {
		name         string
		upto, worker int
		expected     []int
	}{
		{
			name:     "Test 1",
			upto:     10,
			worker:   3,
			expected: []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81},
		},
		{
			name:     "Test 2",
			upto:     4,
			worker:   4,
			expected: []int{0, 1, 4, 9},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, tt.upto)
			for i := range input {
				input[i] = i
			}
			squares, err := processJobs(ctx, input, tt.worker)
			if err != nil {
				t.Errorf("Unexpected Error %v", err)
			}
			if !slices.Equal(squares, tt.expected) {
				t.Errorf("Expected Result %v but got %v", tt.expected, squares)
			}

		})
	}
}
