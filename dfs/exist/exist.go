package exist

func Exist(board [][]byte, word string) bool {

	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[0]); j++ {
			if Dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func Dfs(board [][]byte, i, j, k int, word string) bool {
	if k == len(word) { return true }
	if i<0 || j<0 || i==len(board) || j==len(board[0]) { return false }

	if board[i][j] == word[k] {
		tmp := word[k]
		board[i][j] = ' '
		if Dfs(board, i-1, j, k+1, word) || Dfs(board, i+1, j, k+1, word) || Dfs(board, i, j-1, k+1, word) || Dfs(board, i, j+1, k+1, word) {
			return true
		} else {
			board[i][j] = tmp
		}
	}
	return false
}
