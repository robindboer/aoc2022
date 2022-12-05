package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	splittedStrings := strings.Fields(strings.TrimSpace(string(input)))

	part1, part2 := 0, 0
	for i, contents := range splittedStrings {
		part1 += commonItemTypeValue(contents[:len(contents)/2], contents[len(contents)/2:])

		if i%3 == 0 {
			part2 += commonItemTypeValue(splittedStrings[i : i+3]...)
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func commonItemTypeValue(strs ...string) int {
loop:
	for _, r := range strs[0] {
		for _, s := range strs[1:] {
			if !strings.ContainsRune(s, r) {
				continue loop
			}
		}
		return strings.IndexRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r) + 1
	}
	return 0
}
