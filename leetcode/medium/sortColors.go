package main

func sortColors(colors []int) {
	for i := len(colors) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if colors[j] > colors[j+1] {
				temp := colors[j+1]
				colors[j+1] = colors[j]
				colors[j] = temp
			}
		}
	}
}

func main() {
	sortColors([]int{2, 1, 0})
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
