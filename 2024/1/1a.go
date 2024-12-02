package d1

import (
	"bytes"
	"fmt"
	"io"
	"slices"
)

var newline = []byte{0xA}

func RunA(r io.Reader) int {
	p := make([]byte, 32_000)
	read, _ := r.Read(p)
	if read == 32_000 {
		panic("too much data")
	}

	p = p[:read]
	rows := bytes.Split(p, newline)

	size := len(rows)
	locs1 := make([]int, size)
	locs2 := make([]int, size)

	for i, row := range rows {
		split := bytes.Split(row, []byte{' ', ' ', ' '})
		if len(split) < 2 {
			continue
		}
		first := split[0]
		second := split[1]

		var nr int
		fmt.Sscanf(string(first), "%d", &nr)
		locs1[i] = nr
		fmt.Sscanf(string(second), "%d", &nr)
		locs2[i] = nr
	}
	slices.SortFunc(locs1, func(x, y int) int { return x - y })
	slices.SortFunc(locs2, func(x, y int) int { return x - y })

	result := 0
	for i, x := range locs1 {
		y := locs2[i]
		result += abs(x - y)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
