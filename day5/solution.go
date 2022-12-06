package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	splittedLines := strings.Split(string(input), "\n\n")
	crates := strings.Split(splittedLines[0], "\n")
	keys := crates[len(crates)-1]

	heap1, heap2 := map[rune]string{}, map[rune]string{}
	for _, s := range crates {
		for i, r := range s {
			if unicode.IsLetter(r) {
				heap1[rune(keys[i])] += string(r)
				heap2[rune(keys[i])] += string(r)
			}
		}
	}

	for _, s := range strings.Split(strings.TrimSpace(splittedLines[1]), "\n") {
		var quantity int
		var from, to rune
		fmt.Sscanf(s, "move %d from %c to %c", &quantity, &from, &to)

		for i := 0; i < quantity; i++ {
			heap1[to] = heap1[from][:1] + heap1[to]
			heap1[from] = heap1[from][1:]
		}

		heap2[to] = heap2[from][:quantity] + heap2[to]
		heap2[from] = heap2[from][quantity:]
	}

	part1, part2 := "", ""
	for _, r := range strings.ReplaceAll(keys, " ", "") {
		part1 += heap1[r][:1]
		part2 += heap2[r][:1]
	}
	fmt.Println(part1)
	fmt.Println(part2)

}
