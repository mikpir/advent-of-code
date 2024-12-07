package main

import (
	"fmt"
	"os"

	d1 "github.com/mikpir/advent-of-code/1"
	d2 "github.com/mikpir/advent-of-code/2"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arg provided")
		return
	}
	if os.Args[1] == "1" {
		file, err := os.Open("./1/data")
		if err != nil {
			fmt.Println("Could not open file")
			return
		}
		resultA := d1.RunA(file)
		fmt.Printf("A: %d\n", resultA)

		file.Seek(0, 0)
		resultB := d1.RunB(file)
		fmt.Printf("B: %d\n", resultB)
		return
	}
	if os.Args[1] == "2" {
		file, err := os.Open("./2/data")
		if err != nil {
			fmt.Println("Could not open file")
			return
		}
		resultB := d2.RunB(file)
		fmt.Printf("B: %d\n", resultB)
		return

	}

	fmt.Println("No match")
}
