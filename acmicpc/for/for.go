package main

import (
	"fmt"
	"strconv"
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

func prob1924() {
	var m, d, total int
	fmt.Scanf("%d %d", &m, &d)
	for i := 1; i < m; i++ {
		if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
			total += 31
		} else if i == 2 {
			total += 28
		} else {
			total += 30
		}
	}

	total += d

	remainder := total % 7

	if remainder == 0 {
		fmt.Printf("SUN\n")
	} else if remainder == 1 {
		fmt.Printf("MON\n")
	} else if remainder == 2 {
    fmt.Printf("TUE\n")
  } else if remainder == 3 {
    fmt.Printf("WED\n")
  } else if remainder == 4 {
    fmt.Printf("THU\n")
  } else if remainder == 5 {
    fmt.Printf("FRI\n")
  } else if remainder == 6 {
    fmt.Printf("SAT\n")
  }
}

func prob8393() {
	var x, result int
	fmt.Scanf("%d", &x)
	for i := 1; i <= x; i++ {
		result += i
	}
	fmt.Printf("%d\n", result)
}

func prob11720() {
	var n, result int64
	var x string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &x)

	for i :=  0; i < len(x); i++ {
		tmp, _ := strconv.ParseInt(string(x[i]), 10, 32)
		result += tmp
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
	// prob1924()
	// prob8393()
	prob11720()
	// prob11721()
}
