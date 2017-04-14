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

func main() {
	// prob9498()
	prob10871()
}
