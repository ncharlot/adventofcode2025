package main

import (
	"fmt"
	"os"
	"strings"
)

func GetInvalids(from int, to int) []int {
	var invalids []int

	for id := from; id <= to; id++ {
		id_as_str := fmt.Sprintf("%d", id)
		length := len(id_as_str)

		if length%2 != 0 {
			continue
		}

		if id_as_str[0:length/2] == id_as_str[length/2:] {
			invalids = append(invalids, id)
		}
	}

	return invalids
}

func main() {
	result := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	line := string(data)
	for ran := range strings.SplitSeq(line, ",") {
		fromStr, toStr, _ := strings.Cut(ran, "-")
		var from, to int
		fmt.Sscanf(fromStr, "%d", &from)
		fmt.Sscanf(toStr, "%d", &to)
		for _, invalid := range GetInvalids(from, to) {
			result += invalid
		}
	}

	fmt.Println(result)
}
