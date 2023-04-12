package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

// two pointers pattern uses two pointers to iterate over an array or list until the conditions of the problem are satisfied
func main() {
	// nums := []int{3, 7, 1, 2, 8, 4, 5}
	// // 1,2,3,4,5,7,8
	// fmt.Println(findSumOfThree(nums, 10))
	// var a byte = ' '
	// fmt.Println(a)
	// fmt.Println(reverseWords("Hello   world"))
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))
	// fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	// fmt.Println(reverseWords2("God Ding"))
	// reversePrefix("abcdefd", byte('d'))
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
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

func isPalindrome(inputString string) bool {
	// abcba
	left := 0
	right := len(inputString) - 1

	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func twoSum(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
	low := 0
	high := len(nums) - 1
	for low < high {
		sum := nums[low] + nums[high]
		if sum == target {
			return true
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return false
}

// func arithmeticTriplets(nums []int, diff int) int {
// 	//Input: nums = [0,1,4,6,7,10], diff = 3
// 	sort.Sort(sort.IntSlice(nums))
// 	count := 0
// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println(i)
// 		j := i + 1
// 		k := len(nums) - 1
// 		for j < k {
// 			ji := nums[j] - nums[i]
// 			if ji > diff {
// 				break
// 			}
// 			if ji == diff {
// 				kj := nums[k] - nums[j]
// 				if kj < diff {
// 					break
// 				}
// 				if kj == diff {
// 					count++
// 					k--
// 				} else if kj > diff {
// 					k--
// 				}
// 			} else if ji < diff {
// 				j++
// 			}
// 		}
// 	}
// 	return count
// }

func arithmeticTriplets(nums []int, diff int) int {
	//Input: nums = [0,1,4,6,7,10], diff = 3
	sort.Sort(sort.IntSlice(nums))
	count := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
		j := i + 1
		for j < len(nums)-2 {
			ji := nums[j] - nums[i]
			if ji > diff {
				break
			}
			if ji == diff {
				for k := j + 1; k < len(nums); k++ {
					kj := nums[k] - nums[j]
					if kj == diff {
						count++
					} else if kj > diff {
						break
					}
				}
				j++
			} else {
				j++
			}
		}
	}
	return count
}

func reverseWords2(s string) string {
	sByte := []byte(s)
	start, end := 0, 0
	for {
		for start < len(sByte) && sByte[start] == ' ' {
			start++
		}
		if start == len(sByte) {
			break
		}
		end = start + 1
		for end < len(sByte) && sByte[end] != ' ' {
			end++
		}
		strRev(sByte, start, end-1)
		start = end
	}
	return string(sByte)
}

func flipAndInvertImage(image [][]int) [][]int {
	//image = [[1,1,0],[1,0,1],[0,0,0]]
	for i := 0; i < len(image); i++ {
		img := image[i]
		left := 0
		right := len(img) - 1
		for left < right {
			temp := img[left]
			img[left] = img[right]
			img[right] = temp
			left++
			right--
		}
		invertImg(img)
	}
	return image
}

func invertImg(img []int) []int {
	for i := 0; i < len(img); i++ {
		if img[i] == 0 {
			img[i] = 1
		} else if img[i] == 1 {
			img[i] = 0
		}
	}
	return img
}

func reversePrefix(word string, ch byte) string {
	//Input: word = "abcdefd", ch = "d"
	//Output: "dcbaefd"
	wordByte := []byte(word)
	start, end := 0, 0
	for i := 0; i < len(wordByte); i++ {
		if wordByte[i] == ch {
			end = i
			break
		}
	}
	for start < end {
		temp := wordByte[start]
		wordByte[start] = wordByte[end]
		wordByte[end] = temp
		start++
		end--
	}
	return string(wordByte)
}

func firstPalindrome(words []string) string {
	for i := 0; i < len(words); i++ {
		w := words[i]
		left := 0
		right := len(w) - 1
		isPalindrome := true
		for left < right {
			if w[left] != w[right] {
				isPalindrome = false
			}
			left++
			right--
		}
		if isPalindrome {
			return w
		}
	}
	return ""
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}

func sortArrayByParity(nums []int) []int {
	//Input: nums = [3,1,2,4]
	//              {-4, -1, 0, 3, 10}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]%2 != 0 && nums[right]%2 == 0 {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
		if nums[left]%2 == 0 {
			left++
		}
		if nums[right]%2 != 0 {
			right--
		}
	}
	return nums
}

func sortedSquares(nums []int) []int {
	//Input: nums = [-5,-3,-2,-1]
	left := 0
	right := len(nums) - 1
	for right > left {
		leftAbs := math.Abs(float64(nums[left]))
		rightAbs := math.Abs(float64(nums[right]))
		if leftAbs >= rightAbs {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			right--
		} else {
			right--
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = int(math.Pow(float64(nums[i]), 2))
	}
	return nums
}
