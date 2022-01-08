package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 Remove Duplicates from Sorted Array
*/
func removeDuplicates(nums []int) int {
	keys := make(map[int]bool)
	list := []int{}

	for _, v := range nums {
		if _, value := keys[v]; !value {
			keys[v] = true
			list = append(list, v)
		}
	}
	sort.Ints(list)
	copy(nums, list)

	return len(list)
}

/*
 Best Time to Buy and Sell Stock II
*/

func maxProfit(prices []int) int {
	var sum int
	for i := 0; i < len(prices)-1; i++ {
		if (prices[i+1] - prices[i]) > 0 {
			sum = sum + (prices[i+1] - prices[i])
		}
	}

	return sum
}

/*
 Rotate Array
*/

func rotateArray(nums []int, k int) {
	s := make([]int, len(nums))

	for i := 0; i < k; i++ {
		s[0] = nums[len(nums)-1]
		for j := 0; j < len(nums)-1; j++ {
			s[j+1] = nums[j]
		}
		copy(nums, s)
	}
}

/*
 Contains Duplicate
*/

func containsDuplicate(nums []int) bool {
	for k, v := range nums {
		for k2, v2 := range nums {
			if v == v2 && k != k2 {
				return true
			}
		}
	}

	return false
}

/*
 Single Number
*/

func singleNumber(nums []int) int {
	keys := make(map[int]bool)

	for k, v := range nums {
		for k2, v2 := range nums {
			if v == v2 && k != k2 {
				keys[v] = true
			}
		}

		if _, check := keys[v]; !check {
			return v
		}
	}

	return 0
}

/*
 Intersection of Two Arrays II
*/

func intersect(nums1 []int, nums2 []int) []int {
	var s []int
	keys := make(map[int]bool)

	for _, v := range nums2 {
		for _, v2 := range nums1 {
			if v == v2 {
				keys[v] = true
			}
		}
	}

	for k, _ := range keys {
		var count1 int
		var count2 int

		for _, v := range nums1 {
			if v == k {
				count1++
			}
		}

		for _, v := range nums2 {
			if v == k {
				count2++
			}
		}

		if count1 <= count2 {
			for i := 0; i < count1; i++ {
				s = append(s, k)
			}
		} else {
			for i := 0; i < count2; i++ {
				s = append(s, k)
			}
		}
	}

	return s
}

/*
 Plus One
*/

func plusOne(digits []int) []int {
	var s, r []int
	var flag bool
	b := 1

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+b == 10 {
			digits[i] = 0
			b = 1
			if i == 0 {
				flag = true
			}
		} else {
			if i == len(digits)-1 || b == 1 {
				digits[i] = digits[i] + 1
				b = 0
			} else if b == 1 && i == 0 {
				digits[i] = digits[i] + 1
				b = 0
			} else {
				digits[i] = digits[i]
				b = 0
			}
		}
		s = append(s, digits[i])

		if flag {
			s = append(s, 1)
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}

	return r
}

/*
 Move Zeroes
*/

func moveZeroes(nums []int) {
	var keys, rm []int
	var count int

	for k, v := range nums {
		if v == 0 {
			count++
			keys = append(keys, k)
		}
	}

	for k, v := range keys {
		rm = append(rm, v-k)
	}

	for _, v := range rm {
		nums = append(nums[:v], nums[v+1:]...)
	}

	for i := 0; i < len(rm); i++ {
		nums = append(nums, 0)
	}
}

/*
 Two Sum
*/

func twoSum(nums []int, target int) []int {
	var keys []int
	for k, _ := range nums {
		dif := target - nums[k]
		for k2, v2 := range nums {
			if dif == v2 && k != k2 {
				keys = append(keys, k)
			}
		}
	}

	return keys
}

/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
*/

func isValidSudoku(board [][]byte) bool {
	s2 := make([][]int, 9)
	for k, v := range board {
		sc := make(map[int]int)
		for k2, v2 := range v {
			byteToInt, _ := strconv.Atoi(string(v2))
			if k < 3 {
				if k2 < 3 {
					s2[0] = append(s2[0], byteToInt)
				}

				if k2 < 6 && k2 >= 3 {
					s2[1] = append(s2[1], byteToInt)
				}
				if k2 < 9 && k2 >= 6 {
					s2[2] = append(s2[2], byteToInt)
				}
			}

			if k < 6 && k >= 3 {
				if k2 < 3 {
					s2[3] = append(s2[3], byteToInt)
				}
				if k2 < 6 && k2 >= 3 {
					s2[4] = append(s2[4], byteToInt)
				}
				if k2 < 9 && k2 >= 6 {
					s2[5] = append(s2[5], byteToInt)
				}
			}

			if k < 9 && k >= 6 {
				if k2 < 3 {
					s2[6] = append(s2[6], byteToInt)
				}
				if k2 < 6 && k2 >= 3 {
					s2[7] = append(s2[7], byteToInt)
				}
				if k2 < 9 && k2 >= 6 {
					s2[8] = append(s2[8], byteToInt)
				}
			}

			if byteToInt > 0 {
				sc[byteToInt]++
			}

			if sc[byteToInt] > 1 {
				fmt.Println(sc)
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		sc := make(map[int]int)
		for _, v := range board {
			byteToInt, _ := strconv.Atoi(string(v[i]))
			if byteToInt > 0 {
				sc[byteToInt]++
			}

			if sc[byteToInt] > 1 {
				fmt.Println(sc)
				return false
			}
		}
	}

	for _, v := range s2 {
		sc := make(map[int]int)
		for _, v2 := range v {
			byteToInt := v2
			if byteToInt > 0 {
				sc[byteToInt]++
			}

			if sc[byteToInt] > 1 {
				return false
			}
		}
	}

	return true
}

/*
 翻轉陣列
   1, 2, 3         7, 4, 1
   4, 5, 6    =>   8, 5, 2
   7, 8, 9         9, 6, 3
*/

func rotate(matrix [][]int) {
	l := len(matrix)
	newMatrix := make([][]int, l)
	for i := 0; i < l; i++ {
		newMatrix[i] = make([]int, l)
	}

	for k, v := range matrix {
		vl := l - 1 - k
		for k1, v1 := range v {
			newMatrix[k1][vl] = v1
		}
	}

	copy(matrix, newMatrix)
}
