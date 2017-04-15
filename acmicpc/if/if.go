package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan() int {
	var x int
	fmt.Scanf("%d", &x)
	return x
}

func prob9498() {
	point := scan()

	if point >= 90 {
		fmt.Printf("A\n")
	} else if point >= 80 {
		fmt.Printf("B\n")
	} else if point >= 70 {
		fmt.Printf("C\n")
	} else if point >= 60 {
		fmt.Printf("D\n")
	} else {
		fmt.Printf("F\n")
	}
}

func prob10871() {
	var n, x int

	fmt.Scanf("%d %d", &n, &x)
	in, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return
	}
	s := strings.Fields(in)

	for i := 0; i < len(s); i++ {
		parsed, err := strconv.ParseInt(s[i], 10, 64)
		if err != nil {
			return
		}
		if int(parsed) < x {
			fmt.Printf("%d ", int(parsed))
		}
	}
}

func prob1546() {
	var sum, max, x int
	fmt.Scanf("%d", &x)
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	inputs := strings.Split(strings.Replace(input, "\n", "", -1), " ")
	for i := 0; i < x; i ++ {
		temp, _ := strconv.ParseInt(inputs[i], 10, 64)
		sum += int(temp)
		if int(temp) > max {
			max = int(temp)
		}
	}

	fmt.Printf("%.2f\n", float64(sum)*100/float64(max)/float64(x))
}

func prob1110() {
	var input int
	fmt.Scanf("%d", &input)
	first := input / 10
	second := input % 10

	newFirst := second
	newSecond := (first + second) % 10
	result := 1

	for newFirst != first || newSecond != second {
		oldFirst := newFirst
		newFirst = newSecond
		newSecond = (oldFirst + newSecond) % 10
		result += 1
	}

	fmt.Printf("%d\n", result)
}

func main() {
	// prob9498()
	// prob10871()
	// prob1546()
	prob1110()
}
