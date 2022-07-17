package main

//import arrayUtil "gopractice/util"

// https://leetcode.com/problems/interleaving-string/
func memoDfs(
	memo [][]int,
	s1, s2, s3 string,
	s1Idx, s2Idx int,
) bool {
	//arrayUtil.Print(memo)
	if memo[s1Idx][s2Idx] == 0 || memo[s1Idx][s2Idx] == 1 {
		return memo[s1Idx][s2Idx] == 1
	}

	if len(s1) == s1Idx {
		res := s2[s2Idx:] == s3[s1Idx+s2Idx:]
		if res {
			for i := s2Idx; i < len(s2)+1; i++ {
				memo[s1Idx][i] = 1
			}
		} else {
			for i := s2Idx; i < len(s2)+1; i++ {
				memo[s1Idx][i] = 0
			}
		}
		return res
	}
	if len(s2) == s2Idx {
		res := s1[s1Idx:] == s3[s1Idx+s2Idx:]
		if res {
			for i := s1Idx; i < len(s1)+1; i++ {
				memo[i][s2Idx] = 1
			}
		} else {
			for i := s1Idx; i < len(s1)+1; i++ {
				memo[i][s2Idx] = 0
			}
		}
		return res
	}

	s3Idx := s1Idx + s2Idx
	chooseS1, chooseS2 := false, false
	if len(s1) > s1Idx && s1[s1Idx] == s3[s3Idx] {
		chooseS1 = memoDfs(memo, s1, s2, s3, s1Idx+1, s2Idx)
	}
	if len(s2) > s2Idx && s2[s2Idx] == s3[s3Idx] {
		chooseS2 = memoDfs(memo, s1, s2, s3, s1Idx, s2Idx+1)
	}
	res := chooseS1 || chooseS2
	if res {
		memo[s1Idx][s2Idx] = 1
	} else {
		memo[s1Idx][s2Idx] = 0
	}
	return res
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	memo := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		memo[i] = make([]int, len(s2)+1)
		for j := 0; j < len(s2)+1; j++ {
			memo[i][j] = -1
		}
	}
	return memoDfs(memo, s1, s2, s3, 0, 0)
}

/*
func Test() {
	//println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	//println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))

	println(isInterleave(
		"abababababababababababababababababababababababababababababababababababababababababababababababababbb",
		"babababababababababababababababababababababababababababababababababababababababababababababababaaaba",
		"abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb",
	))
}
*/
