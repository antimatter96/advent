package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		input    []string
		p1Answer int
		p2Answer int
	}{
		{
			input: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			p1Answer: 13,
			p2Answer: 1,
		},
		{
			input: []string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			p1Answer: 88,
			p2Answer: 36,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			p1, p2 := Run(tc.input)

			if diff := cmp.Diff(tc.p1Answer, p1); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.p2Answer, p2); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}

			p1, p2 = RunNew(tc.input)

			if diff := cmp.Diff(tc.p1Answer, p1); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.p2Answer, p2); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
