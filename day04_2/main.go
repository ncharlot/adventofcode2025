package main

import (
	"bytes"
	"fmt"
	"os"
)

const EMPTY = 46
const ROLL = 64

func main() {
	result := 0

	diagram, _ := os.ReadFile("input.txt")

	// get line length
	ll := bytes.IndexByte(diagram, 10)

	// remove line returns
	diagram = bytes.ReplaceAll(diagram, []byte{10}, []byte{})

	// get diagram length
	dl := len(diagram)

	for i := 0; i < dl; i++ {
		if diagram[i] == ROLL {
			rolls_arround := 0
			adj_pos := []int{
				i-ll-1, i-ll, i-ll+1,
			       i-1,          i+1,
				i+ll-1, i+ll, i+ll+1,
			}

			for _, a := range adj_pos {
				if a >= 0 && a < dl && diagram[a] == ROLL && (a%ll) - (i%ll) >= -1 && (a%ll) - (i%ll) <= 1 {
					rolls_arround += 1
				}
			}

			if rolls_arround < 4 {
				result += 1
				diagram[i] = EMPTY
				i = max(-1, i - ll - 2)
			}
		}
	}

	fmt.Print(result)
}
