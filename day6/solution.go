package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	data := strings.TrimSpace(string(input))
	fmt.Println(findMarker(data, 4))
	fmt.Println(findMarker(data, 14))
}

func findMarker(characters string, markerSize int) int {
	for i := markerSize; i <= len(characters); i++ {
		m := map[rune]struct{}{}
		for _, r := range characters[i-markerSize : i] {
			m[r] = struct{}{}
		}
		if len(m) >= markerSize {
			return i
		}
	}
	return -1
}
