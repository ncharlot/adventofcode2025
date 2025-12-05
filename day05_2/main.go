package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

type IngredientRange struct {
	from int
   	to int
}

func main() {
	result := 0

	var ingredientRanges []IngredientRange

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])
		ingredientRanges = append(ingredientRanges, IngredientRange{from, to})
	}

	slices.SortFunc(ingredientRanges, func(a, b IngredientRange) int {
		return a.from - b.from
	})

	unionedRanges := []IngredientRange{ingredientRanges[0]}
	ingredientRanges = slices.Delete(ingredientRanges, 0, 1)

	for _, r := range ingredientRanges {
		lastIdx := len(unionedRanges) - 1
		if unionedRanges[lastIdx].to >= r.from {
			if unionedRanges[lastIdx].to < r.to {
				unionedRanges[lastIdx].to = r.to
			}
		} else {
			unionedRanges = append(unionedRanges, r)
		}
	}

	for _, r := range unionedRanges {
		result += r.to - r.from + 1
	}

	fmt.Print(result)
}
