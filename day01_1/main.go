package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dial := 50
	result := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rotation := scanner.Text()

		direction := string(rotation[0])
		distance, _ := strconv.Atoi(rotation[1:])

		if direction == "R" {
			dial += distance
		} else {
			dial -= distance
		}

		dial = dial % 100

		if dial == 0 {
			result += 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Print(result)
}
