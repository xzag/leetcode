package the_k_weakest_rows_in_a_matrix

import "sort"

func soldiersCount(row []int) int {
	cnt := 0
	for _, col := range row {
		if col == 0 {
			break
		}

		cnt++
	}

	return cnt
}

func kWeakestRows(mat [][]int, k int) []int {
	for rowIndex, row := range mat {
		mat[rowIndex] = []int{soldiersCount(row), rowIndex}
	}

	sort.Slice(mat, func(i, j int) bool {
		if mat[i][0] == mat[j][0] {
			return mat[i][1] < mat[j][1]
		}
		return mat[i][0] < mat[j][0]
	})

	result := make([]int, k, k)
	for i := 0; i < k; i++ {
		result[i] = mat[i][1]
	}

	return result
}
