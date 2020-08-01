package main

import (
	"fmt"
	"strconv"
)

func GetList(datalist1, datalist2 [][]string) [][]string {
	n := len(datalist1)
	datalist3 := make([][]string, 0)
	for i:=0; i<n; i++ {
		datalist3 = append(datalist3, []string{datalist1[i][1]})
		sum := 0
		for j:=2; j<7; j+=2 {
			datalist3[i] = append(datalist3[i], datalist2[i][j])
			num, _ := strconv.Atoi(datalist2[i][j])
			sum += num
		}
		datalist3[i] = append(datalist3[i], strconv.Itoa(sum))
	}
	return datalist3
}

func main() {
	datalist1 := [][]string{{"1", "zhangwei01"}}
	datalist2 := [][]string{{"1", "01", "200", "02", "150", "03", "196"}}
	fmt.Println(GetList(datalist1, datalist2))
}