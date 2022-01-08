package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Array  Rotate Image
	var input = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(input)

	// ========================================================================

	// var isValidSudokuInput = [][]string{
	// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }
	var isValidSudokuInput = [][]string{
		{".", ".", ".", ".", "5", ".", ".", "1", "."},
		{".", "4", ".", "3", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "3", ".", ".", "1"},
		{"8", ".", ".", ".", ".", ".", ".", "2", "."},
		{".", ".", "2", ".", "7", ".", ".", ".", "."},
		{".", "1", "5", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "2", ".", ".", "."},
		{".", "2", ".", "9", ".", ".", ".", ".", "."},
		{".", ".", "4", ".", ".", ".", ".", ".", "."}}
	s := make([][]byte, 0)
	for _, v := range isValidSudokuInput {
		stringByte := strings.Join(v, "")
		s = append(s, []byte(stringByte))
	}

	a := isValidSudoku(s)
	fmt.Println(a)
}

/*
 尋找一字串 不重複字母字串 非子集合字串(意思非排列組合)
*/

func lengthOfLongestSubstring(str string) int {
	// pwwkew
	m, max, left := make(map[rune]int), 0, 0
	for idx, c := range str {
		if _, okay := m[c]; okay == true && m[c] >= left {
			if idx-left > max {
				max = idx - left
			}
			left = m[c] + 1
		}
		m[c] = idx
	}
	if len(str)-left > max {
		max = len(str) - left
	}
	return max
}

/*
 尋找多字串 最小相同前綴字串
*/

func longestCommonPrefix(strs []string) string {
	m := make(map[string]int)
	max, min := 0, 0
	for _, v := range strs {
		if len(v) >= max {
			max = len(v)
		}
	}

	for _, v := range strs {
		if len(v) <= max {
			min = len(v)
			max = min
		}
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for _, v := range strs {
		for i := min; i >= 0; i-- {
			str := string([]rune(v)[:i])
			m[str]++
			if m[str] == len(strs) {
				return str
			}
		}
	}

	return "1"
}

/*
 MCMXCIV 羅馬數字轉換
*/

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var (
		ss    string
		count int
	)

	for i := len(s); i >= 1; i-- {
		str := string([]rune(s)[i-1 : i])
		if str == ss {
			count -= m[str]
			ss = ""
			continue
		}

		if str == "V" || str == "X" {
			count += m[str]
			ss = "I"
		}

		if str == "L" || str == "C" {
			count += m[str]
			ss = "X"
		}

		if str == "D" || str == "M" {
			count += m[str]
			ss = "C"
		}

		if str == "I" {
			count += m[str]
		}
	}

	return count
}

/*
 121 -121 字串左->右 是否等於 右->左
*/

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	s1 := make([]string, 0)
	s2 := make([]string, 0)

	for i := len(s); i >= 1; i-- {
		str := string([]rune(s)[i-1 : i])
		s1 = append(s1, str)
	}

	for i := 1; i <= len(s); i++ {
		str := string([]rune(s)[i-1 : i])
		s2 = append(s2, str)
	}

	for i := 0; i < len(s); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
