package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetLargestJoltage(bank string) int {
	largestJoltage := 0

	var batteries []int
	for _, c := range bank {
		batteries = append(batteries, int(c - '0'))
	}

	offset := 0
	for digit := 11; digit >= 0; digit-- {
		highest := 0
		for i := offset; i < len(batteries) - digit; i++ {
			if batteries[i] > highest {
				highest = batteries[i]
				offset = i+1
			}
		}

		if offset > 0 {
			largestJoltage *= 10
		}
		largestJoltage += highest
	}

	return largestJoltage
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
