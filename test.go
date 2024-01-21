package main

import (
	"fmt"
	"os"
)

func checkParam(args []string) bool {
	// ila kan number nta3 slice kter mn 9 aw 9al mn 9
	if len(args) != 9 {
		return false
	}
	// i heya vlaue dyal l3ard katkhalf 9 return false
	for i := range args {
		if len(args[i]) != 9 {
			return false
		}
		// i heya vlaue dyal toll katkhalf 9 return false
		for j := range args[i] {
			l := args[i][j]
			// ila kan cat -e mafihch hadchi return false
			if l != '.' && (l < '1' || l > '9') {
				return false
			}
		}
	}
	// l7alat lokhra returni true
	return true
}

// for parinting
func printGreed(grid [][]byte) {
	// y kan3tiw biha value l tol
	for y := 0; y < 9; y++ {
		// x kan3tiw biha value l l3ard
		for x := 0; x < 9; x++ {
			// la kant lkhana lwla kbar mn 0 we sghar mn tali tba3 space
			if x > 0 || x < x-1 {
				fmt.Print(" ")
			}
			// tb3li kol string whdha ftol we fl3ard
			fmt.Print(string(grid[x][y]))

		}
		fmt.Print("\n")
	}
}
func checkSudoku(walo int, grid [][]byte) bool {
	// hiya l value dyal morba3 kaaml bach i t printa kaml ujiw fih l ar9aam
	x := walo % 9
	y := walo / 9
	// kan3tiw hena 9ima l l3ard we kantchikiw kol clone whdha fl3ard
	for w := 0; w < 9; w++ {
		if w != x && grid[y][w] == grid[y][x] {
			return false
		}
	}
	// kan3tiw hena 9ima l toll we kantchikiw kol clone whdha ftoll
	for h := 0; h < 9; h++ {
		if h != y && grid[h][x] == grid[y][x] {
			return false
		}
	}
	// 9assm sudoku l 9 lkhanat 3 clones ftol we 3 fl3ard
	caseX := x / 3
	caseY := y / 3
	// 
	for h := 0; h < 3; h++ {
		for w := 0; w < 3; w++ {
			xx := caseX*3 + w
			yy := caseY*3 + h
			if (xx != x || yy != y) && grid[yy][xx] == grid[y][x] {
				return false
			}
		}
	}
	return true
}

func backTrack(walo int, grid [][]byte) bool {
	if walo == 81 {
		printGreed(grid)
		return true
	}
	x := walo % 9
	y := walo / 9
	if grid[y][x] != '.' {
		if checkSudoku(walo, grid) && backTrack(walo+1, grid) {
			return true
		}
	} else {
		for i := '1'; i <= '9'; i++ {
			grid[y][x] = byte(i)
			if checkSudoku(walo, grid) && backTrack(walo+1, grid) {
				return true
			}
		}
		grid[y][x] = '.'
	}
	return false
}

func resolveSudoku(args []string) bool {
	var grid [][]byte
	for i := range args {
		grid = append(grid, []byte(args[i]))
	}
	return backTrack(0, grid)
}

func main() {
	args := os.Args
	if !checkParam(args[1:]) || !resolveSudoku(args[1:]) {
		fmt.Println("Error")
	}
}
