package main

import "fmt"

func printGrid(grid [][]int) {
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
}

func checkALL(grid [][]int, row, col int) []int {
	if grid[row][col] != 0 { // We can check if it has a number already. We don't need to do anything if it does
		return []int{}
	}

	// Map to track what numbers are possible for that square
	possibilities := make(map[int]bool)
	for i := 1; i <= 9; i++ {
		possibilities[i] = true
	}

	// Check the row then column
	for i := 0; i < 9; i++ {
		if grid[row][i] != 0 {
			delete(possibilities, grid[row][i])
		}
		if grid[i][col] != 0 {
			delete(possibilities, grid[i][col])
		}
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

func findEmptyCell(grid [][]int) []int { // Meant to go through the grid to find us the coordinates of an empty cell
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				return []int{i, j} // Return the coordinates if we have one.
			}
		}
	}
	return []int{-1, -1} // Return this if all is filled.
}

func solveSudoku(grid [][]int) bool {
	coord := findEmptyCell(grid)

	if coord[0] == -1 && coord[1] == -1 { // If there are no empty cells we should get -1 and -1
		return true
	}

	// Get possibilities
	validAnswers := checkALL(grid, coord[0], coord[1])

	for _, num := range validAnswers {
		grid[coord[0]][coord[1]] = num // Set that number into the coordinates
		if solveSudoku(grid) {
			return true // If it gets us the answer stop
		}
		grid[coord[0]][coord[1]] = 0 // Undo
	}
	return false // If false then we can't solve this
}

func main() {
	grid := [][]int{
		{0, 0, 4, 2, 0, 1, 9, 6, 0},
		{9, 0, 6, 7, 4, 5, 2, 1, 0},
		{3, 0, 1, 9, 0, 0, 4, 0, 0},
		{0, 3, 5, 0, 7, 2, 8, 0, 6},
		{6, 0, 0, 0, 5, 0, 7, 0, 0},
		{2, 0, 8, 3, 9, 6, 1, 0, 0},
		{5, 4, 0, 0, 2, 0, 0, 8, 1},
		{8, 9, 2, 0, 1, 0, 0, 0, 4},
		{0, 0, 0, 4, 0, 7, 0, 2, 0},
	}
	fmt.Println("Starting Sudoku Solver...\n")
	printGrid(grid) // Print initial grid

	if solveSudoku(grid) {
		fmt.Println("\nSudoku Solved!\n")
		printGrid(grid) // Print solved grid
	} else {
		fmt.Println("\nNo solution exists.")
	}
}
