package main

import (
	stringUtil "gopractice/util"
	"math"
	"strconv"
)

func reverseInteger(i int) int {
	sign := ""
	if i < 0 {
		sign = "-"
		i *= -1
	}
	str := strconv.Itoa(i)

	reversed, _ := strconv.Atoi(sign + stringUtil.Reverse(str))
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}

/*
func Test() {
	println(reverseInteger(321))
	println(reverseInteger(-321))
}
*/
