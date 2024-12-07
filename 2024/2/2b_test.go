package d2

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestBrute(t *testing.T) {
	t.Skip()
	s := make([]int, 5)
	i := 0
	for {
		fmt.Printf("%v\n", i)

		for i := 0; i < 5; i++ {
			s[i] = rand.IntN(9) + 1
		}

		brute := isSliceOkBrute(s)
		iter := isSliceSafe(s)
		if brute != iter {
			t.Errorf("%v, brute: %v, iter: %v\n", s, brute, iter)
			return
		}

		i++
	}
}

func TestLineSafe(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"9 4 7 5 3", true},
		{"6 8 5 2 1", true},
		{"7 8 5 1", false},
		{"90 89 86 84 83 79", true},
		{"97 96 93 91 85", true},
		{"29 26 24 25 21", true},
		{"36 37 40 43 47", true},
		{"43 44 47 48 49 54", true},
		{"35 33 31 29 27 25 22 18", true},
		{"77 76 73 70 64", true},
		{"68 65 69 72 74 77 80 83", true},
		{"37 40 42 43 44 47 51", true},
		{"70 73 76 79 86", true},
		{"75 77 72 70 69", true},
		{"4 2 3 5", true},
		{"4 2 3 2", true},
		{"1 1 2 3 4 5", true},
		{"1 2 3 4 5 5", true},
		{"5 1 2 3", true},
		{"1 4 3 2", true},
		{"1 6 7 8", true},
		{"1 2 3 4 3", true},
		{"9 8 7 6 7", true},
		{"7 10 8 10 11", true},
		{"29 28 27 25 26 25 22 20", true},
		{"4 10 11 12", true},
		{"5 6 4 3 2", true},
		{"7 5 6 5 4 3", true},
		{"7 10 8 10 11", true},
		{"10 2 3 4 5", true},
		{"5 6 4 3 2 1", true},
		{"1 10 2 3 4", true},
		{"1 2 10 3 4", true},
		{"1 2 3 4 10", true},
		{"10 9 8 4 6", true},
		{"10 9 3 4 5", false},
		{"8 7 6 7 8", false},
		{"8 9 8 9 10", false},
		{"6 6 6 3 2", false},
		{"1 2 10 3 2", false},
		{"10 10 4 3 2", false},
		{"1 2 2 3 2", false},
		{"1 1 2 2 3 4", false},
		{"2 1 2 1 2 3", false},
		{"10 11 9 8 10", false},
		{"1 10 2 10 3", false},
		{"10 10 2 3", false},
		{"1 10 11 3", false},
	}
	for i, test := range tests {
		got := isLineSafeB(test.in)
		if got != test.out {
			t.Errorf("%v, expected %v, got %v (%v)", test.in, test.out, got, i)
		}
	}
}
