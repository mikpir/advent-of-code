package d2

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var log = true

func RunB(r io.Reader) int {
	s := bufio.NewScanner(r)
	safe := 0
	for s.Scan() {
		if isLineSafeB(s.Text()) {
			fmt.Printf("Line %v is safe\n", s.Text())
			safe++
		} else {
			fmt.Printf("Line %v is unsafe\n", s.Text())
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println(err)
	}

	return safe
}

func isLineSafeB(str string) bool {
	if log {
		fmt.Printf("Line: %v\n", str)
	}
	strs := strings.Split(str, " ")
	nums := make([]int, len(strs))
	for i, s := range strs {
		val, _ := strconv.Atoi(s)
		nums[i] = val
	}

	return isSliceSafe(nums)
}

func isSliceSafe(nums []int) bool {
	skipped := false

	prev3 := 0
	prev2 := 0
	prev := 0
	next := 0
	next2 := 0
	for i, curr := range nums {
		if i == 0 {
			prev = curr
			continue
		}
		if i+1 == len(nums) && !skipped {
			return true
		}

		if i < len(nums)-1 {
			next = nums[i+1]
		} else {
			next = 0
		}
		if i < len(nums)-2 {
			next2 = nums[i+2]
		} else {
			next2 = 0
		}

		if log {
			fmt.Printf("i:%v, p3:%v p2:%v p1:%v c:%v \n", i, prev3, prev2, prev, curr)
		}

		if !isChangeOk(prev2, prev, curr) {
			if log {
				fmt.Printf("change not ok\n")
			}
			if skipped {
				return false
			}
			skipped = true

			if isPrev2Bad(prev3, prev2, prev, curr, next, next2) {
				if log {
					fmt.Printf("Looks like prev2 is bad\n")
				}
				prev2 = prev
				prev = curr
			} else if isPrevBad(prev3, prev2, prev, curr, next, next2) {
				if log {
					fmt.Printf("Looks like prev is bad\n")
				}
				prev = curr
			} else if isCurrBad(prev3, prev2, prev, curr, next, next2) {
				if log {
					fmt.Printf("Looks like curr is bad\n")
				}
				prev2 = prev
			} else {
				if log {
					fmt.Printf("No match found\n")
				}
				return false
			}

			continue
		}

		prev3 = prev2
		prev2 = prev
		prev = curr
	}
	return true
}

func isPrevBad(prev3, prev2, prev, curr, next, next2 int) bool {
	fmt.Printf("Checking prev %v\n", prev)
	return isChangeOk(prev3, prev2, curr) && isChangeOk(prev2, curr, next) && isChangeOk(curr, next, next2)
}

func isPrev2Bad(prev3, prev2, prev, curr, next, next2 int) bool {
	fmt.Printf("Checking prev2 %v\n", prev2)
	return isChangeOk(prev3, prev, curr) && isChangeOk(prev, curr, next) && isChangeOk(curr, next, next2)
}

func isCurrBad(prev3, prev2, prev, curr, next, next2 int) bool {
	fmt.Printf("Checking curr %v\n", curr)
	return isChangeOk(prev3, prev2, prev) && isChangeOk(prev2, prev, next) && isChangeOk(prev, next, next2)
}

func isChangeOk(prev2, prev, curr int) bool {
	result := checkChange(prev2, prev, curr)
	if log {
		fmt.Printf("%v -> %v -> %v = %v\n", prev2, prev, curr, result)
	}
	return result
}

func checkChange(prev2, prev, curr int) bool {
	prevChange := prev - prev2
	if prev2 == 0 {
		prevChange = 0
	}
	if prev2 == 0 && prev == 0 {
		return true
	}
	if curr == 0 {
		return true
	}

	currChange := curr - prev
	if currChange < -3 || currChange > 3 {
		return false
	}
	if currChange < 0 && prevChange > 0 {
		return false
	}
	if currChange > 0 && prevChange < 0 {
		return false
	}
	if currChange == 0 {
		return false
	}
	return true
}
