package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func modulo(nb int, div int) int {
	return ((nb % div) + div) % div
}

func rotate(dial int, rotation string, result int) (int, int) {
	direction := string(rotation[0])
	distance, _ := strconv.Atoi(rotation[1:])

	if direction == "L" {
		if distance >= dial {
			if dial != 0 {
				result += 1 - (dial - distance) / 100
			} else {
				result += - (dial - distance) / 100
			}
		}
		dial -= distance
	} else {
		result += (dial + distance) / 100
		dial += distance
	}

	dial = modulo(dial, 100)

	return dial, result
}

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
		dial, result = rotate(dial, rotation, result)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Print(result)
}
