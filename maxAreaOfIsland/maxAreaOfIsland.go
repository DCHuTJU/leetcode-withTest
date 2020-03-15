package main

import "fmt"

var isv [][]bool                        // 存储每个位置是否访问过
var dx = [4]int{-1,0,1,0}               // 控制遍历方向
var dy = [4]int{0,-1,0,1}
var n,m int                             // 存储地图边界

func judge(grid [][]int, x,y int) bool {// 判断当前位置能不能走
	if x<0 || x>=n || y<0 || y>=m {     // 不能走
		return true
	}
	if grid[x][y] == 0 {
		return true
	}
	if isv[x][y] {
		return true
	}
	return false
}

func dfs(grid [][]int, x,y int) int {   // 从当前位置开始遍历连接的土地
	if judge(grid,x,y) {                // 遇到边界的时候返回
		return 0
	}
	isv[x][y] = true                    // 记录一下当前位置走过了
	curArea := 1
	for i:=0; i<4; i++ {                // 遍历当前位置的四个方向，并累计与这四个方向连接的土地面积
		curArea += dfs(grid, x+dx[i], y+dy[i])
	}
	return curArea
}

func maxAreaOfIsland1(grid [][]int) int {
	MaxArea := 0                        // 记录最大面积
	isv = [][]bool{}
	n,m = len(grid), len(grid[0])
	for range make([]int, n) {          // 初始化空的 isv 二维数组
		isv = append(isv, make([]bool, m))
	}
	for i:=0; i<n; i++ {                // 遍历地图
		for j:=0; j<m; j++ {
			if grid[i][j]==1 && !isv[i][j] {   // 如果该位置没被访问过，且是“土地”，那么计算当前位置所在的土地面积
				curArea := dfs(grid,i,j)
				if curArea > MaxArea {
					MaxArea = curArea
				}
			}
		}
	}
	return MaxArea
}

func maxAreaOfIsland2(grid [][]int) int {
	area := 0
	var traverse func(i, j int) int
	MaxI, MaxJ := len(grid), 0
	if MaxI > 0 {
		MaxJ = len(grid[0])
	}
	traverse = func(i, j int) int {
		var ae = 0
		if i < 0 || i > MaxI-1 {
			return ae
		}
		if j < 0 || j > MaxJ-1 {
			return ae
		}
		v := grid[i][j]
		grid[i][j] = -1
		if v != 1 {
			return ae
		}
		ae = v + traverse(i, j+1) + traverse(i+1, j) + traverse(i-1, j) + traverse(i, j-1)
		return ae
	}
	for i:=0; i<MaxI; i++ {
		for j:=0; j<MaxJ; j++ {
			if grid[i][j] != 1 {
				continue
			}
			tmp := traverse(i, j)
			if tmp > area {
				area = tmp
			}
		}
	}
	return area
}

func main() {
	var area = [][]int {
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}
	rlt := maxAreaOfIsland1(area)
	fmt.Println(rlt)
	rlt = maxAreaOfIsland2(area)
	fmt.Println(rlt)
}
