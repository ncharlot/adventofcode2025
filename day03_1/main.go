package main

import (
	"bufio"
	"fmt"
	"slices"
	"os"
	// "strings"
)

func GetLargestJoltage(bank string) int {
	var joltages []int

	var batteries []int
	for _, c := range bank {
		batteries = append(batteries, int(c - '0'))
	}

	for i := 0; i < len(batteries) - 1; i++ {
		joltages = append(joltages, batteries[i]*10 + slices.Max(batteries[i+1:]))
	}

	return slices.Max(joltages)
}

func main() {
	result := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result += GetLargestJoltage(scanner.Text())
	}

	fmt.Print(result)
}
