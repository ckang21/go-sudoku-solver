package main

import "fmt"

func checkALL(grid [][]int, row, col int) []int {
	if grid[row][col] != 0 { // We can check if it has a number already. We don't need to do anything if it does
		return []int{}
	}

	// Map to track what numbers are possible for that square
	possibilities := make(map[int]bool)
	for i := 1; i <= 8; i++ {
		possibilities[i] = true
	}

	// Check the row then column
	for i := 0; i <= 8; i++ {
		delete(possibilities, grid[row][i]) // Get rid of any found numbers in the row
		delete(possibilities, grid[i][col]) // Get rid of any found numbers in the column
	}

	// Check the 3x3 square
	startRow, startCol := (row/3)*3, (col/3)*3 // We need the right 3x3 square so say we look at coord [8, 4] will get us 6, 3. That will be the top left of the bottom middle square
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			delete(possibilities, grid[startRow+i][startCol+j])
		}
	}

	// Get the possible answers into a slice of answers
	validNumbers := []int{}
	for num := range possibilities {
		validNumbers = append(validNumbers, num)
	}
	return validNumbers
}

func main() {
	grid := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	row, col := 0, 2

	possibileValues := checkALL(grid, row, col)
	fmt.Printf("Possible values for cell (%d, %d): %v\n", row, col, possibileValues)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Print(grid[i][j], " ")
			if (j+1)%3 == 0 && j < 8 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (i+1)%3 == 0 && i < 8 {
			fmt.Println("----------------------------")
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			possibileValues := checkALL(grid, i, j)
			if len(possibileValues) != 0 {
				fmt.Printf("Possible values for cell (%d, %d): %v\n", i, j, possibileValues)
			}
		}
	}
}
