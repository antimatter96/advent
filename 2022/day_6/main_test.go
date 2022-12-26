package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		input    string
		p1Answer int
		p2Answer int
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			p1Answer: 7,
			p2Answer: 19,
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			p1Answer: 5,
			p2Answer: 23,
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			p1Answer: 6,
			p2Answer: 23,
		},
		{
			input:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			p1Answer: 10,
			p2Answer: 29,
		},
		{
			input:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			p1Answer: 11,
			p2Answer: 26,
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
		})
	}
}
