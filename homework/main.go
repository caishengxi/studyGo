package main

import (
	"fmt"
	"sort"
)

func findOnetimenumber(numbers []int) int {

	numberCount := make(map[int]int)
	for _, number := range numbers {
		numberCount[number]++
	}

	for number, count := range numberCount {
		if count == 1 {
			return number
		}
	}

	return -1
}

// 第二题：判断一个字符串是回文数
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// kuohaoMatch 判断括号是否匹配
func kuohaoMatch(s string) bool {
	arraystr := []rune{}
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			arraystr = append(arraystr, char)
		} else {
			if len(arraystr) == 0 {
				return false
			}

			if char == ')' && arraystr[len(arraystr)-1] == '(' {
				arraystr = arraystr[:len(arraystr)-1]
			} else if char == ']' && arraystr[len(arraystr)-1] == '[' {
				arraystr = arraystr[:len(arraystr)-1]
			} else if char == '}' && arraystr[len(arraystr)-1] == '{' {
				arraystr = arraystr[:len(arraystr)-1]
			} else {
				return false
			}

		}

	}
	if len(arraystr) == 0 {
		return true
	}
	return false

}

func findstrcommonprestr(strs []string) string {
	//先找出最短的字符串
	if len(strs) == 0 {
		return ""
	}

	shortest := strs[0]
	for _, str := range strs {
		if len(str) < len(shortest) {
			shortest = str
		}
	}

	for i := 0; i < len(shortest); i++ {
		char := shortest[i]
		for _, str := range strs {
			if str[i] != char {
				return shortest[:i]
			}
		}
	}

	return ""
}

func plusOne(digits []int) int {

	arraylen := len(digits)
	result := 0
	for i := 0; i < arraylen; i++ {
		result = result*10 + digits[i]
	}
	result++
	return result

}

func removeDuplicates(nums []int) int {
	pinteri := 0
	pointerj := 1
	for pointerj < len(nums) {
		if nums[pinteri] != nums[pointerj] {
			nums[pinteri+1] = nums[pointerj]
			pinteri++
		}
		pointerj++
	}
	return pinteri + 1

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 先对区间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			merged = append(merged, interval)
		}
	}
	return merged

}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	fmt.Println("homework 开始！！")

	numbers := []int{1, 20, 3, 3, 4, 20, 1}
	oneTimeNumber := findOnetimenumber(numbers)
	fmt.Println("第一题：")
	fmt.Printf("The one-time number in the array %v is: %d\n", numbers, oneTimeNumber)

	fmt.Println("第二题：")
	isPal := isPalindrome("221")
	fmt.Printf("Is 'isPal palindrome? %t\n", isPal)

	fmt.Println("第三题：")
	isMatch := kuohaoMatch("{([]){()}}")
	fmt.Printf("Is 'isMatch' match? %t\n", isMatch)

	fmt.Println("第四题：")
	commonPrefix := findstrcommonprestr([]string{"flower", "flow", "flight"})
	fmt.Printf("The common prefix of the strings is: %s\n", commonPrefix)

	fmt.Println("第五题：")
	digits := []int{6, 2, 3, 9}
	plusOneResult := plusOne(digits)
	fmt.Printf("The result of adding one to the digits %v is: %d\n", digits, plusOneResult)

	fmt.Println("第六题：")
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 5}
	newLength := removeDuplicates(nums)
	fmt.Printf("The new length of the array after removing duplicates is: %d\n", newLength)
	fmt.Printf("The modified array is: %v\n", nums[:newLength])

	fmt.Println("第七题：")
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergedIntervals := merge(intervals)
	fmt.Printf("The merged intervals of %v are: %v\n", intervals, mergedIntervals)

	fmt.Println("第八题：")
	numsForTwoSum := []int{2, 7, 11, 15, 3, 6}
	target := 10
	twoSumResult := twoSum(numsForTwoSum, target)
	fmt.Printf("The indices of the two numbers that add up to %d in the array %v are: %v\n", target, numsForTwoSum, twoSumResult)
}
