package main

import (
	"bufio"
	"fmt"
	"os"
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

	list_of_fresh := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			list_of_fresh = false
			continue
		} else {
			if list_of_fresh {
				parts := strings.Split(line, "-")
				from, _ := strconv.Atoi(parts[0])
				to, _ := strconv.Atoi(parts[1])
				ingredientRanges = append(ingredientRanges, IngredientRange{from, to})
			} else {
				ingredient_id, _ := strconv.Atoi(line)
				for _, r := range ingredientRanges {
					if ingredient_id >= r.from && ingredient_id <= r.to {
						result += 1
						break
					}
				}
			}
		}
	}

	fmt.Print(result)
}
