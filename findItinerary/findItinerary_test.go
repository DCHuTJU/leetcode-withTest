package main

import "testing"

var tickets = [][]string{
	{"MUC", "LHR"},
	{"JFK", "MUC"},
	{"SFO", "SJC"},
	{"LHR", "SFO"},
}

func BenchmarkFindItinerary(b *testing.B) {
	for i:=0; i<b.N; i++ {
		findItinerary(tickets)
	}
}
