package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetInvalids(t *testing.T) {
    var tests = []struct {
   		from int
     	to int
      	invalids []int
    }{
    	// Example input
     	{        11,         22,     []int{11, 22}},
		{        95,        115,         []int{99}},
		{       998,       1012,       []int{1010}},
		{1188511880, 1188511890, []int{1188511885}},
		{    222220,     222224,     []int{222222}},
		{   1698522,    1698528,           []int{}},
		{    446443,     446449,     []int{446446}},
		{  38593856,   38593862,   []int{38593859}},
		{    565653,     565659,           []int{}},
		{ 824824821,  824824827,           []int{}},
		{2121212118, 2121212124,           []int{}},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d-%d", tt.from, tt.to)
        t.Run(testname, func(t *testing.T) {
	        invalids := GetInvalids(tt.from, tt.to)
	        if !slices.Equal(invalids, tt.invalids) {
				t.Errorf("got %v, want %v", invalids, tt.invalids)
	        }
        })
    }
}
