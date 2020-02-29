### 三角形最小路径和

同样是**动态规划**的问题。因为三角形最后一层中的最短路径是由前一层决定的。

#### 状态定义
同样使用一个与**原三角形一样大小**（这个很关键，因为在这上面可以优化，后面会提到）的数组作为最终取值的存放。
由于题目中有限制，只能移动到下一行中相邻的节点上。
那么状态转换函数应为：
```go
dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
```
但是，这个函数的适用是有条件的，可以考虑一下，当函数的边界上会怎么样？
* 因为用二维数组存储，肯定第一列就是一条边界，那么这条边界上的值只会收其同列上的前一个值得影响，即：
```go
dp[i][j] = dp[i-1][j] + triangle[i][j]
```
* 当处于另一条边界上，即数组每一行数组的最后一个值时，还会面临相似的问题，此时该边界上的值只能是收受左前方元素的影响，即：
```go
dp[i][j] = dp[i-1][j-1] + triangle[i][j]
```
最后对最后一行进行遍历，即可获得最终结果，最小值。
按照这种思路把代码顺下来没问题，但是前面提到过，`dp`数组是一个大小与**原三角形大小一样**的数组，那么既然前面的值我们后面会用到，但`triangle`本身的值在进行计算后就不再需要了，那么就可以考虑让原来的`triangle`既承担`triangle`的功能，同时也承担起`dp`的功能。
即不需要新建立`dp`数组，而在原来的`triangle`数组上操作，可以大大降低空间复杂的。
```go
func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}
	result := 1 << 31-1
	triangle[1][0] = triangle[1][0] + triangle[0][0]
	triangle[1][1] = triangle[1][1] + triangle[0][0]
	for i:=2; i<len(triangle); i++ {
		for j:=0; j<len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}
		}
	}
	for _, k := range triangle[len(triangle)-1] {
		result = min(result, k)
	}
	return result
}
```

