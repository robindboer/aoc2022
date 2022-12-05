package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	splittedLines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	calories := make([]int, len(splittedLines))
	for i, s := range splittedLines {
		for _, s := range strings.Fields(s) {
			cal, _ := strconv.Atoi(s)
			calories[i] += cal
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println(calories[0])
	fmt.Println(calories[0] + calories[1] + calories[2])
}
