package main

import "testing"

func BenchmarkGetList(b *testing.B) {
	datalist1 := [][]string{{"1", "zhangwei01"}}
	datalist2 := [][]string{{"1", "01", "200", "02", "150", "03", "196"}}
	for i:=0; i<b.N; i++ {
		GetList(datalist1, datalist2)
	}
}
