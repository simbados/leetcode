package main

import "fmt"

type KeyVal [10]byte

func isValidSudoku(board [][]byte) bool {
	columnMap := make(map[int]KeyVal)
	matrix := make(map[int]KeyVal)
	for index, row := range board {
		var key KeyVal
		for i, element := range row {
			if element != '.' {
				// Row
				if ok := key['9'-element]; ok >= 1 {
					return false
				}
				key['9'-element]++
				// Column
				columnDone := handleMap(columnMap, i, element)
				if columnDone {
					return false
				}
				// Matrix
				matrixIndex := ((i / 3) + 1) + ((index / 3) * 3)
				matrixDone := handleMap(matrix, matrixIndex, element)
				if matrixDone {
					return false
				}
			}
		}
	}
	fmt.Println(columnMap)
	return true
}

func handleMap(matrix map[int]KeyVal, matrixIndex int, element byte) bool {
	if val, ok := matrix[matrixIndex]; ok {
		if ok := val['9'-element]; ok >= 1 {
			return true
		}
		val['9'-element]++
		matrix[matrixIndex] = val
	} else {
		var matrixKey KeyVal
		matrixKey['9'-element]++
		matrix[matrixIndex] = matrixKey
	}
	return false
}
