package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(f)

	var heights []int

	for reader.Scan() {
		text := reader.Text()
		readHeight, err := strconv.Atoi(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not convert:%s", text)
			continue
		}

		heights = append(heights, readHeight)
	}

	numberOfIncreasingHeights := 0

	for idx := 1; idx < len(heights); idx++ {
		if heights[idx] > heights[idx-1] {
			numberOfIncreasingHeights++
		}
	}

	fmt.Printf("Number of increasing heights: %d\n", numberOfIncreasingHeights)

}
