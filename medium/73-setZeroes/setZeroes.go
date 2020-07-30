package _3_setZeroes

func SetZeroes(matrix [][]int)  {
	row, col := len(matrix), len(matrix[0])
	rowSet := map[int]int{}
	colSet := map[int]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = i
				colSet[j] = j
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if _, rok := rowSet[i]; rok{
				matrix[i][j] = 0
			} else if _, cok := colSet[j]; cok{
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroes(matrix [][]int)  {
	row, col := len(matrix), len(matrix[0])
	rowSet, colSet := map[int]int{}, map[int]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = i
				colSet[j] = j
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if _, ok := rowSet[i]; ok {
				matrix[i][j] = 0
			} else if _, ok := colSet[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}