package main

import (
	"fmt"
	"slices"
)

func checkBlock(block []int) bool {
	copiedBlock := append([]int(nil), block...)
	slices.Sort(copiedBlock)
	for i, c := range copiedBlock {
		if i+1 != c {
			return false
		}
	}
	return true
}

func buildColumn(grid [][]int, colNumber int) []int {
	res := []int{}
	for i := 0; i < 9; i++ {
		res = append(res, grid[i][colNumber])
	}
	return res
}

func buildSubgrid(grid [][]int, rowStart, colStart int) []int {
	res := []int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res = append(res, grid[rowStart+i][colStart+j])
		}
	}
	return res
}

func isValidSudoku(grid [][]int) bool {
	for i := range grid{
		if (!checkBlock(buildColumn(grid, i))){
			return false
		}
	}
	for _, r := range grid{
		if !checkBlock(r) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !checkBlock(buildSubgrid(grid, i, j)) {
				return false
			}
		}
	}

	return true
}

func main() {
	goodGrid := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	badGrid := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 1},
	}

	fmt.Println("Good grid valid:", isValidSudoku(goodGrid))
	fmt.Println("Bad grid valid:", isValidSudoku(badGrid))
}
