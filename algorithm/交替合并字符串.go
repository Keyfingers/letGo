package main

import (
	"fmt"
	"strings"
)

/*
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
返回 合并后的字符串 。

输入：word1 = "ab", word2 = "pqrs"
输出："apbqrs"
解释：注意，word2 比 word1 长，"rs" 需要追加到合并后字符串的末尾。
word1：  a   b
word2：    p   q   r   s
合并后：  a p b q   r   s
*/
func main() {
	str := mergeStrings("ab", "efd")
	fmt.Println(str)
}

func mergeStrings(word1 string, word2 string) string {
	var result strings.Builder
	i, j := 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			result.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			result.WriteByte(word2[j])
			j++
		}
	}

	return result.String()
}
