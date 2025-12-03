package main

import (
	"fmt"
	"testing"
)

func TestGetLargestJoltage(t *testing.T) {
    var tests = []struct {
  		bank string
   		joltage int
    }{
    	// Example input
    	{"987654321111111", 987654321111},
     	{"811111111111119", 811111111119},
      	{"234234234234278", 434234234278},
       	{"818181911112111", 888911112111},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%s", tt.bank)
        t.Run(testname, func(t *testing.T) {
	        joltage := GetLargestJoltage(tt.bank)
	        if joltage != tt.joltage {
				t.Errorf("got %d, want %d", joltage, tt.joltage)
	        }
        })
    }
}
