package d2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func RunA(r io.Reader) int {
	s := bufio.NewScanner(r)
	safe := 0
	for s.Scan() {
		if isLineSafe(s.Text()) {
			safe++
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println(err)
	}

	return safe
}

func isLineSafe(s string) bool {
	nums := strings.Split(s, " ")

	prev := 0
	prevChange := 0
	for i, n := range nums {
		current, _ := strconv.Atoi(n)

		if i == 0 {
			prev = current
			continue
		}

		change := current - prev
		if change < -3 || change > 3 {
			return false
		}
		if change < 0 && prevChange > 0 {
			return false
		}
		if change > 0 && prevChange < 0 {
			return false
		}
		if change == 0 {
			return false
		}
		prev = current
		prevChange = change
	}
	return true
}
