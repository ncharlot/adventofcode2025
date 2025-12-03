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
    	{"987654321111111", 98},
     	{"811111111111119", 89},
      	{"234234234234278", 78},
       	{"818181911112111", 92},

        // More cases
        {"3342222252323515234222627222212413112534241272422423223227282252222254242224333544215112522943221225", 95},
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
