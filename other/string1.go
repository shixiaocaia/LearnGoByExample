package main

import "strings"

// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

func isUniqueString(str string) bool {
	if len(str) > 3000 {
		return false
	}

	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}

	return true
}

func isUniqueString2(str string) bool {
	if len(str) > 3000 {
		return false
	}

	for _, v := range str {
		if v > 127 {
			return false
		}
		// 判断str中包含v的个数
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}

func main() {
	t1 := "abcde"
	t2 := "abcda"

	println(isUniqueString(t1))
	println(isUniqueString2(t2))
}
