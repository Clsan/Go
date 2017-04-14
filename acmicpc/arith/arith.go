package main

import (
	"fmt"
)

func prob10998() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%d", x*y)
}

func prob1008() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%g", float64(x)/float64(y))
}

func prob10869() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%d\n", x+y)
	fmt.Printf("%d\n", x-y)
	fmt.Printf("%d\n", x*y)
	fmt.Printf("%d\n", x/y)
	fmt.Printf("%d", x%y)
}

func prob10430() {
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)
	fmt.Printf("%d\n", (x+y)%z)
	fmt.Printf("%d\n", (x%z+y%z)%z)
	fmt.Printf("%d\n", (x*y)%z)
	fmt.Printf("%d", (x%z*y%z)%z)
}

func prob2558() {
	var x, y int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Printf("%d", x+y)
}

func prob2839() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 4 || n == 7 {
		fmt.Printf("%d", -1)
		return
	}

	if n%5 == 0 {
		fmt.Printf("%d", n/5)
	} else if n%5 == 1 || n%5 == 3 {
		fmt.Printf("%d", n/5+1)
	} else if n%5 == 2 || n%5 == 4 {
		fmt.Printf("%d", n/5+2)
	}
}

func main() {
	// prob10998()
	// prob1008()
	// prob10869()
	// prob10430()
	// prob2558()
	prob2839()
}
