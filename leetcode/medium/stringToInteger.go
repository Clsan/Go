package main

import (
	arrayUtil "gopractice/util"
	"math"
	"strconv"
)

// https://leetcode.com/problems/string-to-integer-atoi
// Lots of edge cases.
func stringToInteger(s string) int {
	sign := 1
	result := ""
	metSign := false
	metNumber := false
	for _, c := range s {
		if arrayUtil.Contains([]rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}, c) {
			result += string(c)
			metNumber = true
		} else {
			if metNumber {
				break
			} else {
				if c == ' ' || c == '-' || c == '+' {
					if metSign || metNumber {
						return 0
					}
					if c == '-' || c == '+' {
						if c == '-' {
							sign = -1
						}
						metSign = true
					}
				} else {
					return 0
				}
			}
		}
	}
	if result != "" {
		num, _ := strconv.Atoi(result)
		if num*sign >= math.MaxInt32 {
			return math.MaxInt32
		}
		if num*sign <= math.MinInt32 {
			return math.MinInt32
		}
		return num * sign
	}
	return 0
}

func Test() {
	println(stringToInteger("42"))              // 42
	println(stringToInteger("   -42"))          // -42
	println(stringToInteger("4193 with words")) // 4193
	println(stringToInteger("004193"))          // 4193
	println(stringToInteger("words and 987"))   // 0
	println(stringToInteger("-91283472332"))    // -2147483648
	println(stringToInteger("+1"))              // 1
	println(stringToInteger("+-12"))            // 0
	println(stringToInteger("+   12"))          // 0
	println(stringToInteger("  12   2"))        // 12

}
