package d3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var log = false

func RunA(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	result := 0
	for scanner.Scan() {
		matches := getMatches(scanner.Text())
		if log {
			fmt.Printf("Matches: %v\n", matches)
		}
		for _, match := range matches {
			result += match[0] * match[1]
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading", err)
	}
	return result
}

func getMatches(s string) [][]int {
	if log {
		fmt.Printf("input: %v\n", s)
	}
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	if log {
		fmt.Printf("matches: %v\n", matches)
	}
	results := make([][]int, len(matches))
	for i, match := range matches {
		result := make([]int, 2)
		result[0], _ = strconv.Atoi(match[1])
		result[1], _ = strconv.Atoi(match[2])
		results[i] = result
	}
	return results
}
