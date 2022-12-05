package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	part1, part2 := 0, 0
	for _, s := range strings.Fields(strings.TrimSpace(string(input))) {
		var startSection1, endSection1, startSection2, endSection2 int
		fmt.Sscanf(s, "%d-%d,%d-%d", &startSection1, &endSection1, &startSection2, &endSection2)

		// first check if section 1 contains section 2 or section 2 contains section 1
		if startSection1 <= startSection2 && endSection1 >= endSection2 || startSection1 >= startSection2 && endSection1 <= endSection2 {
			part1++
		}

		// check if there is any overlap at all
		// if section1 starts before section 2 ended or section 2 started before section 1 ended
		if startSection1 <= endSection2 && endSection1 >= startSection2 {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}
