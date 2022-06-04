package main

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
		{1, 1, 2},
	}

	for _, test := range tests {
		if got := Add(test.a, test.b); got != test.want {
			t.Errorf("Add(%d, %d) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}
