package main

import (
	"fmt"
	"math"
)

func prob2741() {
  var x int
	fmt.Scanf("%d", &x)
	for i := 1; i < x; i++ {
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("%d", x)
}

func prob2742() {
  var x int
	fmt.Scanf("%d", &x)
	for i := x; i > 0; i-- {
		fmt.Printf("%d\n", i)
	}
}

func prob2739() {
  var x int
	fmt.Scanf("%d", &x)
	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", x, i, x * i)
	}
}

func prob2438() {
	var x int
	fmt.Scanf("%d", &x)
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func prob2439() {
	var x int
	fmt.Scanf("%d", &x)
	for i := 1; i <= x; i++ {
		for j := 1; j <= x - i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("*")	
		}
		fmt.Printf("\n")
	}
}

func prob2440() {
	var x int
	fmt.Scanf("%d", &x)
	for i := x; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func prob2441() {
  var x int
	fmt.Scanf("%d", &x)
	for i := 1; i <= x; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf(" ")
		}
		for j := i; j <= x; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

// 100자리 숫자까지 들어올 수 있음
// Big Int 찾아보아야 함
func prob11720() {
	var n, x int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &x)

	if n == 0  || x / int(math.Pow(float64(10), float64(n))) != 0 {
		return
	}

	var div, result int
	for i := n - 1; i >= 0; i-- {
		div = int(math.Pow(float64(10), float64(i)))
		result += x / div

		x = x % div
	}

	fmt.Printf("%d\n", result)
}

func prob11721() {
  var str string
	fmt.Scanf("%s", &str)
	for i := 0; i < len(str); i++ {
		if i != 0 && i % 10 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%c", str[i])
	}
}

func main() {
	// prob2741()
	// prob2742()
	// prob2739()
	// prob2438()
	// prob2439()
	// prob2440()
	// prob2441()
	prob11720()
	// prob11721()
}
