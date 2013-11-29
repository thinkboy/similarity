package similarity

func Similarity(str1, str2 string) int {
	cost := 0
	var (
		str1Tmp []string
		str2Tmp []string
	)

	for _, s := range str1 {
		str1Tmp = append(str1Tmp, string(s))
	}

	for _, s := range str2 {
		str2Tmp = append(str2Tmp, string(s))
	}

	len1 := len(str1Tmp)
	len2 := len(str2Tmp)

	matrix := make([][]int, 0, len1+1)
	for i := 0; i < len1+1; i++ {
		matrix = append(matrix, make([]int, len2+1, len2+1))
	}

	for i := 1; i <= len1; i++ {
		matrix[i][0] = i
	}

	for i := 1; i <= len2; i++ {
		matrix[0][i] = i
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if str1Tmp[i-1] != str2Tmp[j-1] {
				cost = 1
			}

			left := matrix[i][j-1]
			top := matrix[i-1][j]
			left_top := matrix[i-1][j-1]
			min := top
			if left < top {
				min = left
			}

			matrix[i][j] = min + 1
			if left_top <= min {
				matrix[i][j] = left_top + cost
			}
			cost = 0
		}
	}

	dis := matrix[len1-1][len2-1]

	div := len2
	if len1 > len2 {
		div = len1
	}

	return 100 - dis*100/div
}
