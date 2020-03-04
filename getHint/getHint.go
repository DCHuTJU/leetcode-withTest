package main

import (
	"fmt"
	gsf "github.com/dengchengH/go-simple-func/compare"
	"strconv"
)

func getHint(secret string, guess string) string {
	a, b := 0, 0
	mapS, mapG := make([]int, 10), make([]int, 10)
	for i := range secret {
		// 使用的是ASCII
		tmp, charGuess := secret[i], guess[i]
		if tmp == guess[i] {
			a++
		} else {
			mapS[tmp-'0']++
			mapG[charGuess-'0']++
		}
	}
	for i:=0; i<10; i++ {
		// 找到重叠的
		b += gsf.Min(mapS[i], mapG[i])
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func main() {
	secret := "1807"
	guess := "7810"
	rlt := getHint(secret, guess)
	fmt.Println(rlt)
}
