package d3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

const (
	Enable  = 1
	Disable = 2
	Mul     = 3
)

type ScanResult struct {
	variant int
	mul1    int
	mul2    int
}

func (sr *ScanResult) getResult() int {
	return sr.mul1 * sr.mul2
}

func RunB(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	result := 0

	enabled := true
	offset := 0
	for scanner.Scan() {
		text := scanner.Text()
		for {
			sr, newOffset, err := getResult(text, offset)
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
					result += sr.getResult()
				}
			}
		}
	}
	return result
}

func (sr *ScanResult) String() string {
	if sr == nil {
		return "ENDED"
	}
	switch sr.variant {
	case Enable:
		return "ENABLE"
	case Disable:
		return "DISABLE"
	}
	return fmt.Sprintf("MULTIPLY %v*%v", sr.mul1, sr.mul2)
}

func getResult(str string, offset int) (*ScanResult, int, error) {
	do := "do()"
	doi := 0
	dont := "don't()"
	donti := 0
	mul := "mul(%,%)"
	muli := 0
	num1 := make([]byte, 0, 8)
	num2 := make([]byte, 0, 8)
	num := 1
	str = str[offset:]
	for i := range str {
		curr := str[i]
		if doi != 0 && do[doi] != curr {
			doi = 0
		}
		if do[doi] == curr {
			doi++
			if doi == len(do) {
				return &ScanResult{variant: Enable}, offset + i + 1, nil
			}
		}

		if donti != 0 && dont[donti] != curr {
			donti = 0
		}
		if dont[donti] == curr {
			donti++
			if donti == len(dont) {
				return &ScanResult{variant: Disable}, offset + i + 1, nil
			}
		}

		if mul[muli] == '%' {
			if isNumeric(curr) {
				if num == 1 {
					num1 = append(num1, curr)
				} else {
					num2 = append(num2, curr)
				}
			} else if len(num1) != 0 && len(num2) == 0 && curr == ',' {
				muli += 2
				num = 2
			} else if len(num2) != 0 && curr == ')' {
				i1, _ := strconv.Atoi(string(num1))
				i2, _ := strconv.Atoi(string(num2))
				return &ScanResult{Mul, i1, i2}, offset + i + 1, nil
			} else {
				muli = 0
				num = 1
				num1 = []byte{}
				num2 = []byte{}

				if curr == 'm' {
					muli = 1
				}
				continue
			}
		} else if muli != 0 && mul[muli] != curr {
			muli = 0
		} else if mul[muli] == curr {
			// fmt.Printf("i %v, curr %v\n", muli, string(curr))
			muli++
		}
	}
	return nil, 0, errors.New("Done")
}

func isNumeric(b byte) bool {
	return b >= '0' && b <= '9'
}
