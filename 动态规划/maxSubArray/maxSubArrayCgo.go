package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int isPalidrome(char* s, int start, int end)
{
	for(; start<end; ++start, --end) {
		printf("This is string: %s\n", s);
		if(s[start] != s[end]) {
			printf("It doesn't has a subarray");
			return 0;
		}
	}
	return 1;
}
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func isHasSubArray(s string) int {
	str := C.CString(s)
	defer C.free(unsafe.Pointer(str))
	rlt := C.isPalidrome(str, 0, C.int(len(s)-1))
	return int(rlt)
}

func main() {
	var s = "aba"
	rlt := isHasSubArray(s)
	fmt.Println(rlt)
}