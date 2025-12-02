package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dlclark/regexp2"
)

func GetInvalids(from int, to int) []int {
	var invalids []int

	re := regexp2.MustCompile(`^(\d+)\1+$`, 0)

	for id := from; id <= to; id++ {
		if matched, _ := re.MatchString(fmt.Sprintf("%d", id)); matched {
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
