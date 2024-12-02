package d1

import (
	"bytes"
	"fmt"
	"io"
)

func RunB(r io.Reader) int {
	p := make([]byte, 32_000)
	read, _ := r.Read(p)
	if read == 32_000 {
		panic("too much data")
	}

	p = p[:read]
	rows := bytes.Split(p, newline)

	size := len(rows)
	locs1 := make([]int, size)
	locs2 := make(map[int]int, int(size/2))

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

		previous, ok := locs2[nr]
		if !ok {
			locs2[nr] = 1
		} else {
			locs2[nr] = previous + 1
		}
	}

	result := 0

	for _, loc1 := range locs1 {
		hits := locs2[loc1]
		result += loc1 * hits
	}

	return result
}
