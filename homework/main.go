package main

import (
	"fmt"
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

func main() {
	fmt.Println("homework 开始！！")

	numbers := []int{1, 20, 3, 3, 4, 20, 1}
	oneTimeNumber := findOnetimenumber(numbers)
	fmt.Println("第一题：")
	fmt.Printf("The one-time number in the array %v is: %d\n", numbers, oneTimeNumber)

	fmt.Println("第二题：")
	isPal := isPalindrome("221")
	fmt.Printf("Is 'isPala palindrome? %t\n", isPal)

}
