package main

import "testing"

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		str string
		rlt int
	}{
		{
			"4193 with words",
			4193,
		},
		{
			"words and 987",
			0,
		},
		{
			"-91283472332",
			-2147483648,
		},
	}
	for _, test := range tests {
		if actual := myAtoi(test.str); actual != test.rlt {
			t.Errorf("Test algorithm wrong.")
		}
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	var test = struct {
		str string
		rlt int
	}{
		"4193 with words",
		4193,
	}
	for i:=0; i<b.N; i++ {
		if actual := myAtoi(test.str); actual != test.rlt {
			b.Errorf("Benchmark algorithm is wrong.")
		}
	}
}