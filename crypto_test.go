package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSlice(t *testing.T) {
	tcs := []struct {
		s      string
		result []int
	}{
		{
			s:      "aAbBcC",
			result: []int{0, 1, 0, 1, 0, 1},
		},
		{
			s:      "AAbbCC",
			result: []int{1, 1, 0, 0, 1, 1},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.s, func(t *testing.T) {
			bits := bitSlice(tc.s)
			assert.Equal(t, tc.result, bits)
		})
	}
}

func TestBitsToNums(t *testing.T) {
	tcs := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{0, 1, 0, 1, 0},
			out: []int{10},
		},
		{
			in:  []int{0, 1, 0, 1, 1},
			out: []int{11},
		},
		{
			in:  []int{1, 1, 1, 1, 1},
			out: []int{31},
		},
		{
			in:  []int{1, 1, 1, 1, 1, 1},
			out: nil,
		},
		{
			in:  []int{},
			out: []int{},
		},
		{
			in:  []int{1, 1, 1, 1, 1, 0, 1, 0, 1, 0},
			out: []int{31, 10},
		},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc.in), func(t *testing.T) {
			nums := bitsToNums(tc.in)
			assert.Equal(t, tc.out, nums)
		})
	}
}

func TestNumsToS(t *testing.T) {
	tcs := []struct {
		nums   []int
		result string
	}{
		{
			nums:   []int{0, 1, 2, 3},
			result: "abcd",
		},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			s := numsToS(tc.nums)
			assert.Equal(t, tc.result, s)
		})
	}
}
