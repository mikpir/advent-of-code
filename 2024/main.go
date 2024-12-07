package main

import (
	"fmt"
	"io"
	"log"
	"os"

	d1 "github.com/mikpir/advent-of-code/1"
	d2 "github.com/mikpir/advent-of-code/2"
	d3 "github.com/mikpir/advent-of-code/3"
)

type Day struct {
	afunc func(r io.Reader) int
	bfunc func(r io.Reader) int
	file  string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arg provided")
		return
	}

	days := map[string]Day{
		"1": {d1.RunA, d1.RunB, "./1/data"},
		"2": {d2.RunA, d2.RunB, "./2/data"},
		"3": {d3.RunA, d3.RunB, "./3/data"},
	}
	arg := os.Args[1]
	day, ok := days[arg]
	if !ok {
		fmt.Println("Day not found")
		return
	}

	file, err := os.Open(day.file)
	if err != nil {
		log.Fatal(err)
	}

	resultA := day.afunc(file)
	fmt.Printf("A: %d\n", resultA)

	file.Seek(0, 0)
	resultB := day.bfunc(file)
	fmt.Printf("B: %d\n", resultB)

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
