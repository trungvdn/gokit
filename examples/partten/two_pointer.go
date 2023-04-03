package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func two_pointer() {
	nums := []int{3, 7, 1, 2, 8, 4, 5}
	// 1,2,3,4,5,7,8
	fmt.Println(findSumOfThree(nums, 10))
	var a byte = ' '
	fmt.Println(a)
	fmt.Println(reverseWords("Hello   world"))
}

func findSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))
	for i := 0; i < len(nums); i++ {
		low := i + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			if sum == target {
				return true
			} else if sum < target {
				low++
			} else if sum > target {
				high--
			}
		}
	}
	return false
}

func reverseWords(sentence string) string {
	// hello world
	// dlrow olleh
	// world hello
	// /We need to convert the input strings to lists of characters as strings are immutable in Go
	trimedSentense := trimString(sentence)
	sentenceBytes := []byte(trimedSentense)
	strLen := len(sentenceBytes)
	sentenceBytes = strRev(sentenceBytes, 0, strLen-1)
	strLen = len(sentenceBytes)
	start, end := 0, 0
	for {
		// find end

		for start < len(sentenceBytes) && sentenceBytes[start] == ' ' {
			start += 1
		}

		if start == strLen {
			break
		}

		end = start + 1
		for end < strLen && sentenceBytes[end] != ' ' {
			end += 1
		}

		sentenceBytes = strRev(sentenceBytes, start, end-1)
		start = end
	}
	return string(sentenceBytes)
}

func strRev(str []byte, startRev int, endRev int) []byte {
	for startRev < endRev {
		temp := str[startRev]
		str[startRev] = str[endRev]
		str[endRev] = temp
		startRev++
		endRev--
	}
	return str
}

func trimString(str string) string {
	str = strings.TrimSpace(str)
	regex := regexp.MustCompile("\\s+")
	str = regex.ReplaceAllString(str, " ")
	return str
}
