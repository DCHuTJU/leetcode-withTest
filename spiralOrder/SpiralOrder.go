package main

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 {
		return result
	}
	left, right, up, down := 0, len(matrix[0])-1, 0, len(matrix)-1
	var x, y int
	for left <= right && up <= down {
		for y = left; y <= right && avoid(left, right, up, down); y++ {
			result = append(result, matrix[x][y])
		}
		y--
		up++
		for x = up; x <= down && avoid(left, right, up, down); x++ {
			result = append(result, matrix[x][y])
		}
		x--
		right--
		for y = right; y >= left && avoid(left, right, up, down); y-- {
			result = append(result, matrix[x][y])
		}
		y++
		down--
		for x = down; x >= up && avoid(left, right, up, down); x-- {
			result = append(result, matrix[x][y])
		}
		x++
		left++
	}
	return result
}


func avoid(left, right, up, down int) bool {
	return left <= right && up <= down
}

func main() {
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result := SpiralOrder(matrix)
	fmt.Println(result)
}

/*
代码覆盖率
F:\algorithm\spiralOrder>go test SpiralOrder_test.go SpiralOrder.go -coverprofile=c.out
ok      command-line-arguments  0.048s  coverage: 85.2% of statements

F:\algorithm\spiralOrder>go tool cover -html=c.out
*/