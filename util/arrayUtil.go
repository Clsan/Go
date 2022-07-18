package util

func Print(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			print(arr[i][j])
		}
		println()
	}
	println()
}

func Contains[T comparable](arr []T, val T) bool {
	for _, elem := range arr {
		if elem == val {
			return true
		}
	}
	return false
}
