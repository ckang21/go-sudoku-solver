# Go Sudoku Solver

This is a beginner-friendly Sudoku solver written in Go. It uses **basic solving techniques** to fill in missing numbers while ensuring that each row, column, and 3x3 box follows Sudoku rules.

## 🔍 How It Works
- The program represents the Sudoku grid as a **9x9 slice of slices (`[][]int`)**.
- It follows a **simple backtracking algorithm** to find a solution.
- If a solution exists, the program prints the solved Sudoku board.

## 🖥️ Example Input
- 5 3 _ | _ 7 _ | _ _ _ | 6 _ _ | 1 9 5 | _ _ _ | _ 9 8 | _ _ _ | _ 6 _
- 8 _ _ | _ 6 _ | _ _ 3 | 4 _ _ | 8 _ 3 | _ _ 1 | 7 _ _ | _ 2 _ | _ _ 6
- _ 6 _ | _ _ _ | 2 8 _ | _ _ _ | 4 1 9 | _ _ 5 | _ _ _ | _ 8 _ | _ 7 9
- _(Underscores `_` represent empty spaces that need to be filled.)_

## 🎯 Features
- ✅ Solves standard **9x9 Sudoku puzzles**.
- ✅ Uses **basic logic** without advanced algorithms (beginner-friendly).
- ✅ Outputs the solved grid in a **readable format**.

## 🚀 How to Run
1. **Clone the repository:**
   ```sh
   git clone https://github.com/ckang21/go-sudoku-solver.git
   cd go-sudoku-solver
2. **Run the Program:**
    go run sudoku.go


## Future Improvements
- ✅ Add support for inputting a custom Sudoku grid
- ✅ Allow the user to choose between multiple puzzles
- ✅ Implement a visual interface using Go graphics

## Author
Built by Christian Kang as part of a coding challenge to implement a Sudoku solver in Go.