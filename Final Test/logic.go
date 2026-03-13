package main

import (
	"fmt"
	"strings"
)

func findpair(arr []int, sum int) []int {
	for i := 0; i < len(arr); i++ {
		arr1 := arr[i]

		for j := i + 1; j < len(arr); j++ {
			arr2 := arr[j]
			if arr1+arr2 == sum {
				return []int{arr1, arr2}
			}
		}
	}
	return []int{}
}

func isPalindrome(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")

	left := 0
	right := len(word) - 1

	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func wordCount(word string, wordtarget string) int {
	return strings.Count(word, wordtarget)
}

func incrementDigit(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func numTriangle(input int, start int) {
	for i := 1; i <= input; i++ {
		num := start + (i - 1)
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}

		fmt.Println()
	}
}

func main() {
	//1
	// fmt.Println(findpair([]int{2, 3, 5, 7, 11, 13}, 9))
	// fmt.Println(findpair([]int{2, 3, 5, 7, 11, 13}, 12))

	//2
	// fmt.Println(isPalindrome("MalaM"))
	// fmt.Println(isPalindrome("levEl"))
	// fmt.Println(isPalindrome("kaSur iNi ruSak"))
	// fmt.Println(isPalindrome("saYur"))

	//3
	// fmt.Println(wordCount("xhixhix", "x"))
	// fmt.Println(wordCount("xhixhix", "hi"))
	// fmt.Println(wordCount("mic", "mic"))
	// fmt.Println(wordCount("haha", "ho"))
	// fmt.Println(wordCount("xxxxyz", "xx"))

	//4
	// fmt.Println(incrementDigit([]int{1, 2, 3, 4}))
	// fmt.Println(incrementDigit([]int{1, 4, 8, 9}))
	// fmt.Println(incrementDigit([]int{9, 9, 9, 9}))

	//5
	numTriangle(7, 1)
	numTriangle(7, 5)
}
