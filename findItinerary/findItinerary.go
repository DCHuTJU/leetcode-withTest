package main

import (
	"fmt"
	"sort"
	_ "net/http/pprof"
)

func findItinerary(tickets [][]string) []string {
	var (
		m = map[string][]string{}
		res = []string{}
	)

	for _, ticket := range tickets {
		str, end := ticket[0], ticket[1]
		m[str] = append(m[str], end)
	}

	for key := range m {
		sort.Strings(m[key])
	}

	var dfs func(curr string)

	dfs = func(curr string) {
		for {
			if v, ok := m[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := m[curr][0]
			m[curr] = m[curr][1:]
			dfs(tmp)
		}
		res = append([]string{curr}, res...)
	}
	dfs("JFK")

	return res
}

func main() {
	var tickets = [][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}
	fmt.Println(findItinerary(tickets))
}
