package main

import (
  "fmt"
	"bufio"
	"os"
)

func prob1000() int {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	return x + y
}

func prob1001() int {
  var x, y int
	fmt.Scanf("%d %d", &x, &y)
	return x - y
}

// failed! why?
func prob7287() {
	fmt.Printf("%d\nclsan", 2)
}

func prob10172() {
	fmt.Println("|\\_/|")
  fmt.Println("|q p|   /}")
  fmt.Println("( 0 )\"\"\"\\")
  fmt.Println("|\"^\"`    |")
  fmt.Printf("||_/=\\\\__|")
}

func prob10718() {
	fmt.Println("강한친구 대한육군")
	fmt.Print("강한친구 대한육군")
}

// same as prob11719()
func prob11718() {
	in := bufio.NewReader(os.Stdin)
	var str string
	var err error

	for err == nil {
		str, err = in.ReadString('\n')
		// fmt.Printf(str)
		fmt.Printf("", err == nil, str)
	}
}

func main() {
  // fmt.Printf("%d", prob1000())
	// fmt.Printf("%d", prob1001())
	prob7287()
	// prob10172()
	// prob10718()
	// prob11718()
}
