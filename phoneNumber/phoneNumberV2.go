package main

//func main() {
//	var str string
//	fmt.Scanf("%s", &str)
//
//	rlt := letterCombinationsv2(str)
//	fmt.Println(rlt)
//}

func letterCombinationsv2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	combs := []string{""}
	for i := 0; i < len(digits); i++ {
		combsLen := len(combs)
		for j := 0; j < combsLen; j++ {
			for _, letter := range m[digits[i]-'2'] {
				combs = append(combs, combs[j]+string(letter))
			}
		}
		combs = combs[combsLen:]
	}

	return combs
}