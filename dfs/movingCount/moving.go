package movingCount

func GetStep(a, b int) int {
	res := 0
	for a > 0 {
		res += a % 10
		a /= 10
	}
	for b > 0 {
		res += b % 10
		b /= 10
	}
	return res
}

func MovingCount(m, n, k int) int {
	res := 1
	flag := make([][]int, m)
	for i:=0; i<m; i++ {
		flag[i] = make([]int, n)
	}
	flag[0][0] = 1
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			// 只能从上面和左边过来
			if (i-1>=0 && flag[i-1][j]==1) || (j-1>=0 && flag[i][j-1]==1) {
				if GetStep(i, j) <= k {
					res++
					flag[i][j] = 1
				}
			}
		}
	}
	return res
}