package main

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	maxLength := 1
	lastNonZeroDiff1 := 0
	for i := 1; i < len(nums)-1; i++ {
		diff1 := nums[i] - nums[i-1]
		if diff1 != 0 {
			lastNonZeroDiff1 = diff1
		}
		diff2 := nums[i+1] - nums[i]
		if diff1*diff2 < 0 {
			maxLength += 1
		} else {
			if lastNonZeroDiff1*diff2 < 0 {
				maxLength += 1
			}
		}
	}

	if maxLength == 1 && nums[0] == nums[len(nums)-1] {
		return maxLength
	} else {
		return maxLength + 1
	}
}

func main() {
	/*
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		nums := []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
		nums := []int{1, 7, 4, 9, 2, 5}
		nums := []int{1, 2}
	*/
	nums := []int{1, 1}
	print(wiggleMaxLength(nums))
}
