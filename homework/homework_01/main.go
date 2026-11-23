package main

import (
	"fmt"
	"sort"
)

func findOnetimenumber(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	numberCount := make(map[int]int)
	for _, number := range numbers {
		numberCount[number]++
	}

	uniqueNumber := 0
	foundUnique := false
	for number, count := range numberCount {
		if count == 1 {
			if foundUnique {
				return -1
			}
			uniqueNumber = number
			foundUnique = true
		} else if count != 2 {
			return -1
		}
	}

	if !foundUnique {
		return -1
	}
	return uniqueNumber
}

// 第二题：判断一个字符串是回文数
func isPalindrome(s string) bool {
	characters := []rune(s)
	left, right := 0, len(characters)-1
	for left < right {
		if characters[left] != characters[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// kuohaoMatch 判断括号是否匹配
func kuohaoMatch(s string) bool {
	arraystr := make([]rune, 0, len(s))
	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			arraystr = append(arraystr, char)
		} else {
			if char != ')' && char != ']' && char != '}' {
				return false
			}
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
	if len(strs) == 0 {
		return ""
	}

	shortest := []rune(strs[0])
	for _, str := range strs[1:] {
		characters := []rune(str)
		if len(characters) < len(shortest) {
			shortest = characters
		}
	}

	for i, character := range shortest {
		for _, str := range strs {
			characters := []rune(str)
			if characters[i] != character {
				return string(shortest[:i])
			}
		}
	}

	return string(shortest)
}

func plusOne(digits []int) int {
	if len(digits) == 0 {
		return -1
	}

	maxInt := int(^uint(0) >> 1)
	result := 0
	for _, digit := range digits {
		if digit < 0 || digit > 9 || result > (maxInt-digit)/10 {
			return -1
		}
		result = result*10 + digit
	}
	if result == maxInt {
		return -1
	}
	return result + 1
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return 0
		}
	}

	uniqueCount := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueCount-1] {
			nums[uniqueCount] = nums[i]
			uniqueCount++
		}
	}
	return uniqueCount
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	for _, interval := range intervals {
		if len(interval) != 2 || interval[0] > interval[1] {
			return [][]int{}
		}
	}

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
	if len(nums) < 2 {
		return []int{}
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if (nums[j] <= 0 || sum >= nums[i]) && (nums[j] >= 0 || sum <= nums[i]) && sum == target {
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
