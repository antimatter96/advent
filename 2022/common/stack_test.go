package common

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEmpty(t *testing.T) {
	testCases := []struct {
		stack    Stack[int]
		expected bool
	}{
		{
			expected: false,
			stack:    Stack[int]{arr: []int{0}},
		},
		{
			expected: true,
			stack:    Stack[int]{arr: []int{}},
		},
		{
			expected: true,
			stack:    Stack[int]{arr: nil},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := tc.stack.Empty()

			if diff := cmp.Diff(tc.expected, res); diff != "" {
				t.Errorf("twoSum() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTop(t *testing.T) {
	testCases := []struct {
		stack     Stack[int]
		expected  int
		wantPanic bool
	}{
		{
			expected:  3,
			stack:     Stack[int]{arr: []int{3}},
			wantPanic: false,
		},
		{
			stack:     Stack[int]{arr: []int{}},
			wantPanic: true,
		},
		{
			stack:     Stack[int]{arr: nil},
			wantPanic: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			defer func() {
				r := recover()
				if diff := cmp.Diff(r != nil, tc.wantPanic); diff != "" {
					t.Errorf("Top() mismatch (-want +got):\n%s", diff)
				}
			}()
			res := tc.stack.Top()

			if !tc.wantPanic {
				if diff := cmp.Diff(tc.expected, res); diff != "" {
					t.Errorf("Top() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		stack     Stack[int]
		expected  int
		wantPanic bool
	}{
		{
			expected:  7,
			stack:     Stack[int]{arr: []int{7}},
			wantPanic: false,
		},
		{
			stack:     Stack[int]{arr: []int{}},
			wantPanic: true,
		},
		{
			stack:     Stack[int]{arr: nil},
			wantPanic: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			defer func() {
				r := recover()
				if diff := cmp.Diff(r != nil, tc.wantPanic); diff != "" {
					t.Errorf("Top() mismatch (-want +got):\n%s", diff)
				}
			}()
			res := tc.stack.Pop()

			if !tc.wantPanic {
				if diff := cmp.Diff(tc.expected, res); diff != "" {
					t.Errorf("Top() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		stack    Stack[int]
		expected Stack[int]
	}{
		{
			expected: Stack[int]{arr: []int{7}},
			stack:    Stack[int]{arr: []int{7}},
		},
		{
			stack:    Stack[int]{arr: []int{}},
			expected: Stack[int]{arr: []int{}},
		},
		{
			stack:    Stack[int]{arr: nil},
			expected: Stack[int]{arr: nil},
		},
		{
			stack:    Stack[int]{arr: []int{1, 2, 3}},
			expected: Stack[int]{arr: []int{3, 2, 1}},
		},
		{
			stack:    Stack[int]{arr: []int{1, 2}},
			expected: Stack[int]{arr: []int{2, 1}},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {

			tc.stack.Reverse()

			if diff := cmp.Diff(tc.expected, tc.stack, cmp.Comparer(func(x, y Stack[int]) bool {
				return tc.stack.Equals(&tc.expected)
			})); diff != "" {
				t.Errorf("Top() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
