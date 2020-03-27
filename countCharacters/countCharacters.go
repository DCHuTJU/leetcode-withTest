package main

import (
	"fmt"
	"strings"
)

func countCharacters(words []string, chars string) int {
	// chars中的只能用一次
	var wholeLength int = 0
	// 合理使用 goto 语句
outloop:
	for _, v1 := range words {
		// 拿每一个单词。
		for _, v2 := range v1 {
			// 取字母
			if strings.Count(v1, string(v2)) > strings.Count(chars, string(v2)) {
				// 因为每一个word只能用一次chars中的元素
				// 如果在单词中某个字母的数量大于词库中的数量，说明不成
				continue outloop
			}
		}
		// 如果这个循环走完了，能走到这里，说明这个单词里的字母都在chars中
		wholeLength += len(v1)
	}
	return wholeLength
}

func main() {
	var words = []string{"cat", "bt", "hat", "tree"}
	var chars = "atach"
	rlt := countCharacters(words, chars)
	fmt.Println(rlt)
}
