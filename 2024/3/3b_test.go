package d3

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	in := "mul(411,270,.)"
	offset := 0
	result := 0
	enabled := true
	for {
		sr, newOffset, err := getResult(in, offset)
		fmt.Printf("%v\n", sr)
		offset = newOffset
		if err != nil {
			break
		}
		switch sr.variant {
		case Enable:
			enabled = true
		case Disable:
			enabled = false
		case Mul:
			if enabled {
				result += sr.mul1 * sr.mul2
			}
		}
	}
}
