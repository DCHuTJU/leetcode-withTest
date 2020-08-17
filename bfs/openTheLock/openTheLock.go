package main

import "fmt"

var (
	deadends = []string{"8888"}
	target = "0009"
)

func openLock(deadends []string, target string) int {
	deads := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		deads[deadends[i]] = true
	}
	// 防止走回头路
	visited := make(map[string]bool)
	queue := make([]string, 0)
	step := 0
	start := "0000"
	queue = append(queue, start)
	visited[start] = true

	for len(queue) != 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if deads[cur] == true {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				up := PlusOne(cur, j)
				if deads[up] == false {
					queue = append(queue, up)
					visited[up] = true
				}
				down := MinueOne(cur, j)
				if deads[down] == false {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}

		step++
	}
	return -1
}

func PlusOne(cur string, pos int) string {
	curByte := []byte(cur)
	if curByte[pos] == '9' {
		curByte[pos] = '0'
	} else {
		curByte[pos] += 1
	}
	return string(curByte)
}

func MinueOne(cur string, pos int) string {
	curByte := []byte(cur)
	if curByte[pos] == '0' {
		curByte[pos] = '9'
	} else {
		curByte[pos] -= 1
	}
	return string(curByte)
}

func main() {
	rlt := openLock(deadends, target)
	fmt.Println("The result is: ", rlt)
}