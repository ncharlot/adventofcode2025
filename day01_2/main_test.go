package main

import (
	"fmt"
	"testing"
)

func TestGetInvalids(t *testing.T) {
    var tests = []struct {
   		from int
     	to int
      	invalids []int
    }{
    	// Example input
	    {50, "L68", 0, 82, 1},
	    {82, "L30", 1, 52, 1},
	    {52, "R48", 1,  0, 2},
	    { 0, "L5",  2, 95, 2},
	    {95, "R60", 2, 55, 3},
	    {55, "L55", 3,  0, 4},
	    { 0, "L1",  4, 99, 4},
	    {99, "L99", 4,  0, 5},
	    { 0, "R14", 5, 14, 5},
	    {14, "L82", 5, 32, 6},

		// More cases
    	{50,   "L0", 0, 50, 0},
        {50,   "R0", 0, 50, 0},
       	{50, "L100", 0, 50, 1},
        {50, "R100", 0, 50, 1},
       	{50, "L200", 0, 50, 2},
        {50, "R200", 0, 50, 2},
        {0,  "L100", 0,  0, 1},
        {0,  "R100", 0,  0, 1},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%s,%d", tt.dial, tt.rotation, tt.result)
        t.Run(testname, func(t *testing.T) {
	        new_dial, new_result := rotate(tt.dial, tt.rotation, tt.result)
	        if new_dial != tt.new_dial || new_result != tt.new_result {
	            t.Errorf("got %d,%d, want %d,%d", new_dial, new_result, tt.new_dial, tt.new_result)
	        }
        })
    }
}
