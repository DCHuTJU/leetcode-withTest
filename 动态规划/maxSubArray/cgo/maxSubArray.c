#include <stdio.h>

int longestPalindrome(char[] s)
{
	int len = sizeof(s) / sizeof(s[0]), maxLen = 1;
	int start = 0, i, j;

	// 遍历字符串所有的子串，若子串为回文字符串则更新最长回文的长度
	for(i = 0; i< len - 1; i++) {
		for(j = i+1; j < len; j++) {
			if(isPalidrome(s, i, j)) {
				int pLen = j - i + 1;
				if(pLen > maxLen) {
					start = i;
					maxLen = pLen;
				}
			}
		}
	}
	// 从 start 开始一直到后面 maxLen 结束
	return maxLen;
}

int longestPalindromeDP(char[] s)
{
	int = s.length();
	int longestBegin = 0, maxLen = 1;
	int **P;
	int i;

	// 构造二维数组
	P = (int **)calloc(n, sizeof(int *));
	for(i = 0; i < n; i++) {
		P[i] = (int *)calloc(n, sizeof(int));
	}

	for(i = 0; i < n; i++) {
		P[i][j] = 1;
	}

	for(int i = 0; i < n-1; i++) {
		if(s[i] == s[i+1]) {
            P[i][i+1] = 1;
            longestBegin = i;
            maxLen = 2;
        }
    }

	//依次求P[i][i+2]...P[i][i+n-1]等
	int len = 3;
	for (; len <= n; ++len) {
		for (i = 0; i < n-len+1; ++i) {
			int j = i + len - 1;
			if (s[i] == s[j] && P[i+1][j-1]) {
				P[i][j] = 1;
				longestBegin = i;
				maxLen = len;
			}
		}
	}
	//释放内存
	for (i = 0; i< n; i++)
		free(P[i]);
	free(P);

	return s.substr(longestBegin, maxLen);
}